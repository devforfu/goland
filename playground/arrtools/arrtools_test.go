package arrtools

import (
    "fmt"
    "testing"
)

func TestReverse(t *testing.T) {
    testCases := []struct {
        original []int
        expected []int
    }{
        {[]int{1}, []int{1}},
        {[]int{0, 0, 0}, []int{0, 0, 0}},
        {[]int{1, 2, 3, 4, 5}, []int{5, 4, 3, 2, 1}},
    }

    for _, item := range testCases {
        x, y := item.original, item.expected
        Reverse(x[:])
        if !equalSlices(x, y) {
            t.Error(fmt.Sprintf("%v != %v", x, y))
        }
    }
}

func equalSlices(x, y []int) bool {
    if len(x) != len(y) {
        return false
    }
    for i := range x {
        if x[i] != y[i] {
            return false
        }
    }
    return true
}