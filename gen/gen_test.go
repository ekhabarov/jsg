package gen_test

import (
	"bytes"
	"errors"
	"io/ioutil"

	"github.com/ekhabarov/jsg/ast"
	"github.com/ekhabarov/jsg/gen"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
)

var _ = Describe("Gen", func() {

	Context("SchemaName", func() {

		DescribeTable("IDs",
			func(id, expName string, expErr error) {
				got, err := gen.SchemaName(id)

				if expErr != nil {
					Expect(err).To(MatchError(expErr))
				} else {
					Expect(err).NotTo(HaveOccurred())
				}

				Expect(got).To(Equal(expName))
			},

			Entry("", "https://example.com/name.json", "Name", nil),
			Entry("", "https://example.com/a/b/c/test.json", "Test", nil),
			Entry("", "https://example.com/a/b/c/query.json?a=1", "Query", nil),
			Entry("", "https://example.com/a/b/c/anchor.json#anchor1", "Anchor", nil),
			Entry("", "https://example.com", "", errors.New("invalid schema name")),
		)
	})

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
		)

	})

})
