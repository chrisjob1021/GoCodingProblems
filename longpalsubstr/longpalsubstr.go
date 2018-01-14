// longest palindromic substring
// manacher's algorithm
//
// time complexity: O(n)
// space complexity: O(n)
//
// references
// https://leetcode.com/problems/longest-palindromic-substring/description/
// https://github.com/mission-peace/interview/blob/master/src/com/interview/string/LongestPalindromeSubstring.java
//
// acknowledgements: Tushar Roy @mission-peace
// his video explaining this algorithm: https://youtu.be/V-sEwsca1ak
package longpalsubstr

// don't want to convert to float64, so let's avoid using math.Min
// instead, we'll create our own min() that will return lowest value between 2 integers
func min(x, y int) int {
	if x < y {
		return x
	} else {
		return y
	}
}

// return an output similiar to python's enumerate()
// return 2 values, the current slice position and the current max value
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
	// handle case like input "a" or ""
	if len(s) <= 1 {
		return s
	}

	// we are going to preprocess the string from "abc" to "$a$b$c$" or "abcd" to "$a$b$c$d$"
	// the resulting string will be 2*N + 1
	// this is so that our apporach works for both odd and even inputs
	// since strings are immutable in Go, we'll create an entirely new string
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

	// we are evaluating four cases when we're picking a new middle on our path to find the longest palindromic substring
	// 1. bad - current position is entirely contained within the current palindrome
	// 2. bad - current position is at the end of the input. we should break the loop.
	// 3. good - current position is proper suffix of current middle's palindrome and its left side mirror is a proper prefix
	// 4. bad - current position is proper suffix of current palindrome, but its left side mirror extends beyond the left side of the current middle's palindrome. selecting this as the new middle would not extend at all.

	// i = current center
	for i < len(newStr) {
		// start and end positions to length of current palindrome at center i
		for start > 0 && end < len(newStr)-1 && newStr[start-1] == newStr[end+1] {
			start--
			end++
		}

		// if end - start = 0, then +1 gives us a floor of 1
		// if end - start > 0, then +1 accounts for a floor of 3
		T[i] = end - start + 1

		// case 2 mentioned above
		if end == len(newStr)-1 {
			break
		}

		// set possible new center based on start + end
		// if even, add 1 to move off of '$' (our preprocess character)
		// if odd, add 0 since we're already on an original character
		if i%2 == 0 {
			newCenter = end + 1
		} else {
			newCenter = end
		}

		// is there a better center?
		// evaluate as we mirror the left side of the current middle
		for j := i + 1; j <= end; j++ {
			// pick either left side T[i-(j-i)] of current middle i
			// or j to the end (handles case where left side mirror extends beyond right edge of current palindrome
			T[j] = min(T[i-(j-i)], (end-j)*2+1)

			// evaluate the following criteria:
			//    does point to potential new enter (j)
			//    plus half of the left side mirror (we're already accounting for the other half in j)
			//    equal the end of current palindrome?
			// if yes, then we've found case 3 mentioned above
			if j+T[i-(j-i)]/2 == end {
				newCenter = j
				break
			}
		}

		// set i as the new center
		i = newCenter

		// start = move backward from i based on length of current center
		start = i - T[i]/2
		// end = move backward from i based on length of current center
		end = i + T[i]/2
	}

	centerIndex, maxLen := enumerate(T)

	// centerIndex/2 =
	// remember that T is based on preprocessed index
	// so, we need to halve the index returned to get the actual position in the input string

	// maxLen/2/2 =
	// we then we want to move back half of the length of the result
	// but again need to halve it to get the actual length of the input string
	start = centerIndex/2 - maxLen/2/2

	// substring from start to start plus length of result (again halved to convert to size of input string)
	return s[start : start+maxLen/2]

}
