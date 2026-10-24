// Задание 6. Количество положительных и отрицательных чисел.
package main

import (
	"fmt"

	"go-basics/internal/console"
)

func countSigns(nums []float64) (pos, neg, zero int) {
	for _, v := range nums {
		switch {
		case v > 0:
			pos++
		case v < 0:
			neg++
		default:
			zero++
		}
	}
	return pos, neg, zero
}

func main() {
	n := console.Int("Сколько чисел будет введено (от 1 до 1000): ", 1, 1000)
	nums := console.Floats(n)

	pos, neg, zero := countSigns(nums)
	fmt.Println("Положительных:", pos)
	fmt.Println("Отрицательных:", neg)
	fmt.Println("Нулей:        ", zero)
}
