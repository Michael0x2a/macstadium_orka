// Package orka provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.9.0 DO NOT EDIT.
package orka

// ListLogs1Params defines parameters for ListLogs1.
type ListLogs1Params struct {
	// Limit the total number of logs returned
	Limit *string `json:"limit,omitempty"`
}

// CommitChangesToTheBaseImageJSONBody defines parameters for CommitChangesToTheBaseImage.
type CommitChangesToTheBaseImageJSONBody struct {
	OrkaVmName *string `json:"orka_vm_name,omitempty"`
}

// CopyImageJSONBody defines parameters for CopyImage.
type CopyImageJSONBody struct {
	Image   *string `json:"image,omitempty"`
	NewName *string `json:"new_name,omitempty"`
}

// DeleteImageJSONBody defines parameters for DeleteImage.
type DeleteImageJSONBody struct {
	Image *string `json:"image,omitempty"`
}

// GenerateEmptyStorageJSONBody defines parameters for GenerateEmptyStorage.
type GenerateEmptyStorageJSONBody struct {
	FileName *string `json:"file_name,omitempty"`
	FileSize *string `json:"file_size,omitempty"`
}

// PullImageJSONBody defines parameters for PullImage.
type PullImageJSONBody struct {
	Image   *string `json:"image,omitempty"`
	NewName *string `json:"new_name,omitempty"`
}

// RenameImageJSONBody defines parameters for RenameImage.
type RenameImageJSONBody struct {
	Image   *string `json:"image,omitempty"`
	NewName *string `json:"new_name,omitempty"`
}

// ResizeImageJSONBody defines parameters for ResizeImage.
type ResizeImageJSONBody struct {
	NewImageName *string `json:"new_image_name,omitempty"`
	NewImageSize *string `json:"new_image_size,omitempty"`
	OrkaVmName   *string `json:"orka_vm_name,omitempty"`
	VmPassword   *string `json:"vm_password,omitempty"`
	VmUsername   *string `json:"vm_username,omitempty"`
}

// SaveNewBaseImageFromVmJSONBody defines parameters for SaveNewBaseImageFromVm.
type SaveNewBaseImageFromVmJSONBody struct {
	NewName    *string `json:"new_name,omitempty"`
	OrkaVmName *string `json:"orka_vm_name,omitempty"`
}

// CopyIsoJSONBody defines parameters for CopyIso.
type CopyIsoJSONBody struct {
	Iso     *string `json:"iso,omitempty"`
	NewName *string `json:"new_name,omitempty"`
}

// DeleteIsoJSONBody defines parameters for DeleteIso.
type DeleteIsoJSONBody struct {
	Iso *string `json:"iso,omitempty"`
}

// PullIsoJSONBody defines parameters for PullIso.
type PullIsoJSONBody struct {
	Iso     *string `json:"iso,omitempty"`
	NewName *string `json:"new_name,omitempty"`
}

// RenameIsoJSONBody defines parameters for RenameIso.
type RenameIsoJSONBody struct {
	Iso     *string `json:"iso,omitempty"`
	NewName *string `json:"new_name,omitempty"`
}

// DeleteKubeAccountJSONBody defines parameters for DeleteKubeAccount.
type DeleteKubeAccountJSONBody struct {
	Email *string `json:"email,omitempty"`
}

// ListKubeAccountsJSONBody defines parameters for ListKubeAccounts.
type ListKubeAccountsJSONBody struct {
	Email *string `json:"email,omitempty"`
}

// CreateKubeAccountJSONBody defines parameters for CreateKubeAccount.
type CreateKubeAccountJSONBody struct {
	Email *string `json:"email,omitempty"`
	Name  *string `json:"name,omitempty"`
}

// DownloadKubeconfigJSONBody defines parameters for DownloadKubeconfig.
type DownloadKubeconfigJSONBody struct {
	Email *string `json:"email,omitempty"`
	Name  *string `json:"name,omitempty"`
}

// RegenerateKubeAccountJSONBody defines parameters for RegenerateKubeAccount.
type RegenerateKubeAccountJSONBody struct {
	Email *string `json:"email,omitempty"`
	Name  *string `json:"name,omitempty"`
}

// RemoveNodeDedicationJSONBody defines parameters for RemoveNodeDedication.
type RemoveNodeDedicationJSONBody []string

// DedicateNodeJSONBody defines parameters for DedicateNode.
type DedicateNodeJSONBody []string

// DisableSandboxJSONBody defines parameters for DisableSandbox.
type DisableSandboxJSONBody struct {
	OrkaNodeName *string `json:"orka_node_name,omitempty"`
}

// EnableSandboxJSONBody defines parameters for EnableSandbox.
type EnableSandboxJSONBody struct {
	OrkaNodeName *string `json:"orka_node_name,omitempty"`
}

// UntagNodeJSONBody defines parameters for UntagNode.
type UntagNodeJSONBody struct {
	OrkaNodeName *string `json:"orka_node_name,omitempty"`
}

// TagNodeJSONBody defines parameters for TagNode.
type TagNodeJSONBody struct {
	OrkaNodeName *string `json:"orka_node_name,omitempty"`
}

// AttachDiskJSONBody defines parameters for AttachDisk.
type AttachDiskJSONBody struct {
	ImageName  *string `json:"image_name,omitempty"`
	MountPoint *string `json:"mount_point,omitempty"`
	OrkaVmName *string `json:"orka_vm_name,omitempty"`
}

// CloneVmJSONBody defines parameters for CloneVm.
type CloneVmJSONBody struct {
	CurrentNodeName *string   `json:"current_node_name,omitempty"`
	NewNodes        *[]string `json:"new_nodes,omitempty"`
	OrkaVmName      *string   `json:"orka_vm_name,omitempty"`
}

// SaveVmStateJSONBody defines parameters for SaveVmState.
type SaveVmStateJSONBody struct {
	OrkaVmName *string `json:"orka_vm_name,omitempty"`
}

// CreateVmConfigurationJSONBody defines parameters for CreateVmConfiguration.
type CreateVmConfigurationJSONBody struct {
	AttachedDisk  *string  `json:"attached_disk,omitempty"`
	IsoImage      *string  `json:"iso_image,omitempty"`
	OrkaBaseImage *string  `json:"orka_base_image,omitempty"`
	OrkaCpuCore   *float32 `json:"orka_cpu_core,omitempty"`
	OrkaImage     *string  `json:"orka_image,omitempty"`
	OrkaVmName    *string  `json:"orka_vm_name,omitempty"`
	Scheduler     *string  `json:"scheduler,omitempty"`
	SystemSerial  *string  `json:"system_serial,omitempty"`
	Tag           *string  `json:"tag,omitempty"`
	TagRequired   *bool    `json:"tag_required,omitempty"`
	VcpuCount     *float32 `json:"vcpu_count,omitempty"`
	VncConsole    *bool    `json:"vnc_console,omitempty"`
}

// DeleteUserSVmAdminJSONBody defines parameters for DeleteUserSVmAdmin.
type DeleteUserSVmAdminJSONBody struct {
	OrkaNodeName *string `json:"orka_node_name,omitempty"`
	OrkaVmName   *string `json:"orka_vm_name,omitempty"`
}

// DeployVmConfigurationJSONBody defines parameters for DeployVmConfiguration.
type DeployVmConfigurationJSONBody struct {
	AttachDisk     *bool     `json:"attach_disk,omitempty"`
	AttachedDisk   *string   `json:"attached_disk,omitempty"`
	GpuPassthrough *bool     `json:"gpu_passthrough,omitempty"`
	IsoImage       *string   `json:"iso_image,omitempty"`
	IsoInstall     *bool     `json:"iso_install,omitempty"`
	OrkaNodeName   *string   `json:"orka_node_name,omitempty"`
	OrkaVmName     *string   `json:"orka_vm_name,omitempty"`
	Replicas       *float32  `json:"replicas,omitempty"`
	ReservedPorts  *[]string `json:"reserved_ports,omitempty"`
	Scheduler      *string   `json:"scheduler,omitempty"`
	SystemSerial   *string   `json:"system_serial,omitempty"`
	Tag            *string   `json:"tag,omitempty"`
	TagRequired    *bool     `json:"tag_required,omitempty"`
	VmMetadata     *struct {
		Items *[]struct {
			Key   *string `json:"key,omitempty"`
			Value *string `json:"value,omitempty"`
		} `json:"items,omitempty"`
	} `json:"vm_metadata,omitempty"`
	VncConsole *bool `json:"vnc_console,omitempty"`
}

// ResumeVmJSONBody defines parameters for ResumeVm.
type ResumeVmJSONBody struct {
	OrkaNodeName *string `json:"orka_node_name,omitempty"`
	OrkaVmName   *string `json:"orka_vm_name,omitempty"`
}

// RevertVmJSONBody defines parameters for RevertVm.
type RevertVmJSONBody struct {
	OrkaNodeName *string `json:"orka_node_name,omitempty"`
	OrkaVmName   *string `json:"orka_vm_name,omitempty"`
}

// StartVmJSONBody defines parameters for StartVm.
type StartVmJSONBody struct {
	OrkaNodeName *string `json:"orka_node_name,omitempty"`
	OrkaVmName   *string `json:"orka_vm_name,omitempty"`
}

// StopVmJSONBody defines parameters for StopVm.
type StopVmJSONBody struct {
	OrkaNodeName *string `json:"orka_node_name,omitempty"`
	OrkaVmName   *string `json:"orka_vm_name,omitempty"`
}

// SuspendVmJSONBody defines parameters for SuspendVm.
type SuspendVmJSONBody struct {
	OrkaNodeName *string `json:"orka_node_name,omitempty"`
	OrkaVmName   *string `json:"orka_vm_name,omitempty"`
}

// ListDisksJSONBody defines parameters for ListDisks.
type ListDisksJSONBody struct {
	OrkaVmName *string `json:"orka_vm_name,omitempty"`
}

// MigrateVmJSONBody defines parameters for MigrateVm.
type MigrateVmJSONBody struct {
	CurrentNodeName *string   `json:"current_node_name,omitempty"`
	NewNodes        *[]string `json:"new_nodes,omitempty"`
	OrkaVmName      *string   `json:"orka_vm_name,omitempty"`
}

// PurgeUserSVMsAdminJSONBody defines parameters for PurgeUserSVMsAdmin.
type PurgeUserSVMsAdminJSONBody struct {
	OrkaVmName *string `json:"orka_vm_name,omitempty"`
}

// UnScaleVmJSONBody defines parameters for UnScaleVm.
type UnScaleVmJSONBody struct {
	OrkaVmName *string  `json:"orka_vm_name,omitempty"`
	Replicas   *float32 `json:"replicas,omitempty"`
}

// CreateTokenJSONBody defines parameters for CreateToken.
type CreateTokenJSONBody struct {
	Email    *string `json:"email,omitempty"`
	Password *string `json:"password,omitempty"`
}

// CreateUserJSONBody defines parameters for CreateUser.
type CreateUserJSONBody struct {
	Email    *string `json:"email,omitempty"`
	Group    *string `json:"group,omitempty"`
	Password *string `json:"password,omitempty"`
}

// UpdateUserJSONBody defines parameters for UpdateUser.
type UpdateUserJSONBody struct {
	Email    *string `json:"email,omitempty"`
	Password *string `json:"password,omitempty"`
}

// UngroupUsersJSONBody defines parameters for UngroupUsers.
type UngroupUsersJSONBody []string

// GroupUsersJSONBody defines parameters for GroupUsers.
type GroupUsersJSONBody []string

// ResetPasswordAdminJSONBody defines parameters for ResetPasswordAdmin.
type ResetPasswordAdminJSONBody struct {
	Email    *string `json:"email,omitempty"`
	Password *string `json:"password,omitempty"`
}

// ValidateLicenseKeyJSONBody defines parameters for ValidateLicenseKey.
type ValidateLicenseKeyJSONBody struct {
	LicenseKey *string `json:"licenseKey,omitempty"`
}

// CommitChangesToTheBaseImageJSONRequestBody defines body for CommitChangesToTheBaseImage for application/json ContentType.
type CommitChangesToTheBaseImageJSONRequestBody CommitChangesToTheBaseImageJSONBody

// CopyImageJSONRequestBody defines body for CopyImage for application/json ContentType.
type CopyImageJSONRequestBody CopyImageJSONBody

// DeleteImageJSONRequestBody defines body for DeleteImage for application/json ContentType.
type DeleteImageJSONRequestBody DeleteImageJSONBody

// GenerateEmptyStorageJSONRequestBody defines body for GenerateEmptyStorage for application/json ContentType.
type GenerateEmptyStorageJSONRequestBody GenerateEmptyStorageJSONBody

// PullImageJSONRequestBody defines body for PullImage for application/json ContentType.
type PullImageJSONRequestBody PullImageJSONBody

// RenameImageJSONRequestBody defines body for RenameImage for application/json ContentType.
type RenameImageJSONRequestBody RenameImageJSONBody

// ResizeImageJSONRequestBody defines body for ResizeImage for application/json ContentType.
type ResizeImageJSONRequestBody ResizeImageJSONBody

// SaveNewBaseImageFromVmJSONRequestBody defines body for SaveNewBaseImageFromVm for application/json ContentType.
type SaveNewBaseImageFromVmJSONRequestBody SaveNewBaseImageFromVmJSONBody

// CopyIsoJSONRequestBody defines body for CopyIso for application/json ContentType.
type CopyIsoJSONRequestBody CopyIsoJSONBody

// DeleteIsoJSONRequestBody defines body for DeleteIso for application/json ContentType.
type DeleteIsoJSONRequestBody DeleteIsoJSONBody

// PullIsoJSONRequestBody defines body for PullIso for application/json ContentType.
type PullIsoJSONRequestBody PullIsoJSONBody

// RenameIsoJSONRequestBody defines body for RenameIso for application/json ContentType.
type RenameIsoJSONRequestBody RenameIsoJSONBody

// DeleteKubeAccountJSONRequestBody defines body for DeleteKubeAccount for application/json ContentType.
type DeleteKubeAccountJSONRequestBody DeleteKubeAccountJSONBody

// ListKubeAccountsJSONRequestBody defines body for ListKubeAccounts for application/json ContentType.
type ListKubeAccountsJSONRequestBody ListKubeAccountsJSONBody

// CreateKubeAccountJSONRequestBody defines body for CreateKubeAccount for application/json ContentType.
type CreateKubeAccountJSONRequestBody CreateKubeAccountJSONBody

// DownloadKubeconfigJSONRequestBody defines body for DownloadKubeconfig for application/json ContentType.
type DownloadKubeconfigJSONRequestBody DownloadKubeconfigJSONBody

// RegenerateKubeAccountJSONRequestBody defines body for RegenerateKubeAccount for application/json ContentType.
type RegenerateKubeAccountJSONRequestBody RegenerateKubeAccountJSONBody

// RemoveNodeDedicationJSONRequestBody defines body for RemoveNodeDedication for application/json ContentType.
type RemoveNodeDedicationJSONRequestBody RemoveNodeDedicationJSONBody

// DedicateNodeJSONRequestBody defines body for DedicateNode for application/json ContentType.
type DedicateNodeJSONRequestBody DedicateNodeJSONBody

// DisableSandboxJSONRequestBody defines body for DisableSandbox for application/json ContentType.
type DisableSandboxJSONRequestBody DisableSandboxJSONBody

// EnableSandboxJSONRequestBody defines body for EnableSandbox for application/json ContentType.
type EnableSandboxJSONRequestBody EnableSandboxJSONBody

// UntagNodeJSONRequestBody defines body for UntagNode for application/json ContentType.
type UntagNodeJSONRequestBody UntagNodeJSONBody

// TagNodeJSONRequestBody defines body for TagNode for application/json ContentType.
type TagNodeJSONRequestBody TagNodeJSONBody

// AttachDiskJSONRequestBody defines body for AttachDisk for application/json ContentType.
type AttachDiskJSONRequestBody AttachDiskJSONBody

// CloneVmJSONRequestBody defines body for CloneVm for application/json ContentType.
type CloneVmJSONRequestBody CloneVmJSONBody

// SaveVmStateJSONRequestBody defines body for SaveVmState for application/json ContentType.
type SaveVmStateJSONRequestBody SaveVmStateJSONBody

// CreateVmConfigurationJSONRequestBody defines body for CreateVmConfiguration for application/json ContentType.
type CreateVmConfigurationJSONRequestBody CreateVmConfigurationJSONBody

// DeleteUserSVmAdminJSONRequestBody defines body for DeleteUserSVmAdmin for application/json ContentType.
type DeleteUserSVmAdminJSONRequestBody DeleteUserSVmAdminJSONBody

// DeployVmConfigurationJSONRequestBody defines body for DeployVmConfiguration for application/json ContentType.
type DeployVmConfigurationJSONRequestBody DeployVmConfigurationJSONBody

// ResumeVmJSONRequestBody defines body for ResumeVm for application/json ContentType.
type ResumeVmJSONRequestBody ResumeVmJSONBody

// RevertVmJSONRequestBody defines body for RevertVm for application/json ContentType.
type RevertVmJSONRequestBody RevertVmJSONBody

// StartVmJSONRequestBody defines body for StartVm for application/json ContentType.
type StartVmJSONRequestBody StartVmJSONBody

// StopVmJSONRequestBody defines body for StopVm for application/json ContentType.
type StopVmJSONRequestBody StopVmJSONBody

// SuspendVmJSONRequestBody defines body for SuspendVm for application/json ContentType.
type SuspendVmJSONRequestBody SuspendVmJSONBody

// ListDisksJSONRequestBody defines body for ListDisks for application/json ContentType.
type ListDisksJSONRequestBody ListDisksJSONBody

// MigrateVmJSONRequestBody defines body for MigrateVm for application/json ContentType.
type MigrateVmJSONRequestBody MigrateVmJSONBody

// PurgeUserSVMsAdminJSONRequestBody defines body for PurgeUserSVMsAdmin for application/json ContentType.
type PurgeUserSVMsAdminJSONRequestBody PurgeUserSVMsAdminJSONBody

// UnScaleVmJSONRequestBody defines body for UnScaleVm for application/json ContentType.
type UnScaleVmJSONRequestBody UnScaleVmJSONBody

// CreateTokenJSONRequestBody defines body for CreateToken for application/json ContentType.
type CreateTokenJSONRequestBody CreateTokenJSONBody

// CreateUserJSONRequestBody defines body for CreateUser for application/json ContentType.
type CreateUserJSONRequestBody CreateUserJSONBody

// UpdateUserJSONRequestBody defines body for UpdateUser for application/json ContentType.
type UpdateUserJSONRequestBody UpdateUserJSONBody

// UngroupUsersJSONRequestBody defines body for UngroupUsers for application/json ContentType.
type UngroupUsersJSONRequestBody UngroupUsersJSONBody

// GroupUsersJSONRequestBody defines body for GroupUsers for application/json ContentType.
type GroupUsersJSONRequestBody GroupUsersJSONBody

// ResetPasswordAdminJSONRequestBody defines body for ResetPasswordAdmin for application/json ContentType.
type ResetPasswordAdminJSONRequestBody ResetPasswordAdminJSONBody

// ValidateLicenseKeyJSONRequestBody defines body for ValidateLicenseKey for application/json ContentType.
type ValidateLicenseKeyJSONRequestBody ValidateLicenseKeyJSONBody
