// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package compat

import (
	"encoding/json"
	storage "github.com/Azure/azure-service-operator/v2/api/containerservice/v1api20231001/storage"
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

func Test_ServiceMeshProfile_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from ServiceMeshProfile to ServiceMeshProfile via AssignProperties_To_ServiceMeshProfile & AssignProperties_From_ServiceMeshProfile returns original",
		prop.ForAll(RunPropertyAssignmentTestForServiceMeshProfile, ServiceMeshProfileGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForServiceMeshProfile tests if a specific instance of ServiceMeshProfile can be assigned to storage and back losslessly
func RunPropertyAssignmentTestForServiceMeshProfile(subject ServiceMeshProfile) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other storage.ServiceMeshProfile
	err := copied.AssignProperties_To_ServiceMeshProfile(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual ServiceMeshProfile
	err = actual.AssignProperties_From_ServiceMeshProfile(&other)
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

func Test_ServiceMeshProfile_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ServiceMeshProfile via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForServiceMeshProfile, ServiceMeshProfileGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForServiceMeshProfile runs a test to see if a specific instance of ServiceMeshProfile round trips to JSON and back losslessly
func RunJSONSerializationTestForServiceMeshProfile(subject ServiceMeshProfile) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ServiceMeshProfile
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

// Generator of ServiceMeshProfile instances for property testing - lazily instantiated by ServiceMeshProfileGenerator()
var serviceMeshProfileGenerator gopter.Gen

// ServiceMeshProfileGenerator returns a generator of ServiceMeshProfile instances for property testing.
// We first initialize serviceMeshProfileGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func ServiceMeshProfileGenerator() gopter.Gen {
	if serviceMeshProfileGenerator != nil {
		return serviceMeshProfileGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForServiceMeshProfile(generators)
	serviceMeshProfileGenerator = gen.Struct(reflect.TypeOf(ServiceMeshProfile{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForServiceMeshProfile(generators)
	AddRelatedPropertyGeneratorsForServiceMeshProfile(generators)
	serviceMeshProfileGenerator = gen.Struct(reflect.TypeOf(ServiceMeshProfile{}), generators)

	return serviceMeshProfileGenerator
}

// AddIndependentPropertyGeneratorsForServiceMeshProfile is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForServiceMeshProfile(gens map[string]gopter.Gen) {
	gens["Mode"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForServiceMeshProfile is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForServiceMeshProfile(gens map[string]gopter.Gen) {
	gens["Istio"] = gen.PtrOf(IstioServiceMeshGenerator())
}

func Test_IstioServiceMesh_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from IstioServiceMesh to IstioServiceMesh via AssignProperties_To_IstioServiceMesh & AssignProperties_From_IstioServiceMesh returns original",
		prop.ForAll(RunPropertyAssignmentTestForIstioServiceMesh, IstioServiceMeshGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForIstioServiceMesh tests if a specific instance of IstioServiceMesh can be assigned to storage and back losslessly
func RunPropertyAssignmentTestForIstioServiceMesh(subject IstioServiceMesh) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other storage.IstioServiceMesh
	err := copied.AssignProperties_To_IstioServiceMesh(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual IstioServiceMesh
	err = actual.AssignProperties_From_IstioServiceMesh(&other)
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

func Test_IstioServiceMesh_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of IstioServiceMesh via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForIstioServiceMesh, IstioServiceMeshGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForIstioServiceMesh runs a test to see if a specific instance of IstioServiceMesh round trips to JSON and back losslessly
func RunJSONSerializationTestForIstioServiceMesh(subject IstioServiceMesh) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual IstioServiceMesh
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

// Generator of IstioServiceMesh instances for property testing - lazily instantiated by IstioServiceMeshGenerator()
var istioServiceMeshGenerator gopter.Gen

// IstioServiceMeshGenerator returns a generator of IstioServiceMesh instances for property testing.
func IstioServiceMeshGenerator() gopter.Gen {
	if istioServiceMeshGenerator != nil {
		return istioServiceMeshGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForIstioServiceMesh(generators)
	istioServiceMeshGenerator = gen.Struct(reflect.TypeOf(IstioServiceMesh{}), generators)

	return istioServiceMeshGenerator
}

// AddRelatedPropertyGeneratorsForIstioServiceMesh is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForIstioServiceMesh(gens map[string]gopter.Gen) {
	gens["Components"] = gen.PtrOf(IstioComponentsGenerator())
}

func Test_IstioComponents_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from IstioComponents to IstioComponents via AssignProperties_To_IstioComponents & AssignProperties_From_IstioComponents returns original",
		prop.ForAll(RunPropertyAssignmentTestForIstioComponents, IstioComponentsGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForIstioComponents tests if a specific instance of IstioComponents can be assigned to storage and back losslessly
func RunPropertyAssignmentTestForIstioComponents(subject IstioComponents) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other storage.IstioComponents
	err := copied.AssignProperties_To_IstioComponents(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual IstioComponents
	err = actual.AssignProperties_From_IstioComponents(&other)
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

func Test_IstioComponents_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of IstioComponents via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForIstioComponents, IstioComponentsGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForIstioComponents runs a test to see if a specific instance of IstioComponents round trips to JSON and back losslessly
func RunJSONSerializationTestForIstioComponents(subject IstioComponents) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual IstioComponents
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

// Generator of IstioComponents instances for property testing - lazily instantiated by IstioComponentsGenerator()
var istioComponentsGenerator gopter.Gen

// IstioComponentsGenerator returns a generator of IstioComponents instances for property testing.
func IstioComponentsGenerator() gopter.Gen {
	if istioComponentsGenerator != nil {
		return istioComponentsGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForIstioComponents(generators)
	istioComponentsGenerator = gen.Struct(reflect.TypeOf(IstioComponents{}), generators)

	return istioComponentsGenerator
}

// AddRelatedPropertyGeneratorsForIstioComponents is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForIstioComponents(gens map[string]gopter.Gen) {
	gens["IngressGateways"] = gen.SliceOf(IstioIngressGatewayGenerator())
}

func Test_IstioIngressGateway_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from IstioIngressGateway to IstioIngressGateway via AssignProperties_To_IstioIngressGateway & AssignProperties_From_IstioIngressGateway returns original",
		prop.ForAll(RunPropertyAssignmentTestForIstioIngressGateway, IstioIngressGatewayGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForIstioIngressGateway tests if a specific instance of IstioIngressGateway can be assigned to storage and back losslessly
func RunPropertyAssignmentTestForIstioIngressGateway(subject IstioIngressGateway) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other storage.IstioIngressGateway
	err := copied.AssignProperties_To_IstioIngressGateway(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual IstioIngressGateway
	err = actual.AssignProperties_From_IstioIngressGateway(&other)
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

func Test_IstioIngressGateway_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of IstioIngressGateway via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForIstioIngressGateway, IstioIngressGatewayGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForIstioIngressGateway runs a test to see if a specific instance of IstioIngressGateway round trips to JSON and back losslessly
func RunJSONSerializationTestForIstioIngressGateway(subject IstioIngressGateway) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual IstioIngressGateway
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

// Generator of IstioIngressGateway instances for property testing - lazily instantiated by
// IstioIngressGatewayGenerator()
var istioIngressGatewayGenerator gopter.Gen

// IstioIngressGatewayGenerator returns a generator of IstioIngressGateway instances for property testing.
func IstioIngressGatewayGenerator() gopter.Gen {
	if istioIngressGatewayGenerator != nil {
		return istioIngressGatewayGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForIstioIngressGateway(generators)
	istioIngressGatewayGenerator = gen.Struct(reflect.TypeOf(IstioIngressGateway{}), generators)

	return istioIngressGatewayGenerator
}

// AddIndependentPropertyGeneratorsForIstioIngressGateway is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForIstioIngressGateway(gens map[string]gopter.Gen) {
	gens["Enabled"] = gen.PtrOf(gen.Bool())
	gens["Mode"] = gen.PtrOf(gen.AlphaString())
}
