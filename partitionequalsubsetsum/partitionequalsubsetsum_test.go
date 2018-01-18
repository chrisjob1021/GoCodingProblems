package partitionequalsubsetsum

import "testing"

func TestCanPartition(t *testing.T) {
	var tests = []struct {
		nums   []int
		expect bool
	}{
		{[]int{1, 5, 11, 5}, true},
		{[]int{1, 2, 3, 5}, false},
		{[]int{3, 3, 3, 4, 5}, true},
	}

	for _, test := range tests {
		if got, _ := canPartition(test.nums); got != test.expect {
			t.Errorf("canPartition(%v) = %v (expected %v)", test.nums, got, test.expect)
		}
	}
}

func TestPrintPartition(t *testing.T) {
	var tests = []struct {
		nums   []int
		expect string
	}{
		{[]int{1, 5, 11, 5}, "[[1 5 5] [11]]"},
		{[]int{3, 3, 3, 4, 5}, "[[5 4] [3 3 3]]"},
	}

	for _, test := range tests {
		_, T := canPartition(test.nums)
		if got := printParition(test.nums, T); got != test.expect {
			t.Errorf("printPartition(%v, T) = %v (expected %v)", test.nums, got, test.expect)
		}
	}
}
