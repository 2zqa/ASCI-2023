package solution

import (
	"strings"
	"testing"
)

func BenchmarkSolution3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Solve("input.txt")
	}
}

func TestSplitString(t *testing.T) {
	testCases := []struct {
		input    string
		expected [2]string
	}{
		{"abcdef", [2]string{"abc", "def"}},
		{"xyz123", [2]string{"xyz", "123"}},
		{"abc1234", [2]string{"abc", "1234"}},
	}

	for _, tc := range testCases {
		result1, result2 := SplitString(tc.input)
		if result1 != tc.expected[0] || result2 != tc.expected[1] {
			t.Errorf("SplitString(%q) = (%q, %q), expected (%q, %q)", tc.input, result1, result2, tc.expected[0], tc.expected[1])
		}
	}
}
func TestGetPriority(t *testing.T) {
	testCases := []struct {
		input    rune
		expected int
	}{
		{'a', 1},
		{'z', 26},
		{'A', 27},
		{'Z', 52},
	}

	for _, tc := range testCases {
		result := GetPriority(tc.input)
		if result != tc.expected {
			t.Errorf("GetPriority(%q) = %d, expected %d", tc.input, result, tc.expected)
		}
	}
}

func TestGetSharedLetters(t *testing.T) {
	testCases := []struct {
		partA    string
		partB    string
		expected string
	}{
		{
			partA:    "abc",
			partB:    "def",
			expected: "",
		},
		{
			partA:    "abc",
			partB:    "bcd",
			expected: "bc",
		},
		{
			partA:    "abc",
			partB:    "abc",
			expected: "abc",
		},
	}

	for _, tc := range testCases {
		result := GetSharedLetters(tc.partA, tc.partB)
		if !checkSharedLetters(result, tc.expected) {
			t.Errorf("GetSharedLetters(%q, %q) = %q, expected %q", tc.partA, tc.partB, result, tc.expected)
		}

	}
}

func checkSharedLetters(got, want string) bool {
	if got == want {
		return true
	}

	if len(got) != len(want) {
		return false
	}

	for _, c := range want {
		if !strings.ContainsRune(got, c) {
			return false
		}
	}

	return true
}
