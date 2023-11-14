package solution

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type parsedPassword struct {
	min, max int
	letter   byte
	password string
}

func (p parsedPassword) String() string {
	return fmt.Sprintf("%d-%d %c: %s", p.min, p.max, p.letter, p.password)
}

func Solve(filePath string) (string, error) {
	inputBytes, err := os.ReadFile(filePath)
	if err != nil {
		return "", err
	}

	input := string(inputBytes)

	// Split the input into lines
	lines := strings.Split(input, "\n")

	// Iterate over the lines
	validPasswordCount := 0
	var parsingErr error
	for _, line := range lines {
		var pw parsedPassword
		pw, parsingErr = parseLine(line)
		if err != nil {
			// fmt.Println(err)
			continue
		}
		if isValidPassword(pw) {
			validPasswordCount++
			// } else {
			// 	fmt.Printf("Invalid pw: %v\n", pw)
		}
	}

	return strconv.Itoa(validPasswordCount), parsingErr
}

func isValidPassword(p parsedPassword) bool {
	letterCount := strings.Count(p.password, string(p.letter))
	if letterCount < p.min || letterCount > p.max {
		return false
	}
	return true
}

func parseLine(line string) (parsedPassword, error) {
	// e.g.: `3-16 e: leurrtgbeeur`
	i := strings.IndexRune(line, '-')
	if i == -1 {
		return parsedPassword{}, fmt.Errorf("invalid syntax")
	}

	j := strings.IndexRune(line, ' ')
	if j == -1 {
		return parsedPassword{}, fmt.Errorf("invalid syntax")
	}

	min := line[:i]
	max := line[i+1 : j]
	letter := line[j+1]
	password := line[j+4:]

	parsedMin, err := strconv.Atoi(min)
	if err != nil {
		return parsedPassword{}, fmt.Errorf("invalid syntax")
	}
	parsedMax, err := strconv.Atoi(max)
	if err != nil {
		return parsedPassword{}, fmt.Errorf("invalid syntax")
	}

	return parsedPassword{
		min:      parsedMin,
		max:      parsedMax,
		letter:   letter,
		password: password,
	}, nil
}
