package accounts

type CheckingAccount struct {
	balance int64
	name    string
}

func NewCheckingAccount(balance int64) *CheckingAccount {
	account := CheckingAccount{balance: balance}
	account.name = "Checking Account"
	return &account
}

func (c *CheckingAccount) Deposit(amount int64) {
	c.balance += amount
}

func (c *CheckingAccount) Withdraw(amount int64) {
	c.balance -= amount
}

func (c CheckingAccount) Balance() int64 {
	return c.balance
}
