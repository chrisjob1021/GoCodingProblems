// Longest Palindromic Substring
// Manacher's Algorithm
//
// Time complexity: O(n)
// Space complexity: O(n)
//
// References:
// https://leetcode.com/problems/longest-palindromic-substring/description/
// https://github.com/mission-peace/interview/blob/master/src/com/interview/string/LongestPalindromeSubstring.java
//
// Acknowledgements:
// Tushar Roy @mission-peace
// His video explaining this algorithm: https://youtu.be/V-sEwsca1ak
package longpalsubstr

// Don't want to convert to float64, so let's avoid using math.Min.
// Instead, we'll create our own min() that will return lowest value between two integers.
func min(x, y int) int {
	if x < y {
		return x
	} else {
		return y
	}
}

// Return an output similiar to Python's enumerate().
// Return two values: the current slice position and the current max value.
func enumerate(T []int) (int, int) {
	var position, max int
	for i, val := range T {
		if val > curMax {
			cur, curMax = i, val
		}
	}

	return position, max
}

func longestPalindrome(s string) string {
	// Handle case like input "a" or "".
	if len(s) <= 1 {
		return s
	}

	// We are going to preprocess the string from "abc" to "$a$b$c$" or "abcd" to "$a$b$c$d$".
	// The resulting string will be 2*N + 1.
	// This is so that our apporach works for both odd and even inputs.
	// Since strings are immutable in Go, we'll create an entirely new string.
	newStr, T := make([]string, len(s)*2+1), make([]int, len(s)*2+1)
	orgStrCount := 0

	for i := 0; i < len(newStr); i++ {
		if i%2 == 0 {
			newStr[i] = "$"
		} else {
			newStr[i] = string(s[orgStrCount])
			orgStrCount++
		}
	}

	start, end, i, newCenter := 0, 0, 0, 0

	// We are evaluating four cases when we're picking a new center:
	// 1. Bad - Current position is entirely contained within the current palindrome.
	// 2. Bad - Current position is at the end of the input. We should break the loop in this case.
	// 3. Good - Current position is proper suffix of current middle's palindrome and its left side mirror is a proper prefix.
	// 4. Bad - Current position is proper suffix of current palindrome, but its left side mirror extends beyond the left side of the current middle's palindrome. Selecting this as the new middle would not extend at all.

	// i = current center
	for i < len(newStr) {
		// Start and end positions represent length of current palindrome at center "i".
		for start > 0 && end < len(newStr)-1 && newStr[start-1] == newStr[end+1] {
			start--
			end++
		}

		// If end - start = 0, then +1 gives us a floor of 1.
		// If end - start > 0, then +1 accounts for a floor of 3.
		T[i] = end - start + 1

		// Case 2, as mentioned above.
		if end == len(newStr)-1 {
			break
		}

		// Set possible new center based on start + end.
		// If even, add 1 to move off of '$' (our preprocess character).
		// If odd, add 0 since we're already on an original character.
		if i%2 == 0 {
			newCenter = end + 1
		} else {
			newCenter = end
		}

		// Is there a better center?
		// Evaluate as we mirror the left side of the current middle in our solution array.
		for j := i + 1; j <= end; j++ {
			// Pick either left side T[i-(j-i)] of current middle "i" or "j" to the end.
			// (end-j)*2+1 This handles case where left side mirror extends beyond right edge of current palindrome
			T[j] = min(T[i-(j-i)], (end-j)*2+1)

			// Evaluate the following criteria:
			// Does point to potential new enter (j) plus half of the left side mirror (we're already accounting for the other half in "j") equal the end of current palindrome?
			// If yes, then we've found case 3 mentioned above.
			if j+T[i-(j-i)]/2 == end {
				newCenter = j
				break
			}
		}

		// Set "i" as the new center.
		i = newCenter

		// Move backward from "i" based on length of current center to determine "start".
		start = i - T[i]/2
		// Move backward from "i" based on length of current center to determine "end".
		end = i + T[i]/2
	}

	centerIndex, maxLen := enumerate(T)

	// centerIndex/2
	// Remember that T is based on preprocessed string.
	// We need to halve the index to get the actual position in the original input string.

	// maxLen/2/2
	// We then we want to move back half of the length.
	// Again, need to halve it to get the actual length of the original input string.
	start = centerIndex/2 - maxLen/2/2

	// Substring from start to start plus length of result (again halved to convert to size of input string).
	return s[start : start+maxLen/2]

}
