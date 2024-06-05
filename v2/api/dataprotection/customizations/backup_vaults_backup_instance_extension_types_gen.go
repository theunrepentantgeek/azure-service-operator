// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package customizations

import (
	v20231101 "github.com/Azure/azure-service-operator/v2/api/dataprotection/v1api20231101"
	storage "github.com/Azure/azure-service-operator/v2/api/dataprotection/v1api20231101/storage"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
)

type BackupVaultsBackupInstanceExtension struct {
}

// GetExtendedResources Returns the KubernetesResource slice for Resource versions
func (extension *BackupVaultsBackupInstanceExtension) GetExtendedResources() []genruntime.KubernetesResource {
	return []genruntime.KubernetesResource{
		&v20231101.BackupVaultsBackupInstance{},
		&storage.BackupVaultsBackupInstance{}}
}
