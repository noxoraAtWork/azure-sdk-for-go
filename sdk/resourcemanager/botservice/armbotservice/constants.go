//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armbotservice

const (
	moduleName    = "armbotservice"
	moduleVersion = "v1.0.0"
)

type ChannelName string

const (
	ChannelNameAlexaChannel            ChannelName = "AlexaChannel"
	ChannelNameFacebookChannel         ChannelName = "FacebookChannel"
	ChannelNameEmailChannel            ChannelName = "EmailChannel"
	ChannelNameKikChannel              ChannelName = "KikChannel"
	ChannelNameTelegramChannel         ChannelName = "TelegramChannel"
	ChannelNameSlackChannel            ChannelName = "SlackChannel"
	ChannelNameMsTeamsChannel          ChannelName = "MsTeamsChannel"
	ChannelNameSkypeChannel            ChannelName = "SkypeChannel"
	ChannelNameWebChatChannel          ChannelName = "WebChatChannel"
	ChannelNameDirectLineChannel       ChannelName = "DirectLineChannel"
	ChannelNameSmsChannel              ChannelName = "SmsChannel"
	ChannelNameLineChannel             ChannelName = "LineChannel"
	ChannelNameDirectLineSpeechChannel ChannelName = "DirectLineSpeechChannel"
	ChannelNameOutlookChannel          ChannelName = "OutlookChannel"
	ChannelNameOmnichannel             ChannelName = "Omnichannel"
	ChannelNameTelephonyChannel        ChannelName = "TelephonyChannel"
	ChannelNameAcsChatChannel          ChannelName = "AcsChatChannel"
	ChannelNameSearchAssistant         ChannelName = "SearchAssistant"
	ChannelNameM365Extensions          ChannelName = "M365Extensions"
)

// PossibleChannelNameValues returns the possible values for the ChannelName const type.
func PossibleChannelNameValues() []ChannelName {
	return []ChannelName{
		ChannelNameAlexaChannel,
		ChannelNameFacebookChannel,
		ChannelNameEmailChannel,
		ChannelNameKikChannel,
		ChannelNameTelegramChannel,
		ChannelNameSlackChannel,
		ChannelNameMsTeamsChannel,
		ChannelNameSkypeChannel,
		ChannelNameWebChatChannel,
		ChannelNameDirectLineChannel,
		ChannelNameSmsChannel,
		ChannelNameLineChannel,
		ChannelNameDirectLineSpeechChannel,
		ChannelNameOutlookChannel,
		ChannelNameOmnichannel,
		ChannelNameTelephonyChannel,
		ChannelNameAcsChatChannel,
		ChannelNameSearchAssistant,
		ChannelNameM365Extensions,
	}
}

// EmailChannelAuthMethod - Email channel auth method. 0 Password (Default); 1 Graph.
type EmailChannelAuthMethod float32

const (
	// EmailChannelAuthMethodPassword - Basic authentication.
	EmailChannelAuthMethodPassword EmailChannelAuthMethod = 0
	// EmailChannelAuthMethodGraph - Modern authentication.
	EmailChannelAuthMethodGraph EmailChannelAuthMethod = 1
)

// PossibleEmailChannelAuthMethodValues returns the possible values for the EmailChannelAuthMethod const type.
func PossibleEmailChannelAuthMethodValues() []EmailChannelAuthMethod {
	return []EmailChannelAuthMethod{
		EmailChannelAuthMethodPassword,
		EmailChannelAuthMethodGraph,
	}
}

// Key - Determines which key is to be regenerated
type Key string

const (
	KeyKey1 Key = "key1"
	KeyKey2 Key = "key2"
)

// PossibleKeyValues returns the possible values for the Key const type.
func PossibleKeyValues() []Key {
	return []Key{
		KeyKey1,
		KeyKey2,
	}
}

// Kind - Indicates the type of bot service
type Kind string

const (
	KindAzurebot Kind = "azurebot"
	KindBot      Kind = "bot"
	KindDesigner Kind = "designer"
	KindFunction Kind = "function"
	KindSdk      Kind = "sdk"
)

// PossibleKindValues returns the possible values for the Kind const type.
func PossibleKindValues() []Kind {
	return []Kind{
		KindAzurebot,
		KindBot,
		KindDesigner,
		KindFunction,
		KindSdk,
	}
}

// MsaAppType - Microsoft App Type for the bot
type MsaAppType string

const (
	MsaAppTypeMultiTenant     MsaAppType = "MultiTenant"
	MsaAppTypeSingleTenant    MsaAppType = "SingleTenant"
	MsaAppTypeUserAssignedMSI MsaAppType = "UserAssignedMSI"
)

// PossibleMsaAppTypeValues returns the possible values for the MsaAppType const type.
func PossibleMsaAppTypeValues() []MsaAppType {
	return []MsaAppType{
		MsaAppTypeMultiTenant,
		MsaAppTypeSingleTenant,
		MsaAppTypeUserAssignedMSI,
	}
}

// OperationResultStatus - The status of the operation being performed.
type OperationResultStatus string

const (
	OperationResultStatusCanceled  OperationResultStatus = "Canceled"
	OperationResultStatusFailed    OperationResultStatus = "Failed"
	OperationResultStatusRequested OperationResultStatus = "Requested"
	OperationResultStatusRunning   OperationResultStatus = "Running"
	OperationResultStatusSucceeded OperationResultStatus = "Succeeded"
)

// PossibleOperationResultStatusValues returns the possible values for the OperationResultStatus const type.
func PossibleOperationResultStatusValues() []OperationResultStatus {
	return []OperationResultStatus{
		OperationResultStatusCanceled,
		OperationResultStatusFailed,
		OperationResultStatusRequested,
		OperationResultStatusRunning,
		OperationResultStatusSucceeded,
	}
}

// PrivateEndpointConnectionProvisioningState - The current provisioning state.
type PrivateEndpointConnectionProvisioningState string

const (
	PrivateEndpointConnectionProvisioningStateCreating  PrivateEndpointConnectionProvisioningState = "Creating"
	PrivateEndpointConnectionProvisioningStateDeleting  PrivateEndpointConnectionProvisioningState = "Deleting"
	PrivateEndpointConnectionProvisioningStateFailed    PrivateEndpointConnectionProvisioningState = "Failed"
	PrivateEndpointConnectionProvisioningStateSucceeded PrivateEndpointConnectionProvisioningState = "Succeeded"
)

// PossiblePrivateEndpointConnectionProvisioningStateValues returns the possible values for the PrivateEndpointConnectionProvisioningState const type.
func PossiblePrivateEndpointConnectionProvisioningStateValues() []PrivateEndpointConnectionProvisioningState {
	return []PrivateEndpointConnectionProvisioningState{
		PrivateEndpointConnectionProvisioningStateCreating,
		PrivateEndpointConnectionProvisioningStateDeleting,
		PrivateEndpointConnectionProvisioningStateFailed,
		PrivateEndpointConnectionProvisioningStateSucceeded,
	}
}

// PrivateEndpointServiceConnectionStatus - The private endpoint connection status.
type PrivateEndpointServiceConnectionStatus string

const (
	PrivateEndpointServiceConnectionStatusApproved PrivateEndpointServiceConnectionStatus = "Approved"
	PrivateEndpointServiceConnectionStatusPending  PrivateEndpointServiceConnectionStatus = "Pending"
	PrivateEndpointServiceConnectionStatusRejected PrivateEndpointServiceConnectionStatus = "Rejected"
)

// PossiblePrivateEndpointServiceConnectionStatusValues returns the possible values for the PrivateEndpointServiceConnectionStatus const type.
func PossiblePrivateEndpointServiceConnectionStatusValues() []PrivateEndpointServiceConnectionStatus {
	return []PrivateEndpointServiceConnectionStatus{
		PrivateEndpointServiceConnectionStatusApproved,
		PrivateEndpointServiceConnectionStatusPending,
		PrivateEndpointServiceConnectionStatusRejected,
	}
}

// PublicNetworkAccess - Whether the bot is in an isolated network
type PublicNetworkAccess string

const (
	PublicNetworkAccessDisabled PublicNetworkAccess = "Disabled"
	PublicNetworkAccessEnabled  PublicNetworkAccess = "Enabled"
)

// PossiblePublicNetworkAccessValues returns the possible values for the PublicNetworkAccess const type.
func PossiblePublicNetworkAccessValues() []PublicNetworkAccess {
	return []PublicNetworkAccess{
		PublicNetworkAccessDisabled,
		PublicNetworkAccessEnabled,
	}
}

type RegenerateKeysChannelName string

const (
	RegenerateKeysChannelNameWebChatChannel    RegenerateKeysChannelName = "WebChatChannel"
	RegenerateKeysChannelNameDirectLineChannel RegenerateKeysChannelName = "DirectLineChannel"
)

// PossibleRegenerateKeysChannelNameValues returns the possible values for the RegenerateKeysChannelName const type.
func PossibleRegenerateKeysChannelNameValues() []RegenerateKeysChannelName {
	return []RegenerateKeysChannelName{
		RegenerateKeysChannelNameWebChatChannel,
		RegenerateKeysChannelNameDirectLineChannel,
	}
}

// SKUName - The name of SKU.
type SKUName string

const (
	SKUNameF0 SKUName = "F0"
	SKUNameS1 SKUName = "S1"
)

// PossibleSKUNameValues returns the possible values for the SKUName const type.
func PossibleSKUNameValues() []SKUName {
	return []SKUName{
		SKUNameF0,
		SKUNameS1,
	}
}

// SKUTier - Gets the sku tier. This is based on the SKU name.
type SKUTier string

const (
	SKUTierFree     SKUTier = "Free"
	SKUTierStandard SKUTier = "Standard"
)

// PossibleSKUTierValues returns the possible values for the SKUTier const type.
func PossibleSKUTierValues() []SKUTier {
	return []SKUTier{
		SKUTierFree,
		SKUTierStandard,
	}
}