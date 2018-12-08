package codility

import "testing"

func TestShift(t *testing.T) {
    testCases := []struct {
        in []int
        out []int
        k   int
    }{
        {[]int{1, 2, 3}, []int{1, 2, 3}, 0},
        {[]int{1, 2, 3}, []int{2, 3, 1}, 1},
        {[]int{1, 2, 3}, []int{1, 2, 3}, 3},
        {[]int{1, 2, 3, 4, 5}, []int{5, 1, 2, 3, 4}, 1},
    }

    for _, testCase := range testCases {
        inputArray := testCase.in
        shift(&inputArray, testCase.k)
        if !equalArrays(inputArray, testCase.out) {
            t.Errorf("Arrays are not equal: %v != %v", inputArray, testCase.out)
        }
    }
}

func equalArrays(a []int, b[]int) bool {
    if len(a) != len(b) { return false }
    for i := 0; i < len(a); i++ {
        if a[i] != b[i] { return false }
    }
    return true
}

