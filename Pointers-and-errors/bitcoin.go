package pointers

import "fmt"

// Bitcoin type to be used in wallets
type Bitcoin int

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}
