// Longest Common Substring
// Dynamic Programming
//
// Time complexity: O(len(s1) * len(s2))
// Space complexity: O(len(s1) * len(s2))
//
// References:
// https://www.geeksforgeeks.org/longest-common-substring/
// https://github.com/mission-peace/interview/blob/master/src/com/interview/dynamic/LongestCommonSubstring.java
//
// Acknowledgements:
// Tushar Roy @mission-peace
// His video explaining this algorithm: https://www.youtube.com/watch?v=BysNXJHzCEs
package longcommsubstr

import "fmt"

func longestCommonSubstring(s1, s2 string) (int, [][]int) {
	var max int
	// Create a 2D slice len(s1)+1 * len(s2)+1. We'll index the slice at 1 to make our approach easier.
	T := make([][]int, len(s1)+1)
	for i, _ := range T {
		T[i] = make([]int, len(s2)+1)
	}

	for i := 1; i <= len(s1); i++ {
		for j := 1; j <= len(s2); j++ {
			if s1[i-1] == s2[j-1] {
				// If the character is the same, we take the previous character's value and add one. T[i-1][j-1] represents the previous character in both strings.
				// As you can recall from above, the strings are indexed at 0, but our slice is indexed at 1, hence subtracting 1.
				T[i][j] = T[i-1][j-1] + 1
				// Keep track of the max as we go, so that we can return it at the end.
				if T[i][j] > max {
					max = T[i][j]
				}
			} else {
				// If it doesn't match, we're still going to take the previous character's value, but not add 1.
				T[i][j] = T[i-1][j-1]
			}
		}
	}

	// Return T so that we can pass it to another function to print the longest common substring.
	return max, T
}

func printLongestCommonSubstring(s1 string, max int, T [][]int) string {
	// Allocate a slice of bytes that's equal to the answer.
	res := make([]byte, max)

	// Initialize i, j outside of the for loop so that we can use it another for loop.
	i, j := 1, 1

	// Setup a break label so that we can break out of both loops when we find the position of the end of the longest common substring.
FindPosition:
	for ; i < len(T); i++ {
		for ; j < len(T[0]); j++ {
			if T[i][j] == max {
				break FindPosition
			}
		}
		// Since i and j's scope is outside the above for loop, we need to set j back to 1 at the end of the loop.
		j = 1
	}

	// While max is not 0, put characters into our result slice from end to beginning of the longest common substring.
	for max > 0 {
		res[max-1] = s1[i-1]
		i--
		max--
	}

	return fmt.Sprintf("%s", res)

}
