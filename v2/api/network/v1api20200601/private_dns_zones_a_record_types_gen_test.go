// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20200601

import (
	"encoding/json"
	storage "github.com/Azure/azure-service-operator/v2/api/network/v1api20200601/storage"
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

func Test_PrivateDnsZonesARecord_WhenConvertedToHub_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	parameters.MinSuccessfulTests = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from PrivateDnsZonesARecord to hub returns original",
		prop.ForAll(RunResourceConversionTestForPrivateDnsZonesARecord, PrivateDnsZonesARecordGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunResourceConversionTestForPrivateDnsZonesARecord tests if a specific instance of PrivateDnsZonesARecord round trips to the hub storage version and back losslessly
func RunResourceConversionTestForPrivateDnsZonesARecord(subject PrivateDnsZonesARecord) string {
	// Copy subject to make sure conversion doesn't modify it
	copied := subject.DeepCopy()

	// Convert to our hub version
	var hub storage.PrivateDnsZonesARecord
	err := copied.ConvertTo(&hub)
	if err != nil {
		return err.Error()
	}

	// Convert from our hub version
	var actual PrivateDnsZonesARecord
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

func Test_PrivateDnsZonesARecord_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from PrivateDnsZonesARecord to PrivateDnsZonesARecord via AssignProperties_To_PrivateDnsZonesARecord & AssignProperties_From_PrivateDnsZonesARecord returns original",
		prop.ForAll(RunPropertyAssignmentTestForPrivateDnsZonesARecord, PrivateDnsZonesARecordGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForPrivateDnsZonesARecord tests if a specific instance of PrivateDnsZonesARecord can be assigned to storage and back losslessly
func RunPropertyAssignmentTestForPrivateDnsZonesARecord(subject PrivateDnsZonesARecord) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other storage.PrivateDnsZonesARecord
	err := copied.AssignProperties_To_PrivateDnsZonesARecord(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual PrivateDnsZonesARecord
	err = actual.AssignProperties_From_PrivateDnsZonesARecord(&other)
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

func Test_PrivateDnsZonesARecord_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 20
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of PrivateDnsZonesARecord via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForPrivateDnsZonesARecord, PrivateDnsZonesARecordGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForPrivateDnsZonesARecord runs a test to see if a specific instance of PrivateDnsZonesARecord round trips to JSON and back losslessly
func RunJSONSerializationTestForPrivateDnsZonesARecord(subject PrivateDnsZonesARecord) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual PrivateDnsZonesARecord
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

// Generator of PrivateDnsZonesARecord instances for property testing - lazily instantiated by
// PrivateDnsZonesARecordGenerator()
var privateDnsZonesARecordGenerator gopter.Gen

// PrivateDnsZonesARecordGenerator returns a generator of PrivateDnsZonesARecord instances for property testing.
func PrivateDnsZonesARecordGenerator() gopter.Gen {
	if privateDnsZonesARecordGenerator != nil {
		return privateDnsZonesARecordGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForPrivateDnsZonesARecord(generators)
	privateDnsZonesARecordGenerator = gen.Struct(reflect.TypeOf(PrivateDnsZonesARecord{}), generators)

	return privateDnsZonesARecordGenerator
}

// AddRelatedPropertyGeneratorsForPrivateDnsZonesARecord is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForPrivateDnsZonesARecord(gens map[string]gopter.Gen) {
	gens["Spec"] = PrivateDnsZones_A_SpecGenerator()
	gens["Status"] = PrivateDnsZones_A_STATUSGenerator()
}

func Test_PrivateDnsZones_A_Spec_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from PrivateDnsZones_A_Spec to PrivateDnsZones_A_Spec via AssignProperties_To_PrivateDnsZones_A_Spec & AssignProperties_From_PrivateDnsZones_A_Spec returns original",
		prop.ForAll(RunPropertyAssignmentTestForPrivateDnsZones_A_Spec, PrivateDnsZones_A_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForPrivateDnsZones_A_Spec tests if a specific instance of PrivateDnsZones_A_Spec can be assigned to storage and back losslessly
func RunPropertyAssignmentTestForPrivateDnsZones_A_Spec(subject PrivateDnsZones_A_Spec) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other storage.PrivateDnsZones_A_Spec
	err := copied.AssignProperties_To_PrivateDnsZones_A_Spec(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual PrivateDnsZones_A_Spec
	err = actual.AssignProperties_From_PrivateDnsZones_A_Spec(&other)
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

func Test_PrivateDnsZones_A_Spec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of PrivateDnsZones_A_Spec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForPrivateDnsZones_A_Spec, PrivateDnsZones_A_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForPrivateDnsZones_A_Spec runs a test to see if a specific instance of PrivateDnsZones_A_Spec round trips to JSON and back losslessly
func RunJSONSerializationTestForPrivateDnsZones_A_Spec(subject PrivateDnsZones_A_Spec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual PrivateDnsZones_A_Spec
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

// Generator of PrivateDnsZones_A_Spec instances for property testing - lazily instantiated by
// PrivateDnsZones_A_SpecGenerator()
var privateDnsZones_A_SpecGenerator gopter.Gen

// PrivateDnsZones_A_SpecGenerator returns a generator of PrivateDnsZones_A_Spec instances for property testing.
// We first initialize privateDnsZones_A_SpecGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func PrivateDnsZones_A_SpecGenerator() gopter.Gen {
	if privateDnsZones_A_SpecGenerator != nil {
		return privateDnsZones_A_SpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForPrivateDnsZones_A_Spec(generators)
	privateDnsZones_A_SpecGenerator = gen.Struct(reflect.TypeOf(PrivateDnsZones_A_Spec{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForPrivateDnsZones_A_Spec(generators)
	AddRelatedPropertyGeneratorsForPrivateDnsZones_A_Spec(generators)
	privateDnsZones_A_SpecGenerator = gen.Struct(reflect.TypeOf(PrivateDnsZones_A_Spec{}), generators)

	return privateDnsZones_A_SpecGenerator
}

// AddIndependentPropertyGeneratorsForPrivateDnsZones_A_Spec is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForPrivateDnsZones_A_Spec(gens map[string]gopter.Gen) {
	gens["AzureName"] = gen.AlphaString()
	gens["Etag"] = gen.PtrOf(gen.AlphaString())
	gens["Metadata"] = gen.MapOf(
		gen.AlphaString(),
		gen.AlphaString())
	gens["Ttl"] = gen.PtrOf(gen.Int())
}

// AddRelatedPropertyGeneratorsForPrivateDnsZones_A_Spec is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForPrivateDnsZones_A_Spec(gens map[string]gopter.Gen) {
	gens["ARecords"] = gen.SliceOf(ARecordGenerator())
	gens["AaaaRecords"] = gen.SliceOf(AaaaRecordGenerator())
	gens["CnameRecord"] = gen.PtrOf(CnameRecordGenerator())
	gens["MxRecords"] = gen.SliceOf(MxRecordGenerator())
	gens["PtrRecords"] = gen.SliceOf(PtrRecordGenerator())
	gens["SoaRecord"] = gen.PtrOf(SoaRecordGenerator())
	gens["SrvRecords"] = gen.SliceOf(SrvRecordGenerator())
	gens["TxtRecords"] = gen.SliceOf(TxtRecordGenerator())
}

func Test_PrivateDnsZones_A_STATUS_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from PrivateDnsZones_A_STATUS to PrivateDnsZones_A_STATUS via AssignProperties_To_PrivateDnsZones_A_STATUS & AssignProperties_From_PrivateDnsZones_A_STATUS returns original",
		prop.ForAll(RunPropertyAssignmentTestForPrivateDnsZones_A_STATUS, PrivateDnsZones_A_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForPrivateDnsZones_A_STATUS tests if a specific instance of PrivateDnsZones_A_STATUS can be assigned to storage and back losslessly
func RunPropertyAssignmentTestForPrivateDnsZones_A_STATUS(subject PrivateDnsZones_A_STATUS) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other storage.PrivateDnsZones_A_STATUS
	err := copied.AssignProperties_To_PrivateDnsZones_A_STATUS(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual PrivateDnsZones_A_STATUS
	err = actual.AssignProperties_From_PrivateDnsZones_A_STATUS(&other)
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

func Test_PrivateDnsZones_A_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of PrivateDnsZones_A_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForPrivateDnsZones_A_STATUS, PrivateDnsZones_A_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForPrivateDnsZones_A_STATUS runs a test to see if a specific instance of PrivateDnsZones_A_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForPrivateDnsZones_A_STATUS(subject PrivateDnsZones_A_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual PrivateDnsZones_A_STATUS
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

// Generator of PrivateDnsZones_A_STATUS instances for property testing - lazily instantiated by
// PrivateDnsZones_A_STATUSGenerator()
var privateDnsZones_A_STATUSGenerator gopter.Gen

// PrivateDnsZones_A_STATUSGenerator returns a generator of PrivateDnsZones_A_STATUS instances for property testing.
// We first initialize privateDnsZones_A_STATUSGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func PrivateDnsZones_A_STATUSGenerator() gopter.Gen {
	if privateDnsZones_A_STATUSGenerator != nil {
		return privateDnsZones_A_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForPrivateDnsZones_A_STATUS(generators)
	privateDnsZones_A_STATUSGenerator = gen.Struct(reflect.TypeOf(PrivateDnsZones_A_STATUS{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForPrivateDnsZones_A_STATUS(generators)
	AddRelatedPropertyGeneratorsForPrivateDnsZones_A_STATUS(generators)
	privateDnsZones_A_STATUSGenerator = gen.Struct(reflect.TypeOf(PrivateDnsZones_A_STATUS{}), generators)

	return privateDnsZones_A_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForPrivateDnsZones_A_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForPrivateDnsZones_A_STATUS(gens map[string]gopter.Gen) {
	gens["Etag"] = gen.PtrOf(gen.AlphaString())
	gens["Fqdn"] = gen.PtrOf(gen.AlphaString())
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["IsAutoRegistered"] = gen.PtrOf(gen.Bool())
	gens["Metadata"] = gen.MapOf(
		gen.AlphaString(),
		gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Ttl"] = gen.PtrOf(gen.Int())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForPrivateDnsZones_A_STATUS is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForPrivateDnsZones_A_STATUS(gens map[string]gopter.Gen) {
	gens["ARecords"] = gen.SliceOf(ARecord_STATUSGenerator())
	gens["AaaaRecords"] = gen.SliceOf(AaaaRecord_STATUSGenerator())
	gens["CnameRecord"] = gen.PtrOf(CnameRecord_STATUSGenerator())
	gens["MxRecords"] = gen.SliceOf(MxRecord_STATUSGenerator())
	gens["PtrRecords"] = gen.SliceOf(PtrRecord_STATUSGenerator())
	gens["SoaRecord"] = gen.PtrOf(SoaRecord_STATUSGenerator())
	gens["SrvRecords"] = gen.SliceOf(SrvRecord_STATUSGenerator())
	gens["TxtRecords"] = gen.SliceOf(TxtRecord_STATUSGenerator())
}
