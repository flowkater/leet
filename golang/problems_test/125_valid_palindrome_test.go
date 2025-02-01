package problems_test

import (
	"leetcode/problems"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Problem125", func() {
	Context("Valid Palindrome", func() {
		It("Example 1: A man, a plan, a canal: Panama", func() {
			s := "A man, a plan, a canal: Panama"
			Expect(problems.IsPalindrome(s)).To(BeTrue())
		})

		It("Example 2: race a car", func() {
			s := "race a car"
			Expect(problems.IsPalindrome(s)).To(BeFalse())
		})

		It("Example 3: empty string with space", func() {
			s := " "
			Expect(problems.IsPalindrome(s)).To(BeTrue())
		})

		It("Example 4: 0P", func() {
			s := "0P"
			Expect(problems.IsPalindrome(s)).To(BeFalse())
		})

		It("Additional Case: with numbers and special characters", func() {
			s := "A1b2c,,c b@a1"
			Expect(problems.IsPalindrome(s)).To(BeFalse())
		})

		It("Additional Case: single character", func() {
			s := "a"
			Expect(problems.IsPalindrome(s)).To(BeTrue())
		})

		It("Additional Case: mixed case palindrome", func() {
			s := "RaCeCaR"
			Expect(problems.IsPalindrome(s)).To(BeTrue())
		})
	})
})
