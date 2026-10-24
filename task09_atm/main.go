// Задание 9. Мини-банкомат.
package main

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"go-basics/internal/console"
)

const maxOperation = 1_000_000_00

var (
	ErrNonPositive  = errors.New("сумма должна быть больше нуля")
	ErrInsufficient = errors.New("недостаточно средств")
	amountPattern   = regexp.MustCompile(`^\d{1,9}(\.\d{1,2})?$`)
)

type Account struct {
	balance int64
	history []string
}

func NewAccount(balance int64) *Account {
	return &Account{balance: balance}
}

func (a *Account) Balance() int64 {
	return a.balance
}

func (a *Account) History() []string {
	return a.history
}

func (a *Account) Deposit(amount int64) error {
	if amount <= 0 {
		return ErrNonPositive
	}
	a.balance += amount
	a.history = append(a.history, "Пополнение: +"+formatMoney(amount))
	return nil
}

func (a *Account) Withdraw(amount int64) error {
	if amount <= 0 {
		return ErrNonPositive
	}
	if amount > a.balance {
		return fmt.Errorf("%w: на счёте %s, запрошено %s", ErrInsufficient, formatMoney(a.balance), formatMoney(amount))
	}
	a.balance -= amount
	a.history = append(a.history, "Снятие: -"+formatMoney(amount))
	return nil
}

func parseMoney(s string) (int64, error) {
	s = strings.ReplaceAll(strings.TrimSpace(s), ",", ".")
	if !amountPattern.MatchString(s) {
		return 0, errors.New("введите сумму в формате 1500 или 1500.50")
	}
	rubles, kopecks, _ := strings.Cut(s, ".")
	if len(kopecks) == 1 {
		kopecks += "0"
	}
	r, _ := strconv.ParseInt(rubles, 10, 64)
	var k int64
	if kopecks != "" {
		k, _ = strconv.ParseInt(kopecks, 10, 64)
	}
	return r*100 + k, nil
}

func parseAmount(s string) (int64, error) {
	amount, err := parseMoney(s)
	if err != nil {
		return 0, err
	}
	if amount == 0 {
		return 0, ErrNonPositive
	}
	if amount > maxOperation {
		return 0, fmt.Errorf("сумма одной операции не может превышать %s", formatMoney(maxOperation))
	}
	return amount, nil
}

func formatMoney(kopecks int64) string {
	return fmt.Sprintf("%d.%02d руб.", kopecks/100, kopecks%100)
}

func readMoney(prompt string, parse func(string) (int64, error)) int64 {
	for {
		amount, err := parse(console.Line(prompt))
		if err == nil {
			return amount
		}
		fmt.Println("Ошибка:", err)
	}
}

func main() {
	acc := NewAccount(readMoney("Введите начальный баланс: ", parseMoney))

	for {
		fmt.Println()
		fmt.Println("1 - Показать баланс")
		fmt.Println("2 - Пополнить счёт")
		fmt.Println("3 - Снять наличные")
		fmt.Println("4 - История операций")
		fmt.Println("0 - Выход")

		switch console.Line("Выберите операцию: ") {
		case "1":
			fmt.Println("Текущий баланс:", formatMoney(acc.Balance()))
		case "2":
			amount := readMoney("Сумма пополнения: ", parseAmount)
			if err := acc.Deposit(amount); err != nil {
				fmt.Println("Ошибка:", err)
			} else {
				fmt.Println("Счёт пополнен. Баланс:", formatMoney(acc.Balance()))
			}
		case "3":
			amount := readMoney("Сумма снятия: ", parseAmount)
			if err := acc.Withdraw(amount); err != nil {
				fmt.Println("Ошибка:", err)
			} else {
				fmt.Printf("Выдано %s Баланс: %s\n", formatMoney(amount), formatMoney(acc.Balance()))
			}
		case "4":
			if len(acc.History()) == 0 {
				fmt.Println("Операций пока не было.")
			}
			for i, h := range acc.History() {
				fmt.Printf("%d. %s\n", i+1, h)
			}
		case "0":
			fmt.Println("Работа завершена.")
			return
		default:
			fmt.Println("Неизвестная операция, повторите выбор.")
		}
	}
}
