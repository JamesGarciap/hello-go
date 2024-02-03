package pointerserrors

import (
	"fmt"
	"testing"
)

func TestWallet(t *testing.T) {

	assetBalance := func(t testing.TB, wallet Wallet, want Chichicoin) {
		t.Helper()
		got := wallet.Balance()

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	}

	t.Run("deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Chichicoin(10))
		fmt.Printf("address of balance in test is %p \n", &wallet.balance)

		assetBalance(t, wallet, Chichicoin(10))
	})

	t.Run("withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Chichicoin(20)}
		wallet.Withdraw(Chichicoin(10))
		assetBalance(t, wallet, Chichicoin(10))
	})
}
