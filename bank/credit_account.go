package bank

import (
	"errors"
)

// sentinel error, harus di export jika ingin dipake module lain
var (
	ErrInvalidAmount = errors.New("invalid amount")
	ErrExceedsLimit = errors.New("exceeds limit")
)

type CreditAccount struct {
	id      string
	balance int
}

func NewCreditAccount(id string, limit int) CreditAccount {
	return CreditAccount{
		id:      id,
		balance: limit,
	}
}

// method - attach function ke struct saving account
// method receiver memakai pointer karena ada mutasi data balance
func (sa *CreditAccount) Withdraw(amount int) error {
	if amount <= 0 {
		return ErrInvalidAmount
	}
	if amount > sa.balance {
		return ErrExceedsLimit
	}
	sa.balance = sa.balance - amount
	return nil
}

func (sa CreditAccount) Balance() int {
	return sa.balance
}
