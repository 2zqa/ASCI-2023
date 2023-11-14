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

func TestIsValidPassword(t *testing.T) {
	testCases := []struct {
		name     string
		password parsedPassword
		want     bool
	}{
		{
			name: "valid password",
			password: parsedPassword{
				min:      1,
				max:      3,
				letter:   'a',
				password: "aaa",
			},
			want: true,
		},
		{
			name: "too many b's",
			password: parsedPassword{
				min:      1,
				max:      3,
				letter:   'b',
				password: "bbbb",
			},
			want: false,
		},
		{
			name: "too few b's",
			password: parsedPassword{
				min:      2,
				max:      9,
				letter:   'b',
				password: "b",
			},
			want: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := isValidPassword(tc.password)
			if got != tc.want {
				t.Errorf("isValidPassword(%v) = %v, want %v", tc.password, got, tc.want)
			}
		})
	}
}
func TestParseLine(t *testing.T) {
	testCases := []struct {
		name    string
		line    string
		want    parsedPassword
		wantErr bool
	}{
		{
			name: "valid line",
			line: "1-3 a: mypassword",
			want: parsedPassword{
				min:      1,
				max:      3,
				letter:   'a',
				password: "abcde",
			},
			wantErr: false,
		},
		{
			name: "empty line",
			line: "",
			wantErr: true,
		},
		{
			name:    "invalid syntax",
			line:    "invalid line",
			wantErr: true,
		},
		{
			name:    "invalid min",
			line:    "a-3 b: abcde",
			wantErr: true,
		},
		{
			name:    "invalid max",
			line:    "1-b c: abcde",
			wantErr: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			_, err := parseLine(tc.line)
			if tc.wantErr {
				if err == nil {
					t.Errorf("parseLine(%v) did not return an error", tc.line)
				}
			} else {
				if err != nil {
					t.Errorf("parseLine(%v) returned error", tc.line)
				}
			}
		})
	}
}
