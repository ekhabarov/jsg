package ast

import (
	"fmt"
	"strings"

	"github.com/ekhabarov/jsg/lib"
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

// GoType returns a Go type mapped to schema type, and imported package name, if
// necessary.
func GoType(st SchemaType, format StringFormat, ref string) (string, string, error) {
	switch st {
	case String:
		switch format {
		case FormatDateTime, FormatDate, FormatTime:
			return "time.Time", "time", nil
		case FormatDuration:
			return "time.Duration", "time", nil
		case FormatIPv4, FormatIPv6:
			return "net.IP", "net", nil
		case FormatRegex:
			return "regexp.Regexp", "regexp", nil
		case FormatUUID:
			return "uuid.UUID", "github.com/gofrs/uuid", nil
		default:
			return "string", "", nil
		}
	case Integer:
		return "int", "", nil
	case Number:
		return "float64", "", nil
	case Boolean:
		return "bool", "", nil
	case Array:
		return "[]interface{}", "", nil
	case Null:
		return "interface{}", "", nil
	}

	if ref != "" {
		r, err := lib.URLName(ref)
		if err != nil {
			return "", "", fmt.Errorf("malformed $ref: %w", err)
		}

		return "*" + r, "", nil
	}

	return "", "", fmt.Errorf("unsupported schema type: %s", st)
}
