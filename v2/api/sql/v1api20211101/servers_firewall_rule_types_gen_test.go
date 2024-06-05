// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20211101

import (
	"encoding/json"
	storage "github.com/Azure/azure-service-operator/v2/api/sql/v1api20211101/storage"
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

func Test_ServersFirewallRule_WhenConvertedToHub_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	parameters.MinSuccessfulTests = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from ServersFirewallRule to hub returns original",
		prop.ForAll(RunResourceConversionTestForServersFirewallRule, ServersFirewallRuleGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunResourceConversionTestForServersFirewallRule tests if a specific instance of ServersFirewallRule round trips to the hub storage version and back losslessly
func RunResourceConversionTestForServersFirewallRule(subject ServersFirewallRule) string {
	// Copy subject to make sure conversion doesn't modify it
	copied := subject.DeepCopy()

	// Convert to our hub version
	var hub storage.ServersFirewallRule
	err := copied.ConvertTo(&hub)
	if err != nil {
		return err.Error()
	}

	// Convert from our hub version
	var actual ServersFirewallRule
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

func Test_ServersFirewallRule_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from ServersFirewallRule to ServersFirewallRule via AssignProperties_To_ServersFirewallRule & AssignProperties_From_ServersFirewallRule returns original",
		prop.ForAll(RunPropertyAssignmentTestForServersFirewallRule, ServersFirewallRuleGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForServersFirewallRule tests if a specific instance of ServersFirewallRule can be assigned to storage and back losslessly
func RunPropertyAssignmentTestForServersFirewallRule(subject ServersFirewallRule) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other storage.ServersFirewallRule
	err := copied.AssignProperties_To_ServersFirewallRule(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual ServersFirewallRule
	err = actual.AssignProperties_From_ServersFirewallRule(&other)
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

func Test_ServersFirewallRule_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 20
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ServersFirewallRule via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForServersFirewallRule, ServersFirewallRuleGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForServersFirewallRule runs a test to see if a specific instance of ServersFirewallRule round trips to JSON and back losslessly
func RunJSONSerializationTestForServersFirewallRule(subject ServersFirewallRule) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ServersFirewallRule
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

// Generator of ServersFirewallRule instances for property testing - lazily instantiated by
// ServersFirewallRuleGenerator()
var serversFirewallRuleGenerator gopter.Gen

// ServersFirewallRuleGenerator returns a generator of ServersFirewallRule instances for property testing.
func ServersFirewallRuleGenerator() gopter.Gen {
	if serversFirewallRuleGenerator != nil {
		return serversFirewallRuleGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForServersFirewallRule(generators)
	serversFirewallRuleGenerator = gen.Struct(reflect.TypeOf(ServersFirewallRule{}), generators)

	return serversFirewallRuleGenerator
}

// AddRelatedPropertyGeneratorsForServersFirewallRule is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForServersFirewallRule(gens map[string]gopter.Gen) {
	gens["Spec"] = Servers_FirewallRule_SpecGenerator()
	gens["Status"] = Servers_FirewallRule_STATUSGenerator()
}

func Test_Servers_FirewallRule_Spec_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from Servers_FirewallRule_Spec to Servers_FirewallRule_Spec via AssignProperties_To_Servers_FirewallRule_Spec & AssignProperties_From_Servers_FirewallRule_Spec returns original",
		prop.ForAll(RunPropertyAssignmentTestForServers_FirewallRule_Spec, Servers_FirewallRule_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForServers_FirewallRule_Spec tests if a specific instance of Servers_FirewallRule_Spec can be assigned to storage and back losslessly
func RunPropertyAssignmentTestForServers_FirewallRule_Spec(subject Servers_FirewallRule_Spec) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other storage.Servers_FirewallRule_Spec
	err := copied.AssignProperties_To_Servers_FirewallRule_Spec(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual Servers_FirewallRule_Spec
	err = actual.AssignProperties_From_Servers_FirewallRule_Spec(&other)
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

func Test_Servers_FirewallRule_Spec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Servers_FirewallRule_Spec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForServers_FirewallRule_Spec, Servers_FirewallRule_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForServers_FirewallRule_Spec runs a test to see if a specific instance of Servers_FirewallRule_Spec round trips to JSON and back losslessly
func RunJSONSerializationTestForServers_FirewallRule_Spec(subject Servers_FirewallRule_Spec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Servers_FirewallRule_Spec
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

// Generator of Servers_FirewallRule_Spec instances for property testing - lazily instantiated by
// Servers_FirewallRule_SpecGenerator()
var servers_FirewallRule_SpecGenerator gopter.Gen

// Servers_FirewallRule_SpecGenerator returns a generator of Servers_FirewallRule_Spec instances for property testing.
func Servers_FirewallRule_SpecGenerator() gopter.Gen {
	if servers_FirewallRule_SpecGenerator != nil {
		return servers_FirewallRule_SpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForServers_FirewallRule_Spec(generators)
	servers_FirewallRule_SpecGenerator = gen.Struct(reflect.TypeOf(Servers_FirewallRule_Spec{}), generators)

	return servers_FirewallRule_SpecGenerator
}

// AddIndependentPropertyGeneratorsForServers_FirewallRule_Spec is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForServers_FirewallRule_Spec(gens map[string]gopter.Gen) {
	gens["AzureName"] = gen.AlphaString()
	gens["EndIpAddress"] = gen.PtrOf(gen.AlphaString())
	gens["StartIpAddress"] = gen.PtrOf(gen.AlphaString())
}

func Test_Servers_FirewallRule_STATUS_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from Servers_FirewallRule_STATUS to Servers_FirewallRule_STATUS via AssignProperties_To_Servers_FirewallRule_STATUS & AssignProperties_From_Servers_FirewallRule_STATUS returns original",
		prop.ForAll(RunPropertyAssignmentTestForServers_FirewallRule_STATUS, Servers_FirewallRule_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForServers_FirewallRule_STATUS tests if a specific instance of Servers_FirewallRule_STATUS can be assigned to storage and back losslessly
func RunPropertyAssignmentTestForServers_FirewallRule_STATUS(subject Servers_FirewallRule_STATUS) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other storage.Servers_FirewallRule_STATUS
	err := copied.AssignProperties_To_Servers_FirewallRule_STATUS(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual Servers_FirewallRule_STATUS
	err = actual.AssignProperties_From_Servers_FirewallRule_STATUS(&other)
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

func Test_Servers_FirewallRule_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Servers_FirewallRule_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForServers_FirewallRule_STATUS, Servers_FirewallRule_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForServers_FirewallRule_STATUS runs a test to see if a specific instance of Servers_FirewallRule_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForServers_FirewallRule_STATUS(subject Servers_FirewallRule_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Servers_FirewallRule_STATUS
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

// Generator of Servers_FirewallRule_STATUS instances for property testing - lazily instantiated by
// Servers_FirewallRule_STATUSGenerator()
var servers_FirewallRule_STATUSGenerator gopter.Gen

// Servers_FirewallRule_STATUSGenerator returns a generator of Servers_FirewallRule_STATUS instances for property testing.
func Servers_FirewallRule_STATUSGenerator() gopter.Gen {
	if servers_FirewallRule_STATUSGenerator != nil {
		return servers_FirewallRule_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForServers_FirewallRule_STATUS(generators)
	servers_FirewallRule_STATUSGenerator = gen.Struct(reflect.TypeOf(Servers_FirewallRule_STATUS{}), generators)

	return servers_FirewallRule_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForServers_FirewallRule_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForServers_FirewallRule_STATUS(gens map[string]gopter.Gen) {
	gens["EndIpAddress"] = gen.PtrOf(gen.AlphaString())
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["StartIpAddress"] = gen.PtrOf(gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}
