package problems_test

import (
	"leetcode/problems"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Problem1", func() {
	It("should return [0,1] when nums is [2,7,11,15] and target is 9", func() {
		nums := []int{2, 7, 11, 15}
		target := 9
		result := problems.TwoSum(nums, target)
		Expect(result).To(ConsistOf(0, 1))
	})

	It("should return [1,2] when nums is [3,2,4] and target is 6", func() {
		nums := []int{3, 2, 4}
		target := 6
		result := problems.TwoSum(nums, target)
		Expect(result).To(ConsistOf(1, 2))
	})

	It("should return [0,1] when nums is [3,3] and target is 6", func() {
		nums := []int{3, 3}
		target := 6
		result := problems.TwoSum(nums, target)
		Expect(result).To(ConsistOf(0, 1))
	})

	// TODO: Add more test cases here
})
