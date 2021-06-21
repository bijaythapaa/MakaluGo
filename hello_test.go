package main

import "testing"

func TestHello(t *testing.T) {
	// got := Hello("Bijay")
	// want := "Hello, Bijay"

	// if got != want {
	// 	t.Errorf("got %q want %q", got, want)
	// }

	assertCorrectMessage := func(t testing.TB, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	}

	t.Run("say hello to people", func(t *testing.T) {
		got := Hello("Bijay", "")
		want := "Hello, Bijay"

		assertCorrectMessage(t, got, want)
		// if got != want {
		// 	t.Errorf("got %q, want %q", got, want)
		// }
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Bijay", "Spanish")
		want := "Hola, Bijay"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in French", func(t *testing.T) {
		got := Hello("Bijay", "French")
		want := "Bonjour, Bijay"
		assertCorrectMessage(t, got, want)
	})

	t.Run("say, 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"

		assertCorrectMessage(t, got, want)
		// if got != want {
		// 	t.Errorf("got %q, want %q", got, want)
		// }
	})
}
