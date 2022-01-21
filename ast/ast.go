package ast

import (
	"encoding/json"
	"fmt"
	"io"
	"strings"
)

// SchemaType defines a type of the schema, where a schema can represent just
// one field or a complex document.
//
// https://json-schema.org/draft/2020-12/json-schema-validation.html#rfc.section.6.1.1
type SchemaType uint8

const (
	String SchemaType = 1 << iota
	Number
	Integer
	Object
	Array
	Boolean
	Null
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

	case string('['):
		types := v[1 : len(v)-1]

		for _, p := range strings.Split(types, ",") {
			t, err := typ(strings.TrimSpace(p))
			if err != nil {
				return err
			}

			*st |= t
		}

		return nil

	default:
		return fmt.Errorf("invalid value: %q, expected string or array", b)
	}
}

func typ(t string) (SchemaType, error) {
	switch t {
	case `"string"`:
		return String, nil
	case `"number"`:
		return Number, nil
	case `"integer"`:
		return Integer, nil
	case `"object"`:
		return Object, nil
	case `"array"`:
		return Array, nil
	case `"boolean"`:
		return Boolean, nil
	case `"null"`:
		return Null, nil
	}

	return SchemaType(0), fmt.Errorf("unsupported type: %s", t)
}

// https://json-schema.org/understanding-json-schema/reference/string.html
type StringFormat uint8

const (
	FormatDateTime StringFormat = iota + 1
	FormatTime
	FormatDate
	FormatDuration
	FormatEmail
	FormatIdnEmail
	FormatHostname
	FormatIdnHostname
	FormatIPv4
	FormatIPv6
	FormatUUID
	FormatURI
	FormatURIReference
	FormatIRI
	FormatIRIReference
	FormatURITemplate
	FormatJSONPointer
	FormatRelativeJSONPointer
	FormatRegex
)

func (sf *StringFormat) UnmarshalJSON(b []byte) error {
	f, err := format(string(b))
	if err != nil {
		return err
	}

	*sf = f

	return nil
}

func format(t string) (StringFormat, error) {
	switch t {
	case `"date-time"`:
		return FormatDateTime, nil
	case `"time"`:
		return FormatTime, nil
	case `"date"`:
		return FormatDate, nil
	case `"duration"`:
		return FormatDuration, nil
	case `"email"`:
		return FormatEmail, nil
	case `"idn-email"`:
		return FormatIdnEmail, nil
	case `"hostname"`:
		return FormatHostname, nil
	case `"idn-hostname"`:
		return FormatIdnHostname, nil
	case `"ipv4"`:
		return FormatIPv4, nil
	case `"ipv6"`:
		return FormatIPv6, nil
	case `"uuid"`:
		return FormatUUID, nil
	case `"uri"`:
		return FormatURI, nil
	case `"uri-reference"`:
		return FormatURIReference, nil
	case `"iri"`:
		return FormatIRI, nil
	case `"iri-reference"`:
		return FormatIRIReference, nil
	case `"uri-template"`:
		return FormatURITemplate, nil
	case `"json-pointer"`:
		return FormatJSONPointer, nil
	case `"relative-json-pointer"`:
		return FormatRelativeJSONPointer, nil
	case `"regex"`:
		return FormatRegex, nil
	}

	return StringFormat(0), fmt.Errorf("unsupported format: %s", t)
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

	// String related options.
	MinLength uint32       `json:"minLength"`
	MaxLength uint32       `json:"maxLength"`
	Pattern   string       `json:"pattern"`
	Format    StringFormat `json:"format"`
}

// Parse parses JSON schema into Abstract Syntax Tree.
func Parse(r io.Reader) (*Schema, error) {
	var sch Schema

	if err := json.NewDecoder(r).Decode(&sch); err != nil {
		return nil, fmt.Errorf("failed to parse schema: %w", err)
	}

	return &sch, nil
}
