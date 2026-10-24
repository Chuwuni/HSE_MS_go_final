package main

import "testing"

func TestMaxMin(t *testing.T) {
	cases := []struct {
		nums           []float64
		maxIdx, minIdx int
	}{
		{[]float64{3, -2.5, 10, 0}, 2, 1},
		{[]float64{42}, 0, 0},
		{[]float64{-1, -7, -3}, 0, 1},
		{[]float64{5, 5, 1, 1}, 0, 2},
	}
	for _, c := range cases {
		maxIdx, minIdx := maxMin(c.nums)
		if maxIdx != c.maxIdx || minIdx != c.minIdx {
			t.Errorf("maxMin(%v) = %d, %d; want %d, %d", c.nums, maxIdx, minIdx, c.maxIdx, c.minIdx)
		}
	}
}
