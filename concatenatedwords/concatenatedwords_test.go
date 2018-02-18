package concatenatedwords

import (
	"fmt"
	"testing"
)

func TestFindAllConcatenatedWordsInADict(t *testing.T) {
	var tests = []struct {
		words  []string
		expect string
	}{
		{[]string{"cat", "cats", "catsdogcats", "dog", "dogcatsdog", "hippopotamuses", "rat", "ratcatdogcat"}, "[dogcatsdog catsdogcats ratcatdogcat]"},
	}

	for _, test := range tests {
		got := fmt.Sprintf("%v", findAllConcatenatedWordsInADict(test.words))
		if got != test.expect {
			t.Errorf("findAllConcatenatedWordsInADict(%v) = %v (expected %v)", test.words, got, test.expect)
		}
	}
}
