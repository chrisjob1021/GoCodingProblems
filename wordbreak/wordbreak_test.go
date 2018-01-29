package wordbreak

import (
	"fmt"
	"testing"
)

func TestWordBreak(t *testing.T) {
	var tests = []struct {
		s        string
		wordDict []string
		expect   bool
	}{
		{"catsanddog", []string{"cat", "cats", "and", "sand", "dog"}, true},
	}

	for _, test := range tests {
		if got := wordBreak(test.s, test.wordDict); got != test.expect {
			t.Errorf("wordBreak(%v, %v) = %v (expected %v)", test.s, test.wordDict, got, test.expect)
		}
	}
}

func TestWordBreakII(t *testing.T) {
	var tests = []struct {
		s        string
		wordDict []string
		expect   string
	}{
		{"catsanddog", []string{"cat", "cats", "and", "sand", "dog"}, "[\"cat sand dog\" \"cats and dog\"]"},
	}

	for _, test := range tests {
		if got := fmt.Sprintf("%q", wordBreakII(test.s, test.wordDict)); got != test.expect {
			t.Errorf("wordBreakII(%v, %v) = %v (expected %v)", test.s, test.wordDict, got, test.expect)
		}
	}
}
