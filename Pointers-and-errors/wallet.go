package pointers

import "errors"

// ErrInsufficientFunds is thrown when there are insufficient funds
var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")

// Wallet type that represents a real world wallet
type Wallet struct {
	balance Bitcoin
}

// Deposit adds the provided amount to the balance
func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

// Balance returns the amount of money in the wallet
func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

// Withdraw bitcoins from the balance
func (w *Wallet) Withdraw(amount Bitcoin) error {

	if amount > w.balance {
		return ErrInsufficientFunds
	}
	w.balance -= amount
	return nil
}
