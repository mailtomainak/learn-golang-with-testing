package main

import "testing"

func TestHello(t *testing.T) {
	assertCorrectMessage := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("SayHelloToPeople", func(t *testing.T) {
		got := Hello("Chris")
		want := "Hello Chris"
		assertCorrectMessage(t, got, want)
	})

	t.Run("Empty string defaults to World", func(t *testing.T) {
		got := Hello("")
		want := "Hello World"
		assertCorrectMessage(t, got, want)
	})
}
