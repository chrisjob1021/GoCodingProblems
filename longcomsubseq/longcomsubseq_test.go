package longcomsubseq

import "testing"

func TestLongestCommonSubsequence(t *testing.T) {
	var tests = []struct {
		str1   string
		str2   string
		expect int
	}{
		{"ABCDGH", "AEDFHR", 3},
		{"AGGTAB", "GXTXAYB", 4},
	}

	for _, test := range tests {
		if got := longestCommonSubsequence(test.str1, test.str2); got != test.expect {
			t.Errorf("longestCommonSubsequence(%v, %v) = %v (expected %v)", test.str1, test.str2, got, test.expect)
		}
	}
}
