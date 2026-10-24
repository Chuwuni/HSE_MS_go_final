package main

import "testing"

func TestSign(t *testing.T) {
	cases := map[int]string{5: "положительное", -5: "отрицательное", 0: "ноль"}
	for n, want := range cases {
		if got := sign(n); got != want {
			t.Errorf("sign(%d) = %q; want %q", n, got, want)
		}
	}
}

func TestParity(t *testing.T) {
	cases := map[int]string{4: "чётное", 7: "нечётное", -3: "нечётное", -8: "чётное", 0: "чётное"}
	for n, want := range cases {
		if got := parity(n); got != want {
			t.Errorf("parity(%d) = %q; want %q", n, got, want)
		}
	}
}

func TestDivisibility(t *testing.T) {
	cases := map[int]string{
		30:  "делится и на 3, и на 5 (кратно 15)",
		-45: "делится и на 3, и на 5 (кратно 15)",
		9:   "делится на 3, но не делится на 5",
		10:  "делится на 5, но не делится на 3",
		7:   "не делится ни на 3, ни на 5",
	}
	for n, want := range cases {
		if got := divisibility(n); got != want {
			t.Errorf("divisibility(%d) = %q; want %q", n, got, want)
		}
	}
}
