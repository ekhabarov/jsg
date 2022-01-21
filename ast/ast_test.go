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

				Expect(*schema).To(MatchAllFields(fields))
			},

			Entry("", `{ "type": "string"}`, Fields{"Type": Equal(ast.String)}),
			Entry("", `{ "type": "number"}`, Fields{"Type": Equal(ast.Number)}),
			Entry("", `{ "type": "integer"}`, Fields{"Type": Equal(ast.Integer)}),
			Entry("", `{ "type": "object"}`, Fields{"Type": Equal(ast.Object)}),
			Entry("", `{ "type": "array"}`, Fields{"Type": Equal(ast.Array)}),
			Entry("", `{ "type": "boolean"}`, Fields{"Type": Equal(ast.Boolean)}),
			Entry("", `{ "type": "null"}`, Fields{"Type": Equal(ast.Null)}),
		)
	})

})
