// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20201201storage

import (
	"encoding/json"
	v20230401s "github.com/Azure/azure-service-operator/v2/api/cache/v1api20230401storage"
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

func Test_RedisFirewallRule_WhenConvertedToHub_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	parameters.MinSuccessfulTests = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from RedisFirewallRule to hub returns original",
		prop.ForAll(RunResourceConversionTestForRedisFirewallRule, RedisFirewallRuleGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunResourceConversionTestForRedisFirewallRule tests if a specific instance of RedisFirewallRule round trips to the hub storage version and back losslessly
func RunResourceConversionTestForRedisFirewallRule(subject RedisFirewallRule) string {
	// Copy subject to make sure conversion doesn't modify it
	copied := subject.DeepCopy()

	// Convert to our hub version
	var hub v20230401s.RedisFirewallRule
	err := copied.ConvertTo(&hub)
	if err != nil {
		return err.Error()
	}

	// Convert from our hub version
	var actual RedisFirewallRule
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

func Test_RedisFirewallRule_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from RedisFirewallRule to RedisFirewallRule via AssignProperties_To_RedisFirewallRule & AssignProperties_From_RedisFirewallRule returns original",
		prop.ForAll(RunPropertyAssignmentTestForRedisFirewallRule, RedisFirewallRuleGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForRedisFirewallRule tests if a specific instance of RedisFirewallRule can be assigned to v1api20230401storage and back losslessly
func RunPropertyAssignmentTestForRedisFirewallRule(subject RedisFirewallRule) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20230401s.RedisFirewallRule
	err := copied.AssignProperties_To_RedisFirewallRule(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual RedisFirewallRule
	err = actual.AssignProperties_From_RedisFirewallRule(&other)
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

func Test_RedisFirewallRule_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 20
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of RedisFirewallRule via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForRedisFirewallRule, RedisFirewallRuleGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForRedisFirewallRule runs a test to see if a specific instance of RedisFirewallRule round trips to JSON and back losslessly
func RunJSONSerializationTestForRedisFirewallRule(subject RedisFirewallRule) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual RedisFirewallRule
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

// Generator of RedisFirewallRule instances for property testing - lazily instantiated by RedisFirewallRuleGenerator()
var redisFirewallRuleGenerator gopter.Gen

// RedisFirewallRuleGenerator returns a generator of RedisFirewallRule instances for property testing.
func RedisFirewallRuleGenerator() gopter.Gen {
	if redisFirewallRuleGenerator != nil {
		return redisFirewallRuleGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForRedisFirewallRule(generators)
	redisFirewallRuleGenerator = gen.Struct(reflect.TypeOf(RedisFirewallRule{}), generators)

	return redisFirewallRuleGenerator
}

// AddRelatedPropertyGeneratorsForRedisFirewallRule is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForRedisFirewallRule(gens map[string]gopter.Gen) {
	gens["Spec"] = Redis_FirewallRule_SpecGenerator()
	gens["Status"] = Redis_FirewallRule_STATUSGenerator()
}

func Test_Redis_FirewallRule_Spec_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from Redis_FirewallRule_Spec to Redis_FirewallRule_Spec via AssignProperties_To_Redis_FirewallRule_Spec & AssignProperties_From_Redis_FirewallRule_Spec returns original",
		prop.ForAll(RunPropertyAssignmentTestForRedis_FirewallRule_Spec, Redis_FirewallRule_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForRedis_FirewallRule_Spec tests if a specific instance of Redis_FirewallRule_Spec can be assigned to v1api20230401storage and back losslessly
func RunPropertyAssignmentTestForRedis_FirewallRule_Spec(subject Redis_FirewallRule_Spec) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20230401s.Redis_FirewallRule_Spec
	err := copied.AssignProperties_To_Redis_FirewallRule_Spec(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual Redis_FirewallRule_Spec
	err = actual.AssignProperties_From_Redis_FirewallRule_Spec(&other)
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

func Test_Redis_FirewallRule_Spec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Redis_FirewallRule_Spec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForRedis_FirewallRule_Spec, Redis_FirewallRule_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForRedis_FirewallRule_Spec runs a test to see if a specific instance of Redis_FirewallRule_Spec round trips to JSON and back losslessly
func RunJSONSerializationTestForRedis_FirewallRule_Spec(subject Redis_FirewallRule_Spec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Redis_FirewallRule_Spec
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

// Generator of Redis_FirewallRule_Spec instances for property testing - lazily instantiated by
// Redis_FirewallRule_SpecGenerator()
var redis_FirewallRule_SpecGenerator gopter.Gen

// Redis_FirewallRule_SpecGenerator returns a generator of Redis_FirewallRule_Spec instances for property testing.
func Redis_FirewallRule_SpecGenerator() gopter.Gen {
	if redis_FirewallRule_SpecGenerator != nil {
		return redis_FirewallRule_SpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForRedis_FirewallRule_Spec(generators)
	redis_FirewallRule_SpecGenerator = gen.Struct(reflect.TypeOf(Redis_FirewallRule_Spec{}), generators)

	return redis_FirewallRule_SpecGenerator
}

// AddIndependentPropertyGeneratorsForRedis_FirewallRule_Spec is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForRedis_FirewallRule_Spec(gens map[string]gopter.Gen) {
	gens["AzureName"] = gen.AlphaString()
	gens["EndIP"] = gen.PtrOf(gen.AlphaString())
	gens["OriginalVersion"] = gen.AlphaString()
	gens["StartIP"] = gen.PtrOf(gen.AlphaString())
}

func Test_Redis_FirewallRule_STATUS_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from Redis_FirewallRule_STATUS to Redis_FirewallRule_STATUS via AssignProperties_To_Redis_FirewallRule_STATUS & AssignProperties_From_Redis_FirewallRule_STATUS returns original",
		prop.ForAll(RunPropertyAssignmentTestForRedis_FirewallRule_STATUS, Redis_FirewallRule_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForRedis_FirewallRule_STATUS tests if a specific instance of Redis_FirewallRule_STATUS can be assigned to v1api20230401storage and back losslessly
func RunPropertyAssignmentTestForRedis_FirewallRule_STATUS(subject Redis_FirewallRule_STATUS) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20230401s.Redis_FirewallRule_STATUS
	err := copied.AssignProperties_To_Redis_FirewallRule_STATUS(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual Redis_FirewallRule_STATUS
	err = actual.AssignProperties_From_Redis_FirewallRule_STATUS(&other)
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

func Test_Redis_FirewallRule_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Redis_FirewallRule_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForRedis_FirewallRule_STATUS, Redis_FirewallRule_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForRedis_FirewallRule_STATUS runs a test to see if a specific instance of Redis_FirewallRule_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForRedis_FirewallRule_STATUS(subject Redis_FirewallRule_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Redis_FirewallRule_STATUS
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

// Generator of Redis_FirewallRule_STATUS instances for property testing - lazily instantiated by
// Redis_FirewallRule_STATUSGenerator()
var redis_FirewallRule_STATUSGenerator gopter.Gen

// Redis_FirewallRule_STATUSGenerator returns a generator of Redis_FirewallRule_STATUS instances for property testing.
func Redis_FirewallRule_STATUSGenerator() gopter.Gen {
	if redis_FirewallRule_STATUSGenerator != nil {
		return redis_FirewallRule_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForRedis_FirewallRule_STATUS(generators)
	redis_FirewallRule_STATUSGenerator = gen.Struct(reflect.TypeOf(Redis_FirewallRule_STATUS{}), generators)

	return redis_FirewallRule_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForRedis_FirewallRule_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForRedis_FirewallRule_STATUS(gens map[string]gopter.Gen) {
	gens["EndIP"] = gen.PtrOf(gen.AlphaString())
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["StartIP"] = gen.PtrOf(gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}
