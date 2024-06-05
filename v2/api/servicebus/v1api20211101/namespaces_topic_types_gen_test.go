// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20211101

import (
	"encoding/json"
	storage "github.com/Azure/azure-service-operator/v2/api/servicebus/v1api20211101/storage"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/kr/pretty"
	"github.com/kylelemons/godebug/diff"
	"github.com/leanovate/gopter"
	"github.com/leanovate/gopter/gen"
	"github.com/leanovate/gopter/prop"
	"os"
	"reflect"
	"testing"
)

func Test_NamespacesTopic_WhenConvertedToHub_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	parameters.MinSuccessfulTests = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from NamespacesTopic to hub returns original",
		prop.ForAll(RunResourceConversionTestForNamespacesTopic, NamespacesTopicGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunResourceConversionTestForNamespacesTopic tests if a specific instance of NamespacesTopic round trips to the hub storage version and back losslessly
func RunResourceConversionTestForNamespacesTopic(subject NamespacesTopic) string {
	// Copy subject to make sure conversion doesn't modify it
	copied := subject.DeepCopy()

	// Convert to our hub version
	var hub storage.NamespacesTopic
	err := copied.ConvertTo(&hub)
	if err != nil {
		return err.Error()
	}

	// Convert from our hub version
	var actual NamespacesTopic
	err = actual.ConvertFrom(&hub)
	if err != nil {
		return err.Error()
	}

	// Compare actual with what we started with
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

func Test_NamespacesTopic_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from NamespacesTopic to NamespacesTopic via AssignProperties_To_NamespacesTopic & AssignProperties_From_NamespacesTopic returns original",
		prop.ForAll(RunPropertyAssignmentTestForNamespacesTopic, NamespacesTopicGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForNamespacesTopic tests if a specific instance of NamespacesTopic can be assigned to storage and back losslessly
func RunPropertyAssignmentTestForNamespacesTopic(subject NamespacesTopic) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other storage.NamespacesTopic
	err := copied.AssignProperties_To_NamespacesTopic(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual NamespacesTopic
	err = actual.AssignProperties_From_NamespacesTopic(&other)
	if err != nil {
		return err.Error()
	}

	// Check for a match
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

func Test_NamespacesTopic_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 20
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of NamespacesTopic via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForNamespacesTopic, NamespacesTopicGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForNamespacesTopic runs a test to see if a specific instance of NamespacesTopic round trips to JSON and back losslessly
func RunJSONSerializationTestForNamespacesTopic(subject NamespacesTopic) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual NamespacesTopic
	err = json.Unmarshal(bin, &actual)
	if err != nil {
		return err.Error()
	}

	// Check for outcome
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

// Generator of NamespacesTopic instances for property testing - lazily instantiated by NamespacesTopicGenerator()
var namespacesTopicGenerator gopter.Gen

// NamespacesTopicGenerator returns a generator of NamespacesTopic instances for property testing.
func NamespacesTopicGenerator() gopter.Gen {
	if namespacesTopicGenerator != nil {
		return namespacesTopicGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForNamespacesTopic(generators)
	namespacesTopicGenerator = gen.Struct(reflect.TypeOf(NamespacesTopic{}), generators)

	return namespacesTopicGenerator
}

// AddRelatedPropertyGeneratorsForNamespacesTopic is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForNamespacesTopic(gens map[string]gopter.Gen) {
	gens["Spec"] = Namespaces_Topic_SpecGenerator()
	gens["Status"] = Namespaces_Topic_STATUSGenerator()
}

func Test_Namespaces_Topic_Spec_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from Namespaces_Topic_Spec to Namespaces_Topic_Spec via AssignProperties_To_Namespaces_Topic_Spec & AssignProperties_From_Namespaces_Topic_Spec returns original",
		prop.ForAll(RunPropertyAssignmentTestForNamespaces_Topic_Spec, Namespaces_Topic_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForNamespaces_Topic_Spec tests if a specific instance of Namespaces_Topic_Spec can be assigned to storage and back losslessly
func RunPropertyAssignmentTestForNamespaces_Topic_Spec(subject Namespaces_Topic_Spec) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other storage.Namespaces_Topic_Spec
	err := copied.AssignProperties_To_Namespaces_Topic_Spec(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual Namespaces_Topic_Spec
	err = actual.AssignProperties_From_Namespaces_Topic_Spec(&other)
	if err != nil {
		return err.Error()
	}

	// Check for a match
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

func Test_Namespaces_Topic_Spec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Namespaces_Topic_Spec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForNamespaces_Topic_Spec, Namespaces_Topic_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForNamespaces_Topic_Spec runs a test to see if a specific instance of Namespaces_Topic_Spec round trips to JSON and back losslessly
func RunJSONSerializationTestForNamespaces_Topic_Spec(subject Namespaces_Topic_Spec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Namespaces_Topic_Spec
	err = json.Unmarshal(bin, &actual)
	if err != nil {
		return err.Error()
	}

	// Check for outcome
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

// Generator of Namespaces_Topic_Spec instances for property testing - lazily instantiated by
// Namespaces_Topic_SpecGenerator()
var namespaces_Topic_SpecGenerator gopter.Gen

// Namespaces_Topic_SpecGenerator returns a generator of Namespaces_Topic_Spec instances for property testing.
func Namespaces_Topic_SpecGenerator() gopter.Gen {
	if namespaces_Topic_SpecGenerator != nil {
		return namespaces_Topic_SpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForNamespaces_Topic_Spec(generators)
	namespaces_Topic_SpecGenerator = gen.Struct(reflect.TypeOf(Namespaces_Topic_Spec{}), generators)

	return namespaces_Topic_SpecGenerator
}

// AddIndependentPropertyGeneratorsForNamespaces_Topic_Spec is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForNamespaces_Topic_Spec(gens map[string]gopter.Gen) {
	gens["AutoDeleteOnIdle"] = gen.PtrOf(gen.AlphaString())
	gens["AzureName"] = gen.AlphaString()
	gens["DefaultMessageTimeToLive"] = gen.PtrOf(gen.AlphaString())
	gens["DuplicateDetectionHistoryTimeWindow"] = gen.PtrOf(gen.AlphaString())
	gens["EnableBatchedOperations"] = gen.PtrOf(gen.Bool())
	gens["EnableExpress"] = gen.PtrOf(gen.Bool())
	gens["EnablePartitioning"] = gen.PtrOf(gen.Bool())
	gens["MaxMessageSizeInKilobytes"] = gen.PtrOf(gen.Int())
	gens["MaxSizeInMegabytes"] = gen.PtrOf(gen.Int())
	gens["RequiresDuplicateDetection"] = gen.PtrOf(gen.Bool())
	gens["SupportOrdering"] = gen.PtrOf(gen.Bool())
}

func Test_Namespaces_Topic_STATUS_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from Namespaces_Topic_STATUS to Namespaces_Topic_STATUS via AssignProperties_To_Namespaces_Topic_STATUS & AssignProperties_From_Namespaces_Topic_STATUS returns original",
		prop.ForAll(RunPropertyAssignmentTestForNamespaces_Topic_STATUS, Namespaces_Topic_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForNamespaces_Topic_STATUS tests if a specific instance of Namespaces_Topic_STATUS can be assigned to storage and back losslessly
func RunPropertyAssignmentTestForNamespaces_Topic_STATUS(subject Namespaces_Topic_STATUS) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other storage.Namespaces_Topic_STATUS
	err := copied.AssignProperties_To_Namespaces_Topic_STATUS(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual Namespaces_Topic_STATUS
	err = actual.AssignProperties_From_Namespaces_Topic_STATUS(&other)
	if err != nil {
		return err.Error()
	}

	// Check for a match
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

func Test_Namespaces_Topic_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Namespaces_Topic_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForNamespaces_Topic_STATUS, Namespaces_Topic_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForNamespaces_Topic_STATUS runs a test to see if a specific instance of Namespaces_Topic_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForNamespaces_Topic_STATUS(subject Namespaces_Topic_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Namespaces_Topic_STATUS
	err = json.Unmarshal(bin, &actual)
	if err != nil {
		return err.Error()
	}

	// Check for outcome
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

// Generator of Namespaces_Topic_STATUS instances for property testing - lazily instantiated by
// Namespaces_Topic_STATUSGenerator()
var namespaces_Topic_STATUSGenerator gopter.Gen

// Namespaces_Topic_STATUSGenerator returns a generator of Namespaces_Topic_STATUS instances for property testing.
// We first initialize namespaces_Topic_STATUSGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func Namespaces_Topic_STATUSGenerator() gopter.Gen {
	if namespaces_Topic_STATUSGenerator != nil {
		return namespaces_Topic_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForNamespaces_Topic_STATUS(generators)
	namespaces_Topic_STATUSGenerator = gen.Struct(reflect.TypeOf(Namespaces_Topic_STATUS{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForNamespaces_Topic_STATUS(generators)
	AddRelatedPropertyGeneratorsForNamespaces_Topic_STATUS(generators)
	namespaces_Topic_STATUSGenerator = gen.Struct(reflect.TypeOf(Namespaces_Topic_STATUS{}), generators)

	return namespaces_Topic_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForNamespaces_Topic_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForNamespaces_Topic_STATUS(gens map[string]gopter.Gen) {
	gens["AccessedAt"] = gen.PtrOf(gen.AlphaString())
	gens["AutoDeleteOnIdle"] = gen.PtrOf(gen.AlphaString())
	gens["CreatedAt"] = gen.PtrOf(gen.AlphaString())
	gens["DefaultMessageTimeToLive"] = gen.PtrOf(gen.AlphaString())
	gens["DuplicateDetectionHistoryTimeWindow"] = gen.PtrOf(gen.AlphaString())
	gens["EnableBatchedOperations"] = gen.PtrOf(gen.Bool())
	gens["EnableExpress"] = gen.PtrOf(gen.Bool())
	gens["EnablePartitioning"] = gen.PtrOf(gen.Bool())
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["MaxMessageSizeInKilobytes"] = gen.PtrOf(gen.Int())
	gens["MaxSizeInMegabytes"] = gen.PtrOf(gen.Int())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["RequiresDuplicateDetection"] = gen.PtrOf(gen.Bool())
	gens["SizeInBytes"] = gen.PtrOf(gen.Int())
	gens["Status"] = gen.PtrOf(gen.OneConstOf(
		EntityStatus_STATUS_Active,
		EntityStatus_STATUS_Creating,
		EntityStatus_STATUS_Deleting,
		EntityStatus_STATUS_Disabled,
		EntityStatus_STATUS_ReceiveDisabled,
		EntityStatus_STATUS_Renaming,
		EntityStatus_STATUS_Restoring,
		EntityStatus_STATUS_SendDisabled,
		EntityStatus_STATUS_Unknown))
	gens["SubscriptionCount"] = gen.PtrOf(gen.Int())
	gens["SupportOrdering"] = gen.PtrOf(gen.Bool())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
	gens["UpdatedAt"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForNamespaces_Topic_STATUS is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForNamespaces_Topic_STATUS(gens map[string]gopter.Gen) {
	gens["CountDetails"] = gen.PtrOf(MessageCountDetails_STATUSGenerator())
	gens["SystemData"] = gen.PtrOf(SystemData_STATUSGenerator())
}
