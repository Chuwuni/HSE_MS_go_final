// Package console — чтение и проверка пользовательского ввода, форматирование чисел.
package console

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
)

type Scanner struct {
	r *bufio.Reader
	w io.Writer
}

func New(r io.Reader, w io.Writer) *Scanner {
	return &Scanner{r: bufio.NewReader(r), w: w}
}

func (s *Scanner) Line(prompt string) (string, error) {
	fmt.Fprint(s.w, prompt)
	line, err := s.r.ReadString('\n')
	if err != nil && line == "" {
		return "", err
	}
	return strings.TrimSpace(line), nil
}

func (s *Scanner) Int(prompt string, min, max int) (int, error) {
	for {
		line, err := s.Line(prompt)
		if err != nil {
			return 0, err
		}
		v, convErr := strconv.Atoi(line)
		switch {
		case convErr != nil:
			fmt.Fprintln(s.w, "Ошибка: введите целое число.")
		case v < min || v > max:
			fmt.Fprintf(s.w, "Ошибка: число должно быть в диапазоне от %d до %d.\n", min, max)
		default:
			return v, nil
		}
	}
}

func (s *Scanner) Float(prompt string) (float64, error) {
	for {
		line, err := s.Line(prompt)
		if err != nil {
			return 0, err
		}
		v, convErr := ParseFloat(line)
		if convErr == nil {
			return v, nil
		}
		fmt.Fprintln(s.w, "Ошибка: введите число.")
	}
}

func ParseFloat(s string) (float64, error) {
	v, err := strconv.ParseFloat(strings.ReplaceAll(s, ",", "."), 64)
	if err != nil || math.IsNaN(v) || math.IsInf(v, 0) {
		return 0, errors.New("некорректное число")
	}
	return v, nil
}

func FormatFloat(x float64) string {
	if x == 0 {
		return "0"
	}
	return strconv.FormatFloat(x, 'g', 15, 64)
}

var std = New(os.Stdin, os.Stdout)

func exitOnEOF(err error) {
	if err != nil {
		fmt.Println("\nВвод завершён.")
		os.Exit(0)
	}
}

func Line(prompt string) string {
	v, err := std.Line(prompt)
	exitOnEOF(err)
	return v
}

func Int(prompt string, min, max int) int {
	v, err := std.Int(prompt, min, max)
	exitOnEOF(err)
	return v
}

func Float(prompt string) float64 {
	v, err := std.Float(prompt)
	exitOnEOF(err)
	return v
}

func Floats(n int) []float64 {
	nums := make([]float64, n)
	for i := range nums {
		nums[i] = Float(fmt.Sprintf("Число %d: ", i+1))
	}
	return nums
}

func Ints(n, min, max int) []int {
	nums := make([]int, n)
	for i := range nums {
		nums[i] = Int(fmt.Sprintf("Число %d: ", i+1), min, max)
	}
	return nums
}
