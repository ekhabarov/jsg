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
	// 6.1.1. type
	//
	// The value of this keyword MUST be either a string or an array. If it is an
	// array, elements of the array MUST be strings and MUST be unique. String
	// values MUST be one of the six primitive types ("null", "boolean",
	// "object", "array", "number", or "string"), or "integer" which matches any
	// number with a zero fractional part. An instance validates if and only if
	// the instance is in any of the sets listed for this keyword.
	//
	// https://json-schema.org/draft/2020-12/json-schema-validation.html#rfc.section.6.1.1
	Type SchemaType `json:"type"`

	// 6.2. Validation Keywords for Numeric Instances (number and integer)

	// 6.2.1. multipleOf
	//
	// The value of "multipleOf" MUST be a number, strictly greater than 0. A
	// numeric instance is valid only if division by this keyword's value results
	// in an integer.
	//
	// https://json-schema.org/draft/2020-12/json-schema-validation.html#rfc.section.6.2.1
	MultipleOf float64 `json:"multipleOf"`

	// 6.2.2. maximum
	//
	// The value of "maximum" MUST be a number, representing an inclusive upper
	// limit for a numeric instance. If the instance is a number, then this
	// keyword validates only if the instance is less than or exactly equal to
	// "maximum".
	//
	// https://json-schema.org/draft/2020-12/json-schema-validation.html#rfc.section.6.2.2
	Maximum float64 `json:"maximum"`

	// 6.2.3. exclusiveMaximum
	//
	// The value of "exclusiveMaximum" MUST be a number, representing an
	// exclusive upper limit for a numeric instance. If the instance is a number,
	// then the instance is valid only if it has a value strictly less than (not
	// equal to) "exclusiveMaximum".
	//
	// https://json-schema.org/draft/2020-12/json-schema-validation.html#rfc.section.6.2.3
	ExclusiveMaximum float64 `json:"exclusiveMaximum"`

	// 6.2.4. minimum
	//
	// The value of "minimum" MUST be a number, representing an inclusive lower
	// limit for a numeric instance. If the instance is a number, then this
	// keyword validates only if the instance is greater than or exactly equal to
	// "minimum".
	//
	// https://json-schema.org/draft/2020-12/json-schema-validation.html#rfc.section.6.2.4
	Minimum float64 `json:"minimum"`

	// 6.2.5. exclusiveMinimum
	//
	// The value of "exclusiveMinimum" MUST be a number, representing an
	// exclusive lower limit for a numeric instance. If the instance is a number,
	// then the instance is valid only if it has a value strictly greater than
	// (not equal to) "exclusiveMinimum".
	//
	// https://json-schema.org/draft/2020-12/json-schema-validation.html#rfc.section.6.2.5
	ExclusiveMinimum float64 `json:"exclusiveMinimum"`

	// 6.3. Validation Keywords for Strings

	// 6.3.1. maxLength
	//
	// The value of this keyword MUST be a non-negative integer. A string
	// instance is valid against this keyword if its length is less than, or
	// equal to, the value of this keyword. The length of a string instance is
	// defined as the number of its characters as defined by RFC 8259.
	//
	// https://json-schema.org/draft/2020-12/json-schema-validation.html#rfc.section.6.3.1
	MaxLength uint32 `json:"maxLength"`

	// 6.3.2. minLength
	//
	// The value of this keyword MUST be a non-negative integer. A string
	// instance is valid against this keyword if its length is greater than, or
	// equal to, the value of this keyword. The length of a string instance is
	// defined as the number of its characters as defined by RFC 8259. Omitting
	// this keyword has the same behavior as a value of 0.
	//
	// https://json-schema.org/draft/2020-12/json-schema-validation.html#rfc.section.6.3.2
	MinLength uint32 `json:"minLength"`

	// 6.3.3. pattern
	//
	// The value of this keyword MUST be a string. This string SHOULD be a valid
	// regular expression, according to the ECMA-262 regular expression dialect.
	// A string instance is considered valid if the regular expression matches
	// the instance successfully. Recall: regular expressions are not implicitly
	// anchored.
	//
	// https://json-schema.org/draft/2020-12/json-schema-validation.html#rfc.section.6.3.3
	Pattern string `json:"pattern"`

	// https://json-schema.org/draft/2020-12/json-schema-validation.html#rfc.section.7.3
	Format StringFormat `json:"format"`

	// 10.3.2. Keywords for Applying Subschemas to Objects

	// 10.3.2.1. properties
	//
	// The value of "properties" MUST be an object. Each value of this object
	// MUST be a valid JSON Schema. Validation succeeds if, for each name that
	// appears in both the instance and as a name within this keyword's value,
	// the child instance for that name successfully validates against the
	// corresponding schema. The annotation result of this keyword is the set of
	// instance property names matched by this keyword. Omitting this keyword has
	// the same assertion behavior as an empty object.
	//
	// https://json-schema.org/draft/2020-12/json-schema-core.html#rfc.section.10.3.2.1
	Properties map[string]Schema `json:"properties"`
}

// Parse parses JSON schema into Abstract Syntax Tree.
func Parse(r io.Reader) (*Schema, error) {
	var sch Schema

	if err := json.NewDecoder(r).Decode(&sch); err != nil {
		return nil, fmt.Errorf("failed to parse schema: %w", err)
	}

	return &sch, nil
}
