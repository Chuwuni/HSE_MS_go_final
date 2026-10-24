package main

import "testing"

func TestSumAndAverage(t *testing.T) {
	cases := []struct {
		n       int
		sum     int
		average float64
	}{
		{1, 1, 1},
		{2, 3, 1.5},
		{10, 55, 5.5},
		{100, 5050, 50.5},
	}
	for _, c := range cases {
		sum, avg := sumAndAverage(c.n)
		if sum != c.sum || avg != c.average {
			t.Errorf("sumAndAverage(%d) = %d, %v; want %d, %v", c.n, sum, avg, c.sum, c.average)
		}
	}
}

func TestSumMatchesFormula(t *testing.T) {
	for _, n := range []int{7, 999, maxN} {
		if sum, _ := sumAndAverage(n); sum != n*(n+1)/2 {
			t.Errorf("sumAndAverage(%d) = %d; формула даёт %d", n, sum, n*(n+1)/2)
		}
	}
}
