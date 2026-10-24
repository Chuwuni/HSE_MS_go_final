// Задание 4. Таблица умножения через функцию.
package main

import (
	"fmt"
	"io"
	"os"
	"strconv"

	"go-basics/internal/console"
)

func printTable(w io.Writer, n, limit int) {
	width := len(strconv.Itoa(limit))
	fmt.Fprintf(w, "Таблица умножения для числа %d:\n", n)
	for i := 1; i <= limit; i++ {
		fmt.Fprintf(w, "%d x %*d = %d\n", n, width, i, n*i)
	}
}

func main() {
	n := console.Int("Введите число: ", -1_000_000, 1_000_000)
	limit := console.Int("До какого множителя выводить (от 1 до 100): ", 1, 100)
	printTable(os.Stdout, n, limit)
}
