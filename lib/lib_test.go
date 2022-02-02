package lib_test

import (
	"errors"

	"github.com/ekhabarov/jsg/lib"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
)

var _ = Describe("Lib", func() {

	Context("URLName", func() {

		DescribeTable("IDs",
			func(id, expName string, expErr error) {
				got, err := lib.URLName(id)

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
			Entry("", "https://example.com", "", errors.New("invalid url")),
		)
	})

})
