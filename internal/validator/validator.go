package validator

import (
	"encoding/json"
	"fmt"

	"github.com/MayR-Labs/envdoc-go/internal/parser"
)

// Schema represents a JSON schema for environment variables
type Schema struct {
	Schema     string              `json:"$schema"`
	Type       string              `json:"type"`
	Properties map[string]Property `json:"properties"`
	Required   []string            `json:"required,omitempty"`
}

// Property represents a property in the schema
type Property struct {
	Type        string `json:"type"`
	Description string `json:"description,omitempty"`
}

// GenerateSchema generates a JSON schema from environment variables
func GenerateSchema(envVars []parser.EnvVar) (string, error) {
	properties := make(map[string]Property)
	var required []string

	for _, envVar := range envVars {
		description := envVar.Comment
		if description != "" && len(description) > 2 {
			description = description[2:] // Remove "# " prefix
		}

		properties[envVar.Key] = Property{
			Type:        "string",
			Description: description,
		}
		required = append(required, envVar.Key)
	}

	schema := Schema{
		Schema:     "http://json-schema.org/draft-07/schema#",
		Type:       "object",
		Properties: properties,
		Required:   required,
	}

	jsonData, err := json.MarshalIndent(schema, "", "  ")
	if err != nil {
		return "", fmt.Errorf("failed to marshal schema: %w", err)
	}

	return string(jsonData), nil
}

// ValidateAgainstSchema validates environment variables against a JSON schema
func ValidateAgainstSchema(envVars []parser.EnvVar, schemaJSON string) ([]string, error) {
	var schema Schema
	if err := json.Unmarshal([]byte(schemaJSON), &schema); err != nil {
		return nil, fmt.Errorf("failed to parse schema: %w", err)
	}

	var errors []string
	envKeys := make(map[string]bool)

	for _, envVar := range envVars {
		envKeys[envVar.Key] = true
	}

	// Check for missing required keys
	for _, required := range schema.Required {
		if !envKeys[required] {
			errors = append(errors, fmt.Sprintf("Missing required key: %s", required))
		}
	}

	// Check for extra keys not in schema
	for key := range envKeys {
		if _, exists := schema.Properties[key]; !exists {
			errors = append(errors, fmt.Sprintf("Key not in schema: %s", key))
		}
	}

	return errors, nil
}
