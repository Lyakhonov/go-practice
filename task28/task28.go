package main

import (
	"errors"
	"fmt"
)

type BankAccount struct {
	AccountNumber string
	OwnerName     string
	Balance       float64
}

func NewBankAccount(accountNumber, ownerName string, initialBalance float64) *BankAccount {
	return &BankAccount{
		AccountNumber: accountNumber,
		OwnerName:     ownerName,
		Balance:       initialBalance,
	}
}

func (acc *BankAccount) Deposit(amount float64) {
	if amount > 0 {
		acc.Balance += amount
		fmt.Printf("Счет %s: пополнение на %.2f\n", acc.AccountNumber, amount)
	} else {
		fmt.Println("Ошибка: сумма пополнения должна быть положительной")
	}
}

func (acc *BankAccount) Withdraw(amount float64) error {
	if amount <= 0 {
		return errors.New("сумма снятия должна быть положительной")
	}
	if amount > acc.Balance {
		return fmt.Errorf("недостаточно средств. Текущий баланс: %.2f", acc.Balance)
	}
	acc.Balance -= amount
	fmt.Printf("Счет %s: снятие %.2f\n", acc.AccountNumber, amount)
	return nil
}

func (acc *BankAccount) CheckBalance() float64 {
	return acc.Balance
}

func (acc *BankAccount) PrintInfo() {
	fmt.Printf("\nИнформация о счете:\n")
	fmt.Printf("Номер счета: %s\n", acc.AccountNumber)
	fmt.Printf("Владелец: %s\n", acc.OwnerName)
	fmt.Printf("Баланс: %.2f\n", acc.CheckBalance())
}

func main() {
	account := NewBankAccount("1234567890", "Иван Петров", 1000.0)
	account.PrintInfo()

	fmt.Println("\n===== Операции =====")
	account.Deposit(500.0)
	account.Deposit(-100.0)

	err := account.Withdraw(300.0)
	if err != nil {
		fmt.Println("Ошибка при снятии:", err)
	}

	err = account.Withdraw(2000.0)
	if err != nil {
		fmt.Println("Ошибка при снятии:", err)
	}

	err = account.Withdraw(-100.0)
	if err != nil {
		fmt.Println("Ошибка при снятии:", err)
	}
	account.PrintInfo()
}
