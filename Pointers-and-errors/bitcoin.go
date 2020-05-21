package pointers

import "fmt"

// Bitcoin type to be used in wallets
type Bitcoin int

// String method that appends BTC to the value
func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}
