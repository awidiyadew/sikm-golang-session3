package bank

import (
	"errors"
)

type SavingAccount struct {
	id      string
	balance int
}

func NewSavingAccount(id string) SavingAccount {
	return SavingAccount{
		id:      id,
		balance: 0,
	}
}

func (sa *SavingAccount) Deposit(amount int) error {
	if amount <= 0 {
		return errors.New("invalid amount")
	}
	sa.balance += amount
	return nil
}

// method - attach function ke struct saving account
func (sa *SavingAccount) Withdraw(amount int) error {
	if amount > sa.balance {
		// ini dipakai kalau temen2 perlu formatiing message dengan variable yg dimiliki
		// err2 := fmt.Errorf("insufficient balance %v %v", sa.balance, amount)

		// ini kalo pake hardcoded string
		return errors.New("insufficient balance")
	}
	sa.balance = sa.balance - amount
	return nil
}

func (sa SavingAccount) Balance() int {
	return sa.balance
}
