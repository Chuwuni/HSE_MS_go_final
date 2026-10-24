// Задание 5. Максимум из N чисел.
package main

import (
	"fmt"

	"go-basics/internal/console"
)

func maxMin(nums []float64) (maxIdx, minIdx int) {
	for i, v := range nums {
		if v > nums[maxIdx] {
			maxIdx = i
		}
		if v < nums[minIdx] {
			minIdx = i
		}
	}
	return maxIdx, minIdx
}

func main() {
	n := console.Int("Сколько чисел будет введено (от 1 до 1000): ", 1, 1000)
	nums := console.Floats(n)

	maxIdx, minIdx := maxMin(nums)
	fmt.Printf("Максимальное значение: %s (число №%d)\n", console.FormatFloat(nums[maxIdx]), maxIdx+1)
	fmt.Printf("Минимальное значение:  %s (число №%d)\n", console.FormatFloat(nums[minIdx]), minIdx+1)
}
