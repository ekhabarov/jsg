package gen_test

import (
	"bytes"
	"io/ioutil"

	"github.com/ekhabarov/jsg/ast"
	"github.com/ekhabarov/jsg/gen"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
)

var _ = Describe("Gen", func() {

	Context("Generate", func() {

		DescribeTable("Call",
			func(schema ast.Schema, expFile string) {
				w := bytes.NewBuffer([]byte{})
				err := gen.Generate(w, &schema)
				Expect(err).NotTo(HaveOccurred())

				data, err := ioutil.ReadFile("./testdata/" + expFile + ".golden")
				Expect(err).NotTo(HaveOccurred())

				Expect(w.String()).To(Equal(string(data)))
			},

			Entry("Schema without ID", ast.Schema{}, "nil_schema.go"),

			Entry("Simple model", ast.Schema{
				ID: "https://example.com/model.json",
				Properties: map[string]ast.Schema{
					"String":  {Type: ast.String},
					"Integer": {Type: ast.Integer},
					"Number":  {Type: ast.Number},
					"Boolean": {Type: ast.Boolean},
					"Array":   {Type: ast.Array},
					"Null":    {Type: ast.Null},
				},
			}, "model_struct.go"),

			Entry("String formats", ast.Schema{
				ID: "https://example.com/string_formats.json",
				Properties: map[string]ast.Schema{
					"NoFormat":            {Type: ast.String},
					"DateTime":            {Type: ast.String, Format: ast.FormatDateTime},
					"Time":                {Type: ast.String, Format: ast.FormatTime},
					"Date":                {Type: ast.String, Format: ast.FormatDate},
					"Duration":            {Type: ast.String, Format: ast.FormatDuration},
					"Email":               {Type: ast.String, Format: ast.FormatEmail},
					"IdnEmail":            {Type: ast.String, Format: ast.FormatIdnEmail},
					"Hostname":            {Type: ast.String, Format: ast.FormatHostname},
					"IdnHostname":         {Type: ast.String, Format: ast.FormatIdnHostname},
					"IPv4":                {Type: ast.String, Format: ast.FormatIPv4},
					"IPv6":                {Type: ast.String, Format: ast.FormatIPv6},
					"UUID":                {Type: ast.String, Format: ast.FormatUUID},
					"URI":                 {Type: ast.String, Format: ast.FormatURI},
					"URIReference":        {Type: ast.String, Format: ast.FormatURIReference},
					"IRI":                 {Type: ast.String, Format: ast.FormatIRI},
					"IRIReference":        {Type: ast.String, Format: ast.FormatIRIReference},
					"JSONPointer":         {Type: ast.String, Format: ast.FormatJSONPointer},
					"RelativeJSONPointer": {Type: ast.String, Format: ast.FormatRelativeJSONPointer},
					"Regexp":              {Type: ast.String, Format: ast.FormatRegex},
				},
			}, "string_formats.go"),

			Entry("Struct with ref", ast.Schema{
				ID: "https://example.com/ref.json",
				Properties: map[string]ast.Schema{
					"String": {Type: ast.String},
					"Sub":    {Ref: "https://example.com/inner.json"},
				},
			}, "struct_with_ref.go"),
		)

	})

})
