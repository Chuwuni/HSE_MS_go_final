package main

import (
	"errors"
	"testing"
)

func TestCalculate(t *testing.T) {
	cases := []struct {
		op   string
		a, b float64
		want float64
	}{
		{"+", 2, 3, 5},
		{"-", 2, 3, -1},
		{"*", -4, 2.5, -10},
		{"/", 7, 2, 3.5},
		{"%", 7, 3, 1},
		{"^", 2, 10, 1024},
		{"^", 9, 0.5, 3},
		{"sq", -3, 0, 9},
		{"sqrt", 16, 0, 4},
		{"sqrt", 0, 0, 0},
	}
	for _, c := range cases {
		got, err := calculate(c.op, c.a, c.b)
		if err != nil || got != c.want {
			t.Errorf("calculate(%q, %v, %v) = %v, %v; want %v", c.op, c.a, c.b, got, err, c.want)
		}
	}
}

func TestCalculateErrors(t *testing.T) {
	cases := []struct {
		op   string
		a, b float64
	}{
		{"/", 1, 0},
		{"%", 1, 0},
		{"sqrt", -4, 0},
		{"^", -8, 1.0 / 3},
		{"^", 10, 400},
		{"*", 1e300, 1e300},
		{"&", 1, 1},
	}
	for _, c := range cases {
		if _, err := calculate(c.op, c.a, c.b); err == nil {
			t.Errorf("calculate(%q, %v, %v): ожидалась ошибка", c.op, c.a, c.b)
		}
	}
	if _, err := calculate("/", 5, 0); !errors.Is(err, errDivByZero) {
		t.Errorf("деление на ноль должно возвращать errDivByZero, получено %v", err)
	}
}

func TestExpression(t *testing.T) {
	cases := []struct {
		op   string
		a, b float64
		want string
	}{
		{"+", 1, -2, "1 + (-2)"},
		{"sq", -3, 0, "(-3)^2"},
		{"sqrt", 2, 0, "sqrt(2)"},
	}
	for _, c := range cases {
		if got := expression(c.op, c.a, c.b); got != c.want {
			t.Errorf("expression(%q, %v, %v) = %q; want %q", c.op, c.a, c.b, got, c.want)
		}
	}
}
