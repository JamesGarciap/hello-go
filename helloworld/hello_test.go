package main

import "testing"

func TestHello(t *testing.T) {
	// t.Run correspond to a subtest
	t.Run("Say 'Hello, World! if no name is passed", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"

		assertCorrectMessage(t, got, want)
	})

	t.Run("Say hello with a name if a name is parameter", func(t *testing.T) {
		got := Hello("Chris", "")
		want := "Hello, Chris"

		assertCorrectMessage(t, got, want)
	})

	t.Run("In Spanish", func(t *testing.T) {
		got := Hello("Eliane", "Spanish")
		want := "Hola, Eliane"

		assertCorrectMessage(t, got, want)
	})

	t.Run("In Swedish", func(t *testing.T) {
		got := Hello("Jonas", "Swedish")
		want := "Hej, Jonas"

		assertCorrectMessage(t, got, want)
	})
}

func TestGreetingPrefix(t *testing.T) {
	got := greetingPrefix("Spanish")
	want := "Hola, "

	assertCorrectMessage(t, got, want)
}

func assertCorrectMessage(t testing.TB, got string, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q wants %q", got, want)
	}
}
