package pointersanderrors

import (
  "fmt"
  "errors"
)

var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")

type Bitcoin float64

type Stringer interface {
  String() string
}

type Wallet struct {
  balance Bitcoin
}

func (w *Wallet) Deposit(value Bitcoin) {
  w.balance += value
}

func (w *Wallet) Balance() Bitcoin {
  return w.balance
}

func (w *Wallet) Withdraw(value Bitcoin) error {
  if value > w.balance {
    return ErrInsufficientFunds
  }
  w.balance -= value
  return nil
}

func (b Bitcoin) String() string {
  return fmt.Sprintf("%f BTC", b)
}
