package main

import "testing"

func TestIsPrime(t *testing.T) {
	primes := []int{2, 3, 5, 7, 13, 17, 97, 7919, 1_000_000_007, 2_147_483_647}
	for _, n := range primes {
		if !isPrime(n) {
			t.Errorf("isPrime(%d) = false; want true", n)
		}
	}
	composites := []int{-7, -1, 0, 1, 4, 9, 15, 25, 49, 100, 7917, 1_000_000_007 * 3, 9_223_372_036_854_775_807}
	for _, n := range composites {
		if isPrime(n) {
			t.Errorf("isPrime(%d) = true; want false", n)
		}
	}
}

func TestSmallestDivisor(t *testing.T) {
	cases := map[int]int{4: 2, 9: 3, 35: 5, 49: 7, 97: 97, 9_223_372_036_854_775_807: 7}
	for n, want := range cases {
		if got := smallestDivisor(n); got != want {
			t.Errorf("smallestDivisor(%d) = %d; want %d", n, got, want)
		}
	}
}
