// Longest Common Subsequence
// Dynamic Programming
//
// Time complexity: O(len(str1) * len(str2))
// Space complexity: O(len(str1) * len(str2))
//
// References:
// https://www.geeksforgeeks.org/longest-common-subsequence/
// https://github.com/mission-peace/interview/blob/master/src/com/interview/dynamic/LongestCommonSubsequence.java
//
// Acknowledgements:
// Tushar Roy @mission-peace
// His video explaining this algorithm: https://www.youtube.com/watch?v=NnD96abizww
package longcomsubseq

// Helper max() function to avoid converting int to float64 to use math.Max().
func max(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}

func longestCommonSubsequence(str1, str2 string) int {
	// Create a result 2D matrix that is the length of str1 * length of str2.
	// Our result will index each of the strings at 1.
	// Index 0 will represent the longest common subsequence if there were no characters.
	T := make([][]int, len(str1)+1)
	for i, _ := range T {
		T[i] = make([]int, len(str2)+1)
	}

	// For the purpose of our comments, let's use example inputs “ABCDGH” and “AEDFHR”.
	for i := 1; i <= len(str1); i++ {
		for j := 1; j <= len(str2); j++ {
			// If the string is the same, then we should take the previous result and add one.
			// If we're starting at the first character, the previous result was 0. This is represeted at T[i-1][j-1].
			if str1[i-1] == str2[j-1] {
				T[i][j] = T[i-1][j-1] + 1
			} else {
				// Otherwise, we can take the maximum if we deleted a character from either input string.
				// For example, "AB" != "AE". Take the maximum of either string T[i-1][j] or T[i][j-1]. In this case, the value is the same.
				T[i][j] = max(T[i-1][j], T[i][j-1])
			}
		}
	}

	// The result is stored at the end of both strings in the bottom right of the 2D matrix.
	return T[len(str1)][len(str2)]
}
