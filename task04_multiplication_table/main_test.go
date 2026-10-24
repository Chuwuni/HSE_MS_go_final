package main

import (
	"strings"
	"testing"
)

func TestPrintTable(t *testing.T) {
	var b strings.Builder
	printTable(&b, 7, 10)
	want := `Таблица умножения для числа 7:
7 x  1 = 7
7 x  2 = 14
7 x  3 = 21
7 x  4 = 28
7 x  5 = 35
7 x  6 = 42
7 x  7 = 49
7 x  8 = 56
7 x  9 = 63
7 x 10 = 70
`
	if b.String() != want {
		t.Errorf("printTable(7, 10):\n%s\nожидалось:\n%s", b.String(), want)
	}
}

func TestPrintTableNegative(t *testing.T) {
	var b strings.Builder
	printTable(&b, -3, 2)
	want := "Таблица умножения для числа -3:\n-3 x 1 = -3\n-3 x 2 = -6\n"
	if b.String() != want {
		t.Errorf("printTable(-3, 2) = %q; want %q", b.String(), want)
	}
}
