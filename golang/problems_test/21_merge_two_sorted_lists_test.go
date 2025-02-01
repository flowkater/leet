package problems_test

import (
	"leetcode/problems"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Problem21", func() {
	Context("Merge Two Sorted Lists", func() {
		It("Example 1", func() {
			list1 := createLinkedList([]int{1, 2, 4})
			list2 := createLinkedList([]int{1, 3, 4})
			Expect(linkedListToSlice(problems.MergeTwoLists(list1, list2))).To(Equal([]int{1, 1, 2, 3, 4, 4}))
		})

		It("Example 2", func() {
			list1 := createLinkedList([]int{})
			list2 := createLinkedList([]int{})
			Expect(linkedListToSlice(problems.MergeTwoLists(list1, list2))).To(BeEmpty())
		})

		It("Example 3", func() {
			list1 := createLinkedList([]int{})
			list2 := createLinkedList([]int{0})
			Expect(linkedListToSlice(problems.MergeTwoLists(list1, list2))).To(Equal([]int{0}))
		})

		It("Additional Case: one empty list", func() {
			list1 := createLinkedList([]int{1, 3, 5})
			list2 := createLinkedList([]int{})
			Expect(linkedListToSlice(problems.MergeTwoLists(list1, list2))).To(Equal([]int{1, 3, 5}))
		})
	})
})

func createLinkedList(nums []int) *problems.ListNode {
	if len(nums) == 0 {
		return nil
	}

	head := &problems.ListNode{Val: nums[0]}
	current := head
	for _, num := range nums[1:] {
		current.Next = &problems.ListNode{Val: num}
		current = current.Next
	}
	return head
}

func linkedListToSlice(head *problems.ListNode) []int {
	var result []int
	current := head
	for current != nil {
		result = append(result, current.Val)
		current = current.Next
	}
	return result
}
