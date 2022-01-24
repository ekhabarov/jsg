package ast_test

import (
	"strings"

	"github.com/ekhabarov/jsg/ast"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
	. "github.com/onsi/gomega/gstruct"
)

var _ = Describe("Ast", func() {

	Context("Parse", func() {

		DescribeTable("SchemaType",
			func(data string, fields Fields) {
				schema, err := ast.Parse(strings.NewReader(data))
				Expect(err).NotTo(HaveOccurred())

				Expect(*schema).To(MatchFields(IgnoreExtras, fields))
			},

			// Numbers

			// Number props

			Entry("", `{
				"type": "number",
				"multipleOf": 10,
				"maximum" : 100,
				"exclusiveMaximum": 101,
				"minimum": 50,
				"exclusiveMinimum": 49
			}`, Fields{
				"Type":             Equal(ast.Number),
				"MultipleOf":       Equal(10.0),
				"Maximum":          Equal(100.0),
				"ExclusiveMaximum": Equal(101.0),
				"Minimum":          Equal(50.0),
				"ExclusiveMinimum": Equal(49.0),
			}),

			Entry("", `{"type": "integer"}`, Fields{"Type": Equal(ast.Integer)}),

			// String type

			// String length

			Entry("String: length", `{"type": "string", "minLength": 3, "maxLength": 5}`, Fields{
				"Type":      Equal(ast.String),
				"MinLength": Equal(uint32(3)),
				"MaxLength": Equal(uint32(5)),
			}),

			// String pattern

			Entry("String: pattern", `{"type": "string", "pattern": "[A-Za-z0-9]"}`, Fields{
				"Type":    Equal(ast.String),
				"Pattern": Equal("[A-Za-z0-9]"),
			}),

			// String Format

			Entry("String: format date-time", `{"type": "string", "format": "date-time"}`, Fields{
				"Type":   Equal(ast.String),
				"Format": Equal(ast.FormatDateTime),
			}),

			Entry("String: format time", `{"type": "string", "format": "time"}`, Fields{
				"Type":   Equal(ast.String),
				"Format": Equal(ast.FormatTime),
			}),

			Entry("String: format date", `{"type": "string", "format": "date"}`, Fields{
				"Type":   Equal(ast.String),
				"Format": Equal(ast.FormatDate),
			}),

			Entry("String: format duration", `{"type": "string", "format": "duration"}`, Fields{
				"Type":   Equal(ast.String),
				"Format": Equal(ast.FormatDuration),
			}),

			Entry("String: format email", `{"type": "string", "format": "email"}`, Fields{
				"Type":   Equal(ast.String),
				"Format": Equal(ast.FormatEmail),
			}),

			Entry("String: format idn-email", `{"type": "string", "format": "idn-email"}`, Fields{
				"Type":   Equal(ast.String),
				"Format": Equal(ast.FormatIdnEmail),
			}),

			Entry("String: format hostname", `{"type": "string", "format": "hostname"}`, Fields{
				"Type":   Equal(ast.String),
				"Format": Equal(ast.FormatHostname),
			}),

			Entry("String: format idn-hostname", `{"type": "string", "format": "idn-hostname"}`, Fields{
				"Type":   Equal(ast.String),
				"Format": Equal(ast.FormatIdnHostname),
			}),

			Entry("String: format ipv4", `{"type": "string", "format": "ipv4"}`, Fields{
				"Type":   Equal(ast.String),
				"Format": Equal(ast.FormatIPv4),
			}),

			Entry("String: format ipv6", `{"type": "string", "format": "ipv6"}`, Fields{
				"Type":   Equal(ast.String),
				"Format": Equal(ast.FormatIPv6),
			}),

			Entry("String: format uuid", `{"type": "string", "format": "uuid"}`, Fields{
				"Type":   Equal(ast.String),
				"Format": Equal(ast.FormatUUID),
			}),

			Entry("String: format uri", `{"type": "string", "format": "uri"}`, Fields{
				"Type":   Equal(ast.String),
				"Format": Equal(ast.FormatURI),
			}),

			Entry("String: format uri-reference", `{"type": "string", "format": "uri-reference"}`, Fields{
				"Type":   Equal(ast.String),
				"Format": Equal(ast.FormatURIReference),
			}),

			Entry("String: format iri", `{"type": "string", "format": "iri"}`, Fields{
				"Type":   Equal(ast.String),
				"Format": Equal(ast.FormatIRI),
			}),

			Entry("String: format iri-reference", `{"type": "string", "format": "iri-reference"}`, Fields{
				"Type":   Equal(ast.String),
				"Format": Equal(ast.FormatIRIReference),
			}),

			Entry("String: format uri-template", `{"type": "string", "format": "uri-template"}`, Fields{
				"Type":   Equal(ast.String),
				"Format": Equal(ast.FormatURITemplate),
			}),

			Entry("String: format json-pointer", `{"type": "string", "format": "json-pointer"}`, Fields{
				"Type":   Equal(ast.String),
				"Format": Equal(ast.FormatJSONPointer),
			}),

			Entry("String: format relative-json-pointer", `{"type": "string", "format": "relative-json-pointer"}`, Fields{
				"Type":   Equal(ast.String),
				"Format": Equal(ast.FormatRelativeJSONPointer),
			}),

			Entry("String: format regex", `{"type": "string", "format": "regex"}`, Fields{
				"Type":   Equal(ast.String),
				"Format": Equal(ast.FormatRegex),
			}),

			// Object

			Entry("", `{"type": "object"}`, Fields{"Type": Equal(ast.Object)}),

			Entry("", `{
				"type": "object",
				"properties": {
					"s": {"type": "string"},
					"i": {"type": "integer"},
					"n": {"type": "number"},
					"b": {"type": "boolean"},
					"a": {"type": "array"}
				}
			}`, Fields{
				"Type": Equal(ast.Object),
				"Properties": MatchAllKeys(Keys{
					"s": MatchFields(IgnoreExtras, Fields{"Type": Equal(ast.String)}),
					"i": MatchFields(IgnoreExtras, Fields{"Type": Equal(ast.Integer)}),
					"n": MatchFields(IgnoreExtras, Fields{"Type": Equal(ast.Number)}),
					"b": MatchFields(IgnoreExtras, Fields{"Type": Equal(ast.Boolean)}),
					"a": MatchFields(IgnoreExtras, Fields{"Type": Equal(ast.Array)}),
				}),
			}),

			// Array
			Entry("", `{"type": "array"}`, Fields{"Type": Equal(ast.Array)}),
			Entry("", `{"type": "boolean"}`, Fields{"Type": Equal(ast.Boolean)}),
			Entry("", `{"type": "null"}`, Fields{"Type": Equal(ast.Null)}),

			Entry("", `{"type": ["string", "number", "boolean"]}`, Fields{
				"Type": Equal(ast.String | ast.Number | ast.Boolean),
			}),
		)
	})

})
