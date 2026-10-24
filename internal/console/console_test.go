package console

import (
	"errors"
	"io"
	"math"
	"strings"
	"testing"
)

func TestIntRetriesUntilValid(t *testing.T) {
	var out strings.Builder
	s := New(strings.NewReader("abc\n500\n42\n"), &out)
	v, err := s.Int("> ", 1, 100)
	if err != nil || v != 42 {
		t.Fatalf("Int() = %d, %v; want 42, nil", v, err)
	}
	if !strings.Contains(out.String(), "введите целое число") || !strings.Contains(out.String(), "в диапазоне") {
		t.Errorf("нет сообщений об ошибке ввода: %q", out.String())
	}
}

func TestFloatRejectsInfAndNaN(t *testing.T) {
	s := New(strings.NewReader("inf\nNaN\n1e400\n2,5\n"), io.Discard)
	v, err := s.Float("> ")
	if err != nil || v != 2.5 {
		t.Fatalf("Float() = %v, %v; want 2.5, nil", v, err)
	}
}

func TestEOF(t *testing.T) {
	s := New(strings.NewReader(""), io.Discard)
	if _, err := s.Int("> ", 0, 10); !errors.Is(err, io.EOF) {
		t.Fatalf("ожидалась io.EOF, получено %v", err)
	}
}

func TestLastLineWithoutNewline(t *testing.T) {
	s := New(strings.NewReader("7"), io.Discard)
	if v, err := s.Int("> ", 0, 10); err != nil || v != 7 {
		t.Fatalf("Int() = %d, %v; want 7, nil", v, err)
	}
}

func TestFormatFloat(t *testing.T) {
	cases := []struct {
		in   float64
		want string
	}{
		{1e6, "1000000"},
		{0.1 + 0.2, "0.3"},
		{-2.5, "-2.5"},
		{0, "0"},
		{math.Copysign(0, -1), "0"},
		{1e20, "1e+20"},
		{1.0 / 3.0, "0.333333333333333"},
		{123456789.5, "123456789.5"},
	}
	for _, c := range cases {
		if got := FormatFloat(c.in); got != c.want {
			t.Errorf("FormatFloat(%v) = %q; want %q", c.in, got, c.want)
		}
	}
}
