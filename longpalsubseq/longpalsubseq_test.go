package longpalsubseq

import "testing"

func TestLongestPalindromeSubseq(t *testing.T) {
	var tests = []struct {
		input  string
		expect int
	}{
		{"bbbab", 4},
		{"cbbd", 2},
		{"agbdba", 5},
	}

	for _, test := range tests {
		if got, _ := longestPalindromeSubseq(test.input); got != test.expect {
			t.Errorf("longestPalindromeSubseq(%v) = %v (expected %v)", test.input, got, test.expect)
		}
	}
}

func TestPrintLongestPalindromeSubseq(t *testing.T) {
	var tests = []struct {
		input  string
		expect string
	}{
		{"bbbab", "bbbb"},
		{"cbbd", "bb"},
		{"agbdba", "abdba"},
	}

	for _, test := range tests {
		_, T := longestPalindromeSubseq(test.input)
		if got := printLongestPalindromeSubseq(test.input, T); got != test.expect {
			t.Errorf("printLongestPalindromeSubseq(%v) = %v (expected %v)", test.input, got, test.expect)
		}
	}
}
