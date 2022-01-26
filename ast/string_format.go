package ast

import "fmt"

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
