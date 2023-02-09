//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armrecoveryservicesbackup

// BMSPrepareDataMoveOperationResultClientGetResponse contains the response from method BMSPrepareDataMoveOperationResultClient.Get.
type BMSPrepareDataMoveOperationResultClientGetResponse struct {
	VaultStorageConfigOperationResultResponseClassification
}

// UnmarshalJSON implements the json.Unmarshaller interface for type BMSPrepareDataMoveOperationResultClientGetResponse.
func (b *BMSPrepareDataMoveOperationResultClientGetResponse) UnmarshalJSON(data []byte) error {
	res, err := unmarshalVaultStorageConfigOperationResultResponseClassification(data)
	if err != nil {
		return err
	}
	b.VaultStorageConfigOperationResultResponseClassification = res
	return nil
}

// BackupEnginesClientGetResponse contains the response from method BackupEnginesClient.Get.
type BackupEnginesClientGetResponse struct {
	BackupEngineBaseResource
}

// BackupEnginesClientListResponse contains the response from method BackupEnginesClient.NewListPager.
type BackupEnginesClientListResponse struct {
	BackupEngineBaseResourceList
}

// BackupJobsClientListResponse contains the response from method BackupJobsClient.NewListPager.
type BackupJobsClientListResponse struct {
	JobResourceList
}

// BackupOperationResultsClientGetResponse contains the response from method BackupOperationResultsClient.Get.
type BackupOperationResultsClientGetResponse struct {
	// placeholder for future response values
}

// BackupOperationStatusesClientGetResponse contains the response from method BackupOperationStatusesClient.Get.
type BackupOperationStatusesClientGetResponse struct {
	OperationStatus
}

// BackupPoliciesClientListResponse contains the response from method BackupPoliciesClient.NewListPager.
type BackupPoliciesClientListResponse struct {
	ProtectionPolicyResourceList
}

// BackupProtectableItemsClientListResponse contains the response from method BackupProtectableItemsClient.NewListPager.
type BackupProtectableItemsClientListResponse struct {
	WorkloadProtectableItemResourceList
}

// BackupProtectedItemsClientListResponse contains the response from method BackupProtectedItemsClient.NewListPager.
type BackupProtectedItemsClientListResponse struct {
	ProtectedItemResourceList
}

// BackupProtectionContainersClientListResponse contains the response from method BackupProtectionContainersClient.NewListPager.
type BackupProtectionContainersClientListResponse struct {
	ProtectionContainerResourceList
}

// BackupProtectionIntentClientListResponse contains the response from method BackupProtectionIntentClient.NewListPager.
type BackupProtectionIntentClientListResponse struct {
	ProtectionIntentResourceList
}

// BackupResourceEncryptionConfigsClientGetResponse contains the response from method BackupResourceEncryptionConfigsClient.Get.
type BackupResourceEncryptionConfigsClientGetResponse struct {
	BackupResourceEncryptionConfigExtendedResource
}

// BackupResourceEncryptionConfigsClientUpdateResponse contains the response from method BackupResourceEncryptionConfigsClient.Update.
type BackupResourceEncryptionConfigsClientUpdateResponse struct {
	// placeholder for future response values
}

// BackupResourceStorageConfigsNonCRRClientGetResponse contains the response from method BackupResourceStorageConfigsNonCRRClient.Get.
type BackupResourceStorageConfigsNonCRRClientGetResponse struct {
	BackupResourceConfigResource
}

// BackupResourceStorageConfigsNonCRRClientPatchResponse contains the response from method BackupResourceStorageConfigsNonCRRClient.Patch.
type BackupResourceStorageConfigsNonCRRClientPatchResponse struct {
	// placeholder for future response values
}

// BackupResourceStorageConfigsNonCRRClientUpdateResponse contains the response from method BackupResourceStorageConfigsNonCRRClient.Update.
type BackupResourceStorageConfigsNonCRRClientUpdateResponse struct {
	BackupResourceConfigResource
}

// BackupResourceVaultConfigsClientGetResponse contains the response from method BackupResourceVaultConfigsClient.Get.
type BackupResourceVaultConfigsClientGetResponse struct {
	BackupResourceVaultConfigResource
}

// BackupResourceVaultConfigsClientPutResponse contains the response from method BackupResourceVaultConfigsClient.Put.
type BackupResourceVaultConfigsClientPutResponse struct {
	BackupResourceVaultConfigResource
}

// BackupResourceVaultConfigsClientUpdateResponse contains the response from method BackupResourceVaultConfigsClient.Update.
type BackupResourceVaultConfigsClientUpdateResponse struct {
	BackupResourceVaultConfigResource
}

// BackupStatusClientGetResponse contains the response from method BackupStatusClient.Get.
type BackupStatusClientGetResponse struct {
	BackupStatusResponse
}

// BackupUsageSummariesClientListResponse contains the response from method BackupUsageSummariesClient.NewListPager.
type BackupUsageSummariesClientListResponse struct {
	BackupManagementUsageList
}

// BackupWorkloadItemsClientListResponse contains the response from method BackupWorkloadItemsClient.NewListPager.
type BackupWorkloadItemsClientListResponse struct {
	WorkloadItemResourceList
}

// BackupsClientTriggerResponse contains the response from method BackupsClient.Trigger.
type BackupsClientTriggerResponse struct {
	// placeholder for future response values
}

// ClientBMSPrepareDataMoveResponse contains the response from method Client.BeginBMSPrepareDataMove.
type ClientBMSPrepareDataMoveResponse struct {
	// placeholder for future response values
}

// ClientBMSTriggerDataMoveResponse contains the response from method Client.BeginBMSTriggerDataMove.
type ClientBMSTriggerDataMoveResponse struct {
	// placeholder for future response values
}

// ClientGetOperationStatusResponse contains the response from method Client.GetOperationStatus.
type ClientGetOperationStatusResponse struct {
	OperationStatus
}

// ClientMoveRecoveryPointResponse contains the response from method Client.BeginMoveRecoveryPoint.
type ClientMoveRecoveryPointResponse struct {
	// placeholder for future response values
}

// DeletedProtectionContainersClientListResponse contains the response from method DeletedProtectionContainersClient.NewListPager.
type DeletedProtectionContainersClientListResponse struct {
	ProtectionContainerResourceList
}

// ExportJobsOperationResultsClientGetResponse contains the response from method ExportJobsOperationResultsClient.Get.
type ExportJobsOperationResultsClientGetResponse struct {
	OperationResultInfoBaseResource
}

// FeatureSupportClientValidateResponse contains the response from method FeatureSupportClient.Validate.
type FeatureSupportClientValidateResponse struct {
	AzureVMResourceFeatureSupportResponse
}

// ItemLevelRecoveryConnectionsClientProvisionResponse contains the response from method ItemLevelRecoveryConnectionsClient.Provision.
type ItemLevelRecoveryConnectionsClientProvisionResponse struct {
	// placeholder for future response values
}

// ItemLevelRecoveryConnectionsClientRevokeResponse contains the response from method ItemLevelRecoveryConnectionsClient.Revoke.
type ItemLevelRecoveryConnectionsClientRevokeResponse struct {
	// placeholder for future response values
}

// JobCancellationsClientTriggerResponse contains the response from method JobCancellationsClient.Trigger.
type JobCancellationsClientTriggerResponse struct {
	// placeholder for future response values
}

// JobDetailsClientGetResponse contains the response from method JobDetailsClient.Get.
type JobDetailsClientGetResponse struct {
	JobResource
}

// JobOperationResultsClientGetResponse contains the response from method JobOperationResultsClient.Get.
type JobOperationResultsClientGetResponse struct {
	// placeholder for future response values
}

// JobsClientExportResponse contains the response from method JobsClient.Export.
type JobsClientExportResponse struct {
	// placeholder for future response values
}

// OperationClientValidateResponse contains the response from method OperationClient.Validate.
type OperationClientValidateResponse struct {
	ValidateOperationsResponse
}

// OperationsClientListResponse contains the response from method OperationsClient.NewListPager.
type OperationsClientListResponse struct {
	ClientDiscoveryResponse
}

// PrivateEndpointClientGetOperationStatusResponse contains the response from method PrivateEndpointClient.GetOperationStatus.
type PrivateEndpointClientGetOperationStatusResponse struct {
	OperationStatus
}

// PrivateEndpointConnectionClientDeleteResponse contains the response from method PrivateEndpointConnectionClient.BeginDelete.
type PrivateEndpointConnectionClientDeleteResponse struct {
	// placeholder for future response values
}

// PrivateEndpointConnectionClientGetResponse contains the response from method PrivateEndpointConnectionClient.Get.
type PrivateEndpointConnectionClientGetResponse struct {
	PrivateEndpointConnectionResource
}

// PrivateEndpointConnectionClientPutResponse contains the response from method PrivateEndpointConnectionClient.BeginPut.
type PrivateEndpointConnectionClientPutResponse struct {
	PrivateEndpointConnectionResource
}

// ProtectableContainersClientListResponse contains the response from method ProtectableContainersClient.NewListPager.
type ProtectableContainersClientListResponse struct {
	ProtectableContainerResourceList
}

// ProtectedItemOperationResultsClientGetResponse contains the response from method ProtectedItemOperationResultsClient.Get.
type ProtectedItemOperationResultsClientGetResponse struct {
	ProtectedItemResource
}

// ProtectedItemOperationStatusesClientGetResponse contains the response from method ProtectedItemOperationStatusesClient.Get.
type ProtectedItemOperationStatusesClientGetResponse struct {
	OperationStatus
}

// ProtectedItemsClientCreateOrUpdateResponse contains the response from method ProtectedItemsClient.CreateOrUpdate.
type ProtectedItemsClientCreateOrUpdateResponse struct {
	ProtectedItemResource
}

// ProtectedItemsClientDeleteResponse contains the response from method ProtectedItemsClient.Delete.
type ProtectedItemsClientDeleteResponse struct {
	// placeholder for future response values
}

// ProtectedItemsClientGetResponse contains the response from method ProtectedItemsClient.Get.
type ProtectedItemsClientGetResponse struct {
	ProtectedItemResource
}

// ProtectionContainerOperationResultsClientGetResponse contains the response from method ProtectionContainerOperationResultsClient.Get.
type ProtectionContainerOperationResultsClientGetResponse struct {
	ProtectionContainerResource
}

// ProtectionContainerRefreshOperationResultsClientGetResponse contains the response from method ProtectionContainerRefreshOperationResultsClient.Get.
type ProtectionContainerRefreshOperationResultsClientGetResponse struct {
	// placeholder for future response values
}

// ProtectionContainersClientGetResponse contains the response from method ProtectionContainersClient.Get.
type ProtectionContainersClientGetResponse struct {
	ProtectionContainerResource
}

// ProtectionContainersClientInquireResponse contains the response from method ProtectionContainersClient.Inquire.
type ProtectionContainersClientInquireResponse struct {
	// placeholder for future response values
}

// ProtectionContainersClientRefreshResponse contains the response from method ProtectionContainersClient.Refresh.
type ProtectionContainersClientRefreshResponse struct {
	// placeholder for future response values
}

// ProtectionContainersClientRegisterResponse contains the response from method ProtectionContainersClient.Register.
type ProtectionContainersClientRegisterResponse struct {
	ProtectionContainerResource
}

// ProtectionContainersClientUnregisterResponse contains the response from method ProtectionContainersClient.Unregister.
type ProtectionContainersClientUnregisterResponse struct {
	// placeholder for future response values
}

// ProtectionIntentClientCreateOrUpdateResponse contains the response from method ProtectionIntentClient.CreateOrUpdate.
type ProtectionIntentClientCreateOrUpdateResponse struct {
	ProtectionIntentResource
}

// ProtectionIntentClientDeleteResponse contains the response from method ProtectionIntentClient.Delete.
type ProtectionIntentClientDeleteResponse struct {
	// placeholder for future response values
}

// ProtectionIntentClientGetResponse contains the response from method ProtectionIntentClient.Get.
type ProtectionIntentClientGetResponse struct {
	ProtectionIntentResource
}

// ProtectionIntentClientValidateResponse contains the response from method ProtectionIntentClient.Validate.
type ProtectionIntentClientValidateResponse struct {
	PreValidateEnableBackupResponse
}

// ProtectionPoliciesClientCreateOrUpdateResponse contains the response from method ProtectionPoliciesClient.CreateOrUpdate.
type ProtectionPoliciesClientCreateOrUpdateResponse struct {
	ProtectionPolicyResource
}

// ProtectionPoliciesClientDeleteResponse contains the response from method ProtectionPoliciesClient.BeginDelete.
type ProtectionPoliciesClientDeleteResponse struct {
	// placeholder for future response values
}

// ProtectionPoliciesClientGetResponse contains the response from method ProtectionPoliciesClient.Get.
type ProtectionPoliciesClientGetResponse struct {
	ProtectionPolicyResource
}

// ProtectionPolicyOperationResultsClientGetResponse contains the response from method ProtectionPolicyOperationResultsClient.Get.
type ProtectionPolicyOperationResultsClientGetResponse struct {
	ProtectionPolicyResource
}

// ProtectionPolicyOperationStatusesClientGetResponse contains the response from method ProtectionPolicyOperationStatusesClient.Get.
type ProtectionPolicyOperationStatusesClientGetResponse struct {
	OperationStatus
}

// RecoveryPointsClientGetResponse contains the response from method RecoveryPointsClient.Get.
type RecoveryPointsClientGetResponse struct {
	RecoveryPointResource
}

// RecoveryPointsClientListResponse contains the response from method RecoveryPointsClient.NewListPager.
type RecoveryPointsClientListResponse struct {
	RecoveryPointResourceList
}

// RecoveryPointsRecommendedForMoveClientListResponse contains the response from method RecoveryPointsRecommendedForMoveClient.NewListPager.
type RecoveryPointsRecommendedForMoveClientListResponse struct {
	RecoveryPointResourceList
}

// ResourceGuardProxiesClientGetResponse contains the response from method ResourceGuardProxiesClient.NewGetPager.
type ResourceGuardProxiesClientGetResponse struct {
	ResourceGuardProxyBaseResourceList
}

// ResourceGuardProxyClientDeleteResponse contains the response from method ResourceGuardProxyClient.Delete.
type ResourceGuardProxyClientDeleteResponse struct {
	// placeholder for future response values
}

// ResourceGuardProxyClientGetResponse contains the response from method ResourceGuardProxyClient.Get.
type ResourceGuardProxyClientGetResponse struct {
	ResourceGuardProxyBaseResource
}

// ResourceGuardProxyClientPutResponse contains the response from method ResourceGuardProxyClient.Put.
type ResourceGuardProxyClientPutResponse struct {
	ResourceGuardProxyBaseResource
}

// ResourceGuardProxyClientUnlockDeleteResponse contains the response from method ResourceGuardProxyClient.UnlockDelete.
type ResourceGuardProxyClientUnlockDeleteResponse struct {
	UnlockDeleteResponse
}

// RestoresClientTriggerResponse contains the response from method RestoresClient.BeginTrigger.
type RestoresClientTriggerResponse struct {
	// placeholder for future response values
}

// SecurityPINsClientGetResponse contains the response from method SecurityPINsClient.Get.
type SecurityPINsClientGetResponse struct {
	TokenInformation
}

// ValidateOperationClientTriggerResponse contains the response from method ValidateOperationClient.BeginTrigger.
type ValidateOperationClientTriggerResponse struct {
	// placeholder for future response values
}

// ValidateOperationResultsClientGetResponse contains the response from method ValidateOperationResultsClient.Get.
type ValidateOperationResultsClientGetResponse struct {
	ValidateOperationsResponse
}

// ValidateOperationStatusesClientGetResponse contains the response from method ValidateOperationStatusesClient.Get.
type ValidateOperationStatusesClientGetResponse struct {
	OperationStatus
}