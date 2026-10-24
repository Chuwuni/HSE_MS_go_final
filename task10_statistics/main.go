// Задание 10. Статистика последовательности.
package main

import (
	"fmt"
	"math"
	"sort"

	"go-basics/internal/console"
)

type Stats struct {
	Count    int
	Sum      int
	Mean     float64
	Max      int
	Min      int
	Even     int
	Odd      int
	Range    int
	Median   float64
	Variance float64
	StdDev   float64
}

func calcStats(nums []int) Stats {
	s := Stats{Count: len(nums), Max: nums[0], Min: nums[0]}
	for _, v := range nums {
		s.Sum += v
		if v > s.Max {
			s.Max = v
		}
		if v < s.Min {
			s.Min = v
		}
		if v%2 == 0 {
			s.Even++
		} else {
			s.Odd++
		}
	}
	s.Mean = float64(s.Sum) / float64(s.Count)
	s.Range = s.Max - s.Min

	for _, v := range nums {
		d := float64(v) - s.Mean
		s.Variance += d * d
	}
	s.Variance /= float64(s.Count)
	s.StdDev = math.Sqrt(s.Variance)

	sorted := make([]int, len(nums))
	copy(sorted, nums)
	sort.Ints(sorted)
	mid := len(sorted) / 2
	if len(sorted)%2 == 0 {
		s.Median = float64(sorted[mid-1]+sorted[mid]) / 2
	} else {
		s.Median = float64(sorted[mid])
	}
	return s
}

func main() {
	n := console.Int("Количество чисел (от 1 до 1000): ", 1, 1000)
	nums := console.Ints(n, -1_000_000_000, 1_000_000_000)
	s := calcStats(nums)
	f := console.FormatFloat
	row := func(label string, value any) { fmt.Printf("%-25s %v\n", label+":", value) }

	fmt.Println()
	row("Сумма", s.Sum)
	row("Среднее значение", f(s.Mean))
	row("Максимальное значение", s.Max)
	row("Минимальное значение", s.Min)
	row("Количество чётных", s.Even)
	row("Количество нечётных", s.Odd)
	fmt.Println()
	fmt.Println("Дополнительно:")
	row("Количество чисел", s.Count)
	row("Размах", s.Range)
	row("Медиана", f(s.Median))
	row("Дисперсия", fmt.Sprintf("%.4f", s.Variance))
	row("Стандартное отклонение", fmt.Sprintf("%.4f", s.StdDev))
}
