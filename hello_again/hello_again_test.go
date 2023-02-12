package main

import "testing"

func TestHelloAgain(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := HelloAgain("Wan")
		want := "Hello, Wan"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
	t.Run("say 'Hello' when an empty string is supplied", func(t *testing.T) {
		got := HelloAgain("")
		want := "Hello, World"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
}
