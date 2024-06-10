package accounts

import (
	"fmt"
	"gobank/customers"
)

type CheckingAccount struct {
	AccountHolder customers.Holder
	BranchNumber  int
	AccountNumber int
	balance       float64
}

func (account *CheckingAccount) Withdraw(amount float64) (string, float64) {
	var canWithdraw = account.balance < amount && amount > 0

	if canWithdraw {
		fmt.Println("Impossible to withdraw.")
		return "Impossible to withdraw.", account.balance
	}

	account.balance -= amount
	fmt.Println("Withdrawal successful")
	return "Withdrawal successful", account.balance
}

func (account *CheckingAccount) Deposit(amount float64) (string, float64) {
	var canDeposit = amount > 0

	if canDeposit {
		account.balance += amount
		fmt.Println("Deposit successful")
		return "Deposit successful", account.balance
	}

	fmt.Println("Impossible to deposit.")
	return "Impossible to deposit.", account.balance
}

func (account *CheckingAccount) Transfer(amount float64, destinationAccount *CheckingAccount) (string, float64) {
	var canTransfer = account.balance < amount && amount > 0

	if canTransfer {
		fmt.Println("Impossible to transfer.")
		return "Impossible to transfer.", account.balance
	}

	account.balance -= amount
	destinationAccount.balance += amount
	fmt.Println("Transfer successful")
	return "Transfer successful", account.balance
}
