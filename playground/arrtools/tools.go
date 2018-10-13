package arrtools

// Reverse reverses the order of elements in array.
func Reverse(arr []int) {
    for i, j := 0, len(arr) - 1; i < j; i, j = i+1, j-1 {
        arr[i], arr[j] = arr[j], arr[i]
    }
}

// EqualSlices asserts that two slices are of equal length and same content.
func EqualSlices(x, y []int) bool {
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

// NonEmpty filters empty strings from slice inplace.
func NonEmpty(strings []string) []string {
    i := 0
    for _, s := range strings {
        if s != "" {
            strings[i] = s
            i++
        }
    }
    return strings[:i]
}