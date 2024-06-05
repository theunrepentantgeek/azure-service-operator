// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20211101

import (
	"fmt"
	storage "github.com/Azure/azure-service-operator/v2/api/sql/v1api20211101/storage"
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
// - Generated from: /sql/resource-manager/Microsoft.Sql/stable/2021-11-01/IPv6FirewallRules.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/ipv6FirewallRules/{firewallRuleName}
type ServersIPV6FirewallRule struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              Servers_Ipv6FirewallRule_Spec   `json:"spec,omitempty"`
	Status            Servers_Ipv6FirewallRule_STATUS `json:"status,omitempty"`
}

var _ conditions.Conditioner = &ServersIPV6FirewallRule{}

// GetConditions returns the conditions of the resource
func (rule *ServersIPV6FirewallRule) GetConditions() conditions.Conditions {
	return rule.Status.Conditions
}

// SetConditions sets the conditions on the resource status
func (rule *ServersIPV6FirewallRule) SetConditions(conditions conditions.Conditions) {
	rule.Status.Conditions = conditions
}

var _ conversion.Convertible = &ServersIPV6FirewallRule{}

// ConvertFrom populates our ServersIPV6FirewallRule from the provided hub ServersIPV6FirewallRule
func (rule *ServersIPV6FirewallRule) ConvertFrom(hub conversion.Hub) error {
	source, ok := hub.(*storage.ServersIPV6FirewallRule)
	if !ok {
		return fmt.Errorf("expected sql/v1api20211101/storage/ServersIPV6FirewallRule but received %T instead", hub)
	}

	return rule.AssignProperties_From_ServersIPV6FirewallRule(source)
}

// ConvertTo populates the provided hub ServersIPV6FirewallRule from our ServersIPV6FirewallRule
func (rule *ServersIPV6FirewallRule) ConvertTo(hub conversion.Hub) error {
	destination, ok := hub.(*storage.ServersIPV6FirewallRule)
	if !ok {
		return fmt.Errorf("expected sql/v1api20211101/storage/ServersIPV6FirewallRule but received %T instead", hub)
	}

	return rule.AssignProperties_To_ServersIPV6FirewallRule(destination)
}

// +kubebuilder:webhook:path=/mutate-sql-azure-com-v1api20211101-serversipv6firewallrule,mutating=true,sideEffects=None,matchPolicy=Exact,failurePolicy=fail,groups=sql.azure.com,resources=serversipv6firewallrules,verbs=create;update,versions=v1api20211101,name=default.v1api20211101.serversipv6firewallrules.sql.azure.com,admissionReviewVersions=v1

var _ admission.Defaulter = &ServersIPV6FirewallRule{}

// Default applies defaults to the ServersIPV6FirewallRule resource
func (rule *ServersIPV6FirewallRule) Default() {
	rule.defaultImpl()
	var temp any = rule
	if runtimeDefaulter, ok := temp.(genruntime.Defaulter); ok {
		runtimeDefaulter.CustomDefault()
	}
}

// defaultAzureName defaults the Azure name of the resource to the Kubernetes name
func (rule *ServersIPV6FirewallRule) defaultAzureName() {
	if rule.Spec.AzureName == "" {
		rule.Spec.AzureName = rule.Name
	}
}

// defaultImpl applies the code generated defaults to the ServersIPV6FirewallRule resource
func (rule *ServersIPV6FirewallRule) defaultImpl() { rule.defaultAzureName() }

var _ genruntime.ImportableResource = &ServersIPV6FirewallRule{}

// InitializeSpec initializes the spec for this resource from the given status
func (rule *ServersIPV6FirewallRule) InitializeSpec(status genruntime.ConvertibleStatus) error {
	if s, ok := status.(*Servers_Ipv6FirewallRule_STATUS); ok {
		return rule.Spec.Initialize_From_Servers_Ipv6FirewallRule_STATUS(s)
	}

	return fmt.Errorf("expected Status of type Servers_Ipv6FirewallRule_STATUS but received %T instead", status)
}

var _ genruntime.KubernetesResource = &ServersIPV6FirewallRule{}

// AzureName returns the Azure name of the resource
func (rule *ServersIPV6FirewallRule) AzureName() string {
	return rule.Spec.AzureName
}

// GetAPIVersion returns the ARM API version of the resource. This is always "2021-11-01"
func (rule ServersIPV6FirewallRule) GetAPIVersion() string {
	return string(APIVersion_Value)
}

// GetResourceScope returns the scope of the resource
func (rule *ServersIPV6FirewallRule) GetResourceScope() genruntime.ResourceScope {
	return genruntime.ResourceScopeResourceGroup
}

// GetSpec returns the specification of this resource
func (rule *ServersIPV6FirewallRule) GetSpec() genruntime.ConvertibleSpec {
	return &rule.Spec
}

// GetStatus returns the status of this resource
func (rule *ServersIPV6FirewallRule) GetStatus() genruntime.ConvertibleStatus {
	return &rule.Status
}

// GetSupportedOperations returns the operations supported by the resource
func (rule *ServersIPV6FirewallRule) GetSupportedOperations() []genruntime.ResourceOperation {
	return []genruntime.ResourceOperation{
		genruntime.ResourceOperationDelete,
		genruntime.ResourceOperationGet,
		genruntime.ResourceOperationPut,
	}
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.Sql/servers/ipv6FirewallRules"
func (rule *ServersIPV6FirewallRule) GetType() string {
	return "Microsoft.Sql/servers/ipv6FirewallRules"
}

// NewEmptyStatus returns a new empty (blank) status
func (rule *ServersIPV6FirewallRule) NewEmptyStatus() genruntime.ConvertibleStatus {
	return &Servers_Ipv6FirewallRule_STATUS{}
}

// Owner returns the ResourceReference of the owner
func (rule *ServersIPV6FirewallRule) Owner() *genruntime.ResourceReference {
	group, kind := genruntime.LookupOwnerGroupKind(rule.Spec)
	return rule.Spec.Owner.AsResourceReference(group, kind)
}

// SetStatus sets the status of this resource
func (rule *ServersIPV6FirewallRule) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
	if st, ok := status.(*Servers_Ipv6FirewallRule_STATUS); ok {
		rule.Status = *st
		return nil
	}

	// Convert status to required version
	var st Servers_Ipv6FirewallRule_STATUS
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return errors.Wrap(err, "failed to convert status")
	}

	rule.Status = st
	return nil
}

// +kubebuilder:webhook:path=/validate-sql-azure-com-v1api20211101-serversipv6firewallrule,mutating=false,sideEffects=None,matchPolicy=Exact,failurePolicy=fail,groups=sql.azure.com,resources=serversipv6firewallrules,verbs=create;update,versions=v1api20211101,name=validate.v1api20211101.serversipv6firewallrules.sql.azure.com,admissionReviewVersions=v1

var _ admission.Validator = &ServersIPV6FirewallRule{}

// ValidateCreate validates the creation of the resource
func (rule *ServersIPV6FirewallRule) ValidateCreate() (admission.Warnings, error) {
	validations := rule.createValidations()
	var temp any = rule
	if runtimeValidator, ok := temp.(genruntime.Validator); ok {
		validations = append(validations, runtimeValidator.CreateValidations()...)
	}
	return genruntime.ValidateCreate(validations)
}

// ValidateDelete validates the deletion of the resource
func (rule *ServersIPV6FirewallRule) ValidateDelete() (admission.Warnings, error) {
	validations := rule.deleteValidations()
	var temp any = rule
	if runtimeValidator, ok := temp.(genruntime.Validator); ok {
		validations = append(validations, runtimeValidator.DeleteValidations()...)
	}
	return genruntime.ValidateDelete(validations)
}

// ValidateUpdate validates an update of the resource
func (rule *ServersIPV6FirewallRule) ValidateUpdate(old runtime.Object) (admission.Warnings, error) {
	validations := rule.updateValidations()
	var temp any = rule
	if runtimeValidator, ok := temp.(genruntime.Validator); ok {
		validations = append(validations, runtimeValidator.UpdateValidations()...)
	}
	return genruntime.ValidateUpdate(old, validations)
}

// createValidations validates the creation of the resource
func (rule *ServersIPV6FirewallRule) createValidations() []func() (admission.Warnings, error) {
	return []func() (admission.Warnings, error){rule.validateResourceReferences, rule.validateOwnerReference}
}

// deleteValidations validates the deletion of the resource
func (rule *ServersIPV6FirewallRule) deleteValidations() []func() (admission.Warnings, error) {
	return nil
}

// updateValidations validates the update of the resource
func (rule *ServersIPV6FirewallRule) updateValidations() []func(old runtime.Object) (admission.Warnings, error) {
	return []func(old runtime.Object) (admission.Warnings, error){
		func(old runtime.Object) (admission.Warnings, error) {
			return rule.validateResourceReferences()
		},
		rule.validateWriteOnceProperties,
		func(old runtime.Object) (admission.Warnings, error) {
			return rule.validateOwnerReference()
		},
	}
}

// validateOwnerReference validates the owner field
func (rule *ServersIPV6FirewallRule) validateOwnerReference() (admission.Warnings, error) {
	return genruntime.ValidateOwner(rule)
}

// validateResourceReferences validates all resource references
func (rule *ServersIPV6FirewallRule) validateResourceReferences() (admission.Warnings, error) {
	refs, err := reflecthelpers.FindResourceReferences(&rule.Spec)
	if err != nil {
		return nil, err
	}
	return genruntime.ValidateResourceReferences(refs)
}

// validateWriteOnceProperties validates all WriteOnce properties
func (rule *ServersIPV6FirewallRule) validateWriteOnceProperties(old runtime.Object) (admission.Warnings, error) {
	oldObj, ok := old.(*ServersIPV6FirewallRule)
	if !ok {
		return nil, nil
	}

	return genruntime.ValidateWriteOnceProperties(oldObj, rule)
}

// AssignProperties_From_ServersIPV6FirewallRule populates our ServersIPV6FirewallRule from the provided source ServersIPV6FirewallRule
func (rule *ServersIPV6FirewallRule) AssignProperties_From_ServersIPV6FirewallRule(source *storage.ServersIPV6FirewallRule) error {

	// ObjectMeta
	rule.ObjectMeta = *source.ObjectMeta.DeepCopy()

	// Spec
	var spec Servers_Ipv6FirewallRule_Spec
	err := spec.AssignProperties_From_Servers_Ipv6FirewallRule_Spec(&source.Spec)
	if err != nil {
		return errors.Wrap(err, "calling AssignProperties_From_Servers_Ipv6FirewallRule_Spec() to populate field Spec")
	}
	rule.Spec = spec

	// Status
	var status Servers_Ipv6FirewallRule_STATUS
	err = status.AssignProperties_From_Servers_Ipv6FirewallRule_STATUS(&source.Status)
	if err != nil {
		return errors.Wrap(err, "calling AssignProperties_From_Servers_Ipv6FirewallRule_STATUS() to populate field Status")
	}
	rule.Status = status

	// No error
	return nil
}

// AssignProperties_To_ServersIPV6FirewallRule populates the provided destination ServersIPV6FirewallRule from our ServersIPV6FirewallRule
func (rule *ServersIPV6FirewallRule) AssignProperties_To_ServersIPV6FirewallRule(destination *storage.ServersIPV6FirewallRule) error {

	// ObjectMeta
	destination.ObjectMeta = *rule.ObjectMeta.DeepCopy()

	// Spec
	var spec storage.Servers_Ipv6FirewallRule_Spec
	err := rule.Spec.AssignProperties_To_Servers_Ipv6FirewallRule_Spec(&spec)
	if err != nil {
		return errors.Wrap(err, "calling AssignProperties_To_Servers_Ipv6FirewallRule_Spec() to populate field Spec")
	}
	destination.Spec = spec

	// Status
	var status storage.Servers_Ipv6FirewallRule_STATUS
	err = rule.Status.AssignProperties_To_Servers_Ipv6FirewallRule_STATUS(&status)
	if err != nil {
		return errors.Wrap(err, "calling AssignProperties_To_Servers_Ipv6FirewallRule_STATUS() to populate field Status")
	}
	destination.Status = status

	// No error
	return nil
}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (rule *ServersIPV6FirewallRule) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: rule.Spec.OriginalVersion(),
		Kind:    "ServersIPV6FirewallRule",
	}
}

// +kubebuilder:object:root=true
// Generator information:
// - Generated from: /sql/resource-manager/Microsoft.Sql/stable/2021-11-01/IPv6FirewallRules.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/ipv6FirewallRules/{firewallRuleName}
type ServersIPV6FirewallRuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ServersIPV6FirewallRule `json:"items"`
}

type Servers_Ipv6FirewallRule_Spec struct {
	// AzureName: The name of the resource in Azure. This is often the same as the name of the resource in Kubernetes but it
	// doesn't have to be.
	AzureName string `json:"azureName,omitempty"`

	// EndIPv6Address: The end IP address of the firewall rule. Must be IPv6 format. Must be greater than or equal to
	// startIpAddress.
	EndIPv6Address *string `json:"endIPv6Address,omitempty"`

	// +kubebuilder:validation:Required
	// Owner: The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also
	// controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a
	// reference to a sql.azure.com/Server resource
	Owner *genruntime.KnownResourceReference `group:"sql.azure.com" json:"owner,omitempty" kind:"Server"`

	// StartIPv6Address: The start IP address of the firewall rule. Must be IPv6 format.
	StartIPv6Address *string `json:"startIPv6Address,omitempty"`
}

var _ genruntime.ARMTransformer = &Servers_Ipv6FirewallRule_Spec{}

// ConvertToARM converts from a Kubernetes CRD object to an ARM object
func (rule *Servers_Ipv6FirewallRule_Spec) ConvertToARM(resolved genruntime.ConvertToARMResolvedDetails) (interface{}, error) {
	if rule == nil {
		return nil, nil
	}
	result := &Servers_Ipv6FirewallRule_Spec_ARM{}

	// Set property "Name":
	result.Name = resolved.Name

	// Set property "Properties":
	if rule.EndIPv6Address != nil || rule.StartIPv6Address != nil {
		result.Properties = &IPv6ServerFirewallRuleProperties_ARM{}
	}
	if rule.EndIPv6Address != nil {
		endIPv6Address := *rule.EndIPv6Address
		result.Properties.EndIPv6Address = &endIPv6Address
	}
	if rule.StartIPv6Address != nil {
		startIPv6Address := *rule.StartIPv6Address
		result.Properties.StartIPv6Address = &startIPv6Address
	}
	return result, nil
}

// NewEmptyARMValue returns an empty ARM value suitable for deserializing into
func (rule *Servers_Ipv6FirewallRule_Spec) NewEmptyARMValue() genruntime.ARMResourceStatus {
	return &Servers_Ipv6FirewallRule_Spec_ARM{}
}

// PopulateFromARM populates a Kubernetes CRD object from an Azure ARM object
func (rule *Servers_Ipv6FirewallRule_Spec) PopulateFromARM(owner genruntime.ArbitraryOwnerReference, armInput interface{}) error {
	typedInput, ok := armInput.(Servers_Ipv6FirewallRule_Spec_ARM)
	if !ok {
		return fmt.Errorf("unexpected type supplied for PopulateFromARM() function. Expected Servers_Ipv6FirewallRule_Spec_ARM, got %T", armInput)
	}

	// Set property "AzureName":
	rule.SetAzureName(genruntime.ExtractKubernetesResourceNameFromARMName(typedInput.Name))

	// Set property "EndIPv6Address":
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.EndIPv6Address != nil {
			endIPv6Address := *typedInput.Properties.EndIPv6Address
			rule.EndIPv6Address = &endIPv6Address
		}
	}

	// Set property "Owner":
	rule.Owner = &genruntime.KnownResourceReference{
		Name:  owner.Name,
		ARMID: owner.ARMID,
	}

	// Set property "StartIPv6Address":
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.StartIPv6Address != nil {
			startIPv6Address := *typedInput.Properties.StartIPv6Address
			rule.StartIPv6Address = &startIPv6Address
		}
	}

	// No error
	return nil
}

var _ genruntime.ConvertibleSpec = &Servers_Ipv6FirewallRule_Spec{}

// ConvertSpecFrom populates our Servers_Ipv6FirewallRule_Spec from the provided source
func (rule *Servers_Ipv6FirewallRule_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	src, ok := source.(*storage.Servers_Ipv6FirewallRule_Spec)
	if ok {
		// Populate our instance from source
		return rule.AssignProperties_From_Servers_Ipv6FirewallRule_Spec(src)
	}

	// Convert to an intermediate form
	src = &storage.Servers_Ipv6FirewallRule_Spec{}
	err := src.ConvertSpecFrom(source)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertSpecFrom()")
	}

	// Update our instance from src
	err = rule.AssignProperties_From_Servers_Ipv6FirewallRule_Spec(src)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertSpecFrom()")
	}

	return nil
}

// ConvertSpecTo populates the provided destination from our Servers_Ipv6FirewallRule_Spec
func (rule *Servers_Ipv6FirewallRule_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	dst, ok := destination.(*storage.Servers_Ipv6FirewallRule_Spec)
	if ok {
		// Populate destination from our instance
		return rule.AssignProperties_To_Servers_Ipv6FirewallRule_Spec(dst)
	}

	// Convert to an intermediate form
	dst = &storage.Servers_Ipv6FirewallRule_Spec{}
	err := rule.AssignProperties_To_Servers_Ipv6FirewallRule_Spec(dst)
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

// AssignProperties_From_Servers_Ipv6FirewallRule_Spec populates our Servers_Ipv6FirewallRule_Spec from the provided source Servers_Ipv6FirewallRule_Spec
func (rule *Servers_Ipv6FirewallRule_Spec) AssignProperties_From_Servers_Ipv6FirewallRule_Spec(source *storage.Servers_Ipv6FirewallRule_Spec) error {

	// AzureName
	rule.AzureName = source.AzureName

	// EndIPv6Address
	rule.EndIPv6Address = genruntime.ClonePointerToString(source.EndIPv6Address)

	// Owner
	if source.Owner != nil {
		owner := source.Owner.Copy()
		rule.Owner = &owner
	} else {
		rule.Owner = nil
	}

	// StartIPv6Address
	rule.StartIPv6Address = genruntime.ClonePointerToString(source.StartIPv6Address)

	// No error
	return nil
}

// AssignProperties_To_Servers_Ipv6FirewallRule_Spec populates the provided destination Servers_Ipv6FirewallRule_Spec from our Servers_Ipv6FirewallRule_Spec
func (rule *Servers_Ipv6FirewallRule_Spec) AssignProperties_To_Servers_Ipv6FirewallRule_Spec(destination *storage.Servers_Ipv6FirewallRule_Spec) error {
	// Create a new property bag
	propertyBag := genruntime.NewPropertyBag()

	// AzureName
	destination.AzureName = rule.AzureName

	// EndIPv6Address
	destination.EndIPv6Address = genruntime.ClonePointerToString(rule.EndIPv6Address)

	// OriginalVersion
	destination.OriginalVersion = rule.OriginalVersion()

	// Owner
	if rule.Owner != nil {
		owner := rule.Owner.Copy()
		destination.Owner = &owner
	} else {
		destination.Owner = nil
	}

	// StartIPv6Address
	destination.StartIPv6Address = genruntime.ClonePointerToString(rule.StartIPv6Address)

	// Update the property bag
	if len(propertyBag) > 0 {
		destination.PropertyBag = propertyBag
	} else {
		destination.PropertyBag = nil
	}

	// No error
	return nil
}

// Initialize_From_Servers_Ipv6FirewallRule_STATUS populates our Servers_Ipv6FirewallRule_Spec from the provided source Servers_Ipv6FirewallRule_STATUS
func (rule *Servers_Ipv6FirewallRule_Spec) Initialize_From_Servers_Ipv6FirewallRule_STATUS(source *Servers_Ipv6FirewallRule_STATUS) error {

	// EndIPv6Address
	rule.EndIPv6Address = genruntime.ClonePointerToString(source.EndIPv6Address)

	// StartIPv6Address
	rule.StartIPv6Address = genruntime.ClonePointerToString(source.StartIPv6Address)

	// No error
	return nil
}

// OriginalVersion returns the original API version used to create the resource.
func (rule *Servers_Ipv6FirewallRule_Spec) OriginalVersion() string {
	return GroupVersion.Version
}

// SetAzureName sets the Azure name of the resource
func (rule *Servers_Ipv6FirewallRule_Spec) SetAzureName(azureName string) { rule.AzureName = azureName }

type Servers_Ipv6FirewallRule_STATUS struct {
	// Conditions: The observed state of the resource
	Conditions []conditions.Condition `json:"conditions,omitempty"`

	// EndIPv6Address: The end IP address of the firewall rule. Must be IPv6 format. Must be greater than or equal to
	// startIpAddress.
	EndIPv6Address *string `json:"endIPv6Address,omitempty"`

	// Id: Resource ID.
	Id *string `json:"id,omitempty"`

	// Name: Resource name.
	Name *string `json:"name,omitempty"`

	// StartIPv6Address: The start IP address of the firewall rule. Must be IPv6 format.
	StartIPv6Address *string `json:"startIPv6Address,omitempty"`

	// Type: Resource type.
	Type *string `json:"type,omitempty"`
}

var _ genruntime.ConvertibleStatus = &Servers_Ipv6FirewallRule_STATUS{}

// ConvertStatusFrom populates our Servers_Ipv6FirewallRule_STATUS from the provided source
func (rule *Servers_Ipv6FirewallRule_STATUS) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	src, ok := source.(*storage.Servers_Ipv6FirewallRule_STATUS)
	if ok {
		// Populate our instance from source
		return rule.AssignProperties_From_Servers_Ipv6FirewallRule_STATUS(src)
	}

	// Convert to an intermediate form
	src = &storage.Servers_Ipv6FirewallRule_STATUS{}
	err := src.ConvertStatusFrom(source)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertStatusFrom()")
	}

	// Update our instance from src
	err = rule.AssignProperties_From_Servers_Ipv6FirewallRule_STATUS(src)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertStatusFrom()")
	}

	return nil
}

// ConvertStatusTo populates the provided destination from our Servers_Ipv6FirewallRule_STATUS
func (rule *Servers_Ipv6FirewallRule_STATUS) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	dst, ok := destination.(*storage.Servers_Ipv6FirewallRule_STATUS)
	if ok {
		// Populate destination from our instance
		return rule.AssignProperties_To_Servers_Ipv6FirewallRule_STATUS(dst)
	}

	// Convert to an intermediate form
	dst = &storage.Servers_Ipv6FirewallRule_STATUS{}
	err := rule.AssignProperties_To_Servers_Ipv6FirewallRule_STATUS(dst)
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

var _ genruntime.FromARMConverter = &Servers_Ipv6FirewallRule_STATUS{}

// NewEmptyARMValue returns an empty ARM value suitable for deserializing into
func (rule *Servers_Ipv6FirewallRule_STATUS) NewEmptyARMValue() genruntime.ARMResourceStatus {
	return &Servers_Ipv6FirewallRule_STATUS_ARM{}
}

// PopulateFromARM populates a Kubernetes CRD object from an Azure ARM object
func (rule *Servers_Ipv6FirewallRule_STATUS) PopulateFromARM(owner genruntime.ArbitraryOwnerReference, armInput interface{}) error {
	typedInput, ok := armInput.(Servers_Ipv6FirewallRule_STATUS_ARM)
	if !ok {
		return fmt.Errorf("unexpected type supplied for PopulateFromARM() function. Expected Servers_Ipv6FirewallRule_STATUS_ARM, got %T", armInput)
	}

	// no assignment for property "Conditions"

	// Set property "EndIPv6Address":
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.EndIPv6Address != nil {
			endIPv6Address := *typedInput.Properties.EndIPv6Address
			rule.EndIPv6Address = &endIPv6Address
		}
	}

	// Set property "Id":
	if typedInput.Id != nil {
		id := *typedInput.Id
		rule.Id = &id
	}

	// Set property "Name":
	if typedInput.Name != nil {
		name := *typedInput.Name
		rule.Name = &name
	}

	// Set property "StartIPv6Address":
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.StartIPv6Address != nil {
			startIPv6Address := *typedInput.Properties.StartIPv6Address
			rule.StartIPv6Address = &startIPv6Address
		}
	}

	// Set property "Type":
	if typedInput.Type != nil {
		typeVar := *typedInput.Type
		rule.Type = &typeVar
	}

	// No error
	return nil
}

// AssignProperties_From_Servers_Ipv6FirewallRule_STATUS populates our Servers_Ipv6FirewallRule_STATUS from the provided source Servers_Ipv6FirewallRule_STATUS
func (rule *Servers_Ipv6FirewallRule_STATUS) AssignProperties_From_Servers_Ipv6FirewallRule_STATUS(source *storage.Servers_Ipv6FirewallRule_STATUS) error {

	// Conditions
	rule.Conditions = genruntime.CloneSliceOfCondition(source.Conditions)

	// EndIPv6Address
	rule.EndIPv6Address = genruntime.ClonePointerToString(source.EndIPv6Address)

	// Id
	rule.Id = genruntime.ClonePointerToString(source.Id)

	// Name
	rule.Name = genruntime.ClonePointerToString(source.Name)

	// StartIPv6Address
	rule.StartIPv6Address = genruntime.ClonePointerToString(source.StartIPv6Address)

	// Type
	rule.Type = genruntime.ClonePointerToString(source.Type)

	// No error
	return nil
}

// AssignProperties_To_Servers_Ipv6FirewallRule_STATUS populates the provided destination Servers_Ipv6FirewallRule_STATUS from our Servers_Ipv6FirewallRule_STATUS
func (rule *Servers_Ipv6FirewallRule_STATUS) AssignProperties_To_Servers_Ipv6FirewallRule_STATUS(destination *storage.Servers_Ipv6FirewallRule_STATUS) error {
	// Create a new property bag
	propertyBag := genruntime.NewPropertyBag()

	// Conditions
	destination.Conditions = genruntime.CloneSliceOfCondition(rule.Conditions)

	// EndIPv6Address
	destination.EndIPv6Address = genruntime.ClonePointerToString(rule.EndIPv6Address)

	// Id
	destination.Id = genruntime.ClonePointerToString(rule.Id)

	// Name
	destination.Name = genruntime.ClonePointerToString(rule.Name)

	// StartIPv6Address
	destination.StartIPv6Address = genruntime.ClonePointerToString(rule.StartIPv6Address)

	// Type
	destination.Type = genruntime.ClonePointerToString(rule.Type)

	// Update the property bag
	if len(propertyBag) > 0 {
		destination.PropertyBag = propertyBag
	} else {
		destination.PropertyBag = nil
	}

	// No error
	return nil
}

func init() {
	SchemeBuilder.Register(&ServersIPV6FirewallRule{}, &ServersIPV6FirewallRuleList{})
}
