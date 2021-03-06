package parser_test

import (
	"vsphere-cleaner/parser"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Parser", func() {
	var parserObj parser.Parser
	BeforeEach(func() {
		parserObj = parser.NewParser()
	})
	It("runs successfully", func() {
		_, err := parser.NewParser().Parse("test_lock")
		Expect(err).NotTo(HaveOccurred())
	})

	It("parses the test file", func() {
		config, err := parserObj.Parse("test_lock")
		Expect(err).NotTo(HaveOccurred())
		Expect(config.DirectorIP()).To(Equal("10.74.42.44"))
	})

	It("returns error if file does not exist", func() {
		_, err := parserObj.Parse("missing_test_lock")
		Expect(err).To(HaveOccurred())
	})

	It("returns error if the file is not a yaml", func() {
		_, err := parserObj.Parse("parser.go")
		Expect(err).To(HaveOccurred())
	})
})
