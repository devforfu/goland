package main

import "fmt"

func main() {
	var x uint8 = 1 << 1 | 1 << 5
	var y uint8 = 1 << 1 | 1 << 2 | 1 << 7

	values := []struct{
		operation string
		result uint8
	}{
		{"x", x},
		{"y", y},
		{"not(y)", not(y)},
		{"x|y",x|y},
		{"x&y",x&y},
		{"x^y",x^y},
		{"x&^y", x&^y},
		{"x&not(y)",x&not(y)},
	}

	for _, value := range values {
		fmt.Printf("%-8s = %08b\n", value.operation, value.result)
	}
}

func not(b byte) byte {
	return 255 - b
}
