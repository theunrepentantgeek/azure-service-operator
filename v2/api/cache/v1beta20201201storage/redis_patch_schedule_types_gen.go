// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20201201storage

import (
	v20201201s "github.com/Azure/azure-service-operator/v2/api/cache/v1api20201201storage"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
	"github.com/pkg/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"sigs.k8s.io/controller-runtime/pkg/conversion"
)

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="Severity",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].severity"
// +kubebuilder:printcolumn:name="Reason",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].reason"
// +kubebuilder:printcolumn:name="Message",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].message"
// Storage version of v1beta20201201.RedisPatchSchedule
// Deprecated version of RedisPatchSchedule. Use v1api20201201.RedisPatchSchedule instead
type RedisPatchSchedule struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              Redis_PatchSchedule_Spec   `json:"spec,omitempty"`
	Status            Redis_PatchSchedule_STATUS `json:"status,omitempty"`
}

var _ conditions.Conditioner = &RedisPatchSchedule{}

// GetConditions returns the conditions of the resource
func (schedule *RedisPatchSchedule) GetConditions() conditions.Conditions {
	return schedule.Status.Conditions
}

// SetConditions sets the conditions on the resource status
func (schedule *RedisPatchSchedule) SetConditions(conditions conditions.Conditions) {
	schedule.Status.Conditions = conditions
}

var _ conversion.Convertible = &RedisPatchSchedule{}

// ConvertFrom populates our RedisPatchSchedule from the provided hub RedisPatchSchedule
func (schedule *RedisPatchSchedule) ConvertFrom(hub conversion.Hub) error {
	// intermediate variable for conversion
	var source v20201201s.RedisPatchSchedule

	err := source.ConvertFrom(hub)
	if err != nil {
		return errors.Wrap(err, "converting from hub to source")
	}

	err = schedule.AssignProperties_From_RedisPatchSchedule(&source)
	if err != nil {
		return errors.Wrap(err, "converting from source to schedule")
	}

	return nil
}

// ConvertTo populates the provided hub RedisPatchSchedule from our RedisPatchSchedule
func (schedule *RedisPatchSchedule) ConvertTo(hub conversion.Hub) error {
	// intermediate variable for conversion
	var destination v20201201s.RedisPatchSchedule
	err := schedule.AssignProperties_To_RedisPatchSchedule(&destination)
	if err != nil {
		return errors.Wrap(err, "converting to destination from schedule")
	}
	err = destination.ConvertTo(hub)
	if err != nil {
		return errors.Wrap(err, "converting from destination to hub")
	}

	return nil
}

var _ genruntime.KubernetesResource = &RedisPatchSchedule{}

// AzureName returns the Azure name of the resource (always "default")
func (schedule *RedisPatchSchedule) AzureName() string {
	return "default"
}

// GetAPIVersion returns the ARM API version of the resource. This is always "2020-12-01"
func (schedule RedisPatchSchedule) GetAPIVersion() string {
	return string(APIVersion_Value)
}

// GetResourceScope returns the scope of the resource
func (schedule *RedisPatchSchedule) GetResourceScope() genruntime.ResourceScope {
	return genruntime.ResourceScopeResourceGroup
}

// GetSpec returns the specification of this resource
func (schedule *RedisPatchSchedule) GetSpec() genruntime.ConvertibleSpec {
	return &schedule.Spec
}

// GetStatus returns the status of this resource
func (schedule *RedisPatchSchedule) GetStatus() genruntime.ConvertibleStatus {
	return &schedule.Status
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.Cache/redis/patchSchedules"
func (schedule *RedisPatchSchedule) GetType() string {
	return "Microsoft.Cache/redis/patchSchedules"
}

// NewEmptyStatus returns a new empty (blank) status
func (schedule *RedisPatchSchedule) NewEmptyStatus() genruntime.ConvertibleStatus {
	return &Redis_PatchSchedule_STATUS{}
}

// Owner returns the ResourceReference of the owner
func (schedule *RedisPatchSchedule) Owner() *genruntime.ResourceReference {
	group, kind := genruntime.LookupOwnerGroupKind(schedule.Spec)
	return &genruntime.ResourceReference{
		Group: group,
		Kind:  kind,
		Name:  schedule.Spec.Owner.Name,
	}
}

// SetStatus sets the status of this resource
func (schedule *RedisPatchSchedule) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
	if st, ok := status.(*Redis_PatchSchedule_STATUS); ok {
		schedule.Status = *st
		return nil
	}

	// Convert status to required version
	var st Redis_PatchSchedule_STATUS
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return errors.Wrap(err, "failed to convert status")
	}

	schedule.Status = st
	return nil
}

// AssignProperties_From_RedisPatchSchedule populates our RedisPatchSchedule from the provided source RedisPatchSchedule
func (schedule *RedisPatchSchedule) AssignProperties_From_RedisPatchSchedule(source *v20201201s.RedisPatchSchedule) error {

	// ObjectMeta
	schedule.ObjectMeta = *source.ObjectMeta.DeepCopy()

	// Spec
	var spec Redis_PatchSchedule_Spec
	err := spec.AssignProperties_From_Redis_PatchSchedule_Spec(&source.Spec)
	if err != nil {
		return errors.Wrap(err, "calling AssignProperties_From_Redis_PatchSchedule_Spec() to populate field Spec")
	}
	schedule.Spec = spec

	// Status
	var status Redis_PatchSchedule_STATUS
	err = status.AssignProperties_From_Redis_PatchSchedule_STATUS(&source.Status)
	if err != nil {
		return errors.Wrap(err, "calling AssignProperties_From_Redis_PatchSchedule_STATUS() to populate field Status")
	}
	schedule.Status = status

	// Invoke the augmentConversionForRedisPatchSchedule interface (if implemented) to customize the conversion
	var scheduleAsAny any = schedule
	if augmentedSchedule, ok := scheduleAsAny.(augmentConversionForRedisPatchSchedule); ok {
		err := augmentedSchedule.AssignPropertiesFrom(source)
		if err != nil {
			return errors.Wrap(err, "calling augmented AssignPropertiesFrom() for conversion")
		}
	}

	// No error
	return nil
}

// AssignProperties_To_RedisPatchSchedule populates the provided destination RedisPatchSchedule from our RedisPatchSchedule
func (schedule *RedisPatchSchedule) AssignProperties_To_RedisPatchSchedule(destination *v20201201s.RedisPatchSchedule) error {

	// ObjectMeta
	destination.ObjectMeta = *schedule.ObjectMeta.DeepCopy()

	// Spec
	var spec v20201201s.Redis_PatchSchedule_Spec
	err := schedule.Spec.AssignProperties_To_Redis_PatchSchedule_Spec(&spec)
	if err != nil {
		return errors.Wrap(err, "calling AssignProperties_To_Redis_PatchSchedule_Spec() to populate field Spec")
	}
	destination.Spec = spec

	// Status
	var status v20201201s.Redis_PatchSchedule_STATUS
	err = schedule.Status.AssignProperties_To_Redis_PatchSchedule_STATUS(&status)
	if err != nil {
		return errors.Wrap(err, "calling AssignProperties_To_Redis_PatchSchedule_STATUS() to populate field Status")
	}
	destination.Status = status

	// Invoke the augmentConversionForRedisPatchSchedule interface (if implemented) to customize the conversion
	var scheduleAsAny any = schedule
	if augmentedSchedule, ok := scheduleAsAny.(augmentConversionForRedisPatchSchedule); ok {
		err := augmentedSchedule.AssignPropertiesTo(destination)
		if err != nil {
			return errors.Wrap(err, "calling augmented AssignPropertiesTo() for conversion")
		}
	}

	// No error
	return nil
}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (schedule *RedisPatchSchedule) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: schedule.Spec.OriginalVersion,
		Kind:    "RedisPatchSchedule",
	}
}

// +kubebuilder:object:root=true
// Storage version of v1beta20201201.RedisPatchSchedule
// Deprecated version of RedisPatchSchedule. Use v1api20201201.RedisPatchSchedule instead
type RedisPatchScheduleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RedisPatchSchedule `json:"items"`
}

type augmentConversionForRedisPatchSchedule interface {
	AssignPropertiesFrom(src *v20201201s.RedisPatchSchedule) error
	AssignPropertiesTo(dst *v20201201s.RedisPatchSchedule) error
}

// Storage version of v1beta20201201.Redis_PatchSchedule_Spec
type Redis_PatchSchedule_Spec struct {
	OriginalVersion string `json:"originalVersion,omitempty"`

	// +kubebuilder:validation:Required
	// Owner: The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also
	// controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a
	// reference to a cache.azure.com/Redis resource
	Owner           *genruntime.KnownResourceReference `group:"cache.azure.com" json:"owner,omitempty" kind:"Redis"`
	PropertyBag     genruntime.PropertyBag             `json:"$propertyBag,omitempty"`
	ScheduleEntries []ScheduleEntry                    `json:"scheduleEntries,omitempty"`
}

var _ genruntime.ConvertibleSpec = &Redis_PatchSchedule_Spec{}

// ConvertSpecFrom populates our Redis_PatchSchedule_Spec from the provided source
func (schedule *Redis_PatchSchedule_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	src, ok := source.(*v20201201s.Redis_PatchSchedule_Spec)
	if ok {
		// Populate our instance from source
		return schedule.AssignProperties_From_Redis_PatchSchedule_Spec(src)
	}

	// Convert to an intermediate form
	src = &v20201201s.Redis_PatchSchedule_Spec{}
	err := src.ConvertSpecFrom(source)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertSpecFrom()")
	}

	// Update our instance from src
	err = schedule.AssignProperties_From_Redis_PatchSchedule_Spec(src)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertSpecFrom()")
	}

	return nil
}

// ConvertSpecTo populates the provided destination from our Redis_PatchSchedule_Spec
func (schedule *Redis_PatchSchedule_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	dst, ok := destination.(*v20201201s.Redis_PatchSchedule_Spec)
	if ok {
		// Populate destination from our instance
		return schedule.AssignProperties_To_Redis_PatchSchedule_Spec(dst)
	}

	// Convert to an intermediate form
	dst = &v20201201s.Redis_PatchSchedule_Spec{}
	err := schedule.AssignProperties_To_Redis_PatchSchedule_Spec(dst)
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

// AssignProperties_From_Redis_PatchSchedule_Spec populates our Redis_PatchSchedule_Spec from the provided source Redis_PatchSchedule_Spec
func (schedule *Redis_PatchSchedule_Spec) AssignProperties_From_Redis_PatchSchedule_Spec(source *v20201201s.Redis_PatchSchedule_Spec) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(source.PropertyBag)

	// OriginalVersion
	schedule.OriginalVersion = source.OriginalVersion

	// Owner
	if source.Owner != nil {
		owner := source.Owner.Copy()
		schedule.Owner = &owner
	} else {
		schedule.Owner = nil
	}

	// ScheduleEntries
	if source.ScheduleEntries != nil {
		scheduleEntryList := make([]ScheduleEntry, len(source.ScheduleEntries))
		for scheduleEntryIndex, scheduleEntryItem := range source.ScheduleEntries {
			// Shadow the loop variable to avoid aliasing
			scheduleEntryItem := scheduleEntryItem
			var scheduleEntry ScheduleEntry
			err := scheduleEntry.AssignProperties_From_ScheduleEntry(&scheduleEntryItem)
			if err != nil {
				return errors.Wrap(err, "calling AssignProperties_From_ScheduleEntry() to populate field ScheduleEntries")
			}
			scheduleEntryList[scheduleEntryIndex] = scheduleEntry
		}
		schedule.ScheduleEntries = scheduleEntryList
	} else {
		schedule.ScheduleEntries = nil
	}

	// Update the property bag
	if len(propertyBag) > 0 {
		schedule.PropertyBag = propertyBag
	} else {
		schedule.PropertyBag = nil
	}

	// Invoke the augmentConversionForRedis_PatchSchedule_Spec interface (if implemented) to customize the conversion
	var scheduleAsAny any = schedule
	if augmentedSchedule, ok := scheduleAsAny.(augmentConversionForRedis_PatchSchedule_Spec); ok {
		err := augmentedSchedule.AssignPropertiesFrom(source)
		if err != nil {
			return errors.Wrap(err, "calling augmented AssignPropertiesFrom() for conversion")
		}
	}

	// No error
	return nil
}

// AssignProperties_To_Redis_PatchSchedule_Spec populates the provided destination Redis_PatchSchedule_Spec from our Redis_PatchSchedule_Spec
func (schedule *Redis_PatchSchedule_Spec) AssignProperties_To_Redis_PatchSchedule_Spec(destination *v20201201s.Redis_PatchSchedule_Spec) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(schedule.PropertyBag)

	// OriginalVersion
	destination.OriginalVersion = schedule.OriginalVersion

	// Owner
	if schedule.Owner != nil {
		owner := schedule.Owner.Copy()
		destination.Owner = &owner
	} else {
		destination.Owner = nil
	}

	// ScheduleEntries
	if schedule.ScheduleEntries != nil {
		scheduleEntryList := make([]v20201201s.ScheduleEntry, len(schedule.ScheduleEntries))
		for scheduleEntryIndex, scheduleEntryItem := range schedule.ScheduleEntries {
			// Shadow the loop variable to avoid aliasing
			scheduleEntryItem := scheduleEntryItem
			var scheduleEntry v20201201s.ScheduleEntry
			err := scheduleEntryItem.AssignProperties_To_ScheduleEntry(&scheduleEntry)
			if err != nil {
				return errors.Wrap(err, "calling AssignProperties_To_ScheduleEntry() to populate field ScheduleEntries")
			}
			scheduleEntryList[scheduleEntryIndex] = scheduleEntry
		}
		destination.ScheduleEntries = scheduleEntryList
	} else {
		destination.ScheduleEntries = nil
	}

	// Update the property bag
	if len(propertyBag) > 0 {
		destination.PropertyBag = propertyBag
	} else {
		destination.PropertyBag = nil
	}

	// Invoke the augmentConversionForRedis_PatchSchedule_Spec interface (if implemented) to customize the conversion
	var scheduleAsAny any = schedule
	if augmentedSchedule, ok := scheduleAsAny.(augmentConversionForRedis_PatchSchedule_Spec); ok {
		err := augmentedSchedule.AssignPropertiesTo(destination)
		if err != nil {
			return errors.Wrap(err, "calling augmented AssignPropertiesTo() for conversion")
		}
	}

	// No error
	return nil
}

// Storage version of v1beta20201201.Redis_PatchSchedule_STATUS
// Deprecated version of Redis_PatchSchedule_STATUS. Use v1api20201201.Redis_PatchSchedule_STATUS instead
type Redis_PatchSchedule_STATUS struct {
	Conditions      []conditions.Condition `json:"conditions,omitempty"`
	Id              *string                `json:"id,omitempty"`
	Location        *string                `json:"location,omitempty"`
	Name            *string                `json:"name,omitempty"`
	PropertyBag     genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	ScheduleEntries []ScheduleEntry_STATUS `json:"scheduleEntries,omitempty"`
	Type            *string                `json:"type,omitempty"`
}

var _ genruntime.ConvertibleStatus = &Redis_PatchSchedule_STATUS{}

// ConvertStatusFrom populates our Redis_PatchSchedule_STATUS from the provided source
func (schedule *Redis_PatchSchedule_STATUS) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	src, ok := source.(*v20201201s.Redis_PatchSchedule_STATUS)
	if ok {
		// Populate our instance from source
		return schedule.AssignProperties_From_Redis_PatchSchedule_STATUS(src)
	}

	// Convert to an intermediate form
	src = &v20201201s.Redis_PatchSchedule_STATUS{}
	err := src.ConvertStatusFrom(source)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertStatusFrom()")
	}

	// Update our instance from src
	err = schedule.AssignProperties_From_Redis_PatchSchedule_STATUS(src)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertStatusFrom()")
	}

	return nil
}

// ConvertStatusTo populates the provided destination from our Redis_PatchSchedule_STATUS
func (schedule *Redis_PatchSchedule_STATUS) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	dst, ok := destination.(*v20201201s.Redis_PatchSchedule_STATUS)
	if ok {
		// Populate destination from our instance
		return schedule.AssignProperties_To_Redis_PatchSchedule_STATUS(dst)
	}

	// Convert to an intermediate form
	dst = &v20201201s.Redis_PatchSchedule_STATUS{}
	err := schedule.AssignProperties_To_Redis_PatchSchedule_STATUS(dst)
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

// AssignProperties_From_Redis_PatchSchedule_STATUS populates our Redis_PatchSchedule_STATUS from the provided source Redis_PatchSchedule_STATUS
func (schedule *Redis_PatchSchedule_STATUS) AssignProperties_From_Redis_PatchSchedule_STATUS(source *v20201201s.Redis_PatchSchedule_STATUS) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(source.PropertyBag)

	// Conditions
	schedule.Conditions = genruntime.CloneSliceOfCondition(source.Conditions)

	// Id
	schedule.Id = genruntime.ClonePointerToString(source.Id)

	// Location
	schedule.Location = genruntime.ClonePointerToString(source.Location)

	// Name
	schedule.Name = genruntime.ClonePointerToString(source.Name)

	// ScheduleEntries
	if source.ScheduleEntries != nil {
		scheduleEntryList := make([]ScheduleEntry_STATUS, len(source.ScheduleEntries))
		for scheduleEntryIndex, scheduleEntryItem := range source.ScheduleEntries {
			// Shadow the loop variable to avoid aliasing
			scheduleEntryItem := scheduleEntryItem
			var scheduleEntry ScheduleEntry_STATUS
			err := scheduleEntry.AssignProperties_From_ScheduleEntry_STATUS(&scheduleEntryItem)
			if err != nil {
				return errors.Wrap(err, "calling AssignProperties_From_ScheduleEntry_STATUS() to populate field ScheduleEntries")
			}
			scheduleEntryList[scheduleEntryIndex] = scheduleEntry
		}
		schedule.ScheduleEntries = scheduleEntryList
	} else {
		schedule.ScheduleEntries = nil
	}

	// Type
	schedule.Type = genruntime.ClonePointerToString(source.Type)

	// Update the property bag
	if len(propertyBag) > 0 {
		schedule.PropertyBag = propertyBag
	} else {
		schedule.PropertyBag = nil
	}

	// Invoke the augmentConversionForRedis_PatchSchedule_STATUS interface (if implemented) to customize the conversion
	var scheduleAsAny any = schedule
	if augmentedSchedule, ok := scheduleAsAny.(augmentConversionForRedis_PatchSchedule_STATUS); ok {
		err := augmentedSchedule.AssignPropertiesFrom(source)
		if err != nil {
			return errors.Wrap(err, "calling augmented AssignPropertiesFrom() for conversion")
		}
	}

	// No error
	return nil
}

// AssignProperties_To_Redis_PatchSchedule_STATUS populates the provided destination Redis_PatchSchedule_STATUS from our Redis_PatchSchedule_STATUS
func (schedule *Redis_PatchSchedule_STATUS) AssignProperties_To_Redis_PatchSchedule_STATUS(destination *v20201201s.Redis_PatchSchedule_STATUS) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(schedule.PropertyBag)

	// Conditions
	destination.Conditions = genruntime.CloneSliceOfCondition(schedule.Conditions)

	// Id
	destination.Id = genruntime.ClonePointerToString(schedule.Id)

	// Location
	destination.Location = genruntime.ClonePointerToString(schedule.Location)

	// Name
	destination.Name = genruntime.ClonePointerToString(schedule.Name)

	// ScheduleEntries
	if schedule.ScheduleEntries != nil {
		scheduleEntryList := make([]v20201201s.ScheduleEntry_STATUS, len(schedule.ScheduleEntries))
		for scheduleEntryIndex, scheduleEntryItem := range schedule.ScheduleEntries {
			// Shadow the loop variable to avoid aliasing
			scheduleEntryItem := scheduleEntryItem
			var scheduleEntry v20201201s.ScheduleEntry_STATUS
			err := scheduleEntryItem.AssignProperties_To_ScheduleEntry_STATUS(&scheduleEntry)
			if err != nil {
				return errors.Wrap(err, "calling AssignProperties_To_ScheduleEntry_STATUS() to populate field ScheduleEntries")
			}
			scheduleEntryList[scheduleEntryIndex] = scheduleEntry
		}
		destination.ScheduleEntries = scheduleEntryList
	} else {
		destination.ScheduleEntries = nil
	}

	// Type
	destination.Type = genruntime.ClonePointerToString(schedule.Type)

	// Update the property bag
	if len(propertyBag) > 0 {
		destination.PropertyBag = propertyBag
	} else {
		destination.PropertyBag = nil
	}

	// Invoke the augmentConversionForRedis_PatchSchedule_STATUS interface (if implemented) to customize the conversion
	var scheduleAsAny any = schedule
	if augmentedSchedule, ok := scheduleAsAny.(augmentConversionForRedis_PatchSchedule_STATUS); ok {
		err := augmentedSchedule.AssignPropertiesTo(destination)
		if err != nil {
			return errors.Wrap(err, "calling augmented AssignPropertiesTo() for conversion")
		}
	}

	// No error
	return nil
}

type augmentConversionForRedis_PatchSchedule_Spec interface {
	AssignPropertiesFrom(src *v20201201s.Redis_PatchSchedule_Spec) error
	AssignPropertiesTo(dst *v20201201s.Redis_PatchSchedule_Spec) error
}

type augmentConversionForRedis_PatchSchedule_STATUS interface {
	AssignPropertiesFrom(src *v20201201s.Redis_PatchSchedule_STATUS) error
	AssignPropertiesTo(dst *v20201201s.Redis_PatchSchedule_STATUS) error
}

// Storage version of v1beta20201201.ScheduleEntry
// Deprecated version of ScheduleEntry. Use v1api20201201.ScheduleEntry instead
type ScheduleEntry struct {
	DayOfWeek         *string                `json:"dayOfWeek,omitempty"`
	MaintenanceWindow *string                `json:"maintenanceWindow,omitempty"`
	PropertyBag       genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	StartHourUtc      *int                   `json:"startHourUtc,omitempty"`
}

// AssignProperties_From_ScheduleEntry populates our ScheduleEntry from the provided source ScheduleEntry
func (entry *ScheduleEntry) AssignProperties_From_ScheduleEntry(source *v20201201s.ScheduleEntry) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(source.PropertyBag)

	// DayOfWeek
	entry.DayOfWeek = genruntime.ClonePointerToString(source.DayOfWeek)

	// MaintenanceWindow
	entry.MaintenanceWindow = genruntime.ClonePointerToString(source.MaintenanceWindow)

	// StartHourUtc
	entry.StartHourUtc = genruntime.ClonePointerToInt(source.StartHourUtc)

	// Update the property bag
	if len(propertyBag) > 0 {
		entry.PropertyBag = propertyBag
	} else {
		entry.PropertyBag = nil
	}

	// Invoke the augmentConversionForScheduleEntry interface (if implemented) to customize the conversion
	var entryAsAny any = entry
	if augmentedEntry, ok := entryAsAny.(augmentConversionForScheduleEntry); ok {
		err := augmentedEntry.AssignPropertiesFrom(source)
		if err != nil {
			return errors.Wrap(err, "calling augmented AssignPropertiesFrom() for conversion")
		}
	}

	// No error
	return nil
}

// AssignProperties_To_ScheduleEntry populates the provided destination ScheduleEntry from our ScheduleEntry
func (entry *ScheduleEntry) AssignProperties_To_ScheduleEntry(destination *v20201201s.ScheduleEntry) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(entry.PropertyBag)

	// DayOfWeek
	destination.DayOfWeek = genruntime.ClonePointerToString(entry.DayOfWeek)

	// MaintenanceWindow
	destination.MaintenanceWindow = genruntime.ClonePointerToString(entry.MaintenanceWindow)

	// StartHourUtc
	destination.StartHourUtc = genruntime.ClonePointerToInt(entry.StartHourUtc)

	// Update the property bag
	if len(propertyBag) > 0 {
		destination.PropertyBag = propertyBag
	} else {
		destination.PropertyBag = nil
	}

	// Invoke the augmentConversionForScheduleEntry interface (if implemented) to customize the conversion
	var entryAsAny any = entry
	if augmentedEntry, ok := entryAsAny.(augmentConversionForScheduleEntry); ok {
		err := augmentedEntry.AssignPropertiesTo(destination)
		if err != nil {
			return errors.Wrap(err, "calling augmented AssignPropertiesTo() for conversion")
		}
	}

	// No error
	return nil
}

// Storage version of v1beta20201201.ScheduleEntry_STATUS
// Deprecated version of ScheduleEntry_STATUS. Use v1api20201201.ScheduleEntry_STATUS instead
type ScheduleEntry_STATUS struct {
	DayOfWeek         *string                `json:"dayOfWeek,omitempty"`
	MaintenanceWindow *string                `json:"maintenanceWindow,omitempty"`
	PropertyBag       genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	StartHourUtc      *int                   `json:"startHourUtc,omitempty"`
}

// AssignProperties_From_ScheduleEntry_STATUS populates our ScheduleEntry_STATUS from the provided source ScheduleEntry_STATUS
func (entry *ScheduleEntry_STATUS) AssignProperties_From_ScheduleEntry_STATUS(source *v20201201s.ScheduleEntry_STATUS) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(source.PropertyBag)

	// DayOfWeek
	entry.DayOfWeek = genruntime.ClonePointerToString(source.DayOfWeek)

	// MaintenanceWindow
	entry.MaintenanceWindow = genruntime.ClonePointerToString(source.MaintenanceWindow)

	// StartHourUtc
	entry.StartHourUtc = genruntime.ClonePointerToInt(source.StartHourUtc)

	// Update the property bag
	if len(propertyBag) > 0 {
		entry.PropertyBag = propertyBag
	} else {
		entry.PropertyBag = nil
	}

	// Invoke the augmentConversionForScheduleEntry_STATUS interface (if implemented) to customize the conversion
	var entryAsAny any = entry
	if augmentedEntry, ok := entryAsAny.(augmentConversionForScheduleEntry_STATUS); ok {
		err := augmentedEntry.AssignPropertiesFrom(source)
		if err != nil {
			return errors.Wrap(err, "calling augmented AssignPropertiesFrom() for conversion")
		}
	}

	// No error
	return nil
}

// AssignProperties_To_ScheduleEntry_STATUS populates the provided destination ScheduleEntry_STATUS from our ScheduleEntry_STATUS
func (entry *ScheduleEntry_STATUS) AssignProperties_To_ScheduleEntry_STATUS(destination *v20201201s.ScheduleEntry_STATUS) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(entry.PropertyBag)

	// DayOfWeek
	destination.DayOfWeek = genruntime.ClonePointerToString(entry.DayOfWeek)

	// MaintenanceWindow
	destination.MaintenanceWindow = genruntime.ClonePointerToString(entry.MaintenanceWindow)

	// StartHourUtc
	destination.StartHourUtc = genruntime.ClonePointerToInt(entry.StartHourUtc)

	// Update the property bag
	if len(propertyBag) > 0 {
		destination.PropertyBag = propertyBag
	} else {
		destination.PropertyBag = nil
	}

	// Invoke the augmentConversionForScheduleEntry_STATUS interface (if implemented) to customize the conversion
	var entryAsAny any = entry
	if augmentedEntry, ok := entryAsAny.(augmentConversionForScheduleEntry_STATUS); ok {
		err := augmentedEntry.AssignPropertiesTo(destination)
		if err != nil {
			return errors.Wrap(err, "calling augmented AssignPropertiesTo() for conversion")
		}
	}

	// No error
	return nil
}

type augmentConversionForScheduleEntry interface {
	AssignPropertiesFrom(src *v20201201s.ScheduleEntry) error
	AssignPropertiesTo(dst *v20201201s.ScheduleEntry) error
}

type augmentConversionForScheduleEntry_STATUS interface {
	AssignPropertiesFrom(src *v20201201s.ScheduleEntry_STATUS) error
	AssignPropertiesTo(dst *v20201201s.ScheduleEntry_STATUS) error
}

func init() {
	SchemeBuilder.Register(&RedisPatchSchedule{}, &RedisPatchScheduleList{})
}
