package main

import "testing"

// if type is the same for multiple func parameters, just put type after the last one
func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func TestHello(t *testing.T) {
	t.Run(
		"saying hello to people",
		func(t *testing.T) {
			got := Hello("Chris", "")
			want := "Hello, Chris"
			assertCorrectMessage(t, got, want)

		})

	t.Run(
		"saying 'Hello, World' when an empty string is supplied",
		func(t *testing.T) {
			got := Hello("", "")
			want := "Hello, World"
			assertCorrectMessage(t, got, want)
		})

	t.Run(
		"in Spanish",
		func(t *testing.T) {
			got := Hello("Valentina", "Spanish")
			want := "Hola, Valentina"
			assertCorrectMessage(t, got, want)
		})

	t.Run(
		"in French",
		func(t *testing.T) {
			got := Hello("Amélie", "French")
			want := "Bonjour, Amélie"
			assertCorrectMessage(t, got, want)
		})

	t.Run(
		"in Italian",
		func(t *testing.T) {
			got := Hello("Giovanna", "Italian")
			want := "Buongiorno, Giovanna"
			assertCorrectMessage(t, got, want)
		})

	t.Run(
		"in Portuguese",
		func(t *testing.T) {
			got := Hello("Milo", "Portuguese")
			want := "Olá, Milo"
			assertCorrectMessage(t, got, want)
		})
}
