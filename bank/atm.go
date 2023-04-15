package bank

type ATM struct {
	BankName string
}

type Withdrawer interface {
	Withdraw(amount int) error
}

type WithdrawDepsotable interface {
	Withdrawer
	Deposit(amount int) error
}

func (atm ATM) CashWithdraw(card Withdrawer, amount int) {
	card.Withdraw(amount)
}
