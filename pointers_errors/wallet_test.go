package pointerserrors

import (
	"fmt"
	"testing"
)

func TestWallet(t *testing.T) {
	t.Run("deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Chichicoin(10))
		fmt.Printf("address of balance in test is %p \n", &wallet.balance)

		assetBalance(t, wallet, Chichicoin(10))
	})

	t.Run("withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Chichicoin(20)}
		got := wallet.Withdraw(Chichicoin(10))
		assertNoError(t, got)
		assetBalance(t, wallet, Chichicoin(10))
	})

	t.Run("withdraw insufficient funds", func(t *testing.T) {
		starringBalance := Chichicoin(20)
		wallet := Wallet{balance: starringBalance}
		got := wallet.Withdraw(Chichicoin(100))

		assertError(t, got, ErrInsufficientFunds)
		assetBalance(t, wallet, starringBalance)
	})
}

func assertNoError(t testing.TB, got error) {
	t.Helper()
	if got != nil {
		t.Fatal("Got an error but dint't want one")
	}
}

func assetBalance(t testing.TB, wallet Wallet, want Chichicoin) {
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
