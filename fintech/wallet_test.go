package fintech

import (
	"testing"
)

func TestWallet(t *testing.T) {
	//wallet := Wallet{}
	//wallet.Deposit(BitCoin(10))
	//got := wallet.Balance()
	//fmt.Printf("address of balance in test is %v\n", &wallet.balance)
	//want := BitCoin(20)
	//if got != want {
	//	t.Errorf("got: %s, want: %s", got, want)
	//}

	//t.Run("deposits", func(t *testing.T) {
	//	wallet := Wallet{}
	//	wallet.Deposit(BitCoin(200))
	//	got := wallet.Balance()
	//	want := BitCoin(200)
	//	if got != want {
	//		t.Errorf("got: %s, want: %s", got, want)
	//	}
	//})
	//
	//t.Run("withdraws", func(t *testing.T) {
	//	wallet := Wallet{balance: BitCoin(200)}
	//	wallet.WithDraw(BitCoin(90))
	//	got := wallet.Balance()
	//	want := BitCoin(110)
	//	if got != want {
	//		t.Errorf("got: %s, want: %s", got, want)
	//	}
	//})

	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(BitCoin(99))
		//want := BitCoin(99)
		assertBalance(t, wallet, BitCoin(99))
	})

	t.Run("Withdraw", func(t *testing.T) {
		wallet := Wallet{balance: BitCoin(99)}
		err := wallet.WithDraw(BitCoin(9))
		assertBalance(t, wallet, BitCoin(90))
		assertNoError(t, err)
	})

	t.Run("Withdraw insufficient fund", func(t *testing.T) {
		startingBalance := BitCoin(199)
		wallet := Wallet{startingBalance}
		err := wallet.WithDraw(BitCoin(200))

		assertBalance(t, wallet, startingBalance)
		//assertError(t, err, "cannot withdraw, insufficient funds")
		assertError(t, err, ErrInsufficientFunds)
	})

	//t.Run("withdraw insufficient funds", func(t *testing.T) {
	//	startingBalance := BitCoin(20)
	//	wallet := Wallet{startingBalance}
	//	err := wallet.WithDraw(BitCoin(190))
	//
	//	assertBalance(t, wallet, startingBalance)
	//	if err != nil {
	//		t.Errorf("wanted an error but didnt get one")
	//	}
	//})

	//assertBalance := func(t testing.TB, wallet Wallet, want BitCoin) {
	//	t.Helper()
	//	got := wallet.Balance()
	//
	//	if got != want {
	//		t.Errorf("got: %s, want: %s", got, want)
	//	}
	//}
	//
	//assertNoError := func(t testing.TB, got error) {
	//	t.Helper()
	//	if got != nil {
	//		t.Fatal("got an error but didn't want one")
	//	}
	//}
	//
	//assertError := func(t testing.TB, got error, want error) {
	//	t.Helper()
	//	if got == nil {
	//		t.Fatal("didn't get one but wanted an error")
	//	}
	//	if got != want {
	//		t.Errorf("got: %q, want: %q", got, want)
	//	}
	//}
}

func assertBalance(t testing.TB, wallet Wallet, want BitCoin) {
	t.Helper()
	got := wallet.Balance()

	if got != want {
		t.Errorf("got: %s, want: %s", got, want)
	}
}

func assertNoError(t testing.TB, got error) {
	t.Helper()
	if got != nil {
		t.Error("got an error but didnt want one")
	}
}

func assertError(t testing.TB, got, want error) {
	t.Helper()
	if got == nil {
		t.Fatal("didnt get an error but wanted one")
	}
	if got != want {
		t.Errorf("got: %s, want: %s", got, want)
	}
}
