// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20180601

import (
	"encoding/json"
	storage "github.com/Azure/azure-service-operator/v2/api/dbformariadb/v1api20180601/storage"
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

func Test_Database_WhenConvertedToHub_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	parameters.MinSuccessfulTests = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from Database to hub returns original",
		prop.ForAll(RunResourceConversionTestForDatabase, DatabaseGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunResourceConversionTestForDatabase tests if a specific instance of Database round trips to the hub storage version and back losslessly
func RunResourceConversionTestForDatabase(subject Database) string {
	// Copy subject to make sure conversion doesn't modify it
	copied := subject.DeepCopy()

	// Convert to our hub version
	var hub storage.Database
	err := copied.ConvertTo(&hub)
	if err != nil {
		return err.Error()
	}

	// Convert from our hub version
	var actual Database
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

func Test_Database_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from Database to Database via AssignProperties_To_Database & AssignProperties_From_Database returns original",
		prop.ForAll(RunPropertyAssignmentTestForDatabase, DatabaseGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForDatabase tests if a specific instance of Database can be assigned to storage and back losslessly
func RunPropertyAssignmentTestForDatabase(subject Database) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other storage.Database
	err := copied.AssignProperties_To_Database(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual Database
	err = actual.AssignProperties_From_Database(&other)
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

func Test_Database_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 20
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Database via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForDatabase, DatabaseGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForDatabase runs a test to see if a specific instance of Database round trips to JSON and back losslessly
func RunJSONSerializationTestForDatabase(subject Database) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Database
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

// Generator of Database instances for property testing - lazily instantiated by DatabaseGenerator()
var databaseGenerator gopter.Gen

// DatabaseGenerator returns a generator of Database instances for property testing.
func DatabaseGenerator() gopter.Gen {
	if databaseGenerator != nil {
		return databaseGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForDatabase(generators)
	databaseGenerator = gen.Struct(reflect.TypeOf(Database{}), generators)

	return databaseGenerator
}

// AddRelatedPropertyGeneratorsForDatabase is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForDatabase(gens map[string]gopter.Gen) {
	gens["Spec"] = Servers_Database_SpecGenerator()
	gens["Status"] = Servers_Database_STATUSGenerator()
}

func Test_Servers_Database_Spec_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from Servers_Database_Spec to Servers_Database_Spec via AssignProperties_To_Servers_Database_Spec & AssignProperties_From_Servers_Database_Spec returns original",
		prop.ForAll(RunPropertyAssignmentTestForServers_Database_Spec, Servers_Database_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForServers_Database_Spec tests if a specific instance of Servers_Database_Spec can be assigned to storage and back losslessly
func RunPropertyAssignmentTestForServers_Database_Spec(subject Servers_Database_Spec) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other storage.Servers_Database_Spec
	err := copied.AssignProperties_To_Servers_Database_Spec(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual Servers_Database_Spec
	err = actual.AssignProperties_From_Servers_Database_Spec(&other)
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

func Test_Servers_Database_Spec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Servers_Database_Spec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForServers_Database_Spec, Servers_Database_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForServers_Database_Spec runs a test to see if a specific instance of Servers_Database_Spec round trips to JSON and back losslessly
func RunJSONSerializationTestForServers_Database_Spec(subject Servers_Database_Spec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Servers_Database_Spec
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

// Generator of Servers_Database_Spec instances for property testing - lazily instantiated by
// Servers_Database_SpecGenerator()
var servers_Database_SpecGenerator gopter.Gen

// Servers_Database_SpecGenerator returns a generator of Servers_Database_Spec instances for property testing.
func Servers_Database_SpecGenerator() gopter.Gen {
	if servers_Database_SpecGenerator != nil {
		return servers_Database_SpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForServers_Database_Spec(generators)
	servers_Database_SpecGenerator = gen.Struct(reflect.TypeOf(Servers_Database_Spec{}), generators)

	return servers_Database_SpecGenerator
}

// AddIndependentPropertyGeneratorsForServers_Database_Spec is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForServers_Database_Spec(gens map[string]gopter.Gen) {
	gens["AzureName"] = gen.AlphaString()
	gens["Charset"] = gen.PtrOf(gen.AlphaString())
	gens["Collation"] = gen.PtrOf(gen.AlphaString())
}

func Test_Servers_Database_STATUS_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from Servers_Database_STATUS to Servers_Database_STATUS via AssignProperties_To_Servers_Database_STATUS & AssignProperties_From_Servers_Database_STATUS returns original",
		prop.ForAll(RunPropertyAssignmentTestForServers_Database_STATUS, Servers_Database_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForServers_Database_STATUS tests if a specific instance of Servers_Database_STATUS can be assigned to storage and back losslessly
func RunPropertyAssignmentTestForServers_Database_STATUS(subject Servers_Database_STATUS) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other storage.Servers_Database_STATUS
	err := copied.AssignProperties_To_Servers_Database_STATUS(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual Servers_Database_STATUS
	err = actual.AssignProperties_From_Servers_Database_STATUS(&other)
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

func Test_Servers_Database_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Servers_Database_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForServers_Database_STATUS, Servers_Database_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForServers_Database_STATUS runs a test to see if a specific instance of Servers_Database_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForServers_Database_STATUS(subject Servers_Database_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Servers_Database_STATUS
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

// Generator of Servers_Database_STATUS instances for property testing - lazily instantiated by
// Servers_Database_STATUSGenerator()
var servers_Database_STATUSGenerator gopter.Gen

// Servers_Database_STATUSGenerator returns a generator of Servers_Database_STATUS instances for property testing.
func Servers_Database_STATUSGenerator() gopter.Gen {
	if servers_Database_STATUSGenerator != nil {
		return servers_Database_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForServers_Database_STATUS(generators)
	servers_Database_STATUSGenerator = gen.Struct(reflect.TypeOf(Servers_Database_STATUS{}), generators)

	return servers_Database_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForServers_Database_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForServers_Database_STATUS(gens map[string]gopter.Gen) {
	gens["Charset"] = gen.PtrOf(gen.AlphaString())
	gens["Collation"] = gen.PtrOf(gen.AlphaString())
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}
