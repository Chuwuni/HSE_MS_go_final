// Задание 1. Калькулятор с проверкой.
package main

import (
	"errors"
	"fmt"
	"math"
	"strings"

	"go-basics/internal/console"
)

var errDivByZero = errors.New("деление на ноль невозможно")

var unary = map[string]bool{
	"+": false, "-": false, "*": false, "/": false, "%": false, "^": false,
	"sqrt": true, "sq": true,
}

func calculate(op string, a, b float64) (float64, error) {
	var r float64
	switch op {
	case "+":
		r = a + b
	case "-":
		r = a - b
	case "*":
		r = a * b
	case "/":
		if b == 0 {
			return 0, errDivByZero
		}
		r = a / b
	case "%":
		if b == 0 {
			return 0, errDivByZero
		}
		r = math.Mod(a, b)
	case "^":
		r = math.Pow(a, b)
	case "sq":
		r = a * a
	case "sqrt":
		if a < 0 {
			return 0, errors.New("квадратный корень из отрицательного числа не определён")
		}
		r = math.Sqrt(a)
	default:
		return 0, fmt.Errorf("неизвестная операция %q", op)
	}
	if math.IsNaN(r) {
		return 0, errors.New("результат не определён")
	}
	if math.IsInf(r, 0) {
		return 0, errors.New("результат слишком велик")
	}
	return r, nil
}

func expression(op string, a, b float64) string {
	sa, sb := console.FormatFloat(a), console.FormatFloat(b)
	switch op {
	case "sqrt":
		return "sqrt(" + sa + ")"
	case "sq":
		if a < 0 {
			sa = "(" + sa + ")"
		}
		return sa + "^2"
	default:
		if b < 0 {
			sb = "(" + sb + ")"
		}
		return sa + " " + op + " " + sb
	}
}

func readOperation() string {
	for {
		op := strings.ToLower(console.Line("Введите операцию: "))
		if _, ok := unary[op]; ok {
			return op
		}
		fmt.Println("Ошибка: неизвестная операция, повторите ввод.")
	}
}

func main() {
	fmt.Println("Калькулятор")
	fmt.Println("Операции с двумя числами: + - * / % (остаток) ^ (степень)")
	fmt.Println("Операции с одним числом:  sqrt (квадратный корень), sq (квадрат)")

	for {
		fmt.Println()
		a := console.Float("Введите число: ")
		op := readOperation()
		var b float64
		if !unary[op] {
			b = console.Float("Введите второе число: ")
		}

		if r, err := calculate(op, a, b); err != nil {
			fmt.Println("Ошибка:", err)
		} else {
			fmt.Printf("Результат: %s = %s\n", expression(op, a, b), console.FormatFloat(r))
		}

		answer := strings.ToLower(console.Line("Выполнить ещё одно вычисление? (д/н): "))
		if answer != "д" && answer != "да" && answer != "y" && answer != "yes" {
			return
		}
	}
}
