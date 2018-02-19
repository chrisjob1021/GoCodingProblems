package sorting

// Iterate through the input array and always select the minimum element.
// Swap "i" with the minimum element.
//
// Time Complexity: O(n^2)
//
// References:
// https://www.geeksforgeeks.org/selection-sort/
func SelectionSort(input []int) []int {
	for i := 0; i < len(input); i++ {
		min := i
		for j := i + 1; j < len(input); j++ {
			if input[j] < input[min] {
				min = j
			}
		}

		input[i], input[min] = input[min], input[i]
	}

	return input
}

// Sort an input slice like you would sort a card in your hand full of cards.
//
// Time Complexity: O(n^2)
//
// References:
// https://www.geeksforgeeks.org/insertion-sort/
func InsertionSort(input []int) []int {
	for i := 1; i < len(input); i++ {
		// "key" is the card we're trying to place in proper order in our "hand".
		key := input[i]
		j := i - 1

		// We'll use j to track where to put the "cards".
		for ; j >= 0 && input[j] > key; j-- {
			// If current j is greater than key, move it one to the right.
			input[j+1] = input[j]
		}

		// Finally, put the "key" in place.
		// Would not have moved if it's greater than all the cards before it.
		// Otherwise, j+1 will represent the position before all of the cards we moved above.
		input[j+1] = key
	}

	return input
}
