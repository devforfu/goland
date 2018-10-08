package main

import "fmt"

func main() {
    xs := [...]int{1, 2, 3, 4, 5}
    reverse(xs[:])
    fmt.Printf("%v", xs)
}

func reverse(values []int) {
    for i, j := 0, len(values) - 1; i < j; i, j = i + 1, j - 1 {
        values[i], values[j] = values[j], values[i]
    }
}
