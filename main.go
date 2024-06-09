package main

import "fmt"

type CheckingAccount struct {
	accountHolder string
	branchNumber  int
	accountNumber int
	balance       float64
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
}
