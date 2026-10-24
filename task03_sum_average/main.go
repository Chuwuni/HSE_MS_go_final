// Задание 3. Сумма и среднее от 1 до N.
package main

import (
	"fmt"

	"go-basics/internal/console"
)

const maxN = 1_000_000

func sumAndAverage(n int) (int, float64) {
	sum := 0
	for i := 1; i <= n; i++ {
		sum += i
	}
	return sum, float64(sum) / float64(n)
}

func main() {
	n := console.Int(fmt.Sprintf("Введите N (от 1 до %d): ", maxN), 1, maxN)

	sum, avg := sumAndAverage(n)
	fmt.Printf("Сумма чисел от 1 до %d: %d\n", n, sum)
	fmt.Printf("Среднее значение: %s\n", console.FormatFloat(avg))
	fmt.Printf("Проверка по формуле n(n+1)/2: %d\n", n*(n+1)/2)
}
