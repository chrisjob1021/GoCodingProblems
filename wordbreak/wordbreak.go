package wordbreak

import (
	"bytes"
)

// Word Break
// Dynamic Programming
//
// Time complexity: O(len(s) * len(s))
// Space complexity: O(len(s) * len(s))
//
// References:
// https://leetcode.com/problems/word-break/description/
// https://github.com/mission-peace/interview/blob/master/src/com/interview/dynamic/BreakMultipleWordsWithNoSpaceIntoSpace.java
//
// Acknowledgements:
// Tushar Roy @mission-peace
// His video explaining this algorithm: https://www.youtube.com/watch?v=WepWFGxiwRs
func wordBreak(s string, wordDict []string) bool {
	// Put dictionary words into map so that we can check if they exist in constant time, rather than O(N).
	wordMap := make(map[string]bool)
	for _, word := range wordDict {
		wordMap[word] = true
	}

	// Allocate a 2D matrix of N*N where N represents the length of the input string.
	// This will represent slices of the input string. Example: T[0][2] represents s[0...2]
	T := make([][]bool, len(s))
	for i, _ := range T {
		T[i] = make([]bool, len(s))
	}

	// Create a variable "l" that will track the length of our slices.
	for l := 1; l <= len(s); l++ {
		// Create an "i" variable that will track the beginning of our slice.
		for i := 0; i <= len(s)-l; i++ {
			j := i + l - 1
			sliceWord := s[i : j+1]

			// If the word represented by s[i...j] is in our wordMap, set matrix position to true.
			if _, ok := wordMap[sliceWord]; ok {
				T[i][j] = true
				continue
			}

			// Otherwise, we need to iterate through all the mid points of the slice.
			// We are going to determine if the slice represented by "i" and "j" can be broken such that the resuling slices are "true" in our solution matrix.
			// "k" represents the midpoint.
			for k := i + 1; k <= j; k++ {
				if T[i][k-1] && T[k][j] {
					T[i][j] = true
				}
			}
		}
	}

	return T[0][len(s)-1]

}

// Word Break II
// Memoization + DFS
//
// Time complexity: O(len(s) ^ len(s/smallest word in dictionary)
// Space complexity: O(len(s)/len(smallest word in dictionary)) + O(number of words in dictionary)
//
// References:
// https://leetcode.com/problems/word-break-ii/description/
// https://github.com/mission-peace/interview/blob/master/src/com/interview/dynamic/BreakMultipleWordsWithNoSpaceIntoSpace.java
// https://leetcode.com/problems/word-break-ii/discuss/44167
func wordBreakII(s string, wordDict []string) []string {
	// Create a map version of the word dictionary, so we can look up words in constant time rather than O(N).
	// Identify the longest word in the provided dictionary so we can set an upper bound on the length of our search.
	var maxLen int

	wordMap := make(map[string]bool)
	for _, val := range wordDict {
		wordMap[val] = true
		// Leverage helper max() function so that we do not need to convert integers to float64 to leverage math.Max().
		maxLen = max(maxLen, len(val))
	}

	// Create a map representing the start postion of our search as the key and the resulting slice of strings as our value.
	resultMap := make(map[int][]string)

	// Use DFS helper function with memoization (see above, resultMap).
	return wordBreakIIDFSUtil(s, wordMap, resultMap, 0, maxLen)

}

func wordBreakIIDFSUtil(s string, wordMap map[string]bool, resultMap map[int][]string, start int, maxLen int) []string {
	// If we've reached the end of the input string, return an empty string as the result.
	if start == len(s) {
		return []string{""}
	}

	// If we already have the start position in our result map, just return it.
	if _, ok := resultMap[start]; ok {
		return resultMap[start]
	}

	var words []string
	// Iterate from start to length of longest dictionary word and end of string.
	for i := start; i < start+maxLen && i < len(s); i++ {
		// Create substring from the input string from "start" to increasing "i" position (s[start...i]).
		newWord := s[start : i+1]

		// If our new word is not in the provided list of words, we should immediately increment and continue our search.
		if _, ok := wordMap[newWord]; !ok {
			continue
		}

		// Otherwise, we should our depth first search until the end of the input string.
		returnedWords := wordBreakIIDFSUtil(s, wordMap, resultMap, i+1, maxLen)
		for _, word := range returnedWords {
			// Allocate a buffer so that we can efficiently append our words to a result string.
			var b bytes.Buffer
			b.Write([]byte(newWord))
			// Put a space between each word.
			// Checking if length is 0, because we'll return an empty string "" (len 0) when we encounter the end of string.
			if len(word) != 0 {
				b.Write([]byte(" "))
				b.Write([]byte(word))

			}
			words = append(words, b.String())
		}
	}

	// Store our result so that we do not duplicate work.
	resultMap[start] = words
	return words
}

// Helper function to return max of two integers.
func max(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}
