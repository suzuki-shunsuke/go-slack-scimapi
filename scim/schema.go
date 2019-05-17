package scim

import (
	"encoding/json"
)

type (
	Schema struct {
		ID          string      `json:"id"`
		Name        string      `json:"name"`
		Description string      `json:"description"`
		Schema      []string    `json:"schema"`
		Endpoint    string      `json:"endpoint"`
		Attributes  []Attribute `json:"attributes"`
	}

	Attribute struct {
		MultiValued                   bool        `json:"multiValued"`
		ReadOnly                      bool        `json:"readOnly"`
		Required                      bool        `json:"required"`
		CaseExact                     bool        `json:"caseExact"`
		Name                          string      `json:"name"`
		Type                          string      `json:"type"`
		Description                   string      `json:"description"`
		Schema                        string      `json:"schema"`
		MultiValuedAttributeChildName string      `json:"multiValuedAttributeChildName,omitempty"`
		SubAttributes                 []Attribute `json:"subAttributes,omitempty"`
		CanonicalValues               []string    `json:"canonicalValues,omitempty"`
	}

	Meta struct {
		Created      string   `json:"created"`
		LastModified string   `json:"lastModified"`
		Location     string   `json:"location"`
		Version      string   `json:"version"`
		Attributes   []string `json:"attributes"`
	}
)

func (schema *Schema) UnmarshalJSON(b []byte) error {
	type alias Schema
	a := struct {
		Schema json.RawMessage `json:"schema"`
		*alias
	}{
		alias: (*alias)(schema),
	}
	if err := json.Unmarshal(b, &a); err != nil {
		return err
	}
	if len(a.Schema) == 0 {
		return nil
	}

	list := []string{}
	err := json.Unmarshal(a.Schema, &list)
	if err == nil {
		schema.Schema = list
		return nil
	}
	child := ""
	if err := json.Unmarshal(a.Schema, &child); err == nil {
		schema.Schema = []string{child}
		return nil
	}
	return err
}

func (attr *Attribute) UnmarshalJSON(b []byte) error {
	type alias Attribute
	a := struct {
		SubAttributes json.RawMessage `json:"subAttributes"`
		*alias
	}{
		alias: (*alias)(attr),
	}
	if err := json.Unmarshal(b, &a); err != nil {
		return err
	}
	if len(a.SubAttributes) == 0 {
		return nil
	}

	list := []Attribute{}
	err := json.Unmarshal(a.SubAttributes, &list)
	if err == nil {
		attr.SubAttributes = list
		return nil
	}
	child := Attribute{}
	if err := json.Unmarshal(a.SubAttributes, &child); err == nil {
		attr.SubAttributes = []Attribute{child}
		return nil
	}
	return err
}
