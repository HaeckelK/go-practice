package accounts

import "testing"

func TestCheckingAccountDeposit(t *testing.T) {
	// Given an account with a balance
	account := NewCheckingAccount(50)
	// When amount is deposited
	account.Deposit(25)
	// Then new balance agrees
	want := 75
	got := account.Balance()
	if got != int64(want) {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestCheckingAccountWithdraw(t *testing.T) {
	// Given an account with a balance
	account := NewCheckingAccount(50)
	// When amount is withdrawn
	account.Withdraw(25)
	// Then new balance agrees
	want := 25
	got := account.Balance()
	if got != int64(want) {
		t.Errorf("got %q, wanted %q", got, want)
	}
}
