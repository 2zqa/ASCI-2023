package solution

import "testing"

func BenchmarkSolution2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Solve("input.txt")
	}
}

func TestSolution2(t *testing.T) {
	expected := "515"
	result, err := Solve("input.txt")
	if err != nil {
		t.Errorf("Solve returned an error: %v", err)
	}
	if result != expected {
		t.Errorf("Solve = %v, expected %v", result, expected)
	}
}
