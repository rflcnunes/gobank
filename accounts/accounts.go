package accounts

type Account interface {
	Withdraw(amount float64) (string, float64)
	Deposit(amount float64) (string, float64)
}
