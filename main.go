package main

import (
	"fmt"
	"gobank/accounts"
)

func main() {
	johnsAccount := accounts.CheckingAccount{
		AccountHolder: "John Doe",
		BranchNumber:  1234,
		AccountNumber: 5678,
		Balance:       100.89,
	}

	marysAccount := accounts.CheckingAccount{
		AccountHolder: "Mary Jane",
		BranchNumber:  1234,
		AccountNumber: 5679,
		Balance:       238.43,
	}

	fmt.Println("Checking Accounts\n", johnsAccount, "\n", marysAccount)
	johnsAccount.Withdraw(120.00)
	fmt.Println("Checking Accounts\n", johnsAccount, "\n", marysAccount)
	marysAccount.Deposit(100.00)
	fmt.Println("Checking Accounts\n", johnsAccount, "\n", marysAccount)
	johnsAccount.Transfer(50.00, &marysAccount)
	fmt.Println("Checking Accounts\n", johnsAccount, "\n", marysAccount)
}
