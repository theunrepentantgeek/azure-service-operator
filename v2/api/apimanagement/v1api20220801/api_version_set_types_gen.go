// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20220801

import (
	"fmt"
	storage "github.com/Azure/azure-service-operator/v2/api/apimanagement/v1api20220801/storage"
	"github.com/Azure/azure-service-operator/v2/internal/reflecthelpers"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
	"github.com/pkg/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"sigs.k8s.io/controller-runtime/pkg/conversion"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"
)

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="Severity",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].severity"
// +kubebuilder:printcolumn:name="Reason",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].reason"
// +kubebuilder:printcolumn:name="Message",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].message"
// Generator information:
// - Generated from: /apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/apimapiversionsets.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/apiVersionSets/{versionSetId}
type ApiVersionSet struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              Service_ApiVersionSet_Spec   `json:"spec,omitempty"`
	Status            Service_ApiVersionSet_STATUS `json:"status,omitempty"`
}

var _ conditions.Conditioner = &ApiVersionSet{}

// GetConditions returns the conditions of the resource
func (versionSet *ApiVersionSet) GetConditions() conditions.Conditions {
	return versionSet.Status.Conditions
}

// SetConditions sets the conditions on the resource status
func (versionSet *ApiVersionSet) SetConditions(conditions conditions.Conditions) {
	versionSet.Status.Conditions = conditions
}

var _ conversion.Convertible = &ApiVersionSet{}

// ConvertFrom populates our ApiVersionSet from the provided hub ApiVersionSet
func (versionSet *ApiVersionSet) ConvertFrom(hub conversion.Hub) error {
	source, ok := hub.(*storage.ApiVersionSet)
	if !ok {
		return fmt.Errorf("expected apimanagement/v1api20220801/storage/ApiVersionSet but received %T instead", hub)
	}

	return versionSet.AssignProperties_From_ApiVersionSet(source)
}

// ConvertTo populates the provided hub ApiVersionSet from our ApiVersionSet
func (versionSet *ApiVersionSet) ConvertTo(hub conversion.Hub) error {
	destination, ok := hub.(*storage.ApiVersionSet)
	if !ok {
		return fmt.Errorf("expected apimanagement/v1api20220801/storage/ApiVersionSet but received %T instead", hub)
	}

	return versionSet.AssignProperties_To_ApiVersionSet(destination)
}

// +kubebuilder:webhook:path=/mutate-apimanagement-azure-com-v1api20220801-apiversionset,mutating=true,sideEffects=None,matchPolicy=Exact,failurePolicy=fail,groups=apimanagement.azure.com,resources=apiversionsets,verbs=create;update,versions=v1api20220801,name=default.v1api20220801.apiversionsets.apimanagement.azure.com,admissionReviewVersions=v1

var _ admission.Defaulter = &ApiVersionSet{}

// Default applies defaults to the ApiVersionSet resource
func (versionSet *ApiVersionSet) Default() {
	versionSet.defaultImpl()
	var temp any = versionSet
	if runtimeDefaulter, ok := temp.(genruntime.Defaulter); ok {
		runtimeDefaulter.CustomDefault()
	}
}

// defaultAzureName defaults the Azure name of the resource to the Kubernetes name
func (versionSet *ApiVersionSet) defaultAzureName() {
	if versionSet.Spec.AzureName == "" {
		versionSet.Spec.AzureName = versionSet.Name
	}
}

// defaultImpl applies the code generated defaults to the ApiVersionSet resource
func (versionSet *ApiVersionSet) defaultImpl() { versionSet.defaultAzureName() }

var _ genruntime.ImportableResource = &ApiVersionSet{}

// InitializeSpec initializes the spec for this resource from the given status
func (versionSet *ApiVersionSet) InitializeSpec(status genruntime.ConvertibleStatus) error {
	if s, ok := status.(*Service_ApiVersionSet_STATUS); ok {
		return versionSet.Spec.Initialize_From_Service_ApiVersionSet_STATUS(s)
	}

	return fmt.Errorf("expected Status of type Service_ApiVersionSet_STATUS but received %T instead", status)
}

var _ genruntime.KubernetesResource = &ApiVersionSet{}

// AzureName returns the Azure name of the resource
func (versionSet *ApiVersionSet) AzureName() string {
	return versionSet.Spec.AzureName
}

// GetAPIVersion returns the ARM API version of the resource. This is always "2022-08-01"
func (versionSet ApiVersionSet) GetAPIVersion() string {
	return string(APIVersion_Value)
}

// GetResourceScope returns the scope of the resource
func (versionSet *ApiVersionSet) GetResourceScope() genruntime.ResourceScope {
	return genruntime.ResourceScopeResourceGroup
}

// GetSpec returns the specification of this resource
func (versionSet *ApiVersionSet) GetSpec() genruntime.ConvertibleSpec {
	return &versionSet.Spec
}

// GetStatus returns the status of this resource
func (versionSet *ApiVersionSet) GetStatus() genruntime.ConvertibleStatus {
	return &versionSet.Status
}

// GetSupportedOperations returns the operations supported by the resource
func (versionSet *ApiVersionSet) GetSupportedOperations() []genruntime.ResourceOperation {
	return []genruntime.ResourceOperation{
		genruntime.ResourceOperationDelete,
		genruntime.ResourceOperationGet,
		genruntime.ResourceOperationHead,
		genruntime.ResourceOperationPut,
	}
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.ApiManagement/service/apiVersionSets"
func (versionSet *ApiVersionSet) GetType() string {
	return "Microsoft.ApiManagement/service/apiVersionSets"
}

// NewEmptyStatus returns a new empty (blank) status
func (versionSet *ApiVersionSet) NewEmptyStatus() genruntime.ConvertibleStatus {
	return &Service_ApiVersionSet_STATUS{}
}

// Owner returns the ResourceReference of the owner
func (versionSet *ApiVersionSet) Owner() *genruntime.ResourceReference {
	group, kind := genruntime.LookupOwnerGroupKind(versionSet.Spec)
	return versionSet.Spec.Owner.AsResourceReference(group, kind)
}

// SetStatus sets the status of this resource
func (versionSet *ApiVersionSet) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
	if st, ok := status.(*Service_ApiVersionSet_STATUS); ok {
		versionSet.Status = *st
		return nil
	}

	// Convert status to required version
	var st Service_ApiVersionSet_STATUS
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return errors.Wrap(err, "failed to convert status")
	}

	versionSet.Status = st
	return nil
}

// +kubebuilder:webhook:path=/validate-apimanagement-azure-com-v1api20220801-apiversionset,mutating=false,sideEffects=None,matchPolicy=Exact,failurePolicy=fail,groups=apimanagement.azure.com,resources=apiversionsets,verbs=create;update,versions=v1api20220801,name=validate.v1api20220801.apiversionsets.apimanagement.azure.com,admissionReviewVersions=v1

var _ admission.Validator = &ApiVersionSet{}

// ValidateCreate validates the creation of the resource
func (versionSet *ApiVersionSet) ValidateCreate() (admission.Warnings, error) {
	validations := versionSet.createValidations()
	var temp any = versionSet
	if runtimeValidator, ok := temp.(genruntime.Validator); ok {
		validations = append(validations, runtimeValidator.CreateValidations()...)
	}
	return genruntime.ValidateCreate(validations)
}

// ValidateDelete validates the deletion of the resource
func (versionSet *ApiVersionSet) ValidateDelete() (admission.Warnings, error) {
	validations := versionSet.deleteValidations()
	var temp any = versionSet
	if runtimeValidator, ok := temp.(genruntime.Validator); ok {
		validations = append(validations, runtimeValidator.DeleteValidations()...)
	}
	return genruntime.ValidateDelete(validations)
}

// ValidateUpdate validates an update of the resource
func (versionSet *ApiVersionSet) ValidateUpdate(old runtime.Object) (admission.Warnings, error) {
	validations := versionSet.updateValidations()
	var temp any = versionSet
	if runtimeValidator, ok := temp.(genruntime.Validator); ok {
		validations = append(validations, runtimeValidator.UpdateValidations()...)
	}
	return genruntime.ValidateUpdate(old, validations)
}

// createValidations validates the creation of the resource
func (versionSet *ApiVersionSet) createValidations() []func() (admission.Warnings, error) {
	return []func() (admission.Warnings, error){versionSet.validateResourceReferences, versionSet.validateOwnerReference}
}

// deleteValidations validates the deletion of the resource
func (versionSet *ApiVersionSet) deleteValidations() []func() (admission.Warnings, error) {
	return nil
}

// updateValidations validates the update of the resource
func (versionSet *ApiVersionSet) updateValidations() []func(old runtime.Object) (admission.Warnings, error) {
	return []func(old runtime.Object) (admission.Warnings, error){
		func(old runtime.Object) (admission.Warnings, error) {
			return versionSet.validateResourceReferences()
		},
		versionSet.validateWriteOnceProperties,
		func(old runtime.Object) (admission.Warnings, error) {
			return versionSet.validateOwnerReference()
		},
	}
}

// validateOwnerReference validates the owner field
func (versionSet *ApiVersionSet) validateOwnerReference() (admission.Warnings, error) {
	return genruntime.ValidateOwner(versionSet)
}

// validateResourceReferences validates all resource references
func (versionSet *ApiVersionSet) validateResourceReferences() (admission.Warnings, error) {
	refs, err := reflecthelpers.FindResourceReferences(&versionSet.Spec)
	if err != nil {
		return nil, err
	}
	return genruntime.ValidateResourceReferences(refs)
}

// validateWriteOnceProperties validates all WriteOnce properties
func (versionSet *ApiVersionSet) validateWriteOnceProperties(old runtime.Object) (admission.Warnings, error) {
	oldObj, ok := old.(*ApiVersionSet)
	if !ok {
		return nil, nil
	}

	return genruntime.ValidateWriteOnceProperties(oldObj, versionSet)
}

// AssignProperties_From_ApiVersionSet populates our ApiVersionSet from the provided source ApiVersionSet
func (versionSet *ApiVersionSet) AssignProperties_From_ApiVersionSet(source *storage.ApiVersionSet) error {

	// ObjectMeta
	versionSet.ObjectMeta = *source.ObjectMeta.DeepCopy()

	// Spec
	var spec Service_ApiVersionSet_Spec
	err := spec.AssignProperties_From_Service_ApiVersionSet_Spec(&source.Spec)
	if err != nil {
		return errors.Wrap(err, "calling AssignProperties_From_Service_ApiVersionSet_Spec() to populate field Spec")
	}
	versionSet.Spec = spec

	// Status
	var status Service_ApiVersionSet_STATUS
	err = status.AssignProperties_From_Service_ApiVersionSet_STATUS(&source.Status)
	if err != nil {
		return errors.Wrap(err, "calling AssignProperties_From_Service_ApiVersionSet_STATUS() to populate field Status")
	}
	versionSet.Status = status

	// No error
	return nil
}

// AssignProperties_To_ApiVersionSet populates the provided destination ApiVersionSet from our ApiVersionSet
func (versionSet *ApiVersionSet) AssignProperties_To_ApiVersionSet(destination *storage.ApiVersionSet) error {

	// ObjectMeta
	destination.ObjectMeta = *versionSet.ObjectMeta.DeepCopy()

	// Spec
	var spec storage.Service_ApiVersionSet_Spec
	err := versionSet.Spec.AssignProperties_To_Service_ApiVersionSet_Spec(&spec)
	if err != nil {
		return errors.Wrap(err, "calling AssignProperties_To_Service_ApiVersionSet_Spec() to populate field Spec")
	}
	destination.Spec = spec

	// Status
	var status storage.Service_ApiVersionSet_STATUS
	err = versionSet.Status.AssignProperties_To_Service_ApiVersionSet_STATUS(&status)
	if err != nil {
		return errors.Wrap(err, "calling AssignProperties_To_Service_ApiVersionSet_STATUS() to populate field Status")
	}
	destination.Status = status

	// No error
	return nil
}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (versionSet *ApiVersionSet) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: versionSet.Spec.OriginalVersion(),
		Kind:    "ApiVersionSet",
	}
}

// +kubebuilder:object:root=true
// Generator information:
// - Generated from: /apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/apimapiversionsets.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/apiVersionSets/{versionSetId}
type ApiVersionSetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ApiVersionSet `json:"items"`
}

type Service_ApiVersionSet_Spec struct {
	// +kubebuilder:validation:MaxLength=80
	// +kubebuilder:validation:MinLength=1
	// +kubebuilder:validation:Pattern="^[^*#&+:<>?]+$"
	// AzureName: The name of the resource in Azure. This is often the same as the name of the resource in Kubernetes but it
	// doesn't have to be.
	AzureName string `json:"azureName,omitempty"`

	// Description: Description of API Version Set.
	Description *string `json:"description,omitempty"`

	// +kubebuilder:validation:Required
	// +kubebuilder:validation:MaxLength=100
	// +kubebuilder:validation:MinLength=1
	// DisplayName: Name of API Version Set
	DisplayName *string `json:"displayName,omitempty"`

	// +kubebuilder:validation:Required
	// Owner: The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also
	// controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a
	// reference to a apimanagement.azure.com/Service resource
	Owner *genruntime.KnownResourceReference `group:"apimanagement.azure.com" json:"owner,omitempty" kind:"Service"`

	// +kubebuilder:validation:MaxLength=100
	// +kubebuilder:validation:MinLength=1
	// VersionHeaderName: Name of HTTP header parameter that indicates the API Version if versioningScheme is set to `header`.
	VersionHeaderName *string `json:"versionHeaderName,omitempty"`

	// +kubebuilder:validation:MaxLength=100
	// +kubebuilder:validation:MinLength=1
	// VersionQueryName: Name of query parameter that indicates the API Version if versioningScheme is set to `query`.
	VersionQueryName *string `json:"versionQueryName,omitempty"`

	// +kubebuilder:validation:Required
	// VersioningScheme: An value that determines where the API Version identifier will be located in a HTTP request.
	VersioningScheme *ApiVersionSetContractProperties_VersioningScheme `json:"versioningScheme,omitempty"`
}

var _ genruntime.ARMTransformer = &Service_ApiVersionSet_Spec{}

// ConvertToARM converts from a Kubernetes CRD object to an ARM object
func (versionSet *Service_ApiVersionSet_Spec) ConvertToARM(resolved genruntime.ConvertToARMResolvedDetails) (interface{}, error) {
	if versionSet == nil {
		return nil, nil
	}
	result := &Service_ApiVersionSet_Spec_ARM{}

	// Set property "Name":
	result.Name = resolved.Name

	// Set property "Properties":
	if versionSet.Description != nil ||
		versionSet.DisplayName != nil ||
		versionSet.VersionHeaderName != nil ||
		versionSet.VersionQueryName != nil ||
		versionSet.VersioningScheme != nil {
		result.Properties = &ApiVersionSetContractProperties_ARM{}
	}
	if versionSet.Description != nil {
		description := *versionSet.Description
		result.Properties.Description = &description
	}
	if versionSet.DisplayName != nil {
		displayName := *versionSet.DisplayName
		result.Properties.DisplayName = &displayName
	}
	if versionSet.VersionHeaderName != nil {
		versionHeaderName := *versionSet.VersionHeaderName
		result.Properties.VersionHeaderName = &versionHeaderName
	}
	if versionSet.VersionQueryName != nil {
		versionQueryName := *versionSet.VersionQueryName
		result.Properties.VersionQueryName = &versionQueryName
	}
	if versionSet.VersioningScheme != nil {
		versioningScheme := *versionSet.VersioningScheme
		result.Properties.VersioningScheme = &versioningScheme
	}
	return result, nil
}

// NewEmptyARMValue returns an empty ARM value suitable for deserializing into
func (versionSet *Service_ApiVersionSet_Spec) NewEmptyARMValue() genruntime.ARMResourceStatus {
	return &Service_ApiVersionSet_Spec_ARM{}
}

// PopulateFromARM populates a Kubernetes CRD object from an Azure ARM object
func (versionSet *Service_ApiVersionSet_Spec) PopulateFromARM(owner genruntime.ArbitraryOwnerReference, armInput interface{}) error {
	typedInput, ok := armInput.(Service_ApiVersionSet_Spec_ARM)
	if !ok {
		return fmt.Errorf("unexpected type supplied for PopulateFromARM() function. Expected Service_ApiVersionSet_Spec_ARM, got %T", armInput)
	}

	// Set property "AzureName":
	versionSet.SetAzureName(genruntime.ExtractKubernetesResourceNameFromARMName(typedInput.Name))

	// Set property "Description":
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.Description != nil {
			description := *typedInput.Properties.Description
			versionSet.Description = &description
		}
	}

	// Set property "DisplayName":
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.DisplayName != nil {
			displayName := *typedInput.Properties.DisplayName
			versionSet.DisplayName = &displayName
		}
	}

	// Set property "Owner":
	versionSet.Owner = &genruntime.KnownResourceReference{
		Name:  owner.Name,
		ARMID: owner.ARMID,
	}

	// Set property "VersionHeaderName":
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.VersionHeaderName != nil {
			versionHeaderName := *typedInput.Properties.VersionHeaderName
			versionSet.VersionHeaderName = &versionHeaderName
		}
	}

	// Set property "VersionQueryName":
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.VersionQueryName != nil {
			versionQueryName := *typedInput.Properties.VersionQueryName
			versionSet.VersionQueryName = &versionQueryName
		}
	}

	// Set property "VersioningScheme":
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.VersioningScheme != nil {
			versioningScheme := *typedInput.Properties.VersioningScheme
			versionSet.VersioningScheme = &versioningScheme
		}
	}

	// No error
	return nil
}

var _ genruntime.ConvertibleSpec = &Service_ApiVersionSet_Spec{}

// ConvertSpecFrom populates our Service_ApiVersionSet_Spec from the provided source
func (versionSet *Service_ApiVersionSet_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	src, ok := source.(*storage.Service_ApiVersionSet_Spec)
	if ok {
		// Populate our instance from source
		return versionSet.AssignProperties_From_Service_ApiVersionSet_Spec(src)
	}

	// Convert to an intermediate form
	src = &storage.Service_ApiVersionSet_Spec{}
	err := src.ConvertSpecFrom(source)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertSpecFrom()")
	}

	// Update our instance from src
	err = versionSet.AssignProperties_From_Service_ApiVersionSet_Spec(src)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertSpecFrom()")
	}

	return nil
}

// ConvertSpecTo populates the provided destination from our Service_ApiVersionSet_Spec
func (versionSet *Service_ApiVersionSet_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	dst, ok := destination.(*storage.Service_ApiVersionSet_Spec)
	if ok {
		// Populate destination from our instance
		return versionSet.AssignProperties_To_Service_ApiVersionSet_Spec(dst)
	}

	// Convert to an intermediate form
	dst = &storage.Service_ApiVersionSet_Spec{}
	err := versionSet.AssignProperties_To_Service_ApiVersionSet_Spec(dst)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertSpecTo()")
	}

	// Update dst from our instance
	err = dst.ConvertSpecTo(destination)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertSpecTo()")
	}

	return nil
}

// AssignProperties_From_Service_ApiVersionSet_Spec populates our Service_ApiVersionSet_Spec from the provided source Service_ApiVersionSet_Spec
func (versionSet *Service_ApiVersionSet_Spec) AssignProperties_From_Service_ApiVersionSet_Spec(source *storage.Service_ApiVersionSet_Spec) error {

	// AzureName
	versionSet.AzureName = source.AzureName

	// Description
	versionSet.Description = genruntime.ClonePointerToString(source.Description)

	// DisplayName
	if source.DisplayName != nil {
		displayName := *source.DisplayName
		versionSet.DisplayName = &displayName
	} else {
		versionSet.DisplayName = nil
	}

	// Owner
	if source.Owner != nil {
		owner := source.Owner.Copy()
		versionSet.Owner = &owner
	} else {
		versionSet.Owner = nil
	}

	// VersionHeaderName
	if source.VersionHeaderName != nil {
		versionHeaderName := *source.VersionHeaderName
		versionSet.VersionHeaderName = &versionHeaderName
	} else {
		versionSet.VersionHeaderName = nil
	}

	// VersionQueryName
	if source.VersionQueryName != nil {
		versionQueryName := *source.VersionQueryName
		versionSet.VersionQueryName = &versionQueryName
	} else {
		versionSet.VersionQueryName = nil
	}

	// VersioningScheme
	if source.VersioningScheme != nil {
		versioningScheme := *source.VersioningScheme
		versioningSchemeTemp := genruntime.ToEnum(versioningScheme, apiVersionSetContractProperties_VersioningScheme_Values)
		versionSet.VersioningScheme = &versioningSchemeTemp
	} else {
		versionSet.VersioningScheme = nil
	}

	// No error
	return nil
}

// AssignProperties_To_Service_ApiVersionSet_Spec populates the provided destination Service_ApiVersionSet_Spec from our Service_ApiVersionSet_Spec
func (versionSet *Service_ApiVersionSet_Spec) AssignProperties_To_Service_ApiVersionSet_Spec(destination *storage.Service_ApiVersionSet_Spec) error {
	// Create a new property bag
	propertyBag := genruntime.NewPropertyBag()

	// AzureName
	destination.AzureName = versionSet.AzureName

	// Description
	destination.Description = genruntime.ClonePointerToString(versionSet.Description)

	// DisplayName
	if versionSet.DisplayName != nil {
		displayName := *versionSet.DisplayName
		destination.DisplayName = &displayName
	} else {
		destination.DisplayName = nil
	}

	// OriginalVersion
	destination.OriginalVersion = versionSet.OriginalVersion()

	// Owner
	if versionSet.Owner != nil {
		owner := versionSet.Owner.Copy()
		destination.Owner = &owner
	} else {
		destination.Owner = nil
	}

	// VersionHeaderName
	if versionSet.VersionHeaderName != nil {
		versionHeaderName := *versionSet.VersionHeaderName
		destination.VersionHeaderName = &versionHeaderName
	} else {
		destination.VersionHeaderName = nil
	}

	// VersionQueryName
	if versionSet.VersionQueryName != nil {
		versionQueryName := *versionSet.VersionQueryName
		destination.VersionQueryName = &versionQueryName
	} else {
		destination.VersionQueryName = nil
	}

	// VersioningScheme
	if versionSet.VersioningScheme != nil {
		versioningScheme := string(*versionSet.VersioningScheme)
		destination.VersioningScheme = &versioningScheme
	} else {
		destination.VersioningScheme = nil
	}

	// Update the property bag
	if len(propertyBag) > 0 {
		destination.PropertyBag = propertyBag
	} else {
		destination.PropertyBag = nil
	}

	// No error
	return nil
}

// Initialize_From_Service_ApiVersionSet_STATUS populates our Service_ApiVersionSet_Spec from the provided source Service_ApiVersionSet_STATUS
func (versionSet *Service_ApiVersionSet_Spec) Initialize_From_Service_ApiVersionSet_STATUS(source *Service_ApiVersionSet_STATUS) error {

	// Description
	versionSet.Description = genruntime.ClonePointerToString(source.Description)

	// DisplayName
	if source.DisplayName != nil {
		displayName := *source.DisplayName
		versionSet.DisplayName = &displayName
	} else {
		versionSet.DisplayName = nil
	}

	// VersionHeaderName
	if source.VersionHeaderName != nil {
		versionHeaderName := *source.VersionHeaderName
		versionSet.VersionHeaderName = &versionHeaderName
	} else {
		versionSet.VersionHeaderName = nil
	}

	// VersionQueryName
	if source.VersionQueryName != nil {
		versionQueryName := *source.VersionQueryName
		versionSet.VersionQueryName = &versionQueryName
	} else {
		versionSet.VersionQueryName = nil
	}

	// VersioningScheme
	if source.VersioningScheme != nil {
		versioningScheme := ApiVersionSetContractProperties_VersioningScheme(*source.VersioningScheme)
		versionSet.VersioningScheme = &versioningScheme
	} else {
		versionSet.VersioningScheme = nil
	}

	// No error
	return nil
}

// OriginalVersion returns the original API version used to create the resource.
func (versionSet *Service_ApiVersionSet_Spec) OriginalVersion() string {
	return GroupVersion.Version
}

// SetAzureName sets the Azure name of the resource
func (versionSet *Service_ApiVersionSet_Spec) SetAzureName(azureName string) {
	versionSet.AzureName = azureName
}

type Service_ApiVersionSet_STATUS struct {
	// Conditions: The observed state of the resource
	Conditions []conditions.Condition `json:"conditions,omitempty"`

	// Description: Description of API Version Set.
	Description *string `json:"description,omitempty"`

	// DisplayName: Name of API Version Set
	DisplayName *string `json:"displayName,omitempty"`

	// Id: Fully qualified resource ID for the resource. Ex -
	// /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id *string `json:"id,omitempty"`

	// Name: The name of the resource
	Name *string `json:"name,omitempty"`

	// Type: The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string `json:"type,omitempty"`

	// VersionHeaderName: Name of HTTP header parameter that indicates the API Version if versioningScheme is set to `header`.
	VersionHeaderName *string `json:"versionHeaderName,omitempty"`

	// VersionQueryName: Name of query parameter that indicates the API Version if versioningScheme is set to `query`.
	VersionQueryName *string `json:"versionQueryName,omitempty"`

	// VersioningScheme: An value that determines where the API Version identifier will be located in a HTTP request.
	VersioningScheme *ApiVersionSetContractProperties_VersioningScheme_STATUS `json:"versioningScheme,omitempty"`
}

var _ genruntime.ConvertibleStatus = &Service_ApiVersionSet_STATUS{}

// ConvertStatusFrom populates our Service_ApiVersionSet_STATUS from the provided source
func (versionSet *Service_ApiVersionSet_STATUS) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	src, ok := source.(*storage.Service_ApiVersionSet_STATUS)
	if ok {
		// Populate our instance from source
		return versionSet.AssignProperties_From_Service_ApiVersionSet_STATUS(src)
	}

	// Convert to an intermediate form
	src = &storage.Service_ApiVersionSet_STATUS{}
	err := src.ConvertStatusFrom(source)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertStatusFrom()")
	}

	// Update our instance from src
	err = versionSet.AssignProperties_From_Service_ApiVersionSet_STATUS(src)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertStatusFrom()")
	}

	return nil
}

// ConvertStatusTo populates the provided destination from our Service_ApiVersionSet_STATUS
func (versionSet *Service_ApiVersionSet_STATUS) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	dst, ok := destination.(*storage.Service_ApiVersionSet_STATUS)
	if ok {
		// Populate destination from our instance
		return versionSet.AssignProperties_To_Service_ApiVersionSet_STATUS(dst)
	}

	// Convert to an intermediate form
	dst = &storage.Service_ApiVersionSet_STATUS{}
	err := versionSet.AssignProperties_To_Service_ApiVersionSet_STATUS(dst)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertStatusTo()")
	}

	// Update dst from our instance
	err = dst.ConvertStatusTo(destination)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertStatusTo()")
	}

	return nil
}

var _ genruntime.FromARMConverter = &Service_ApiVersionSet_STATUS{}

// NewEmptyARMValue returns an empty ARM value suitable for deserializing into
func (versionSet *Service_ApiVersionSet_STATUS) NewEmptyARMValue() genruntime.ARMResourceStatus {
	return &Service_ApiVersionSet_STATUS_ARM{}
}

// PopulateFromARM populates a Kubernetes CRD object from an Azure ARM object
func (versionSet *Service_ApiVersionSet_STATUS) PopulateFromARM(owner genruntime.ArbitraryOwnerReference, armInput interface{}) error {
	typedInput, ok := armInput.(Service_ApiVersionSet_STATUS_ARM)
	if !ok {
		return fmt.Errorf("unexpected type supplied for PopulateFromARM() function. Expected Service_ApiVersionSet_STATUS_ARM, got %T", armInput)
	}

	// no assignment for property "Conditions"

	// Set property "Description":
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.Description != nil {
			description := *typedInput.Properties.Description
			versionSet.Description = &description
		}
	}

	// Set property "DisplayName":
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.DisplayName != nil {
			displayName := *typedInput.Properties.DisplayName
			versionSet.DisplayName = &displayName
		}
	}

	// Set property "Id":
	if typedInput.Id != nil {
		id := *typedInput.Id
		versionSet.Id = &id
	}

	// Set property "Name":
	if typedInput.Name != nil {
		name := *typedInput.Name
		versionSet.Name = &name
	}

	// Set property "Type":
	if typedInput.Type != nil {
		typeVar := *typedInput.Type
		versionSet.Type = &typeVar
	}

	// Set property "VersionHeaderName":
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.VersionHeaderName != nil {
			versionHeaderName := *typedInput.Properties.VersionHeaderName
			versionSet.VersionHeaderName = &versionHeaderName
		}
	}

	// Set property "VersionQueryName":
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.VersionQueryName != nil {
			versionQueryName := *typedInput.Properties.VersionQueryName
			versionSet.VersionQueryName = &versionQueryName
		}
	}

	// Set property "VersioningScheme":
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.VersioningScheme != nil {
			versioningScheme := *typedInput.Properties.VersioningScheme
			versionSet.VersioningScheme = &versioningScheme
		}
	}

	// No error
	return nil
}

// AssignProperties_From_Service_ApiVersionSet_STATUS populates our Service_ApiVersionSet_STATUS from the provided source Service_ApiVersionSet_STATUS
func (versionSet *Service_ApiVersionSet_STATUS) AssignProperties_From_Service_ApiVersionSet_STATUS(source *storage.Service_ApiVersionSet_STATUS) error {

	// Conditions
	versionSet.Conditions = genruntime.CloneSliceOfCondition(source.Conditions)

	// Description
	versionSet.Description = genruntime.ClonePointerToString(source.Description)

	// DisplayName
	versionSet.DisplayName = genruntime.ClonePointerToString(source.DisplayName)

	// Id
	versionSet.Id = genruntime.ClonePointerToString(source.Id)

	// Name
	versionSet.Name = genruntime.ClonePointerToString(source.Name)

	// Type
	versionSet.Type = genruntime.ClonePointerToString(source.Type)

	// VersionHeaderName
	versionSet.VersionHeaderName = genruntime.ClonePointerToString(source.VersionHeaderName)

	// VersionQueryName
	versionSet.VersionQueryName = genruntime.ClonePointerToString(source.VersionQueryName)

	// VersioningScheme
	if source.VersioningScheme != nil {
		versioningScheme := *source.VersioningScheme
		versioningSchemeTemp := genruntime.ToEnum(versioningScheme, apiVersionSetContractProperties_VersioningScheme_STATUS_Values)
		versionSet.VersioningScheme = &versioningSchemeTemp
	} else {
		versionSet.VersioningScheme = nil
	}

	// No error
	return nil
}

// AssignProperties_To_Service_ApiVersionSet_STATUS populates the provided destination Service_ApiVersionSet_STATUS from our Service_ApiVersionSet_STATUS
func (versionSet *Service_ApiVersionSet_STATUS) AssignProperties_To_Service_ApiVersionSet_STATUS(destination *storage.Service_ApiVersionSet_STATUS) error {
	// Create a new property bag
	propertyBag := genruntime.NewPropertyBag()

	// Conditions
	destination.Conditions = genruntime.CloneSliceOfCondition(versionSet.Conditions)

	// Description
	destination.Description = genruntime.ClonePointerToString(versionSet.Description)

	// DisplayName
	destination.DisplayName = genruntime.ClonePointerToString(versionSet.DisplayName)

	// Id
	destination.Id = genruntime.ClonePointerToString(versionSet.Id)

	// Name
	destination.Name = genruntime.ClonePointerToString(versionSet.Name)

	// Type
	destination.Type = genruntime.ClonePointerToString(versionSet.Type)

	// VersionHeaderName
	destination.VersionHeaderName = genruntime.ClonePointerToString(versionSet.VersionHeaderName)

	// VersionQueryName
	destination.VersionQueryName = genruntime.ClonePointerToString(versionSet.VersionQueryName)

	// VersioningScheme
	if versionSet.VersioningScheme != nil {
		versioningScheme := string(*versionSet.VersioningScheme)
		destination.VersioningScheme = &versioningScheme
	} else {
		destination.VersioningScheme = nil
	}

	// Update the property bag
	if len(propertyBag) > 0 {
		destination.PropertyBag = propertyBag
	} else {
		destination.PropertyBag = nil
	}

	// No error
	return nil
}

// +kubebuilder:validation:Enum={"Header","Query","Segment"}
type ApiVersionSetContractProperties_VersioningScheme string

const (
	ApiVersionSetContractProperties_VersioningScheme_Header  = ApiVersionSetContractProperties_VersioningScheme("Header")
	ApiVersionSetContractProperties_VersioningScheme_Query   = ApiVersionSetContractProperties_VersioningScheme("Query")
	ApiVersionSetContractProperties_VersioningScheme_Segment = ApiVersionSetContractProperties_VersioningScheme("Segment")
)

// Mapping from string to ApiVersionSetContractProperties_VersioningScheme
var apiVersionSetContractProperties_VersioningScheme_Values = map[string]ApiVersionSetContractProperties_VersioningScheme{
	"header":  ApiVersionSetContractProperties_VersioningScheme_Header,
	"query":   ApiVersionSetContractProperties_VersioningScheme_Query,
	"segment": ApiVersionSetContractProperties_VersioningScheme_Segment,
}

type ApiVersionSetContractProperties_VersioningScheme_STATUS string

const (
	ApiVersionSetContractProperties_VersioningScheme_STATUS_Header  = ApiVersionSetContractProperties_VersioningScheme_STATUS("Header")
	ApiVersionSetContractProperties_VersioningScheme_STATUS_Query   = ApiVersionSetContractProperties_VersioningScheme_STATUS("Query")
	ApiVersionSetContractProperties_VersioningScheme_STATUS_Segment = ApiVersionSetContractProperties_VersioningScheme_STATUS("Segment")
)

// Mapping from string to ApiVersionSetContractProperties_VersioningScheme_STATUS
var apiVersionSetContractProperties_VersioningScheme_STATUS_Values = map[string]ApiVersionSetContractProperties_VersioningScheme_STATUS{
	"header":  ApiVersionSetContractProperties_VersioningScheme_STATUS_Header,
	"query":   ApiVersionSetContractProperties_VersioningScheme_STATUS_Query,
	"segment": ApiVersionSetContractProperties_VersioningScheme_STATUS_Segment,
}

func init() {
	SchemeBuilder.Register(&ApiVersionSet{}, &ApiVersionSetList{})
}
