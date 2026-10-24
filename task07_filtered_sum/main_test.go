package main

import "testing"

func TestIsSuitable(t *testing.T) {
	cases := map[int]bool{6: true, 12: true, 18: true, 2: false, 3: false, 9: false, 10: false, 7: false}
	for n, want := range cases {
		if got := isSuitable(n); got != want {
			t.Errorf("isSuitable(%d) = %v; want %v", n, got, want)
		}
	}
}

func TestSumSuitable(t *testing.T) {
	cases := []struct {
		n, sum, count int
	}{
		{5, 0, 0},
		{6, 6, 1},
		{20, 36, 3},
		{30, 90, 5},
	}
	for _, c := range cases {
		sum, count := sumSuitable(c.n)
		if sum != c.sum || count != c.count {
			t.Errorf("sumSuitable(%d) = %d, %d; want %d, %d", c.n, sum, count, c.sum, c.count)
		}
	}
}

func TestSumSuitableMatchesFormula(t *testing.T) {
	n := maxN
	k := n / 6
	if sum, _ := sumSuitable(n); sum != 6*k*(k+1)/2 {
		t.Errorf("sumSuitable(%d) = %d; формула даёт %d", n, sum, 6*k*(k+1)/2)
	}
}
