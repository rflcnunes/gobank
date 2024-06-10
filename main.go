package main

import (
	"fmt"
	"gobank/accounts"
	"gobank/customers"
)

func main() {
	johns := customers.Holder{
		Name:           "John Doe",
		DocumentNumber: "12345678900",
		Job:            "Software Developer I",
	}

	maryJane := customers.Holder{
		Name:           "Mary Jane",
		DocumentNumber: "12345678901",
		Job:            "Software Developer II",
	}

	johnsAccount := accounts.CheckingAccount{
		AccountHolder: johns,
		BranchNumber:  1234,
		AccountNumber: 5678,
	}

	marysAccount := accounts.CheckingAccount{
		AccountHolder: maryJane,
		BranchNumber:  1234,
		AccountNumber: 5679,
	}

	fmt.Println("Accounts\n", johnsAccount, "\n", marysAccount)
}
