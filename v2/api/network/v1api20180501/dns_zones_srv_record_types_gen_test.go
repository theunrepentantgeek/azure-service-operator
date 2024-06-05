// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20180501

import (
	"encoding/json"
	storage "github.com/Azure/azure-service-operator/v2/api/network/v1api20180501/storage"
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

func Test_DnsZonesSRVRecord_WhenConvertedToHub_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	parameters.MinSuccessfulTests = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from DnsZonesSRVRecord to hub returns original",
		prop.ForAll(RunResourceConversionTestForDnsZonesSRVRecord, DnsZonesSRVRecordGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunResourceConversionTestForDnsZonesSRVRecord tests if a specific instance of DnsZonesSRVRecord round trips to the hub storage version and back losslessly
func RunResourceConversionTestForDnsZonesSRVRecord(subject DnsZonesSRVRecord) string {
	// Copy subject to make sure conversion doesn't modify it
	copied := subject.DeepCopy()

	// Convert to our hub version
	var hub storage.DnsZonesSRVRecord
	err := copied.ConvertTo(&hub)
	if err != nil {
		return err.Error()
	}

	// Convert from our hub version
	var actual DnsZonesSRVRecord
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

func Test_DnsZonesSRVRecord_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from DnsZonesSRVRecord to DnsZonesSRVRecord via AssignProperties_To_DnsZonesSRVRecord & AssignProperties_From_DnsZonesSRVRecord returns original",
		prop.ForAll(RunPropertyAssignmentTestForDnsZonesSRVRecord, DnsZonesSRVRecordGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForDnsZonesSRVRecord tests if a specific instance of DnsZonesSRVRecord can be assigned to storage and back losslessly
func RunPropertyAssignmentTestForDnsZonesSRVRecord(subject DnsZonesSRVRecord) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other storage.DnsZonesSRVRecord
	err := copied.AssignProperties_To_DnsZonesSRVRecord(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual DnsZonesSRVRecord
	err = actual.AssignProperties_From_DnsZonesSRVRecord(&other)
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

func Test_DnsZonesSRVRecord_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 20
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of DnsZonesSRVRecord via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForDnsZonesSRVRecord, DnsZonesSRVRecordGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForDnsZonesSRVRecord runs a test to see if a specific instance of DnsZonesSRVRecord round trips to JSON and back losslessly
func RunJSONSerializationTestForDnsZonesSRVRecord(subject DnsZonesSRVRecord) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual DnsZonesSRVRecord
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

// Generator of DnsZonesSRVRecord instances for property testing - lazily instantiated by DnsZonesSRVRecordGenerator()
var dnsZonesSRVRecordGenerator gopter.Gen

// DnsZonesSRVRecordGenerator returns a generator of DnsZonesSRVRecord instances for property testing.
func DnsZonesSRVRecordGenerator() gopter.Gen {
	if dnsZonesSRVRecordGenerator != nil {
		return dnsZonesSRVRecordGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForDnsZonesSRVRecord(generators)
	dnsZonesSRVRecordGenerator = gen.Struct(reflect.TypeOf(DnsZonesSRVRecord{}), generators)

	return dnsZonesSRVRecordGenerator
}

// AddRelatedPropertyGeneratorsForDnsZonesSRVRecord is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForDnsZonesSRVRecord(gens map[string]gopter.Gen) {
	gens["Spec"] = DnsZones_SRV_SpecGenerator()
	gens["Status"] = DnsZones_SRV_STATUSGenerator()
}

func Test_DnsZones_SRV_Spec_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from DnsZones_SRV_Spec to DnsZones_SRV_Spec via AssignProperties_To_DnsZones_SRV_Spec & AssignProperties_From_DnsZones_SRV_Spec returns original",
		prop.ForAll(RunPropertyAssignmentTestForDnsZones_SRV_Spec, DnsZones_SRV_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForDnsZones_SRV_Spec tests if a specific instance of DnsZones_SRV_Spec can be assigned to storage and back losslessly
func RunPropertyAssignmentTestForDnsZones_SRV_Spec(subject DnsZones_SRV_Spec) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other storage.DnsZones_SRV_Spec
	err := copied.AssignProperties_To_DnsZones_SRV_Spec(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual DnsZones_SRV_Spec
	err = actual.AssignProperties_From_DnsZones_SRV_Spec(&other)
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

func Test_DnsZones_SRV_Spec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of DnsZones_SRV_Spec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForDnsZones_SRV_Spec, DnsZones_SRV_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForDnsZones_SRV_Spec runs a test to see if a specific instance of DnsZones_SRV_Spec round trips to JSON and back losslessly
func RunJSONSerializationTestForDnsZones_SRV_Spec(subject DnsZones_SRV_Spec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual DnsZones_SRV_Spec
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

// Generator of DnsZones_SRV_Spec instances for property testing - lazily instantiated by DnsZones_SRV_SpecGenerator()
var dnsZones_SRV_SpecGenerator gopter.Gen

// DnsZones_SRV_SpecGenerator returns a generator of DnsZones_SRV_Spec instances for property testing.
// We first initialize dnsZones_SRV_SpecGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func DnsZones_SRV_SpecGenerator() gopter.Gen {
	if dnsZones_SRV_SpecGenerator != nil {
		return dnsZones_SRV_SpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDnsZones_SRV_Spec(generators)
	dnsZones_SRV_SpecGenerator = gen.Struct(reflect.TypeOf(DnsZones_SRV_Spec{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDnsZones_SRV_Spec(generators)
	AddRelatedPropertyGeneratorsForDnsZones_SRV_Spec(generators)
	dnsZones_SRV_SpecGenerator = gen.Struct(reflect.TypeOf(DnsZones_SRV_Spec{}), generators)

	return dnsZones_SRV_SpecGenerator
}

// AddIndependentPropertyGeneratorsForDnsZones_SRV_Spec is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForDnsZones_SRV_Spec(gens map[string]gopter.Gen) {
	gens["AzureName"] = gen.AlphaString()
	gens["Metadata"] = gen.MapOf(
		gen.AlphaString(),
		gen.AlphaString())
	gens["TTL"] = gen.PtrOf(gen.Int())
}

// AddRelatedPropertyGeneratorsForDnsZones_SRV_Spec is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForDnsZones_SRV_Spec(gens map[string]gopter.Gen) {
	gens["AAAARecords"] = gen.SliceOf(AaaaRecordGenerator())
	gens["ARecords"] = gen.SliceOf(ARecordGenerator())
	gens["CNAMERecord"] = gen.PtrOf(CnameRecordGenerator())
	gens["CaaRecords"] = gen.SliceOf(CaaRecordGenerator())
	gens["MXRecords"] = gen.SliceOf(MxRecordGenerator())
	gens["NSRecords"] = gen.SliceOf(NsRecordGenerator())
	gens["PTRRecords"] = gen.SliceOf(PtrRecordGenerator())
	gens["SOARecord"] = gen.PtrOf(SoaRecordGenerator())
	gens["SRVRecords"] = gen.SliceOf(SrvRecordGenerator())
	gens["TXTRecords"] = gen.SliceOf(TxtRecordGenerator())
	gens["TargetResource"] = gen.PtrOf(SubResourceGenerator())
}

func Test_DnsZones_SRV_STATUS_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from DnsZones_SRV_STATUS to DnsZones_SRV_STATUS via AssignProperties_To_DnsZones_SRV_STATUS & AssignProperties_From_DnsZones_SRV_STATUS returns original",
		prop.ForAll(RunPropertyAssignmentTestForDnsZones_SRV_STATUS, DnsZones_SRV_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForDnsZones_SRV_STATUS tests if a specific instance of DnsZones_SRV_STATUS can be assigned to storage and back losslessly
func RunPropertyAssignmentTestForDnsZones_SRV_STATUS(subject DnsZones_SRV_STATUS) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other storage.DnsZones_SRV_STATUS
	err := copied.AssignProperties_To_DnsZones_SRV_STATUS(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual DnsZones_SRV_STATUS
	err = actual.AssignProperties_From_DnsZones_SRV_STATUS(&other)
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

func Test_DnsZones_SRV_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of DnsZones_SRV_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForDnsZones_SRV_STATUS, DnsZones_SRV_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForDnsZones_SRV_STATUS runs a test to see if a specific instance of DnsZones_SRV_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForDnsZones_SRV_STATUS(subject DnsZones_SRV_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual DnsZones_SRV_STATUS
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

// Generator of DnsZones_SRV_STATUS instances for property testing - lazily instantiated by
// DnsZones_SRV_STATUSGenerator()
var dnsZones_SRV_STATUSGenerator gopter.Gen

// DnsZones_SRV_STATUSGenerator returns a generator of DnsZones_SRV_STATUS instances for property testing.
// We first initialize dnsZones_SRV_STATUSGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func DnsZones_SRV_STATUSGenerator() gopter.Gen {
	if dnsZones_SRV_STATUSGenerator != nil {
		return dnsZones_SRV_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDnsZones_SRV_STATUS(generators)
	dnsZones_SRV_STATUSGenerator = gen.Struct(reflect.TypeOf(DnsZones_SRV_STATUS{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDnsZones_SRV_STATUS(generators)
	AddRelatedPropertyGeneratorsForDnsZones_SRV_STATUS(generators)
	dnsZones_SRV_STATUSGenerator = gen.Struct(reflect.TypeOf(DnsZones_SRV_STATUS{}), generators)

	return dnsZones_SRV_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForDnsZones_SRV_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForDnsZones_SRV_STATUS(gens map[string]gopter.Gen) {
	gens["Etag"] = gen.PtrOf(gen.AlphaString())
	gens["Fqdn"] = gen.PtrOf(gen.AlphaString())
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Metadata"] = gen.MapOf(
		gen.AlphaString(),
		gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["ProvisioningState"] = gen.PtrOf(gen.AlphaString())
	gens["TTL"] = gen.PtrOf(gen.Int())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForDnsZones_SRV_STATUS is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForDnsZones_SRV_STATUS(gens map[string]gopter.Gen) {
	gens["AAAARecords"] = gen.SliceOf(AaaaRecord_STATUSGenerator())
	gens["ARecords"] = gen.SliceOf(ARecord_STATUSGenerator())
	gens["CNAMERecord"] = gen.PtrOf(CnameRecord_STATUSGenerator())
	gens["CaaRecords"] = gen.SliceOf(CaaRecord_STATUSGenerator())
	gens["MXRecords"] = gen.SliceOf(MxRecord_STATUSGenerator())
	gens["NSRecords"] = gen.SliceOf(NsRecord_STATUSGenerator())
	gens["PTRRecords"] = gen.SliceOf(PtrRecord_STATUSGenerator())
	gens["SOARecord"] = gen.PtrOf(SoaRecord_STATUSGenerator())
	gens["SRVRecords"] = gen.SliceOf(SrvRecord_STATUSGenerator())
	gens["TXTRecords"] = gen.SliceOf(TxtRecord_STATUSGenerator())
	gens["TargetResource"] = gen.PtrOf(SubResource_STATUSGenerator())
}
