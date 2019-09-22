package main

import "testing"

func TestHello(t *testing.T) {
	assertCorrectMessage := func(t *testing.T, got, want string) {
		t.Helper()

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("say hello to people", func(t *testing.T) {
		got := Hello("Chris", "")
		want := "Hello, Chris"

		assertCorrectMessage(t, got, want)
	})

	t.Run("say 'Hello, World' for empty string", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"

		assertCorrectMessage(t, got, want)
	})

	t.Run("say สวัสดี in Thai", func(t *testing.T) {
		got := Hello("narze", "Thai")
		want := "สวัสดี, narze"

		assertCorrectMessage(t, got, want)
	})

	t.Run("say Hola in Spanish", func(t *testing.T) {
		got := Hello("narze", "Spanish")
		want := "Hola, narze"

		assertCorrectMessage(t, got, want)
	})

	t.Run("say Bonjour in French", func(t *testing.T) {
		got := Hello("narze", "French")
		want := "Bonjour, narze"

		assertCorrectMessage(t, got, want)
	})
}
