package pointers

import "fmt"

// Bitcoin type to be used in wallets
type Bitcoin int

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

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
