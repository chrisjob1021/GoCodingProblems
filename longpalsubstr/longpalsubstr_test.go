package longpalsubstr

import "testing"

func TestLongestPalindrome(t *testing.T) {
	var tests = []struct {
		input  string
		expect string
	}{
		{"babad", "bab"},
		{"cbbd", "bb"},
		{"a", "a"},
		{"aaaa", "aaaa"},
	}

	for _, test := range tests {
		if got := longestPalindrome(test.input); got != test.expect {
			t.Errorf("longestPalindrome(%v) = %v (expected %v)", test.input, got, test.expect)
		}
	}
}
