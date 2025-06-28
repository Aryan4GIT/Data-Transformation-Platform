package utils

import (
	"data_mapping/models"
	"fmt"
	"strings"
	"time"

	"github.com/antonmedv/expr"
)

func Transform(input map[string]interface{}, rules []models.MappingRule) (map[string]interface{}, error) {
	output := make(map[string]interface{})
	if apps, ok := input["applicantDetails"].([]interface{}); ok {
		transformedApps := []interface{}{}
		for _, app := range apps {
			appMap, ok := app.(map[string]interface{})
			if !ok {
				continue
			}
			transformed := ApplyRules(appMap, rules)
			transformedApps = append(transformedApps, transformed)
		}
		output["applicants"] = transformedApps
	} else {
		output = ApplyRules(input, rules)
	}
	return output, nil
}

func ApplyRules(input map[string]interface{}, rules []models.MappingRule) map[string]interface{} {
	output := make(map[string]interface{})
	for _, rule := range rules {
		val, exists := GetNestedValue(input, rule.SourcePath)
		if !exists {
			continue
		}
		var transformedVal interface{}
		var err error
		if rule.TransformLogic != "" {
			// Evaluate dynamic logic using expr
			params := map[string]interface{}{"value": val, "input": input}
			transformedVal, err = expr.Eval(rule.TransformLogic, params)
		} else {
			transformedVal, err = ApplyTransform(val, rule.TransformType)
		}
		if err == nil {
			SetNestedValue(output, rule.DestinationPath, transformedVal)
		}
	}
	return output
}

func GetNestedValue(data map[string]interface{}, path []string) (interface{}, bool) {
	current := data
	for i, key := range path {
		val, exists := current[key]
		if !exists {
			return nil, false
		}
		if i == len(path)-1 {
			return val, true
		}
		next, ok := val.(map[string]interface{})
		if !ok {
			return nil, false
		}
		current = next
	}
	return nil, false
}

func SetNestedValue(data map[string]interface{}, path []string, value interface{}) {
	current := data
	for i, key := range path {
		if i == len(path)-1 {
			current[key] = value
			return
		}
		if _, exists := current[key]; !exists {
			current[key] = make(map[string]interface{})
		}
		if next, ok := current[key].(map[string]interface{}); ok {
			current = next
		} else {
			newMap := make(map[string]interface{})
			current[key] = newMap
			current = newMap
		}
	}
}

func ApplyTransform(value interface{}, transformType string) (interface{}, error) {
	switch transformType {
	case "copy":
		return value, nil
	case "toString":
		return fmt.Sprintf("%v", value), nil
	case "mapGender":
		if s, ok := value.(string); ok {
			s = strings.ToUpper(s)
			if s == "MALE" {
				return "M", nil
			} else if s == "FEMALE" {
				return "F", nil
			}
			return "O", nil
		}
		return value, nil
	case "toBool":
		if s, ok := value.(string); ok {
			s = strings.ToLower(s)
			return s == "yes" || s == "true", nil
		}
		if b, ok := value.(bool); ok {
			return b, nil
		}
		return false, nil
	case "formatDate":
		if s, ok := value.(string); ok {
			formats := []string{
				"02-January-2006",
				"02-Jan-2006",
				"02/January/2006",
				"02-January-06",
				"2006-01-02",
				time.RFC3339,
			}
			for _, format := range formats {
				t, err := time.Parse(format, s)
				if err == nil {
					return t.Format("2006-01-02"), nil
				}
			}
			return s, nil
		}
		return value, nil
	case "toUpperCase":
		if s, ok := value.(string); ok {
			return strings.ToUpper(s), nil
		}
		return value, nil
	case "toLowerCase":
		if s, ok := value.(string); ok {
			return strings.ToLower(s), nil
		}
		return value, nil
	case "capitalize":
		if s, ok := value.(string); ok {
			if len(s) == 0 {
				return s, nil
			}
			return strings.ToUpper(s[:1]) + strings.ToLower(s[1:]), nil
		}
		return value, nil
	default:
		return value, nil
	}
}
