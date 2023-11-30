// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20210601

import (
	"encoding/json"
	v20210601s "github.com/Azure/azure-service-operator/v2/api/dbforpostgresql/v1api20210601/storage"
	v20221201s "github.com/Azure/azure-service-operator/v2/api/dbforpostgresql/v1api20221201/storage"
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

func Test_FlexibleServersFirewallRule_WhenConvertedToHub_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	parameters.MinSuccessfulTests = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from FlexibleServersFirewallRule to hub returns original",
		prop.ForAll(RunResourceConversionTestForFlexibleServersFirewallRule, FlexibleServersFirewallRuleGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunResourceConversionTestForFlexibleServersFirewallRule tests if a specific instance of FlexibleServersFirewallRule round trips to the hub storage version and back losslessly
func RunResourceConversionTestForFlexibleServersFirewallRule(subject FlexibleServersFirewallRule) string {
	// Copy subject to make sure conversion doesn't modify it
	copied := subject.DeepCopy()

	// Convert to our hub version
	var hub v20221201s.FlexibleServersFirewallRule
	err := copied.ConvertTo(&hub)
	if err != nil {
		return err.Error()
	}

	// Convert from our hub version
	var actual FlexibleServersFirewallRule
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

func Test_FlexibleServersFirewallRule_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from FlexibleServersFirewallRule to FlexibleServersFirewallRule via AssignProperties_To_FlexibleServersFirewallRule & AssignProperties_From_FlexibleServersFirewallRule returns original",
		prop.ForAll(RunPropertyAssignmentTestForFlexibleServersFirewallRule, FlexibleServersFirewallRuleGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForFlexibleServersFirewallRule tests if a specific instance of FlexibleServersFirewallRule can be assigned to storage and back losslessly
func RunPropertyAssignmentTestForFlexibleServersFirewallRule(subject FlexibleServersFirewallRule) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20210601s.FlexibleServersFirewallRule
	err := copied.AssignProperties_To_FlexibleServersFirewallRule(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual FlexibleServersFirewallRule
	err = actual.AssignProperties_From_FlexibleServersFirewallRule(&other)
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

func Test_FlexibleServersFirewallRule_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 20
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of FlexibleServersFirewallRule via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForFlexibleServersFirewallRule, FlexibleServersFirewallRuleGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForFlexibleServersFirewallRule runs a test to see if a specific instance of FlexibleServersFirewallRule round trips to JSON and back losslessly
func RunJSONSerializationTestForFlexibleServersFirewallRule(subject FlexibleServersFirewallRule) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual FlexibleServersFirewallRule
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

// Generator of FlexibleServersFirewallRule instances for property testing - lazily instantiated by
// FlexibleServersFirewallRuleGenerator()
var flexibleServersFirewallRuleGenerator gopter.Gen

// FlexibleServersFirewallRuleGenerator returns a generator of FlexibleServersFirewallRule instances for property testing.
func FlexibleServersFirewallRuleGenerator() gopter.Gen {
	if flexibleServersFirewallRuleGenerator != nil {
		return flexibleServersFirewallRuleGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForFlexibleServersFirewallRule(generators)
	flexibleServersFirewallRuleGenerator = gen.Struct(reflect.TypeOf(FlexibleServersFirewallRule{}), generators)

	return flexibleServersFirewallRuleGenerator
}

// AddRelatedPropertyGeneratorsForFlexibleServersFirewallRule is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForFlexibleServersFirewallRule(gens map[string]gopter.Gen) {
	gens["Spec"] = FlexibleServers_FirewallRule_SpecGenerator()
	gens["Status"] = FlexibleServers_FirewallRule_STATUSGenerator()
}

func Test_FlexibleServers_FirewallRule_Spec_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from FlexibleServers_FirewallRule_Spec to FlexibleServers_FirewallRule_Spec via AssignProperties_To_FlexibleServers_FirewallRule_Spec & AssignProperties_From_FlexibleServers_FirewallRule_Spec returns original",
		prop.ForAll(RunPropertyAssignmentTestForFlexibleServers_FirewallRule_Spec, FlexibleServers_FirewallRule_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForFlexibleServers_FirewallRule_Spec tests if a specific instance of FlexibleServers_FirewallRule_Spec can be assigned to storage and back losslessly
func RunPropertyAssignmentTestForFlexibleServers_FirewallRule_Spec(subject FlexibleServers_FirewallRule_Spec) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20210601s.FlexibleServers_FirewallRule_Spec
	err := copied.AssignProperties_To_FlexibleServers_FirewallRule_Spec(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual FlexibleServers_FirewallRule_Spec
	err = actual.AssignProperties_From_FlexibleServers_FirewallRule_Spec(&other)
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

func Test_FlexibleServers_FirewallRule_Spec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of FlexibleServers_FirewallRule_Spec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForFlexibleServers_FirewallRule_Spec, FlexibleServers_FirewallRule_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForFlexibleServers_FirewallRule_Spec runs a test to see if a specific instance of FlexibleServers_FirewallRule_Spec round trips to JSON and back losslessly
func RunJSONSerializationTestForFlexibleServers_FirewallRule_Spec(subject FlexibleServers_FirewallRule_Spec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual FlexibleServers_FirewallRule_Spec
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

// Generator of FlexibleServers_FirewallRule_Spec instances for property testing - lazily instantiated by
// FlexibleServers_FirewallRule_SpecGenerator()
var flexibleServers_FirewallRule_SpecGenerator gopter.Gen

// FlexibleServers_FirewallRule_SpecGenerator returns a generator of FlexibleServers_FirewallRule_Spec instances for property testing.
func FlexibleServers_FirewallRule_SpecGenerator() gopter.Gen {
	if flexibleServers_FirewallRule_SpecGenerator != nil {
		return flexibleServers_FirewallRule_SpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForFlexibleServers_FirewallRule_Spec(generators)
	flexibleServers_FirewallRule_SpecGenerator = gen.Struct(reflect.TypeOf(FlexibleServers_FirewallRule_Spec{}), generators)

	return flexibleServers_FirewallRule_SpecGenerator
}

// AddIndependentPropertyGeneratorsForFlexibleServers_FirewallRule_Spec is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForFlexibleServers_FirewallRule_Spec(gens map[string]gopter.Gen) {
	gens["AzureName"] = gen.AlphaString()
	gens["EndIpAddress"] = gen.PtrOf(gen.AlphaString())
	gens["StartIpAddress"] = gen.PtrOf(gen.AlphaString())
}

func Test_FlexibleServers_FirewallRule_STATUS_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from FlexibleServers_FirewallRule_STATUS to FlexibleServers_FirewallRule_STATUS via AssignProperties_To_FlexibleServers_FirewallRule_STATUS & AssignProperties_From_FlexibleServers_FirewallRule_STATUS returns original",
		prop.ForAll(RunPropertyAssignmentTestForFlexibleServers_FirewallRule_STATUS, FlexibleServers_FirewallRule_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForFlexibleServers_FirewallRule_STATUS tests if a specific instance of FlexibleServers_FirewallRule_STATUS can be assigned to storage and back losslessly
func RunPropertyAssignmentTestForFlexibleServers_FirewallRule_STATUS(subject FlexibleServers_FirewallRule_STATUS) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20210601s.FlexibleServers_FirewallRule_STATUS
	err := copied.AssignProperties_To_FlexibleServers_FirewallRule_STATUS(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual FlexibleServers_FirewallRule_STATUS
	err = actual.AssignProperties_From_FlexibleServers_FirewallRule_STATUS(&other)
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

func Test_FlexibleServers_FirewallRule_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of FlexibleServers_FirewallRule_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForFlexibleServers_FirewallRule_STATUS, FlexibleServers_FirewallRule_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForFlexibleServers_FirewallRule_STATUS runs a test to see if a specific instance of FlexibleServers_FirewallRule_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForFlexibleServers_FirewallRule_STATUS(subject FlexibleServers_FirewallRule_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual FlexibleServers_FirewallRule_STATUS
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

// Generator of FlexibleServers_FirewallRule_STATUS instances for property testing - lazily instantiated by
// FlexibleServers_FirewallRule_STATUSGenerator()
var flexibleServers_FirewallRule_STATUSGenerator gopter.Gen

// FlexibleServers_FirewallRule_STATUSGenerator returns a generator of FlexibleServers_FirewallRule_STATUS instances for property testing.
// We first initialize flexibleServers_FirewallRule_STATUSGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func FlexibleServers_FirewallRule_STATUSGenerator() gopter.Gen {
	if flexibleServers_FirewallRule_STATUSGenerator != nil {
		return flexibleServers_FirewallRule_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForFlexibleServers_FirewallRule_STATUS(generators)
	flexibleServers_FirewallRule_STATUSGenerator = gen.Struct(reflect.TypeOf(FlexibleServers_FirewallRule_STATUS{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForFlexibleServers_FirewallRule_STATUS(generators)
	AddRelatedPropertyGeneratorsForFlexibleServers_FirewallRule_STATUS(generators)
	flexibleServers_FirewallRule_STATUSGenerator = gen.Struct(reflect.TypeOf(FlexibleServers_FirewallRule_STATUS{}), generators)

	return flexibleServers_FirewallRule_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForFlexibleServers_FirewallRule_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForFlexibleServers_FirewallRule_STATUS(gens map[string]gopter.Gen) {
	gens["EndIpAddress"] = gen.PtrOf(gen.AlphaString())
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["StartIpAddress"] = gen.PtrOf(gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForFlexibleServers_FirewallRule_STATUS is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForFlexibleServers_FirewallRule_STATUS(gens map[string]gopter.Gen) {
	gens["SystemData"] = gen.PtrOf(SystemData_STATUSGenerator())
}
