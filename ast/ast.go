package ast

import (
	"encoding/json"
	"fmt"
	"io"
)

// SchemaType defines a type of the schema, where a schema can represent just
// one field or a complex document.
type SchemaType uint8

const (
	Number SchemaType = 1 << iota
)

func (st *SchemaType) UnmarshalJSON(b []byte) error {
	switch v := string(b); v[:1] {
	case string('"'):
		t, err := typ(v)
		if err != nil {
			return err
		}

		*st = t

		return nil

	default:
		return fmt.Errorf("invalid value: %q, expected string or array", b)
	}
}

func typ(t string) (SchemaType, error) {
	switch t {
	case `"number"`:
		return Number, nil
	}

	return SchemaType(0), fmt.Errorf("unsupported type: %s", t)
}

// Schema is an Abstract Syntax Tree (AST) representation of JSON schema.
type Schema struct {
	// The type keyword is fundamental to JSON Schema. It specifies the data type
	// for a schema.
	//
	// The type keyword may either be a string or an array:
	//
	// * If it’s a string, it is the name of one of the basic types above.
	// * If it is an array, it must be an array of strings, where each string is
	//   the name of one of the basic types, and each element is unique. In this
	//   case, the JSON snippet is valid if it matches any of the given types.
	Type SchemaType `json:"type"`
}

// Parse parses JSON schema into Abstract Syntax Tree.
func Parse(r io.Reader) (*Schema, error) {
	var sch Schema

	if err := json.NewDecoder(r).Decode(&sch); err != nil {
		return nil, fmt.Errorf("failed to parse schema: %w", err)
	}

	return &sch, nil
}
