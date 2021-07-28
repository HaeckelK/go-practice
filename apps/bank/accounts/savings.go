package accounts

type SavingsAccount struct {
	balance int64
	name    string
}

func NewSavingsAccount(balance int64) *SavingsAccount {
	account := SavingsAccount{balance: balance}
	account.name = "Savings Account"
	return &account
}

func (s *SavingsAccount) Deposit(amount int64) {
	s.balance += amount
}

func (s *SavingsAccount) Withdraw(amount int64) {
	s.balance -= amount
}

func (s SavingsAccount) Balance() int64 {
	return s.balance
}
