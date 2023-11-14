package solution

import (
	"os"
	"strconv"
	"strings"

	"github.com/hashicorp/go-set/v2"
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
	for _, line := range lines {
		part1, part2 := SplitString(line)
		sharedLetters := GetSharedLetters(part1, part2)
		sum += GetPrioritiesSum(sharedLetters)
	}
	return strconv.Itoa(sum), nil
}

func GetPrioritiesSum(s string) int {
	sum := 0
	for _, r := range s {
		sum += GetPriority(r)
	}
	return sum
}

// GetPriority returns the priority of a given rune as explained in story.md.
// Only supports letters. Might output weird results for other runes.
func GetPriority(r rune) int {
	return (int(r) - 38) % 58
}

func GetSharedLetters(partA, partB string) string {
	partARuneSet := set.From([]rune(partA))
	partBRuneSet := set.From([]rune(partB))

	return string(partARuneSet.Intersect(partBRuneSet).Slice())
}

func SplitString(s string) (part1, part2 string) {
	part1 = s[:len(s)/2]
	part2 = s[len(s)/2:]
	return
}
