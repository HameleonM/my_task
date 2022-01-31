package main

import (
	"testing"
)

func TestMaxMultiplication(t *testing.T) {

	testTable := []struct {
		list     []int
		expected int
	}{
		{
			list:     []int{1, 3, 8, 5, -8, -10, 11},
			expected: 88,
		},
		{
			list:     []int{0, -2, 10},
			expected: 0,
		},
		{
			list:     []int{},
			expected: 0,
		},
	}

	for _, testCase := range testTable {
		result := max_multiplication(testCase.list)

		t.Logf("Calling max_multiplication(%v), result %d\n",
			testCase.list, result)

		if result != testCase.expected {
			t.Errorf("Incorrect result. Expect %d, got %d",
				testCase.expected, result)
		}
	}
}

