// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20230630

import (
	"encoding/json"
	storage "github.com/Azure/azure-service-operator/v2/api/dbformysql/v1api20230630/storage"
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

func Test_FlexibleServersAdministrator_WhenConvertedToHub_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	parameters.MinSuccessfulTests = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from FlexibleServersAdministrator to hub returns original",
		prop.ForAll(RunResourceConversionTestForFlexibleServersAdministrator, FlexibleServersAdministratorGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunResourceConversionTestForFlexibleServersAdministrator tests if a specific instance of FlexibleServersAdministrator round trips to the hub storage version and back losslessly
func RunResourceConversionTestForFlexibleServersAdministrator(subject FlexibleServersAdministrator) string {
	// Copy subject to make sure conversion doesn't modify it
	copied := subject.DeepCopy()

	// Convert to our hub version
	var hub storage.FlexibleServersAdministrator
	err := copied.ConvertTo(&hub)
	if err != nil {
		return err.Error()
	}

	// Convert from our hub version
	var actual FlexibleServersAdministrator
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

func Test_FlexibleServersAdministrator_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from FlexibleServersAdministrator to FlexibleServersAdministrator via AssignProperties_To_FlexibleServersAdministrator & AssignProperties_From_FlexibleServersAdministrator returns original",
		prop.ForAll(RunPropertyAssignmentTestForFlexibleServersAdministrator, FlexibleServersAdministratorGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForFlexibleServersAdministrator tests if a specific instance of FlexibleServersAdministrator can be assigned to storage and back losslessly
func RunPropertyAssignmentTestForFlexibleServersAdministrator(subject FlexibleServersAdministrator) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other storage.FlexibleServersAdministrator
	err := copied.AssignProperties_To_FlexibleServersAdministrator(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual FlexibleServersAdministrator
	err = actual.AssignProperties_From_FlexibleServersAdministrator(&other)
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

func Test_FlexibleServersAdministrator_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 20
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of FlexibleServersAdministrator via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForFlexibleServersAdministrator, FlexibleServersAdministratorGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForFlexibleServersAdministrator runs a test to see if a specific instance of FlexibleServersAdministrator round trips to JSON and back losslessly
func RunJSONSerializationTestForFlexibleServersAdministrator(subject FlexibleServersAdministrator) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual FlexibleServersAdministrator
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

// Generator of FlexibleServersAdministrator instances for property testing - lazily instantiated by
// FlexibleServersAdministratorGenerator()
var flexibleServersAdministratorGenerator gopter.Gen

// FlexibleServersAdministratorGenerator returns a generator of FlexibleServersAdministrator instances for property testing.
func FlexibleServersAdministratorGenerator() gopter.Gen {
	if flexibleServersAdministratorGenerator != nil {
		return flexibleServersAdministratorGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForFlexibleServersAdministrator(generators)
	flexibleServersAdministratorGenerator = gen.Struct(reflect.TypeOf(FlexibleServersAdministrator{}), generators)

	return flexibleServersAdministratorGenerator
}

// AddRelatedPropertyGeneratorsForFlexibleServersAdministrator is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForFlexibleServersAdministrator(gens map[string]gopter.Gen) {
	gens["Spec"] = FlexibleServers_Administrator_SpecGenerator()
	gens["Status"] = FlexibleServers_Administrator_STATUSGenerator()
}

func Test_FlexibleServers_Administrator_Spec_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from FlexibleServers_Administrator_Spec to FlexibleServers_Administrator_Spec via AssignProperties_To_FlexibleServers_Administrator_Spec & AssignProperties_From_FlexibleServers_Administrator_Spec returns original",
		prop.ForAll(RunPropertyAssignmentTestForFlexibleServers_Administrator_Spec, FlexibleServers_Administrator_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForFlexibleServers_Administrator_Spec tests if a specific instance of FlexibleServers_Administrator_Spec can be assigned to storage and back losslessly
func RunPropertyAssignmentTestForFlexibleServers_Administrator_Spec(subject FlexibleServers_Administrator_Spec) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other storage.FlexibleServers_Administrator_Spec
	err := copied.AssignProperties_To_FlexibleServers_Administrator_Spec(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual FlexibleServers_Administrator_Spec
	err = actual.AssignProperties_From_FlexibleServers_Administrator_Spec(&other)
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

func Test_FlexibleServers_Administrator_Spec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of FlexibleServers_Administrator_Spec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForFlexibleServers_Administrator_Spec, FlexibleServers_Administrator_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForFlexibleServers_Administrator_Spec runs a test to see if a specific instance of FlexibleServers_Administrator_Spec round trips to JSON and back losslessly
func RunJSONSerializationTestForFlexibleServers_Administrator_Spec(subject FlexibleServers_Administrator_Spec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual FlexibleServers_Administrator_Spec
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

// Generator of FlexibleServers_Administrator_Spec instances for property testing - lazily instantiated by
// FlexibleServers_Administrator_SpecGenerator()
var flexibleServers_Administrator_SpecGenerator gopter.Gen

// FlexibleServers_Administrator_SpecGenerator returns a generator of FlexibleServers_Administrator_Spec instances for property testing.
func FlexibleServers_Administrator_SpecGenerator() gopter.Gen {
	if flexibleServers_Administrator_SpecGenerator != nil {
		return flexibleServers_Administrator_SpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForFlexibleServers_Administrator_Spec(generators)
	flexibleServers_Administrator_SpecGenerator = gen.Struct(reflect.TypeOf(FlexibleServers_Administrator_Spec{}), generators)

	return flexibleServers_Administrator_SpecGenerator
}

// AddIndependentPropertyGeneratorsForFlexibleServers_Administrator_Spec is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForFlexibleServers_Administrator_Spec(gens map[string]gopter.Gen) {
	gens["AdministratorType"] = gen.PtrOf(gen.OneConstOf(AdministratorProperties_AdministratorType_ActiveDirectory))
	gens["Login"] = gen.PtrOf(gen.AlphaString())
	gens["Sid"] = gen.PtrOf(gen.AlphaString())
	gens["TenantId"] = gen.PtrOf(gen.AlphaString())
}

func Test_FlexibleServers_Administrator_STATUS_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from FlexibleServers_Administrator_STATUS to FlexibleServers_Administrator_STATUS via AssignProperties_To_FlexibleServers_Administrator_STATUS & AssignProperties_From_FlexibleServers_Administrator_STATUS returns original",
		prop.ForAll(RunPropertyAssignmentTestForFlexibleServers_Administrator_STATUS, FlexibleServers_Administrator_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForFlexibleServers_Administrator_STATUS tests if a specific instance of FlexibleServers_Administrator_STATUS can be assigned to storage and back losslessly
func RunPropertyAssignmentTestForFlexibleServers_Administrator_STATUS(subject FlexibleServers_Administrator_STATUS) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other storage.FlexibleServers_Administrator_STATUS
	err := copied.AssignProperties_To_FlexibleServers_Administrator_STATUS(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual FlexibleServers_Administrator_STATUS
	err = actual.AssignProperties_From_FlexibleServers_Administrator_STATUS(&other)
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

func Test_FlexibleServers_Administrator_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of FlexibleServers_Administrator_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForFlexibleServers_Administrator_STATUS, FlexibleServers_Administrator_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForFlexibleServers_Administrator_STATUS runs a test to see if a specific instance of FlexibleServers_Administrator_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForFlexibleServers_Administrator_STATUS(subject FlexibleServers_Administrator_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual FlexibleServers_Administrator_STATUS
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

// Generator of FlexibleServers_Administrator_STATUS instances for property testing - lazily instantiated by
// FlexibleServers_Administrator_STATUSGenerator()
var flexibleServers_Administrator_STATUSGenerator gopter.Gen

// FlexibleServers_Administrator_STATUSGenerator returns a generator of FlexibleServers_Administrator_STATUS instances for property testing.
// We first initialize flexibleServers_Administrator_STATUSGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func FlexibleServers_Administrator_STATUSGenerator() gopter.Gen {
	if flexibleServers_Administrator_STATUSGenerator != nil {
		return flexibleServers_Administrator_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForFlexibleServers_Administrator_STATUS(generators)
	flexibleServers_Administrator_STATUSGenerator = gen.Struct(reflect.TypeOf(FlexibleServers_Administrator_STATUS{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForFlexibleServers_Administrator_STATUS(generators)
	AddRelatedPropertyGeneratorsForFlexibleServers_Administrator_STATUS(generators)
	flexibleServers_Administrator_STATUSGenerator = gen.Struct(reflect.TypeOf(FlexibleServers_Administrator_STATUS{}), generators)

	return flexibleServers_Administrator_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForFlexibleServers_Administrator_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForFlexibleServers_Administrator_STATUS(gens map[string]gopter.Gen) {
	gens["AdministratorType"] = gen.PtrOf(gen.OneConstOf(AdministratorProperties_AdministratorType_STATUS_ActiveDirectory))
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["IdentityResourceId"] = gen.PtrOf(gen.AlphaString())
	gens["Login"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Sid"] = gen.PtrOf(gen.AlphaString())
	gens["TenantId"] = gen.PtrOf(gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForFlexibleServers_Administrator_STATUS is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForFlexibleServers_Administrator_STATUS(gens map[string]gopter.Gen) {
	gens["SystemData"] = gen.PtrOf(SystemData_STATUSGenerator())
}
