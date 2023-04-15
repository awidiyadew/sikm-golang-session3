package bank

type AccountBulanan struct {
	SavingAccount // anonymous embed
	biayaBulanan  int
}

func (a *AccountBulanan) Withdraw(amount int) string {
	a.Withdraw(amount)
	return ""
}

type AccountBulanan2 struct {
	saving       SavingAccount // embed/composition
	biayaBulanan int
}

func (a *AccountBulanan2) Withdraw(amount int) string {
	a.saving.Withdraw(amount) // call method in saving account
	return ""
}
