package main

import "fmt"

type CheckingAccount struct {
	accountHolder string
	branchNumber  int
	accountNumber int
	balance       float64
}

func (account *CheckingAccount) Withdraw(amount float64) string {
	var canWithdraw bool = account.balance < amount && amount > 0

	if canWithdraw {
		fmt.Println("Impossible to withdraw.")
		return "Impossible to withdraw."
	}

	account.balance -= amount
	fmt.Println("Withdrawal successful")
	return "Withdrawal successful"
}

func main() {
	johnsAccount := CheckingAccount{
		"John Doe",
		1234,
		5678,
		100.89,
	}

	marysAccount := CheckingAccount{
		"Mary Jane",
		1234,
		5679,
		238.43,
	}

	fmt.Println("Checking Accounts\n", johnsAccount, "\n", marysAccount)

	johnsAccount.Withdraw(120.00)

	fmt.Println("Checking Accounts\n", johnsAccount, "\n", marysAccount)
}
