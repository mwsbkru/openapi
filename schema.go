package openapi

import (
	"fmt"
)

// Schema -.
type Schema struct {
	Properties Schemas     `json:"properties,omitempty" yaml:"properties,omitempty"`
	Type       string      `json:"type,omitempty" yaml:"type,omitempty"`
	Format     string      `json:"format,omitempty" yaml:"format,omitempty"`
	Default    interface{} `json:"default,omitempty" yaml:"default,omitempty"`
	Example    interface{} `json:"example,omitempty" yaml:"example,omitempty"`
	Required   []string    `json:"required,omitempty" yaml:"required,omitempty"`
	Items      *Schema     `json:"items,omitempty" yaml:"items,omitempty"`
	Ref        string      `json:"$ref,omitempty" yaml:"$ref,omitempty"`
	Faker      string      `json:"x-faker,omitempty" yaml:"x-faker,omitempty"`
}

// IsRef -.
func (s Schema) IsRef() bool {
	return s.Ref != ""
}

// Schemas -.
type Schemas map[string]*Schema

// SchemaContext -.
type SchemaContext interface {
	LookupSchemaByReference(ref string) (Schema, error)
}

// ResponseByExample -.
func (s Schema) ResponseByExample(schemaContext SchemaContext) (interface{}, error) {
	if s.Ref != "" {
		schema, err := schemaContext.LookupSchemaByReference(s.Ref)
		if err != nil {
			return nil, fmt.Errorf("lookup: %w", err)
		}

		return schema.ResponseByExample(schemaContext)
	}

	if s.Example != nil {
		return ExampleToResponse(s.Example), nil
	}

	return s.propertiesExamples(schemaContext)
}

func (s Schema) propertiesExamples(schemaContext SchemaContext) (interface{}, error) {
	if s.Items != nil {
		resp, err := s.Items.ResponseByExample(schemaContext)
		if err != nil {
			return nil, fmt.Errorf("response from items: %w", err)
		}

		res := []interface{}{resp}

		return res, nil
	}

	res := make(map[string]interface{}, len(s.Properties))

	for key, prop := range s.Properties {
		propResp, err := prop.ResponseByExample(schemaContext)
		if err != nil {
			return nil, fmt.Errorf("response for property %q: %w", key, err)
		}

		res[key] = propResp
	}

	return res, nil
}
