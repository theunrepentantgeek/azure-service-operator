// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20231115

import (
	"encoding/json"
	storage "github.com/Azure/azure-service-operator/v2/api/documentdb/v1api20231115/storage"
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

func Test_SqlDatabaseThroughputSetting_WhenConvertedToHub_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	parameters.MinSuccessfulTests = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from SqlDatabaseThroughputSetting to hub returns original",
		prop.ForAll(RunResourceConversionTestForSqlDatabaseThroughputSetting, SqlDatabaseThroughputSettingGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunResourceConversionTestForSqlDatabaseThroughputSetting tests if a specific instance of SqlDatabaseThroughputSetting round trips to the hub storage version and back losslessly
func RunResourceConversionTestForSqlDatabaseThroughputSetting(subject SqlDatabaseThroughputSetting) string {
	// Copy subject to make sure conversion doesn't modify it
	copied := subject.DeepCopy()

	// Convert to our hub version
	var hub storage.SqlDatabaseThroughputSetting
	err := copied.ConvertTo(&hub)
	if err != nil {
		return err.Error()
	}

	// Convert from our hub version
	var actual SqlDatabaseThroughputSetting
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

func Test_SqlDatabaseThroughputSetting_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from SqlDatabaseThroughputSetting to SqlDatabaseThroughputSetting via AssignProperties_To_SqlDatabaseThroughputSetting & AssignProperties_From_SqlDatabaseThroughputSetting returns original",
		prop.ForAll(RunPropertyAssignmentTestForSqlDatabaseThroughputSetting, SqlDatabaseThroughputSettingGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForSqlDatabaseThroughputSetting tests if a specific instance of SqlDatabaseThroughputSetting can be assigned to storage and back losslessly
func RunPropertyAssignmentTestForSqlDatabaseThroughputSetting(subject SqlDatabaseThroughputSetting) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other storage.SqlDatabaseThroughputSetting
	err := copied.AssignProperties_To_SqlDatabaseThroughputSetting(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual SqlDatabaseThroughputSetting
	err = actual.AssignProperties_From_SqlDatabaseThroughputSetting(&other)
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

func Test_SqlDatabaseThroughputSetting_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 20
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SqlDatabaseThroughputSetting via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSqlDatabaseThroughputSetting, SqlDatabaseThroughputSettingGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSqlDatabaseThroughputSetting runs a test to see if a specific instance of SqlDatabaseThroughputSetting round trips to JSON and back losslessly
func RunJSONSerializationTestForSqlDatabaseThroughputSetting(subject SqlDatabaseThroughputSetting) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SqlDatabaseThroughputSetting
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

// Generator of SqlDatabaseThroughputSetting instances for property testing - lazily instantiated by
// SqlDatabaseThroughputSettingGenerator()
var sqlDatabaseThroughputSettingGenerator gopter.Gen

// SqlDatabaseThroughputSettingGenerator returns a generator of SqlDatabaseThroughputSetting instances for property testing.
func SqlDatabaseThroughputSettingGenerator() gopter.Gen {
	if sqlDatabaseThroughputSettingGenerator != nil {
		return sqlDatabaseThroughputSettingGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForSqlDatabaseThroughputSetting(generators)
	sqlDatabaseThroughputSettingGenerator = gen.Struct(reflect.TypeOf(SqlDatabaseThroughputSetting{}), generators)

	return sqlDatabaseThroughputSettingGenerator
}

// AddRelatedPropertyGeneratorsForSqlDatabaseThroughputSetting is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForSqlDatabaseThroughputSetting(gens map[string]gopter.Gen) {
	gens["Spec"] = DatabaseAccounts_SqlDatabases_ThroughputSetting_SpecGenerator()
	gens["Status"] = DatabaseAccounts_SqlDatabases_ThroughputSetting_STATUSGenerator()
}

func Test_DatabaseAccounts_SqlDatabases_ThroughputSetting_Spec_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from DatabaseAccounts_SqlDatabases_ThroughputSetting_Spec to DatabaseAccounts_SqlDatabases_ThroughputSetting_Spec via AssignProperties_To_DatabaseAccounts_SqlDatabases_ThroughputSetting_Spec & AssignProperties_From_DatabaseAccounts_SqlDatabases_ThroughputSetting_Spec returns original",
		prop.ForAll(RunPropertyAssignmentTestForDatabaseAccounts_SqlDatabases_ThroughputSetting_Spec, DatabaseAccounts_SqlDatabases_ThroughputSetting_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForDatabaseAccounts_SqlDatabases_ThroughputSetting_Spec tests if a specific instance of DatabaseAccounts_SqlDatabases_ThroughputSetting_Spec can be assigned to storage and back losslessly
func RunPropertyAssignmentTestForDatabaseAccounts_SqlDatabases_ThroughputSetting_Spec(subject DatabaseAccounts_SqlDatabases_ThroughputSetting_Spec) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other storage.DatabaseAccounts_SqlDatabases_ThroughputSetting_Spec
	err := copied.AssignProperties_To_DatabaseAccounts_SqlDatabases_ThroughputSetting_Spec(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual DatabaseAccounts_SqlDatabases_ThroughputSetting_Spec
	err = actual.AssignProperties_From_DatabaseAccounts_SqlDatabases_ThroughputSetting_Spec(&other)
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

func Test_DatabaseAccounts_SqlDatabases_ThroughputSetting_Spec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of DatabaseAccounts_SqlDatabases_ThroughputSetting_Spec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForDatabaseAccounts_SqlDatabases_ThroughputSetting_Spec, DatabaseAccounts_SqlDatabases_ThroughputSetting_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForDatabaseAccounts_SqlDatabases_ThroughputSetting_Spec runs a test to see if a specific instance of DatabaseAccounts_SqlDatabases_ThroughputSetting_Spec round trips to JSON and back losslessly
func RunJSONSerializationTestForDatabaseAccounts_SqlDatabases_ThroughputSetting_Spec(subject DatabaseAccounts_SqlDatabases_ThroughputSetting_Spec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual DatabaseAccounts_SqlDatabases_ThroughputSetting_Spec
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

// Generator of DatabaseAccounts_SqlDatabases_ThroughputSetting_Spec instances for property testing - lazily
// instantiated by DatabaseAccounts_SqlDatabases_ThroughputSetting_SpecGenerator()
var databaseAccounts_SqlDatabases_ThroughputSetting_SpecGenerator gopter.Gen

// DatabaseAccounts_SqlDatabases_ThroughputSetting_SpecGenerator returns a generator of DatabaseAccounts_SqlDatabases_ThroughputSetting_Spec instances for property testing.
// We first initialize databaseAccounts_SqlDatabases_ThroughputSetting_SpecGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func DatabaseAccounts_SqlDatabases_ThroughputSetting_SpecGenerator() gopter.Gen {
	if databaseAccounts_SqlDatabases_ThroughputSetting_SpecGenerator != nil {
		return databaseAccounts_SqlDatabases_ThroughputSetting_SpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDatabaseAccounts_SqlDatabases_ThroughputSetting_Spec(generators)
	databaseAccounts_SqlDatabases_ThroughputSetting_SpecGenerator = gen.Struct(reflect.TypeOf(DatabaseAccounts_SqlDatabases_ThroughputSetting_Spec{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDatabaseAccounts_SqlDatabases_ThroughputSetting_Spec(generators)
	AddRelatedPropertyGeneratorsForDatabaseAccounts_SqlDatabases_ThroughputSetting_Spec(generators)
	databaseAccounts_SqlDatabases_ThroughputSetting_SpecGenerator = gen.Struct(reflect.TypeOf(DatabaseAccounts_SqlDatabases_ThroughputSetting_Spec{}), generators)

	return databaseAccounts_SqlDatabases_ThroughputSetting_SpecGenerator
}

// AddIndependentPropertyGeneratorsForDatabaseAccounts_SqlDatabases_ThroughputSetting_Spec is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForDatabaseAccounts_SqlDatabases_ThroughputSetting_Spec(gens map[string]gopter.Gen) {
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Tags"] = gen.MapOf(
		gen.AlphaString(),
		gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForDatabaseAccounts_SqlDatabases_ThroughputSetting_Spec is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForDatabaseAccounts_SqlDatabases_ThroughputSetting_Spec(gens map[string]gopter.Gen) {
	gens["Resource"] = gen.PtrOf(ThroughputSettingsResourceGenerator())
}

func Test_DatabaseAccounts_SqlDatabases_ThroughputSetting_STATUS_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from DatabaseAccounts_SqlDatabases_ThroughputSetting_STATUS to DatabaseAccounts_SqlDatabases_ThroughputSetting_STATUS via AssignProperties_To_DatabaseAccounts_SqlDatabases_ThroughputSetting_STATUS & AssignProperties_From_DatabaseAccounts_SqlDatabases_ThroughputSetting_STATUS returns original",
		prop.ForAll(RunPropertyAssignmentTestForDatabaseAccounts_SqlDatabases_ThroughputSetting_STATUS, DatabaseAccounts_SqlDatabases_ThroughputSetting_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForDatabaseAccounts_SqlDatabases_ThroughputSetting_STATUS tests if a specific instance of DatabaseAccounts_SqlDatabases_ThroughputSetting_STATUS can be assigned to storage and back losslessly
func RunPropertyAssignmentTestForDatabaseAccounts_SqlDatabases_ThroughputSetting_STATUS(subject DatabaseAccounts_SqlDatabases_ThroughputSetting_STATUS) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other storage.DatabaseAccounts_SqlDatabases_ThroughputSetting_STATUS
	err := copied.AssignProperties_To_DatabaseAccounts_SqlDatabases_ThroughputSetting_STATUS(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual DatabaseAccounts_SqlDatabases_ThroughputSetting_STATUS
	err = actual.AssignProperties_From_DatabaseAccounts_SqlDatabases_ThroughputSetting_STATUS(&other)
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

func Test_DatabaseAccounts_SqlDatabases_ThroughputSetting_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of DatabaseAccounts_SqlDatabases_ThroughputSetting_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForDatabaseAccounts_SqlDatabases_ThroughputSetting_STATUS, DatabaseAccounts_SqlDatabases_ThroughputSetting_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForDatabaseAccounts_SqlDatabases_ThroughputSetting_STATUS runs a test to see if a specific instance of DatabaseAccounts_SqlDatabases_ThroughputSetting_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForDatabaseAccounts_SqlDatabases_ThroughputSetting_STATUS(subject DatabaseAccounts_SqlDatabases_ThroughputSetting_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual DatabaseAccounts_SqlDatabases_ThroughputSetting_STATUS
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

// Generator of DatabaseAccounts_SqlDatabases_ThroughputSetting_STATUS instances for property testing - lazily
// instantiated by DatabaseAccounts_SqlDatabases_ThroughputSetting_STATUSGenerator()
var databaseAccounts_SqlDatabases_ThroughputSetting_STATUSGenerator gopter.Gen

// DatabaseAccounts_SqlDatabases_ThroughputSetting_STATUSGenerator returns a generator of DatabaseAccounts_SqlDatabases_ThroughputSetting_STATUS instances for property testing.
// We first initialize databaseAccounts_SqlDatabases_ThroughputSetting_STATUSGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func DatabaseAccounts_SqlDatabases_ThroughputSetting_STATUSGenerator() gopter.Gen {
	if databaseAccounts_SqlDatabases_ThroughputSetting_STATUSGenerator != nil {
		return databaseAccounts_SqlDatabases_ThroughputSetting_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDatabaseAccounts_SqlDatabases_ThroughputSetting_STATUS(generators)
	databaseAccounts_SqlDatabases_ThroughputSetting_STATUSGenerator = gen.Struct(reflect.TypeOf(DatabaseAccounts_SqlDatabases_ThroughputSetting_STATUS{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDatabaseAccounts_SqlDatabases_ThroughputSetting_STATUS(generators)
	AddRelatedPropertyGeneratorsForDatabaseAccounts_SqlDatabases_ThroughputSetting_STATUS(generators)
	databaseAccounts_SqlDatabases_ThroughputSetting_STATUSGenerator = gen.Struct(reflect.TypeOf(DatabaseAccounts_SqlDatabases_ThroughputSetting_STATUS{}), generators)

	return databaseAccounts_SqlDatabases_ThroughputSetting_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForDatabaseAccounts_SqlDatabases_ThroughputSetting_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForDatabaseAccounts_SqlDatabases_ThroughputSetting_STATUS(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Tags"] = gen.MapOf(
		gen.AlphaString(),
		gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForDatabaseAccounts_SqlDatabases_ThroughputSetting_STATUS is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForDatabaseAccounts_SqlDatabases_ThroughputSetting_STATUS(gens map[string]gopter.Gen) {
	gens["Resource"] = gen.PtrOf(ThroughputSettingsGetProperties_Resource_STATUSGenerator())
}
