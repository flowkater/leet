package problems_test

import (
	"leetcode/problems"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("1. Two Sum", func() {
	It("should return indices of the two numbers that add up to target", func() {
		nums := []int{2, 7, 11, 15}
		target := 9
		result := problems.TwoSum(nums, target)
		Expect(result).To(ConsistOf(0, 1))

		nums = []int{3, 2, 4}
		target = 6
		result = problems.TwoSum(nums, target)
		Expect(result).To(ConsistOf(1, 2))

		nums = []int{3, 3}
		target = 6
		result = problems.TwoSum(nums, target)
		Expect(result).To(ConsistOf(0, 1))
	})

	// TODO: Add more test cases here
})
