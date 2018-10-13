package main

import (
    "fmt"
)

func main() {
    arr := []string{"first", "first", "second", "first", "third", "second", "second"}
    dedup(&arr)
    fmt.Println(arr)
}

func dedup(arrPtr *[]string) {
    prev := (*arrPtr)[0]
    i, dups := 1, 0
    for _, s := range (*arrPtr)[1:] {
        if s != prev {
            (*arrPtr)[i] = s
            i++
        } else {
            dups++
        }
        prev = s
    }
    *arrPtr = (*arrPtr)[:len(*arrPtr) - dups]
}
