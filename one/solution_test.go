package solution

import "testing"

func BenchmarkSolve(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Solve("input.txt")
	}
}

func TestSolve(t *testing.T) {
	expected := "68923"
	result, err := Solve("input.txt")
	if err != nil {
		t.Errorf("Solution1 returned an error: %v", err)
	}
	if result != expected {
		t.Errorf("Solution1 = %v, expected %v", result, expected)
	}
}
