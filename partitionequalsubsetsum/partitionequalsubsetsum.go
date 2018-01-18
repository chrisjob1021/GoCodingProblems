// Partition Equal Subset Sum
// Dynamic Programming
//
// Time complexity: O(len(input) * targetSum)
// Space complexity: O(len(input) * targetSum)
//
// References:
// https://leetcode.com/problems/partition-equal-subset-sum/description/
// https://github.com/mission-peace/interview/blob/master/src/com/interview/dynamic/SubsetSum.java
//
// Acknowledgements:
// Tushar Roy @mission-peace
// His video explaining this algorithm: https://www.youtube.com/watch?v=s6FhG--P7z0
package partitionequalsubsetsum

import "fmt"

// Take an input slice and swap end element and element at "i"
// Then return slice without the element at end
func remove(slice []int, i int) []int {
	slice[len(slice)-1], slice[i] = slice[i], slice[len(slice)-1]
	return slice[:len(slice)-1]
}

func printParition(nums []int, T [][]bool) string {
	var res []int
	// We need a copy of the input numbers to work on.
	resNums := nums

	// "i" is set to the slice of the input numbers. Keep in mind, our solution is indexed at 1, so we start at len(nums) rather than len(nums)-1.
	// "j" is set to our target value. Since our solution is indexed at 1, but we need the actual target value, we take the opposite approach to "i" and use len(T[0])-1.
	i, j := len(nums), len(T[0])-1

	// We are going to back track from the starting positions outlined above.
	for j > 0 {
		// If the row above is true...
		if T[i-1][j] {
			// ...back up one row.
			i--
			// As outlined in canPartition(), if true at T[i-1][j-nums[i-1]], then our current number is used as part of the solution.
			if T[i-1][j-nums[i-1]] {
				res = append(res, nums[i-1])
				j = j - nums[i-1]
				resNums = remove(resNums, i-1)
			}
		}
	}

	return fmt.Sprintf("%v", [][]int{resNums, res})
}

// This problem is very similiar to common Subset Sum.
// However, instead of an explicit target value, we need to set our own.
// We need to return true or false if a subset sum exists that results in equal sums for the resulting two partitions (slices).
// The partitions need not be equal in element count, but the rather sum of each partition's elements is equal to the other.
// So, our explicit target can be the midpoint of sum of all elements.
func canPartition(nums []int) (bool, [][]bool) {
	var sum int
	for _, val := range nums {
		sum += val
	}

	// If it is not an equal split, no solution exists and we can return false.
	if sum%2 != 0 {
		return false, [][]bool{}
	}
	target := sum / 2
	//fmt.Println(target)

	// We're going to create a 2D matrix with the usable numbers on the Y-axis and every number up to the target on the X-axis.
	// We're going to create it with one more row and column. It will make our approach all come together, as you will see below.
	T := make([][]bool, len(nums)+1)
	for i, _ := range T {
		T[i] = make([]bool, target+1)
		// Set the first column as "true" because we could always create zero with an empty set.
		T[i][0] = true
	}

	for i := 1; i < len(T); i++ {
		for j := 1; j < len(T[0]); j++ {
			// Remember that i-1 is used because our result array T is indexed starting at 1.
			if j-nums[i-1] >= 0 {
				// Pick between two options. Prefer true values over false. The first option is detailed below.
				// The second option is to take the number above's answer for the targeted value "j" minus our value.
				// Meaning, could we add our current number's value to the above number's (T[i-1]) to make up our target value in "j"?
				T[i][j] = T[i-1][j] || T[i-1][j-nums[i-1]]
			} else {
				// If our current number is greater than j, we can take the above number's boolean answer.
				// This is because say if the above number was "true", we could create the value represented by "j" without using our current number.
				T[i][j] = T[i-1][j]
			}
		}
	}

	// Return T to pass to printPartition()
	return T[len(nums)][target], T
}
