package main

import (
	"errors"
	"testing"
)

func TestParseAmount(t *testing.T) {
	valid := map[string]int64{
		"1500":    1500_00,
		"1500.5":  1500_50,
		"1500,50": 1500_50,
		"0.01":    1,
		" 20 ":    20_00,
	}
	for in, want := range valid {
		if got, err := parseAmount(in); err != nil || got != want {
			t.Errorf("parseAmount(%q) = %d, %v; want %d", in, got, err, want)
		}
	}
	invalid := []string{"", "0", "0.00", "-5", "abc", "inf", "NaN", "1.234", "1e5", "10000000.01", "999999999"}
	for _, in := range invalid {
		if _, err := parseAmount(in); err == nil {
			t.Errorf("parseAmount(%q): ожидалась ошибка", in)
		}
	}
}

func TestParseMoney(t *testing.T) {
	valid := map[string]int64{"0": 0, "0.00": 0, "999999999.99": 999999999_99, "1000": 1000_00}
	for in, want := range valid {
		if got, err := parseMoney(in); err != nil || got != want {
			t.Errorf("parseMoney(%q) = %d, %v; want %d", in, got, err, want)
		}
	}
	for _, in := range []string{"", "-1", "abc", "1.234", "1000000000"} {
		if _, err := parseMoney(in); err == nil {
			t.Errorf("parseMoney(%q): ожидалась ошибка", in)
		}
	}
}

func TestDepositWithdraw(t *testing.T) {
	acc := NewAccount(1000_00)
	if err := acc.Withdraw(10); err != nil {
		t.Fatal(err)
	}
	if err := acc.Deposit(30); err != nil {
		t.Fatal(err)
	}
	if err := acc.Withdraw(1000_20); err != nil {
		t.Fatalf("снятие всего баланса 1000.20 должно пройти: %v", err)
	}
	if acc.Balance() != 0 {
		t.Errorf("баланс = %d; want 0", acc.Balance())
	}
	if len(acc.History()) != 3 {
		t.Errorf("в истории %d записей; want 3", len(acc.History()))
	}
}

func TestWithdrawInsufficient(t *testing.T) {
	acc := NewAccount(100_00)
	err := acc.Withdraw(100_01)
	if !errors.Is(err, ErrInsufficient) {
		t.Fatalf("ожидалась ErrInsufficient, получено %v", err)
	}
	if acc.Balance() != 100_00 {
		t.Errorf("баланс изменился после отказа: %d", acc.Balance())
	}
}

func TestNonPositive(t *testing.T) {
	acc := NewAccount(100_00)
	if !errors.Is(acc.Deposit(0), ErrNonPositive) || !errors.Is(acc.Withdraw(-1), ErrNonPositive) {
		t.Error("для неположительных сумм ожидалась ErrNonPositive")
	}
}

func TestFormatMoney(t *testing.T) {
	cases := map[int64]string{0: "0.00 руб.", 5: "0.05 руб.", 1000_20: "1000.20 руб."}
	for in, want := range cases {
		if got := formatMoney(in); got != want {
			t.Errorf("formatMoney(%d) = %q; want %q", in, got, want)
		}
	}
}
