package main

import "fmt"

type CheckingAccount struct {
	accountHolder string
	branchNumber  int
	accountNumber int
	balance       float64
}

func main() {
	fmt.Println("New Checking Account\n",
		CheckingAccount{
			"John Doe",
			1234,
			5678,
			1000.0})
}
