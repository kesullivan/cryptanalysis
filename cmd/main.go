package main

import (
	"fmt"

	"github.com/kesullivan/cryptanalysis/pkg/vigenere"
)

func main() {
	fmt.Println("This program will attempt to solve encrypted text by decrypting it using multiple different algorithms.")
	fmt.Println(vigenere.Encrypt("attackAtDawn", "lemon"))
	fmt.Println(vigenere.Decrypt("LXFOPVEFRNHR", "lemon"))
}
