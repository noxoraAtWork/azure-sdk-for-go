// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package common

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/Azure/azure-sdk-for-go/tools/generator/autorest"
	"github.com/Azure/azure-sdk-for-go/tools/generator/autorest/model"
	"github.com/Azure/azure-sdk-for-go/tools/generator/cmd/template"
	"github.com/Azure/azure-sdk-for-go/tools/generator/common"
	"github.com/Azure/azure-sdk-for-go/tools/generator/repo"
	"github.com/Azure/azure-sdk-for-go/tools/internal/exports"
	"github.com/Masterminds/semver"
)

type GenerateContext struct {
	SDKPath        string
	SDKRepo        *repo.SDKRepository
	SpecPath       string
	SpecCommitHash string
}

type GenerateResult struct {
	Version        string
	RPName         string
	PackageName    string
	PackageAbsPath string
	Changelog      model.Changelog
	ChangelogMD    string
}

func (ctx GenerateContext) GenerateForAutomation(readme, repo, specFolderName string) ([]GenerateResult, []error) {
	absReadme := filepath.Join(ctx.SpecPath, readme)
	absReadmeGo := filepath.Join(filepath.Dir(absReadme), "readme.go.md")

	var result []GenerateResult
	var errors []error

	log.Printf("Get all namespaces from readme file")
	rpMap, err := ReadV2ModuleNameToGetNamespace(absReadmeGo)
	if err != nil {
		return nil, []error{
			fmt.Errorf("cannot get rp and namespaces from readme '%s': %+v", readme, err),
		}
	}

	for rpName, namespaceNames := range rpMap {
		for _, namespaceName := range namespaceNames {
			log.Printf("Process rp: %s, namespace: %s", rpName, namespaceName)
			singleResult, err := ctx.GenerateForSingleRPNamespace(rpName, namespaceName, "", "", repo, specFolderName)
			if err != nil {
				errors = append(errors, err)
				continue
			}
			result = append(result, *singleResult)
		}
	}
	return result, errors
}

func (ctx GenerateContext) GenerateForSingleRPNamespace(rpName, namespaceName, specficPackageTitle, specficVersion, specficRepoURL, specFolderName string) (*GenerateResult, error) {
	packagePath := filepath.Join(ctx.SDKPath, "sdk", rpName, namespaceName)
	changelogPath := filepath.Join(packagePath, common.ChangelogFilename)

	onBoard := false
	var oriExports exports.Content
	if _, err := os.Stat(changelogPath); os.IsNotExist(err) {
		onBoard = true
		log.Printf("Package '%s' changelog not exist, do onboard process", packagePath)

		if specficPackageTitle == "" {
			specficPackageTitle = strings.Title(rpName)
		}

		log.Printf("Use template to generate new rp folder and basic package files...")
		if err = template.GeneratePackageByTemplate(rpName, namespaceName, template.Flags{
			SDKRoot:      ctx.SDKPath,
			TemplatePath: "tools/generator/template/rpName/packageName",
			PackageTitle: specficPackageTitle,
			Commit:       ctx.SpecCommitHash,
		}); err != nil {
			return nil, err
		}
	} else {
		log.Printf("Package '%s' existed, do update process", packagePath)

		log.Printf("Get ori exports for changelog generation...")
		if err := (*ctx.SDKRepo).Add(fmt.Sprintf("sdk\\%s\\%s", rpName, namespaceName)); err != nil {
			return nil, err
		}
		if err := (*ctx.SDKRepo).Stash(); err != nil {
			log.Printf("failed to stash changes: %+v", err)
		}
		oriExports, err = exports.Get(packagePath)
		if err != nil {
			return nil, err
		}
		if err := (*ctx.SDKRepo).StashPop(); err != nil {
			log.Printf("failed to stash pop changes: %+v", err)
		}

		log.Printf("Remove all the files that start with `zz_generated_`...")
		if err = CleanSDKGeneratedFiles(packagePath); err != nil {
			return nil, err
		}
	}

	// same step for onboard and update
	if ctx.SpecCommitHash == "" {
		log.Printf("Change swagger config in `autorest.md` according to local path...")
		autorestMdPath := filepath.Join(packagePath, "autorest.md")
		if err := ChangeConfigWithLocalPath(autorestMdPath, ctx.SpecPath, specFolderName); err != nil {
			return nil, err
		}
	} else {
		log.Printf("Change swagger config in `autorest.md` according to repo URL and commit ID...")
		autorestMdPath := filepath.Join(packagePath, "autorest.md")
		if err := ChangeConfigWithCommitID(autorestMdPath, specficRepoURL, ctx.SpecCommitHash, specFolderName); err != nil {
			return nil, err
		}
	}

	log.Printf("Run `go generate` to regenerate the code...")
	if err := ExecuteGoGenerate(packagePath); err != nil {
		return nil, err
	}

	log.Printf("Run `goimports` to refine the code import...")
	if err := ExecuteGoimports(packagePath); err != nil {
		return nil, err
	}

	log.Printf("Generate changelog for package...")
	newExports, err := exports.Get(packagePath)
	if err != nil {
		return nil, err
	}
	changelog, err := autorest.GetChangelogForPackage(&oriExports, &newExports)
	if err != nil {
		return nil, err
	}

	if onBoard {
		log.Printf("Replace {{NewClientMethod}} placeholder in the README.md ")
		if err = ReplaceNewClientMethodPlaceholder(packagePath, newExports); err != nil {
			return nil, err
		}

		return &GenerateResult{
			Version:        "0.1.0",
			RPName:         rpName,
			PackageName:    namespaceName,
			PackageAbsPath: packagePath,
			Changelog:      *changelog,
			ChangelogMD:    changelog.ToCompactMarkdown(),
		}, nil
	} else {
		log.Printf("Calculate new version...")
		var version *semver.Version
		if specficVersion == "" {
			version, err = CalculateNewVersion(changelog, packagePath)
			if err != nil {
				return nil, err
			}
		} else {
			log.Printf("Use specfic version: %s", specficVersion)
			version, err = semver.NewVersion(specficVersion)
			if err != nil {
				return nil, err
			}
		}

		log.Printf("Add changelog to file...")
		changelogMd, err := AddChangelogToFile(changelog, version, packagePath)
		if err != nil {
			return nil, err
		}

		log.Printf("Remove all the files that start with `zz_generated_`...")
		if err = CleanSDKGeneratedFiles(packagePath); err != nil {
			return nil, err
		}

		log.Printf("Replace version in autorest.md...")
		if err = ReplaceVersion(packagePath, version.String()); err != nil {
			return nil, err
		}

		log.Printf("Run `go generate` to regenerate the code for new version...")
		if err = ExecuteGoGenerate(packagePath); err != nil {
			return nil, err
		}

		return &GenerateResult{
			Version:        version.String(),
			RPName:         rpName,
			PackageName:    namespaceName,
			PackageAbsPath: packagePath,
			Changelog:      *changelog,
			ChangelogMD:    changelogMd,
		}, nil
	}
}
