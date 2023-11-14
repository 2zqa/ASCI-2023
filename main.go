package main

import (
	"fmt"

	one "github.com/2zqa/ASCI-2023/one"
	three "github.com/2zqa/ASCI-2023/three"
	two "github.com/2zqa/ASCI-2023/two"
)

func main() {
	answer, err := one.Solve("one/input.txt")
	if err != nil {
		fmt.Printf("Something went wrong solving challenge 1: %s\n", err)
	} else {
		fmt.Printf("Solution 1: %s\n", answer)
	}

	answer, err = two.Solve("two/input.txt")
	if err != nil {
		fmt.Printf("Something went wrong solving challenge 2: %s\n", err)
	} else {
		fmt.Printf("Solution 2: %s\n", answer)
	}

	answer, err = three.Solve("three/input.txt")
	if err != nil {
		fmt.Printf("Something went wrong solving challenge 3: %s\n", err)
	} else {
		fmt.Printf("Solution 3: %s\n", answer)
	}
}
