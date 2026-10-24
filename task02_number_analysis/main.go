// Задание 2. Анализ числа.
package main

import (
	"fmt"
	"math"

	"go-basics/internal/console"
)

func sign(n int) string {
	switch {
	case n > 0:
		return "положительное"
	case n < 0:
		return "отрицательное"
	default:
		return "ноль"
	}
}

func parity(n int) string {
	if n%2 == 0 {
		return "чётное"
	}
	return "нечётное"
}

func divisibility(n int) string {
	by3 := n%3 == 0
	by5 := n%5 == 0
	switch {
	case by3 && by5:
		return "делится и на 3, и на 5 (кратно 15)"
	case by3:
		return "делится на 3, но не делится на 5"
	case by5:
		return "делится на 5, но не делится на 3"
	default:
		return "не делится ни на 3, ни на 5"
	}
}

func main() {
	n := console.Int("Введите целое число: ", math.MinInt, math.MaxInt)
	fmt.Printf("%-10s %s\n", "Знак:", sign(n))
	fmt.Printf("%-10s %s\n", "Чётность:", parity(n))
	fmt.Printf("%-10s %s\n", "Делимость:", divisibility(n))
}
