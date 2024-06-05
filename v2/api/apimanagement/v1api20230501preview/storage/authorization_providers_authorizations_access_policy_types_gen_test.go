// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package storage

import (
	"encoding/json"
	storage "github.com/Azure/azure-service-operator/v2/api/apimanagement/v1api20220801/storage"
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

func Test_AuthorizationProvidersAuthorizationsAccessPolicy_WhenConvertedToHub_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	parameters.MinSuccessfulTests = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from AuthorizationProvidersAuthorizationsAccessPolicy to hub returns original",
		prop.ForAll(RunResourceConversionTestForAuthorizationProvidersAuthorizationsAccessPolicy, AuthorizationProvidersAuthorizationsAccessPolicyGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunResourceConversionTestForAuthorizationProvidersAuthorizationsAccessPolicy tests if a specific instance of AuthorizationProvidersAuthorizationsAccessPolicy round trips to the hub storage version and back losslessly
func RunResourceConversionTestForAuthorizationProvidersAuthorizationsAccessPolicy(subject AuthorizationProvidersAuthorizationsAccessPolicy) string {
	// Copy subject to make sure conversion doesn't modify it
	copied := subject.DeepCopy()

	// Convert to our hub version
	var hub storage.AuthorizationProvidersAuthorizationsAccessPolicy
	err := copied.ConvertTo(&hub)
	if err != nil {
		return err.Error()
	}

	// Convert from our hub version
	var actual AuthorizationProvidersAuthorizationsAccessPolicy
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

func Test_AuthorizationProvidersAuthorizationsAccessPolicy_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from AuthorizationProvidersAuthorizationsAccessPolicy to AuthorizationProvidersAuthorizationsAccessPolicy via AssignProperties_To_AuthorizationProvidersAuthorizationsAccessPolicy & AssignProperties_From_AuthorizationProvidersAuthorizationsAccessPolicy returns original",
		prop.ForAll(RunPropertyAssignmentTestForAuthorizationProvidersAuthorizationsAccessPolicy, AuthorizationProvidersAuthorizationsAccessPolicyGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForAuthorizationProvidersAuthorizationsAccessPolicy tests if a specific instance of AuthorizationProvidersAuthorizationsAccessPolicy can be assigned to storage and back losslessly
func RunPropertyAssignmentTestForAuthorizationProvidersAuthorizationsAccessPolicy(subject AuthorizationProvidersAuthorizationsAccessPolicy) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other storage.AuthorizationProvidersAuthorizationsAccessPolicy
	err := copied.AssignProperties_To_AuthorizationProvidersAuthorizationsAccessPolicy(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual AuthorizationProvidersAuthorizationsAccessPolicy
	err = actual.AssignProperties_From_AuthorizationProvidersAuthorizationsAccessPolicy(&other)
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

func Test_AuthorizationProvidersAuthorizationsAccessPolicy_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 20
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of AuthorizationProvidersAuthorizationsAccessPolicy via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForAuthorizationProvidersAuthorizationsAccessPolicy, AuthorizationProvidersAuthorizationsAccessPolicyGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForAuthorizationProvidersAuthorizationsAccessPolicy runs a test to see if a specific instance of AuthorizationProvidersAuthorizationsAccessPolicy round trips to JSON and back losslessly
func RunJSONSerializationTestForAuthorizationProvidersAuthorizationsAccessPolicy(subject AuthorizationProvidersAuthorizationsAccessPolicy) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual AuthorizationProvidersAuthorizationsAccessPolicy
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

// Generator of AuthorizationProvidersAuthorizationsAccessPolicy instances for property testing - lazily instantiated by
// AuthorizationProvidersAuthorizationsAccessPolicyGenerator()
var authorizationProvidersAuthorizationsAccessPolicyGenerator gopter.Gen

// AuthorizationProvidersAuthorizationsAccessPolicyGenerator returns a generator of AuthorizationProvidersAuthorizationsAccessPolicy instances for property testing.
func AuthorizationProvidersAuthorizationsAccessPolicyGenerator() gopter.Gen {
	if authorizationProvidersAuthorizationsAccessPolicyGenerator != nil {
		return authorizationProvidersAuthorizationsAccessPolicyGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForAuthorizationProvidersAuthorizationsAccessPolicy(generators)
	authorizationProvidersAuthorizationsAccessPolicyGenerator = gen.Struct(reflect.TypeOf(AuthorizationProvidersAuthorizationsAccessPolicy{}), generators)

	return authorizationProvidersAuthorizationsAccessPolicyGenerator
}

// AddRelatedPropertyGeneratorsForAuthorizationProvidersAuthorizationsAccessPolicy is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForAuthorizationProvidersAuthorizationsAccessPolicy(gens map[string]gopter.Gen) {
	gens["Spec"] = Service_AuthorizationProviders_Authorizations_AccessPolicy_SpecGenerator()
	gens["Status"] = Service_AuthorizationProviders_Authorizations_AccessPolicy_STATUSGenerator()
}

func Test_Service_AuthorizationProviders_Authorizations_AccessPolicy_Spec_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from Service_AuthorizationProviders_Authorizations_AccessPolicy_Spec to Service_AuthorizationProviders_Authorizations_AccessPolicy_Spec via AssignProperties_To_Service_AuthorizationProviders_Authorizations_AccessPolicy_Spec & AssignProperties_From_Service_AuthorizationProviders_Authorizations_AccessPolicy_Spec returns original",
		prop.ForAll(RunPropertyAssignmentTestForService_AuthorizationProviders_Authorizations_AccessPolicy_Spec, Service_AuthorizationProviders_Authorizations_AccessPolicy_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForService_AuthorizationProviders_Authorizations_AccessPolicy_Spec tests if a specific instance of Service_AuthorizationProviders_Authorizations_AccessPolicy_Spec can be assigned to storage and back losslessly
func RunPropertyAssignmentTestForService_AuthorizationProviders_Authorizations_AccessPolicy_Spec(subject Service_AuthorizationProviders_Authorizations_AccessPolicy_Spec) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other storage.Service_AuthorizationProviders_Authorizations_AccessPolicy_Spec
	err := copied.AssignProperties_To_Service_AuthorizationProviders_Authorizations_AccessPolicy_Spec(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual Service_AuthorizationProviders_Authorizations_AccessPolicy_Spec
	err = actual.AssignProperties_From_Service_AuthorizationProviders_Authorizations_AccessPolicy_Spec(&other)
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

func Test_Service_AuthorizationProviders_Authorizations_AccessPolicy_Spec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Service_AuthorizationProviders_Authorizations_AccessPolicy_Spec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForService_AuthorizationProviders_Authorizations_AccessPolicy_Spec, Service_AuthorizationProviders_Authorizations_AccessPolicy_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForService_AuthorizationProviders_Authorizations_AccessPolicy_Spec runs a test to see if a specific instance of Service_AuthorizationProviders_Authorizations_AccessPolicy_Spec round trips to JSON and back losslessly
func RunJSONSerializationTestForService_AuthorizationProviders_Authorizations_AccessPolicy_Spec(subject Service_AuthorizationProviders_Authorizations_AccessPolicy_Spec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Service_AuthorizationProviders_Authorizations_AccessPolicy_Spec
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

// Generator of Service_AuthorizationProviders_Authorizations_AccessPolicy_Spec instances for property testing - lazily
// instantiated by Service_AuthorizationProviders_Authorizations_AccessPolicy_SpecGenerator()
var service_AuthorizationProviders_Authorizations_AccessPolicy_SpecGenerator gopter.Gen

// Service_AuthorizationProviders_Authorizations_AccessPolicy_SpecGenerator returns a generator of Service_AuthorizationProviders_Authorizations_AccessPolicy_Spec instances for property testing.
func Service_AuthorizationProviders_Authorizations_AccessPolicy_SpecGenerator() gopter.Gen {
	if service_AuthorizationProviders_Authorizations_AccessPolicy_SpecGenerator != nil {
		return service_AuthorizationProviders_Authorizations_AccessPolicy_SpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForService_AuthorizationProviders_Authorizations_AccessPolicy_Spec(generators)
	service_AuthorizationProviders_Authorizations_AccessPolicy_SpecGenerator = gen.Struct(reflect.TypeOf(Service_AuthorizationProviders_Authorizations_AccessPolicy_Spec{}), generators)

	return service_AuthorizationProviders_Authorizations_AccessPolicy_SpecGenerator
}

// AddIndependentPropertyGeneratorsForService_AuthorizationProviders_Authorizations_AccessPolicy_Spec is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForService_AuthorizationProviders_Authorizations_AccessPolicy_Spec(gens map[string]gopter.Gen) {
	gens["AppIds"] = gen.SliceOf(gen.AlphaString())
	gens["AzureName"] = gen.AlphaString()
	gens["ObjectId"] = gen.PtrOf(gen.AlphaString())
	gens["OriginalVersion"] = gen.AlphaString()
	gens["TenantId"] = gen.PtrOf(gen.AlphaString())
}

func Test_Service_AuthorizationProviders_Authorizations_AccessPolicy_STATUS_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from Service_AuthorizationProviders_Authorizations_AccessPolicy_STATUS to Service_AuthorizationProviders_Authorizations_AccessPolicy_STATUS via AssignProperties_To_Service_AuthorizationProviders_Authorizations_AccessPolicy_STATUS & AssignProperties_From_Service_AuthorizationProviders_Authorizations_AccessPolicy_STATUS returns original",
		prop.ForAll(RunPropertyAssignmentTestForService_AuthorizationProviders_Authorizations_AccessPolicy_STATUS, Service_AuthorizationProviders_Authorizations_AccessPolicy_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForService_AuthorizationProviders_Authorizations_AccessPolicy_STATUS tests if a specific instance of Service_AuthorizationProviders_Authorizations_AccessPolicy_STATUS can be assigned to storage and back losslessly
func RunPropertyAssignmentTestForService_AuthorizationProviders_Authorizations_AccessPolicy_STATUS(subject Service_AuthorizationProviders_Authorizations_AccessPolicy_STATUS) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other storage.Service_AuthorizationProviders_Authorizations_AccessPolicy_STATUS
	err := copied.AssignProperties_To_Service_AuthorizationProviders_Authorizations_AccessPolicy_STATUS(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual Service_AuthorizationProviders_Authorizations_AccessPolicy_STATUS
	err = actual.AssignProperties_From_Service_AuthorizationProviders_Authorizations_AccessPolicy_STATUS(&other)
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

func Test_Service_AuthorizationProviders_Authorizations_AccessPolicy_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Service_AuthorizationProviders_Authorizations_AccessPolicy_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForService_AuthorizationProviders_Authorizations_AccessPolicy_STATUS, Service_AuthorizationProviders_Authorizations_AccessPolicy_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForService_AuthorizationProviders_Authorizations_AccessPolicy_STATUS runs a test to see if a specific instance of Service_AuthorizationProviders_Authorizations_AccessPolicy_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForService_AuthorizationProviders_Authorizations_AccessPolicy_STATUS(subject Service_AuthorizationProviders_Authorizations_AccessPolicy_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Service_AuthorizationProviders_Authorizations_AccessPolicy_STATUS
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

// Generator of Service_AuthorizationProviders_Authorizations_AccessPolicy_STATUS instances for property testing -
// lazily instantiated by Service_AuthorizationProviders_Authorizations_AccessPolicy_STATUSGenerator()
var service_AuthorizationProviders_Authorizations_AccessPolicy_STATUSGenerator gopter.Gen

// Service_AuthorizationProviders_Authorizations_AccessPolicy_STATUSGenerator returns a generator of Service_AuthorizationProviders_Authorizations_AccessPolicy_STATUS instances for property testing.
func Service_AuthorizationProviders_Authorizations_AccessPolicy_STATUSGenerator() gopter.Gen {
	if service_AuthorizationProviders_Authorizations_AccessPolicy_STATUSGenerator != nil {
		return service_AuthorizationProviders_Authorizations_AccessPolicy_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForService_AuthorizationProviders_Authorizations_AccessPolicy_STATUS(generators)
	service_AuthorizationProviders_Authorizations_AccessPolicy_STATUSGenerator = gen.Struct(reflect.TypeOf(Service_AuthorizationProviders_Authorizations_AccessPolicy_STATUS{}), generators)

	return service_AuthorizationProviders_Authorizations_AccessPolicy_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForService_AuthorizationProviders_Authorizations_AccessPolicy_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForService_AuthorizationProviders_Authorizations_AccessPolicy_STATUS(gens map[string]gopter.Gen) {
	gens["AppIds"] = gen.SliceOf(gen.AlphaString())
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["ObjectId"] = gen.PtrOf(gen.AlphaString())
	gens["TenantId"] = gen.PtrOf(gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}
