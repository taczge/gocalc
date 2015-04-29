package main

import "testing"

func TestEvaluate(t *testing.T) {
	cases := []struct {
		in   string
		want int
	}{
		{"0", 0},
		{"1 2 +", 3},
		{"2 3 *", 6},
		{"2 3 + 3 *", 15},
	}

	for _, c := range cases {
		got, _ := Evaluate(c.in)

		if got != c.want {
			t.Errorf("in: %s => got: %d, want: %d", c.in, got, c.want)
		}
	}
}
