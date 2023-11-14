package solution

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Solve(filePath string) (string, error) {
	inputBytes, err := os.ReadFile(filePath)
	if err != nil {
		return "", err
	}

	input := string(inputBytes)

	// Split the input into lines
	lines := strings.Split(input, "\n")

	// Iterate over the lines
	var sum int
	record := 0
	for index, line := range lines {
		// empty line means that the end of a group is reached
		if line == "" {
			// Update record
			if sum > record {
				record = sum
			}
			// Reset sum
			sum = 0
			continue
		}

		// Parse line to int
		value, err := strconv.Atoi(line)
		if err != nil {
			fmt.Printf("Warning: could not parse line %d: \"%s\"", index, line)
		}
		sum += value
	}

	return strconv.Itoa(record), nil
}
