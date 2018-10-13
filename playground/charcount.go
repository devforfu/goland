package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "unicode"
    "unicode/utf8"
)

func main() {
    counts := make(map[rune]int)
    categories := make(map[string]int)
    var utflen [utf8.UTFMax + 1]int
    invalid := 0

    in := bufio.NewReader(os.Stdin)
    for {
        r, n, err := in.ReadRune()
        if err == io.EOF {
            break
        }
        if err != nil {
            fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
            os.Exit(1)
        }
        if r == unicode.ReplacementChar && n == 1{
            invalid++
            continue
        }

        if unicode.IsLetter(r) {
            categories["letter"]++
        } else if unicode.IsDigit(r) {
            categories["digit"]++
        } else if unicode.IsPunct(r) {
            categories["punct"]++
        } else {
            categories["other"]++
        }

        counts[r]++
        utflen[n]++
    }

    println()
    println("\nrune\tcount")
    for c, n := range counts {
        fmt.Printf("%q\t%d\n", c, n)
    }

    println("\nlen\tcount")
    for i, n := range utflen {
        if i > 0 {
            fmt.Printf("%d\t%d\n", i, n)
        }
    }

    println("\ncat\tcount")
    for k, v := range categories {
        fmt.Printf("%s\t%d\n", k, v)
    }

    if invalid > 0 {
        fmt.Printf("\n%d invalid UTF-8 characters\n", invalid)
    }
}