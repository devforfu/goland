package tools

import (
	"testing"
	"fmt"
)

func TestSort(t *testing.T) {
	var testCases = []struct {
		inputValue [5]int
		expected [5]int
	}{
		{[...]int{5, 4, 3, 2, 1}, [...]int{1, 2, 3, 4, 5}},
	}
	for _, testCase := range testCases {
		length := len(testCase.inputValue)
		arr := testCase.inputValue[:]
		Sort(arr)
		for i := 0; i < length; i++ {
			inputValue := arr[i]
			expected := testCase.expected[i]
			if inputValue != expected {
				t.Error(fmt.Sprintf("%d != %d", inputValue, expected))
			}
		}
	}
}
