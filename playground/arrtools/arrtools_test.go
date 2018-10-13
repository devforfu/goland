package arrtools

import (
    "fmt"
    "testing"
)

func TestEqualSlices(t *testing.T) {
    testCases := []struct {
        fst, snd []int
        equal bool
    }{
        {[]int{1}, []int{1}, true},
        {[]int{1, 2, 3}, []int{1, 2, 3}, true},
        {[]int{1, 2}, []int{2, 3}, false},
        {[]int{1, 2, 3}, []int{1}, false},
    }
    for _, item := range testCases {
        x, y := item.fst, item.snd
        equal := EqualSlices(x, y)
        sign := []string{"==", "!="}[btoi(equal)]
        if equal != item.equal {
            t.Error(fmt.Sprintf("%v %s %v", x, sign, y))
        }
    }
}

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
        if !EqualSlices(x, y) {
            t.Error(fmt.Sprintf("%v != %v", x, y))
        }
    }
}

func btoi(b bool) int {
    if b { return 1 } else { return 0 }
}