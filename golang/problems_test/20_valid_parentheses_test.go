package problems_test

import (
	"leetcode/problems"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Problem20", func() {
	It("should return true when s is '()'", func() {
		s := "()"
		result := problems.IsValid(s)
		Expect(result).To(BeTrue())
	})

	It("should return true when s is '()[]{}'", func() {
		s := "()[]{}"
		result := problems.IsValid(s)
		Expect(result).To(BeTrue())
	})

	It("should return true when s is '([]){}'", func() {
		s := "([]){}"
		result := problems.IsValid(s)
		Expect(result).To(BeTrue())
	})

	It("should return true when s is '([]{})'", func() {
		s := "([]{})"
		result := problems.IsValid(s)
		Expect(result).To(BeTrue())
	})

	It("should return false when s is '(]'", func() {
		s := "(]"
		result := problems.IsValid(s)
		Expect(result).To(BeFalse())
	})

	It("should return true when s is '([])'", func() {
		s := "([])"
		result := problems.IsValid(s)
		Expect(result).To(BeTrue())
	})
})
