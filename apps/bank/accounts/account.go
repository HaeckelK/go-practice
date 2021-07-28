package accounts

type Account interface {
	Deposit(amount int64)
	Withdraw(amount int64)
	Balance() (amount int64)
}
