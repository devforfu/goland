package main

import (
	"../popcounter"
	"fmt"
	"strconv"
	"strings"
	"os"
)


func main() {
	digits := os.Args[1:]
	println(fmt.Sprintf("Args: %s", strings.Join(digits, ", ")))
	for _, digitStr := range digits {
		number, err := strconv.ParseUint(digitStr,0, 64)
		if err != nil {
			fmt.Printf("Cannot parse argument: %s\n", digitStr)
			continue
		}
		count := popcounter.PopCount(number)
		fmt.Printf("%-3d has %-3d non-zero bits\n", number, count)
	}
}