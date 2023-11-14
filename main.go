package main

import (
	"fmt"

	one "github.com/2zqa/ASCI-2023/one"
)

func main() {
	answer, err := one.Solve("one/input.txt")
	if err != nil {
		fmt.Printf("Something went wrong solving challenge 1: %s\n", err)
	} else {
		fmt.Printf("Solution 1: %s\n", answer)
	}
}
