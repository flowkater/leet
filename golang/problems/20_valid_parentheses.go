package problems

// Given a string s containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.

// An input string is valid if:

// Open brackets must be closed by the same type of brackets.
// Open brackets must be closed in the correct order.
// Every close bracket has a corresponding open bracket of the same type.

// Example 1:
// Input: s = "()"
// Output: true

// Example 2:
// Input: s = "()[]{}"
// Output: true

// Example 3:
// Input: s = "(]"
// Output: false

// Example 4:
// Input: s = "([])"
// Output: true

// Constraints:

// 1 <= s.length <= 104
// s consists of parentheses only '()[]{}'.

// Stack 문제
func IsValid(s string) bool {
	stack, m := []rune{}, map[rune]rune{'(': ')', '[': ']', '{': '}'}

	for _, ch := range s {
		if closing, isOpen := m[ch]; isOpen {
			stack = append(stack, closing)
		} else if len(stack) == 0 || stack[len(stack)-1] != ch {
			return false
		} else {
			stack = stack[:len(stack)-1]
		}
	}

	return len(stack) == 0
}
