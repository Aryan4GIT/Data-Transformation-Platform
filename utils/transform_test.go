package utils

import (
	"data_mapping/models"
	"testing"
)

func TestTransform(t *testing.T) {
	// Test basic transformation
	input := map[string]interface{}{
		"customer": map[string]interface{}{
			"firstName": "John",
			"lastName":  "Doe",
		},
	}

	rules := []models.MappingRule{
		{
			SourcePath:      []string{"customer", "firstName"},
			DestinationPath: []string{"user", "first_name"},
			TransformType:   "copy",
		},
		{
			SourcePath:      []string{"customer", "lastName"},
			DestinationPath: []string{"user", "last_name"},
			TransformType:   "toUpperCase",
		},
	}

	output, err := Transform(input, rules)
	if err != nil {
		t.Fatalf("Transform failed: %v", err)
	}

	user, ok := output["user"].(map[string]interface{})
	if !ok {
		t.Fatal("Expected user object in output")
	}

	if user["first_name"] != "John" {
		t.Errorf("Expected first_name to be 'John', got '%v'", user["first_name"])
	}

	if user["last_name"] != "DOE" {
		t.Errorf("Expected last_name to be 'DOE', got '%v'", user["last_name"])
	}
}

func TestTransformWithRequiredFields(t *testing.T) {
	// Test transformation with required fields and default values
	input := map[string]interface{}{
		"customer": map[string]interface{}{
			"firstName": "Jane",
		},
	}

	rules := []models.MappingRule{
		{
			SourcePath:      []string{"customer", "firstName"},
			DestinationPath: []string{"user", "first_name"},
			TransformType:   "copy",
			Required:        true,
		},
		{
			SourcePath:      []string{"customer", "age"},
			DestinationPath: []string{"user", "age"},
			TransformType:   "copy",
			Required:        true,
			DefaultValue:    "30",
		},
	}

	output, err := Transform(input, rules)
	if err != nil {
		t.Fatalf("Transform failed: %v", err)
	}

	user, ok := output["user"].(map[string]interface{})
	if !ok {
		t.Fatal("Expected user object in output")
	}

	if user["first_name"] != "Jane" {
		t.Errorf("Expected first_name to be 'Jane', got '%v'", user["first_name"])
	}

	if user["age"] != 30 {
		t.Errorf("Expected age to be 30, got '%v'", user["age"])
	}
}

func TestApplyTransform(t *testing.T) {
	tests := []struct {
		value         interface{}
		transformType string
		expected      interface{}
	}{
		{"hello", "toUpperCase", "HELLO"},
		{"WORLD", "toLowerCase", "world"},
		{"john", "capitalize", "John"},
		{"Male", "mapGender", "M"},
		{"Female", "mapGender", "F"},
		{"yes", "toBool", true},
		{"no", "toBool", false},
	}

	for _, test := range tests {
		result, err := ApplyTransform(test.value, test.transformType)
		if err != nil {
			t.Errorf("ApplyTransform failed for %v with %s: %v", test.value, test.transformType, err)
			continue
		}

		if result != test.expected {
			t.Errorf("ApplyTransform(%v, %s) = %v, expected %v", test.value, test.transformType, result, test.expected)
		}
	}
}

func TestDynamicExpressions(t *testing.T) {
	// Test with simple expressions
	expr1 := "value * 2"
	context1 := map[string]interface{}{
		"value": 10,
	}
	_, err := EvaluateExpression(expr1, context1)
	if err != nil {
		t.Errorf("Error evaluating expression %s: %v", expr1, err)
	}
}
