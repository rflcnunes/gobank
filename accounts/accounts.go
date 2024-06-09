package accounts

import (
	"fmt"
	"gobank/customers"
)

type CheckingAccount struct {
	AccountHolder customers.Holder
	BranchNumber  int
	AccountNumber int
	Balance       float64
}

func (account *CheckingAccount) Withdraw(amount float64) (string, float64) {
	var canWithdraw = account.Balance < amount && amount > 0

	if canWithdraw {
		fmt.Println("Impossible to withdraw.")
		return "Impossible to withdraw.", account.Balance
	}

	account.Balance -= amount
	fmt.Println("Withdrawal successful")
	return "Withdrawal successful", account.Balance
}

func (account *CheckingAccount) Deposit(amount float64) (string, float64) {
	var canDeposit = amount > 0

	if canDeposit {
		account.Balance += amount
		fmt.Println("Deposit successful")
		return "Deposit successful", account.Balance
	}

	fmt.Println("Impossible to deposit.")
	return "Impossible to deposit.", account.Balance
}

func (account *CheckingAccount) Transfer(amount float64, destinationAccount *CheckingAccount) (string, float64) {
	var canTransfer = account.Balance < amount && amount > 0

	if canTransfer {
		fmt.Println("Impossible to transfer.")
		return "Impossible to transfer.", account.Balance
	}

	account.Balance -= amount
	destinationAccount.Balance += amount
	fmt.Println("Transfer successful")
	return "Transfer successful", account.Balance
}
