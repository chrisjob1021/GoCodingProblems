// Longest Palindromic Subsequence
// Dynamic Programming
//
// Time complexity: O(n2)
// Space complexity: O(n2)
//
// References:
// https://leetcode.com/problems/longest-palindromic-subsequence/description/
// https://github.com/mission-peace/interview/blob/master/src/com/interview/dynamic/LongestPalindromicSubsequence.java
//
// Acknowledgements:
// Tushar Roy @mission-peace
// His video explaining this algorithm: https://youtu.be/_nCsPn7_OgI
package longpalsubseq

import (
	"fmt"
)

// I wish to avoid converting our integers to float64, just for the sake of using math.Max.
// Instead, let's create a simple helper function to return the max of two integers.
func max(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}

func printLongestPalindromeSubseq(s string, T [][]int) string {
	// Create result slice of bytes that represent the length of the longest palindromic subsequence.
	// In each slice, we will store the bytes representing the character at each position.
	res := make([]byte, T[0][len(s)-1])

	// The length of the longest palindromic subsequence is stored at T[0][N], where N is the length of the input string. We will start there and work backwards.
	i := 0
	j := T[0][len(s)-1]

	// Let's also setup two pointers that represent the left and right side of our resulting string.
	// In most cases, we're going to add strings to both sides simultaneously.
	// "l" represents the left side and "r" represents the right side.
	l := 0
	r := len(res) - 1

	for i <= j {
		// This handles the case where the original string's slice increased the size of the longest palindromic substring.
		// We verify that the characters at the start and end of the slice match.
		if s[i] == s[j] && T[i][j] == T[i+1][j-1]+2 {
			res[l] = s[i]
			// Move l toward the right, since we added an item to the result on the left side.
			l++
			res[r] = s[j]
			// Move r toward the left, since we added an item to the result on the right side.
			r--

			i++
			j--
		} else if T[i][j] == T[i+1][j] {
			// If T[i][j] equals the value of T[i+1][j], then we can assume that the longest palindromic subsequence was contained in the slice represented by T[i+1][j].

			// Add the left side of slice represented by T[i+1][j]
			res[l] = s[i+1]
			// Move l toward the right, since we added an item to the result on the left side.
			l++

			// If T[i+1][j] is greater than 1, then we have to add the left and right side of its represented slice.
			if T[i+1][j] > 1 {
				res[r] = s[j]

				// Move r toward the left, since we added an item to the result on the right side.
				r--
			}

			// Backtrack from T[i][j] to where our solution came from (T[i+1][j])
			i++
		} else if T[i][j] == T[i][j-1] {
			// If T[i][j] equals the value of T[i][j-1], then we can assume that the longest palindromic subsequence was contained in the slice represented by T[i][j-1].

			// If T[i][j-1] is greater than 1, then we have to add the left and right side of its represented slice.
			if T[i][j-1] > 0 {
				res[l] = s[i]
				// Move l toward the right, since we added an item to the result on the left side.
				l++
			}

			// Add the right side of slice represented by T[i][j-1]
			res[r] = s[j-1]
			// Move r toward the left, since we added an item to the result on the right side.
			r--

			// Backtrack from T[i][j] to where our solution came from (T[i][j-1])
			j--
		} else if i == j {
			// We need to handle a case where the longest palindromic subsequence is of odd length.
			// For example "abdba": up to this point, we will have added "ab ba" to our resulting string.
			// We would need to add the single character represented at T[i][j].
			res[len(res)/2] = s[i]
			break
		}
	}

	return fmt.Sprintf("%s", res)

}

func longestPalindromeSubseq(s string) (int, [][]int) {
	// We are going to build a N*N 2D matrix that represents the longest palindromic subsequence for every slice of the input string.
	// N represents the length of the input string.
	T := make([][]int, len(s))
	for i, _ := range T {
		T[i] = make([]int, len(s))
		// While we are creating the result matrix, let's also account for slice of length 1.
		// For example, if input is "bbbab", longest palindromic substring at s[0], s[1], s[2] and so on are all 1.
		T[i][i] = 1
	}

	// "length" represents the length of the slice of the input string we will review.
	// Start at 2, because we already took care of length 1. (Iterate up to length-1 because we've already taken care of length 1.)
	for length := 2; length <= len(s); length++ {
		// Now let's iterate through slices of input string of equal size to length
		for i := 0; i <= len(s)-length; i++ {
			// Example: "ab", "i" = s[0] ("a") and "j" = s[1] ("b")
			// Subtract 1 since resulting array is indexed starting at 0.
			j := i + length - 1
			if length == 2 {
				// Consider example of "aa"
				// A slice of s[0..1] has a palindrome of "aa" or length 2.
				if s[i] == s[j] {
					T[i][j] = 2
				} else {
					// Now consider above example of "ab"
					// A slice of s[0..1] has a palindrome of "a" or "b", equaling length 1.
					T[i][j] = 1
				}
			} else if s[i] == s[j] {
				// Consider a length 3 slice of "abad".
				// If we start with s[0..2] or "aba", the resulting palindrome would be length 3.
				// We can reach that value by adding 2 to the result of s[1..2] ("ba").
				T[i][j] = T[i+1][j-1] + 2
			} else {
				// Consider a length 3 slice of "adbb".
				// Let's move forward to the 2nd of the length 3 slices in this example.
				// Evaluating "dbb", the longest palindrome is of length 2 in that slice.
				// To reach that value, we take either the max of the palindromes at s[1..2] ("db") or s[2..3] ("bb").
				// The first is of length 1 and the latter is of length 2.
				T[i][j] = max(T[i+1][j], T[i][j-1])
			}
		}
	}

	// Returning T so that we can pass it to our function to print the longest palindromic subsequence.
	return T[0][len(s)-1], T
}
