package test

import "testing"

func TestMath_sum(t *testing.T) {

	result := sum(1, 2)

	if result != 3 {
		t.Fatal("wrong")
	}

}

func TestMath_table_sum(t *testing.T) {

	type testCase struct {
		name     string
		a        int
		b        int
		expected int
	}

	tests := []testCase{
		{name: "positive numbers", a: 2, b: 3, expected: 5},
		{name: "negative numbers", a: -1, b: -1, expected: -2},
		{name: "zero value", a: 5, b: 0, expected: 5},
	}

	for _, tt := range tests {
		t.Run(
			tt.name,
			func(t *testing.T) {
				actual := sum(tt.a, tt.b)
				if actual != tt.expected {
					t.Errorf("Sum(%d, %d) = %d; want %d", tt.a, tt.b, actual, tt.expected)
				}
			},
		)
	}

}
