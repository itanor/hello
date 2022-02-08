package pointersanderrors

import (
  "testing"
)

func TestWallet(t *testing.T) {

  t.Run("Deposit", func(t *testing.T) {
    wallet := Wallet{}
    wallet.Deposit(Bitcoin(13.3))
    assertBalance(t, wallet, Bitcoin(13.3))
  })

  t.Run("Withdraw with funds", func(t *testing.T) {
    wallet := Wallet{balance: Bitcoin(20.3)}
    err := wallet.Withdraw(Bitcoin(10.0))
    assertNoError(t, err)
    assertBalance(t, wallet, Bitcoin(10.3))
  })

  t.Run("Withdraw insufficient funds", func(t *testing.T) {
    startingBalance := Bitcoin(20.0)
    wallet := Wallet{startingBalance}

    err := wallet.Withdraw(Bitcoin(100.0))
    assertError(t, err, ErrInsufficientFunds)
    assertBalance(t, wallet, startingBalance)
  })
}

func assertBalance(t testing.TB, wallet Wallet, want Bitcoin) {
  t.Helper()
  got := wallet.Balance()

  if got != want {
    t.Errorf("got %s want %s", got, want)
  }
}

func assertError(t testing.TB, got, want error) {
  t.Helper()
  if got == nil {
    t.Fatal("didn't get an error but wanted one")
  }
  if got != want {
    t.Errorf("got %q, want %q", got, want)
  }
}

func assertNoError(t testing.TB, got error) {
  t.Helper()
  if got != nil {
    t.Fatal("got an error but didn't want one")
  }
}

