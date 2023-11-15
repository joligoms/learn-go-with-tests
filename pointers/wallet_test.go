package pointers

import (
	"testing"
)

func TestWallet(t *testing.T) {
	t.Run("deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))

		got := wallet.Balance()
		want := Bitcoin(10)

		assertCorrect(t, got, want)
	})
}

func assertCorrect(t testing.TB, got, want Bitcoin) {
	t.Helper()

	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}
