package payments

import "gobank/accounts"

func PayBillet(account accounts.Account, amount float64) (string, float64) {
	return account.Withdraw(amount)
}
