package longcommsubstr

import (
	"testing"
)

func TestLongestCommonSubstring(t *testing.T) {
	var tests = []struct {
		s1     string
		s2     string
		expect int
	}{
		{"GeeksforGeeks", "GeeksQuiz", 5},
		{"abcdxyz", "xyzabcd", 4},
		{"zxabcdezy", "yzabcdezx", 6},
	}

	for _, test := range tests {
		if got, _ := longestCommonSubstring(test.s1, test.s2); got != test.expect {
			t.Errorf("longestCommonSubstring(%v, %v) = %v (expected %v)", test.s1, test.s2, got, test.expect)
		}
	}
}

func TestPrintLongestCommonSubstring(t *testing.T) {
	var tests = []struct {
		s1     string
		s2     string
		expect string
	}{
		{"GeeksforGeeks", "GeeksQuiz", "Geeks"},
		{"abcdxyz", "xyzabcd", "abcd"},
		{"zxabcdezy", "yzabcdezx", "abcdez"},
	}

	for _, test := range tests {
		max, T := longestCommonSubstring(test.s1, test.s2)
		if got := printLongestCommonSubstring(test.s1, max, T); got != test.expect {
			t.Errorf("printLongestCommonSubstring(%v, %v) = %v (expected %v)", test.s1, test.s2, got, test.expect)
		}
	}
}
