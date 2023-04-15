package main

import (
	"errors"
	"fmt"
	"kloter1/bank"
)

func main() {
	saving := bank.NewSavingAccount("900945")

	fmt.Println("balance before", saving.Balance())
	err := saving.Deposit(50000)
	if err != nil {
		fmt.Println("gagal deposit", err)
	}
	fmt.Println("balance after", saving.Balance())

	err = saving.Withdraw(1000)
	if err != nil {
		fmt.Println("gagal tarik tunai")
	}
	fmt.Println("after withdraw", saving.Balance())

	fmt.Println("=====================")

	credit := bank.NewCreditAccount("800866", 1000000)

	err = credit.Withdraw(100000)
	if err != nil {
		fmt.Println("gagal tarik CC", err)
	}
	fmt.Println("saldo CC", credit.Balance())

	// exceeds limit
	err = credit.Withdraw(1000000)
	if err != nil {
		// possibly error invalid amount/exceeds limit
		fmt.Println("gagal tarik CC", err)

		// CARA YG TIDAK BAIK NGECEK DG STRING ERR
		// if err.Error() == "invalid amount" {
		// 	fmt.Println("nominal harus lebih dr 0")
		// }

		// // cara yg tidak baik, karena error string suatu saat bisa berubah
		// if err.Error() == "exceeds limit" {
		// 	fmt.Println("limitnya habis ")
		// }

		// GOOD - handling dengan check sentinel error
		if errors.Is(err, bank.ErrInvalidAmount) {
			fmt.Println("nominal harus lebih dr 0")
		}

		if errors.Is(err, bank.ErrExceedsLimit) {
			fmt.Println("limitnya habis ")
		}
	}
	fmt.Println("saldo CC 2", credit.Balance())

	fmt.Println("============== ATM")
	atm := bank.ATM{"BCA"}

	// bad - cara tanpa pakai interface
	// atm.CashWithdraw(&saving, nil, 1000)
	// atm.CashWithdraw(nil, &credit, 1000)

	atm.CashWithdraw(&credit, 1000)
	fmt.Println("atm credit.Balance()", credit.Balance())

	atm.CashWithdraw(&saving, 1000)
	fmt.Println("atm saving.Balance()", saving.Balance())

}
