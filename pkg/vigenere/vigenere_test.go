package vigenere

import "testing"

func TestVigenere(t *testing.T) {
	got := Encrypt()
	want := "Vigenere cipher"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
