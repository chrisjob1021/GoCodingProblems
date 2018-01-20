package coinchange

import "testing"

func TestCoinChange(t *testing.T) {
	var tests = []struct {
		coins  []int
		amount int
		expect int
	}{
		{[]int{1, 2, 5}, 11, 3},
		{[]int{2}, 3, -1},
		{[]int{7, 2, 3, 6}, 13, 2},
		{[]int{3, 2, 4}, 6, 2},
	}

	for _, test := range tests {
		if got, _ := coinChange(test.coins, test.amount); got != test.expect {
			t.Errorf("coinChange(%v, %v) = %v (expected %v)", test.coins, test.amount, got, test.expect)
		}
	}
}

func TestPrintCoinChangeCombination(t *testing.T) {
	var tests = []struct {
		coins  []int
		amount int
		expect string
	}{
		{[]int{1, 2, 5}, 11, "[5 5 1]"},
		{[]int{2}, 3, "[]"},
		{[]int{7, 2, 3, 6}, 13, "[6 7]"},
		{[]int{3, 2, 4}, 6, "[3 3]"},
	}

	for _, test := range tests {
		_, R := coinChange(test.coins, test.amount)
		if got := printCoinChangeCombination(test.coins, R); got != test.expect {
			t.Errorf("coinChange(%v, %v) = %v (expected %s)", test.coins, test.amount, got, test.expect)
		}
	}
}

func TestChange(t *testing.T) {
	var tests = []struct {
		coins  []int
		amount int
		expect int
	}{
		{[]int{1, 2, 5}, 5, 4},
		{[]int{2}, 3, 0},
		{[]int{10}, 10, 1},
	}

	for _, test := range tests {
		if got := change(test.amount, test.coins); got != test.expect {
			t.Errorf("coin(%v) = %v (expected %v)", test.coins, got, test.expect)
		}
	}
}
