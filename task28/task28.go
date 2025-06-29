package main

import (
	"errors"
	"fmt"
)

type BankAccount struct {
	ID      string
	Owner   string
	Balance float64
}

func NewBankAccount(id, owner string, initialAmount float64) *BankAccount {
	return &BankAccount{
		ID:      id,
		Owner:   owner,
		Balance: initialAmount,
	}
}

func (ba *BankAccount) AddFunds(amount float64) {
	if amount <= 0 {
		fmt.Println("Ошибка: сумма для пополнения должна быть больше нуля")
		return
	}
	ba.Balance += amount
	fmt.Printf("Счет %s успешно пополнен на %.2f\n", ba.ID, amount)
}

func (ba *BankAccount) WithdrawFunds(amount float64) error {
	if amount <= 0 {
		return errors.New("сумма для снятия должна быть положительной")
	}
	if amount > ba.Balance {
		return fmt.Errorf("недостаточно средств на счете. Текущий баланс: %.2f", ba.Balance)
	}
	ba.Balance -= amount
	fmt.Printf("Со счета %s снято %.2f\n", ba.ID, amount)
	return nil
}

func (ba *BankAccount) GetBalance() float64 {
	return ba.Balance
}

func (ba *BankAccount) DisplayInfo() {
	fmt.Println("\n--- Информация о банковском счете ---")
	fmt.Printf("Идентификатор счета: %s\n", ba.ID)
	fmt.Printf("Владелец счета: %s\n", ba.Owner)
	fmt.Printf("Текущий баланс: %.2f\n", ba.GetBalance())
}

func main() {
	account := NewBankAccount("1234567890", "Иван Петров", 1000.0)
	account.DisplayInfo()

	fmt.Println("\n=== Выполнение операций ===")
	account.AddFunds(500)
	account.AddFunds(-50)

	if err := account.WithdrawFunds(300); err != nil {
		fmt.Println("Ошибка при снятии:", err)
	}

	if err := account.WithdrawFunds(2000); err != nil {
		fmt.Println("Ошибка при снятии:", err)
	}

	if err := account.WithdrawFunds(-100); err != nil {
		fmt.Println("Ошибка при снятии:", err)
	}

	account.DisplayInfo()
}
