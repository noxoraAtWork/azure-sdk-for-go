//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armmediaservices_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mediaservices/armmediaservices/v3"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/mediaservices/resource-manager/Microsoft.Media/Metadata/stable/2022-08-01/examples/streaming-policies-list.json
func ExampleStreamingPoliciesClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armmediaservices.NewStreamingPoliciesClient("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := client.NewListPager("contoso", "contosomedia", &armmediaservices.StreamingPoliciesClientListOptions{Filter: nil,
		Top:     nil,
		Orderby: nil,
	})
	for pager.More() {
		nextResult, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range nextResult.Value {
			// TODO: use page item
			_ = v
		}
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/mediaservices/resource-manager/Microsoft.Media/Metadata/stable/2022-08-01/examples/streaming-policy-get-by-name.json
func ExampleStreamingPoliciesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armmediaservices.NewStreamingPoliciesClient("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.Get(ctx, "contoso", "contosomedia", "clearStreamingPolicy", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/mediaservices/resource-manager/Microsoft.Media/Metadata/stable/2022-08-01/examples/streaming-policies-create-commonEncryptionCbcs-clearKeyEncryption.json
func ExampleStreamingPoliciesClient_Create_createsAStreamingPolicyWithClearKeyEncryptionInCommonEncryptionCbcs() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armmediaservices.NewStreamingPoliciesClient("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = client.Create(ctx, "contoso", "contosomedia", "UserCreatedSecureStreamingPolicyWithCommonEncryptionCbcsOnly", armmediaservices.StreamingPolicy{
		Properties: &armmediaservices.StreamingPolicyProperties{
			CommonEncryptionCbcs: &armmediaservices.CommonEncryptionCbcs{
				ClearKeyEncryptionConfiguration: &armmediaservices.ClearKeyEncryptionConfiguration{
					CustomKeysAcquisitionURLTemplate: to.Ptr("https://contoso.com/{AlternativeMediaId}/clearkey/"),
				},
				ContentKeys: &armmediaservices.StreamingPolicyContentKeys{
					DefaultKey: &armmediaservices.DefaultKey{
						Label: to.Ptr("cbcsDefaultKey"),
					},
				},
				EnabledProtocols: &armmediaservices.EnabledProtocols{
					Dash:            to.Ptr(false),
					Download:        to.Ptr(false),
					Hls:             to.Ptr(true),
					SmoothStreaming: to.Ptr(false),
				},
			},
			DefaultContentKeyPolicyName: to.Ptr("PolicyWithMultipleOptions"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/mediaservices/resource-manager/Microsoft.Media/Metadata/stable/2022-08-01/examples/streaming-policies-create-commonEncryptionCenc-clearKeyEncryption.json
func ExampleStreamingPoliciesClient_Create_createsAStreamingPolicyWithClearKeyEncryptionInCommonEncryptionCenc() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armmediaservices.NewStreamingPoliciesClient("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = client.Create(ctx, "contoso", "contosomedia", "UserCreatedSecureStreamingPolicyWithCommonEncryptionCencOnly", armmediaservices.StreamingPolicy{
		Properties: &armmediaservices.StreamingPolicyProperties{
			CommonEncryptionCenc: &armmediaservices.CommonEncryptionCenc{
				ClearKeyEncryptionConfiguration: &armmediaservices.ClearKeyEncryptionConfiguration{
					CustomKeysAcquisitionURLTemplate: to.Ptr("https://contoso.com/{AlternativeMediaId}/clearkey/"),
				},
				ClearTracks: []*armmediaservices.TrackSelection{
					{
						TrackSelections: []*armmediaservices.TrackPropertyCondition{
							{
								Operation: to.Ptr(armmediaservices.TrackPropertyCompareOperationEqual),
								Property:  to.Ptr(armmediaservices.TrackPropertyTypeFourCC),
								Value:     to.Ptr("hev1"),
							}},
					}},
				ContentKeys: &armmediaservices.StreamingPolicyContentKeys{
					DefaultKey: &armmediaservices.DefaultKey{
						Label: to.Ptr("cencDefaultKey"),
					},
				},
				EnabledProtocols: &armmediaservices.EnabledProtocols{
					Dash:            to.Ptr(true),
					Download:        to.Ptr(false),
					Hls:             to.Ptr(false),
					SmoothStreaming: to.Ptr(true),
				},
			},
			DefaultContentKeyPolicyName: to.Ptr("PolicyWithPlayReadyOptionAndOpenRestriction"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/mediaservices/resource-manager/Microsoft.Media/Metadata/stable/2022-08-01/examples/streaming-policies-create-clear.json
func ExampleStreamingPoliciesClient_Create_createsAStreamingPolicyWithClearStreaming() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armmediaservices.NewStreamingPoliciesClient("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = client.Create(ctx, "contoso", "contosomedia", "UserCreatedClearStreamingPolicy", armmediaservices.StreamingPolicy{
		Properties: &armmediaservices.StreamingPolicyProperties{
			NoEncryption: &armmediaservices.NoEncryption{
				EnabledProtocols: &armmediaservices.EnabledProtocols{
					Dash:            to.Ptr(true),
					Download:        to.Ptr(true),
					Hls:             to.Ptr(true),
					SmoothStreaming: to.Ptr(true),
				},
			},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/mediaservices/resource-manager/Microsoft.Media/Metadata/stable/2022-08-01/examples/streaming-policies-create-commonEncryptionCbcs-only.json
func ExampleStreamingPoliciesClient_Create_createsAStreamingPolicyWithCommonEncryptionCbcsOnly() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armmediaservices.NewStreamingPoliciesClient("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = client.Create(ctx, "contoso", "contosomedia", "UserCreatedSecureStreamingPolicyWithCommonEncryptionCbcsOnly", armmediaservices.StreamingPolicy{
		Properties: &armmediaservices.StreamingPolicyProperties{
			CommonEncryptionCbcs: &armmediaservices.CommonEncryptionCbcs{
				ContentKeys: &armmediaservices.StreamingPolicyContentKeys{
					DefaultKey: &armmediaservices.DefaultKey{
						Label: to.Ptr("cbcsDefaultKey"),
					},
				},
				Drm: &armmediaservices.CbcsDrmConfiguration{
					FairPlay: &armmediaservices.StreamingPolicyFairPlayConfiguration{
						AllowPersistentLicense:              to.Ptr(true),
						CustomLicenseAcquisitionURLTemplate: to.Ptr("https://contoso.com/{AssetAlternativeId}/fairplay/{ContentKeyId}"),
					},
				},
				EnabledProtocols: &armmediaservices.EnabledProtocols{
					Dash:            to.Ptr(false),
					Download:        to.Ptr(false),
					Hls:             to.Ptr(true),
					SmoothStreaming: to.Ptr(false),
				},
			},
			DefaultContentKeyPolicyName: to.Ptr("PolicyWithMultipleOptions"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/mediaservices/resource-manager/Microsoft.Media/Metadata/stable/2022-08-01/examples/streaming-policies-create-commonEncryptionCenc-only.json
func ExampleStreamingPoliciesClient_Create_createsAStreamingPolicyWithCommonEncryptionCencOnly() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armmediaservices.NewStreamingPoliciesClient("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = client.Create(ctx, "contoso", "contosomedia", "UserCreatedSecureStreamingPolicyWithCommonEncryptionCencOnly", armmediaservices.StreamingPolicy{
		Properties: &armmediaservices.StreamingPolicyProperties{
			CommonEncryptionCenc: &armmediaservices.CommonEncryptionCenc{
				ClearTracks: []*armmediaservices.TrackSelection{
					{
						TrackSelections: []*armmediaservices.TrackPropertyCondition{
							{
								Operation: to.Ptr(armmediaservices.TrackPropertyCompareOperationEqual),
								Property:  to.Ptr(armmediaservices.TrackPropertyTypeFourCC),
								Value:     to.Ptr("hev1"),
							}},
					}},
				ContentKeys: &armmediaservices.StreamingPolicyContentKeys{
					DefaultKey: &armmediaservices.DefaultKey{
						Label: to.Ptr("cencDefaultKey"),
					},
				},
				Drm: &armmediaservices.CencDrmConfiguration{
					PlayReady: &armmediaservices.StreamingPolicyPlayReadyConfiguration{
						CustomLicenseAcquisitionURLTemplate: to.Ptr("https://contoso.com/{AssetAlternativeId}/playready/{ContentKeyId}"),
						PlayReadyCustomAttributes:           to.Ptr("PlayReady CustomAttributes"),
					},
					Widevine: &armmediaservices.StreamingPolicyWidevineConfiguration{
						CustomLicenseAcquisitionURLTemplate: to.Ptr("https://contoso.com/{AssetAlternativeId}/widevine/{ContentKeyId"),
					},
				},
				EnabledProtocols: &armmediaservices.EnabledProtocols{
					Dash:            to.Ptr(true),
					Download:        to.Ptr(false),
					Hls:             to.Ptr(false),
					SmoothStreaming: to.Ptr(true),
				},
			},
			DefaultContentKeyPolicyName: to.Ptr("PolicyWithPlayReadyOptionAndOpenRestriction"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/mediaservices/resource-manager/Microsoft.Media/Metadata/stable/2022-08-01/examples/streaming-policies-create-envelopeEncryption-only.json
func ExampleStreamingPoliciesClient_Create_createsAStreamingPolicyWithEnvelopeEncryptionOnly() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armmediaservices.NewStreamingPoliciesClient("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = client.Create(ctx, "contoso", "contosomedia", "UserCreatedSecureStreamingPolicyWithEnvelopeEncryptionOnly", armmediaservices.StreamingPolicy{
		Properties: &armmediaservices.StreamingPolicyProperties{
			DefaultContentKeyPolicyName: to.Ptr("PolicyWithClearKeyOptionAndTokenRestriction"),
			EnvelopeEncryption: &armmediaservices.EnvelopeEncryption{
				ContentKeys: &armmediaservices.StreamingPolicyContentKeys{
					DefaultKey: &armmediaservices.DefaultKey{
						Label: to.Ptr("aesDefaultKey"),
					},
				},
				CustomKeyAcquisitionURLTemplate: to.Ptr("https://contoso.com/{AssetAlternativeId}/envelope/{ContentKeyId}"),
				EnabledProtocols: &armmediaservices.EnabledProtocols{
					Dash:            to.Ptr(true),
					Download:        to.Ptr(false),
					Hls:             to.Ptr(true),
					SmoothStreaming: to.Ptr(true),
				},
			},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/mediaservices/resource-manager/Microsoft.Media/Metadata/stable/2022-08-01/examples/streaming-policies-create-secure-streaming.json
func ExampleStreamingPoliciesClient_Create_createsAStreamingPolicyWithSecureStreaming() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armmediaservices.NewStreamingPoliciesClient("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = client.Create(ctx, "contoso", "contosomedia", "UserCreatedSecureStreamingPolicy", armmediaservices.StreamingPolicy{
		Properties: &armmediaservices.StreamingPolicyProperties{
			CommonEncryptionCbcs: &armmediaservices.CommonEncryptionCbcs{
				ContentKeys: &armmediaservices.StreamingPolicyContentKeys{
					DefaultKey: &armmediaservices.DefaultKey{
						Label: to.Ptr("cbcsDefaultKey"),
					},
				},
				Drm: &armmediaservices.CbcsDrmConfiguration{
					FairPlay: &armmediaservices.StreamingPolicyFairPlayConfiguration{
						AllowPersistentLicense:              to.Ptr(true),
						CustomLicenseAcquisitionURLTemplate: to.Ptr("https://contoso.com/{AssetAlternativeId}/fairplay/{ContentKeyId}"),
					},
				},
				EnabledProtocols: &armmediaservices.EnabledProtocols{
					Dash:            to.Ptr(false),
					Download:        to.Ptr(false),
					Hls:             to.Ptr(true),
					SmoothStreaming: to.Ptr(false),
				},
			},
			CommonEncryptionCenc: &armmediaservices.CommonEncryptionCenc{
				ClearTracks: []*armmediaservices.TrackSelection{
					{
						TrackSelections: []*armmediaservices.TrackPropertyCondition{
							{
								Operation: to.Ptr(armmediaservices.TrackPropertyCompareOperationEqual),
								Property:  to.Ptr(armmediaservices.TrackPropertyTypeFourCC),
								Value:     to.Ptr("hev1"),
							}},
					}},
				ContentKeys: &armmediaservices.StreamingPolicyContentKeys{
					DefaultKey: &armmediaservices.DefaultKey{
						Label: to.Ptr("cencDefaultKey"),
					},
				},
				Drm: &armmediaservices.CencDrmConfiguration{
					PlayReady: &armmediaservices.StreamingPolicyPlayReadyConfiguration{
						CustomLicenseAcquisitionURLTemplate: to.Ptr("https://contoso.com/{AssetAlternativeId}/playready/{ContentKeyId}"),
						PlayReadyCustomAttributes:           to.Ptr("PlayReady CustomAttributes"),
					},
					Widevine: &armmediaservices.StreamingPolicyWidevineConfiguration{
						CustomLicenseAcquisitionURLTemplate: to.Ptr("https://contoso.com/{AssetAlternativeId}/widevine/{ContentKeyId"),
					},
				},
				EnabledProtocols: &armmediaservices.EnabledProtocols{
					Dash:            to.Ptr(true),
					Download:        to.Ptr(false),
					Hls:             to.Ptr(false),
					SmoothStreaming: to.Ptr(true),
				},
			},
			DefaultContentKeyPolicyName: to.Ptr("PolicyWithMultipleOptions"),
			EnvelopeEncryption: &armmediaservices.EnvelopeEncryption{
				ContentKeys: &armmediaservices.StreamingPolicyContentKeys{
					DefaultKey: &armmediaservices.DefaultKey{
						Label: to.Ptr("aesDefaultKey"),
					},
				},
				CustomKeyAcquisitionURLTemplate: to.Ptr("https://contoso.com/{AssetAlternativeId}/envelope/{ContentKeyId}"),
				EnabledProtocols: &armmediaservices.EnabledProtocols{
					Dash:            to.Ptr(true),
					Download:        to.Ptr(false),
					Hls:             to.Ptr(true),
					SmoothStreaming: to.Ptr(true),
				},
			},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/mediaservices/resource-manager/Microsoft.Media/Metadata/stable/2022-08-01/examples/streaming-policies-delete.json
func ExampleStreamingPoliciesClient_Delete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armmediaservices.NewStreamingPoliciesClient("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = client.Delete(ctx, "contoso", "contosomedia", "secureStreamingPolicyWithCommonEncryptionCbcsOnly", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}
