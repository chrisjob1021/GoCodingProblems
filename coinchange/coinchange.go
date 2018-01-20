package coinchange

import (
	"fmt"
	"math"
)

// Coin Change
// Dynamic Programming
//
// Time complexity: O(len(coins) * amount)
// Space complexity: O(amount)
//
// References:
// https://leetcode.com/problems/coin-change/description/
// https://github.com/mission-peace/interview/blob/master/src/com/interview/dynamic/CoinChangingMinimumCoin.java
//
// Acknowledgements:
// Tushar Roy @mission-peace
// His video explaining this algorithm: https://www.youtube.com/watch?v=NJuKJ8sasGk
func coinChange(coins []int, amount int) (int, []int) {
	// Create two slices. The first will hold the minimum number of coins.
	// amount+1 because slices are indexed at 0 and we want an actual representation of the value.
	T := make([]int, amount+1)
	// The second will hold the coin used to reach each number leading up to final amount.
	R := make([]int, amount+1)

	// Initialize T to +INF-1 (-1 to prevent overflow) since we want to compare that our solution is less than the previous solution.
	// Initialize R to -1, representing no solution.
	for i := 1; i < amount+1; i++ {
		T[i] = math.MaxInt64 - 1
		R[i] = -1
	}

	// "j" represents each coin.
	for j := 0; j < len(coins); j++ {
		// "i" represents each amount.
		for i := 1; i < amount+1; i++ {
			// If the amount is greater than the coin's value, we can do something with it.
			// Else, it would not be a valid combination.
			if i >= coins[j] {
				// If T[amount-coin value]+1 is less than current number of coins at T[i]
				// The point of this is to build up the solution. Use a coin as a stepping stone to the final amount.
				if T[i-coins[j]]+1 < T[i] {
					T[i] = T[i-coins[j]] + 1
					// We will also store the coin's position in []coins used to reach each amount. We will use this later to print the combination of coins.
					R[i] = j
				}
			}
		}
	}

	// If the amount is equal to +INF-1, then we there is no valid combination. Return -1 instead.
	// Else, return the number of coins stored at T[amount].
	// We also return the R slice so that we can pass that to another function to print the combination of coins.
	if T[amount] == math.MaxInt64-1 {
		return -1, []int{}
	} else {
		return T[amount], R
	}
}

func printCoinChangeCombination(coins []int, R []int) string {
	var res []int

	// If the last item in the slice is -1, then we never reached a valid result.
	// Return an empty string.
	if len(R) == 0 {
		return fmt.Sprint(res)
	}

	// We're going to start at the end of the slice and backtrack.
	// Once "start" is 0, we can break and print the result.
	start := len(R) - 1
	for start != 0 {
		// Append coin to our result slice.
		// R contains the indexed position within the []coins slice.
		res = append(res, coins[R[start]])
		// We can then work backward. Our next coin in the combination can be found at start - above coin's amount.
		start = start - coins[R[start]]
	}

	return fmt.Sprint(res)
}

// Coin Change II
// Dynamic Programming
//
// Time complexity: O(len(coins) * amount)
// Space complexity: O(len(coins) * amount)
//
// References:
// https://leetcode.com/problems/coin-change-2/description/
// https://github.com/mission-peace/interview/blob/master/src/com/interview/dynamic/CoinChanging.java
//
// Acknowledgements:
// Tushar Roy @mission-peace
// His video explaining this algorithm: https://www.youtube.com/watch?v=_fgjrs570YE
func change(amount int, coins []int) int {
	// Create a 2D slice of len(coins)+1 * amount+1.
	// +1, because indexing at 1 instead of 0 will make things easier below.
	T := make([][]int, len(coins)+1)

	for i, _ := range T {
		T[i] = make([]int, amount+1)
		// Set the first column to 1 representing that if we find a combination, the minimum number is at least 1.
		T[i][0] = 1
	}

	// "i" will iterate through all the coins.
	for i := 1; i < len(coins)+1; i++ {
		// "j" will iterate through all the different amounts by one.
		// Each row will represent, to this point, how many combinations can made up for this coin "i" at this amount "j".
		for j := 1; j < amount+1; j++ {
			// If the coin's amount is greater than the amount, our current coin can't be used in the combination.
			// Populate current position with above coin's value at this amount.
			if coins[i-1] > j {
				T[i][j] = T[i-1][j]
			} else {
				// Build from the bottom up.
				// Take the above coin's number of combinations at this amount.
				// Plus, the current coin's number of combinations to reach the "j" amount.
				// Given that we're building up along the way, we can find above value at T[current coin][current amount "j" -  current coin's value].
				// Current coin = coins[i-1] because it is indexed at 0 and we are indexing at 1.
				// Same bottom up approach is used in coinChange.
				T[i][j] = T[i][j-coins[i-1]] + T[i-1][j]
			}
		}
	}

	// Our solution is found at very bottom right of our 2D slice.
	return T[len(coins)][amount]
}
