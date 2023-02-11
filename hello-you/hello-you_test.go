package main

import "testing"

func TestHelloYou(t *testing.T) {
	got := HelloYou("zyzyx")
	want := "Hello, zyzyx"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
