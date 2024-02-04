package pointerserrors

import (
	"errors"
	"fmt"
)

type Stringer interface {
	String() string
}

type Chichicoin int

func (c Chichicoin) String() string {
	return fmt.Sprintf("%d CTC", c)
}

type Wallet struct {
	balance Chichicoin
}

// adding * to the type indicates we use a pointer to the instance in use
// This prevents creating a copy of wallet
func (w *Wallet) Deposit(amount Chichicoin) {
	// using & allows to get the pointer
	fmt.Printf("address of balance in Deposit is %p \n", &w.balance)
	w.balance += amount
}

func (w *Wallet) Balance() Chichicoin {
	return w.balance
}

var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")

func (w *Wallet) Withdraw(amount Chichicoin) error {
	if amount > w.balance {
		return ErrInsufficientFunds
	}

	w.balance -= amount
	return nil
}
