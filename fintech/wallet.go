package fintech

import (
	"errors"
	"fmt"
)

type BitCoin int

type Wallet struct {
	balance BitCoin
}

type Stringer interface {
	String() string
}

func (b BitCoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")

func (w *Wallet) Deposit(a BitCoin) {
	//fmt.Printf("address of balance in Deposit is %v\n", &w.balance)
	//(*w).balance += a
	w.balance += a
}

func (w *Wallet) WithDraw(a BitCoin) error {
	if a > w.balance {
		//return errors.New("cannot withdraw, insufficient funds")
		return ErrInsufficientFunds
	}
	w.balance -= a
	return nil
}

func (w *Wallet) Balance() BitCoin {
	//return (*w).balance
	return w.balance
}
