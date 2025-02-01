package problems

import "unicode"

// A phrase is a palindrome if, after converting all uppercase letters into lowercase letters and removing all non-alphanumeric characters, it reads the same forward and backward. Alphanumeric characters include letters and numbers.

// Given a string s, return true if it is a palindrome, or false otherwise.

// Example 1:

// Input: s = "A man, a plan, a canal: Panama"
// Output: true
// Explanation: "amanaplanacanalpanama" is a palindrome.
// Example 2:

// Input: s = "race a car"
// Output: false
// Explanation: "raceacar" is not a palindrome.
// Example 3:

// Input: s = " "
// Output: true
// Explanation: s is an empty string "" after removing non-alphanumeric characters.
// Since an empty string reads the same forward and backward, it is a palindrome.

// Constraints:
// 1 <= s.length <= 2 * 10^5
// s consists only of printable ASCII characters.

func IsPalindrome(s string) bool {
	// re := regexp.MustCompile(`[^a-zA-Z0-9]`)
	// s = strings.ToLower(re.ReplaceAllString(s, ""))
	filtered := make([]rune, 0, len(s))
	for _, char := range s {
		if unicode.IsLetter(char) || unicode.IsDigit(char) {
			if char >= 'A' && char <= 'Z' {
				char = char + ('a' - 'A')
			}
			filtered = append(filtered, char)
		}
	}

	for i, j := 0, len(filtered)-1; i < j; i, j = i+1, j-1 {
		if filtered[i] != filtered[j] {
			return false
		}
	}

	return true
}
