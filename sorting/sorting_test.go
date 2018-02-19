package sorting

import (
	"reflect"
	"testing"
)

func TestSelectionSort(t *testing.T) {
	var tests = []struct {
		input  []int
		expect []int
	}{
		{[]int{8, 4, 9, 11, 10}, []int{4, 8, 9, 10, 11}},
	}

	for _, test := range tests {
		if got := SelectionSort(test.input); !reflect.DeepEqual(got, test.expect) {
			t.Errorf("selectionSort(%v) = %v (expected %v)", test.input, got, test.expect)
		}
	}
}

func TestInsertionSort(t *testing.T) {
	var tests = []struct {
		input  []int
		expect []int
	}{
		{[]int{8, 4, 9, 11, 10}, []int{4, 8, 9, 10, 11}},
	}

	for _, test := range tests {
		if got := InsertionSort(test.input); !reflect.DeepEqual(got, test.expect) {
			t.Errorf("selectionSort(%v) = %v (expected %v)", test.input, got, test.expect)
		}
	}
}
