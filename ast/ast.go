package ast

import (
	"encoding/json"
	"fmt"
	"io"
)

// Schema is an Abstract Syntax Tree (AST) representation of JSON schema.
type Schema struct {

	//   8.2.1. The "$id" Keyword
	//
	// The "$id" keyword identifies a schema resource with its canonical URI.
	//
	// Note that this URI is an identifier and not necessarily a network locator.
	// In the case of a network-addressable URL, a schema need not be
	// downloadable from its canonical URI.
	//
	// If present, the value for this keyword MUST be a string, and MUST
	// represent a valid URI-reference. This URI-reference SHOULD be normalized,
	// and MUST resolve to an absolute-URI (without a fragment). Therefore, "$id"
	// MUST NOT contain a non-empty fragment, and SHOULD NOT contain an empty
	// fragment.
	//
	// Since an empty fragment in the context of the application/schema+json
	// media type refers to the same resource as the base URI without a fragment,
	// an implementation MAY normalize a URI ending with an empty fragment by
	// removing the fragment. However, schema authors SHOULD NOT rely on this
	// behavior across implementations. [CREF3]
	//
	// This URI also serves as the base URI for relative URI-references in
	// keywords within the schema resource, in accordance with RFC 3986 section
	// 5.1.1 regarding base URIs embedded in content.
	//
	// The presence of "$id" in a subschema indicates that the subschema
	// constitutes a distinct schema resource within a single schema document.
	// Furthermore, in accordance with RFC 3986 section 5.1.2 regarding
	// encapsulating entities, if an "$id" in a subschema is a relative
	// URI-reference, the base URI for resolving that reference is the URI of the
	// parent schema resource.
	//
	// If no parent schema object explicitly identifies itself as a resource with
	// "$id", the base URI is that of the entire document, as established by the
	// steps given in the previous section.
	ID string `json:"$id"`

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

	// Order inside Properties.
	Ord uint8
}

// Parse parses JSON schema into Abstract Syntax Tree.
func Parse(r io.Reader) (*Schema, error) {
	var sch Schema

	// TODO(ekhabarov): Add an order to properties.
	if err := json.NewDecoder(r).Decode(&sch); err != nil {
		return nil, fmt.Errorf("failed to parse schema: %w", err)
	}

	return &sch, nil
}
