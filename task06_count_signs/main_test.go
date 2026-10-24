package main

import "testing"

func TestCountSigns(t *testing.T) {
	cases := []struct {
		nums           []float64
		pos, neg, zero int
	}{
		{[]float64{1, -1, 0, 0, 3}, 2, 1, 2},
		{[]float64{0}, 0, 0, 1},
		{[]float64{-0.5, -2}, 0, 2, 0},
		{[]float64{0.001, 7}, 2, 0, 0},
	}
	for _, c := range cases {
		pos, neg, zero := countSigns(c.nums)
		if pos != c.pos || neg != c.neg || zero != c.zero {
			t.Errorf("countSigns(%v) = %d, %d, %d; want %d, %d, %d", c.nums, pos, neg, zero, c.pos, c.neg, c.zero)
		}
	}
}
