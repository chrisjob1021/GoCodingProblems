// Concatenated Words
// Dynamic Programming
//
// References:
// https://leetcode.com/problems/concatenated-words/description/
// https://leetcode.com/problems/concatenated-words/discuss/95652/Java-DP-Solution
//
// Acknowledgements:
// Shan Gao @shawngao
// https://leetcode.com/shawngao/
package concatenatedwords

import "sort"

func min(x, y int) int {
	if x < y {
		return x
	} else {
		return y
	}
}

func findAllConcatenatedWordsInADict(words []string) []string {
	var result []string
	// Using a dictionary so we don't need to iterate through every item to check if it already exists.
	wordMap := make(map[string]bool)

	// We know that the smaller words need to be used to construct the larger concatenated words.
	sort.Slice(words, func(i, j int) bool {
		return len(words[i]) < len(words[j])
	})

	for _, word := range words {
		// Call helper function to determine if it is a concatenated word.
		if checkWord(&word, wordMap) {
			result = append(result, word)
		}

		// Build dictionary as we go so that checkWord will only return true if it is a concatenated word.
		// Input word slice contains only unique words.
		// If we added all words to the dictionary right away, checkWord would return true for single words since the input word would contain dictionary words entirely.
		wordMap[word] = true
	}

	return result
}

// Helper function to check one word at a time, compared to our wordMap that we will build as we go.
func checkWord(word *string, wordMap map[string]bool) bool {
	if len(wordMap) == 0 {
		return false
	}

	// Allocate a result slice that is one longer than the input string.
	T := make([]bool, len(*word)+1)
	// Start out as if we have a valid solution.
	T[0] = true

	for i := 1; i <= len(*word); i++ {
		for j := 0; j < i; j++ {
			// If starting index of our substring is false, immediately continue.
			// This is so that we know our solution contains multiple words in the dictionary.
			if !T[j] {
				continue
			}

			// If substring represented by word[j:i] is contained in the map, set the end position to true so that the above logic works.
			if ok := wordMap[(*word)[j:i]]; ok {
				T[i] = true
				break
			}
		}
	}

	return T[len(*word)]
}
