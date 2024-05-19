package vigenere

import "testing"

func TestVigenereEncrypt(t *testing.T) {
	got := Encrypt("ATTACKATDAWN", "lemon")
	want := "LXFOPVEFRNHR"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func TestVigenereDecrypt(t *testing.T) {
	got := Decrypt("LXFOPVEFRNHR", "lemon")
	want := "ATTACKATDAWN"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
