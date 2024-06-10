package accounts

import "gobank/customers"

type SavingsAccount struct {
	AccountHolder                          customers.Holder
	BranchNumber, AccountNumber, Operation int
	balance                                float64
}

func (account *SavingsAccount) Withdraw(amount float64) (string, float64) {
	var canWithdraw = account.balance < amount && amount > 0

	if canWithdraw {
		return "Impossible to withdraw.", account.balance
	}

	account.balance -= amount
	return "Withdrawal successful", account.balance
}

func (account *SavingsAccount) Deposit(amount float64) (string, float64) {
	var canDeposit = amount > 0

	if canDeposit {
		account.balance += amount
		return "Deposit successful", account.balance
	}

	return "Impossible to deposit.", account.balance
}

func (account *SavingsAccount) GetBalance() float64 {
	return account.balance
}
