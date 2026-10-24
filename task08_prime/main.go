// Задание 8. Проверка простого числа.
package main

import (
	"fmt"
	"math"

	"go-basics/internal/console"
)

func smallestDivisor(n int) int {
	if n%2 == 0 {
		return 2
	}
	for d := 3; d <= n/d; d += 2 {
		if n%d == 0 {
			return d
		}
	}
	return n
}

func isPrime(n int) bool {
	return n >= 2 && smallestDivisor(n) == n
}

func main() {
	n := console.Int("Введите целое число: ", math.MinInt, math.MaxInt)
	switch {
	case n < 2:
		fmt.Printf("Число %d не является простым: простые числа больше 1\n", n)
	case isPrime(n):
		fmt.Printf("Число %d является простым\n", n)
	default:
		fmt.Printf("Число %d не является простым: делится на %d\n", n, smallestDivisor(n))
	}
}
