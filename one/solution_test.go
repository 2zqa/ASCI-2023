package solution

import "testing"

func BenchmarkSolution1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Solution1("input.txt")
	}
}

func TestSolution1(t *testing.T) {
	expected := "68923"
	result, err := Solution1("input.txt")
	if err != nil {
		t.Errorf("Solution1 returned an error: %v", err)
	}
	if result != expected {
		t.Errorf("Solution1 = %v, expected %v", result, expected)
	}
}
