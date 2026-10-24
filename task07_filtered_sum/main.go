// Задание 7. Сумма только подходящих чисел.
package main

import (
	"fmt"

	"go-basics/internal/console"
)

const maxN = 1_000_000

func isSuitable(n int) bool {
	return n%2 == 0 && n%3 == 0
}

func sumSuitable(n int) (sum, count int) {
	for i := 1; i <= n; i++ {
		if isSuitable(i) {
			sum += i
			count++
		}
	}
	return sum, count
}

func main() {
	fmt.Println("Условие отбора: число чётное и делится на 3")
	n := console.Int(fmt.Sprintf("Введите N (от 1 до %d): ", maxN), 1, maxN)

	sum, count := sumSuitable(n)
	if count == 0 {
		fmt.Printf("От 1 до %d подходящих чисел нет, сумма равна 0\n", n)
		return
	}
	fmt.Printf("Подходящих чисел от 1 до %d: %d\n", n, count)
	fmt.Printf("Сумма подходящих чисел: %d\n", sum)
}
