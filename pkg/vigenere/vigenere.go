package vigenere

import (
	"fmt"
	"strings"
)

/*
C_i = (M_i + K_i) mod 26

Text is placed into the same case then brought into range 0-26 for math then brought back to uppercase ASCII range
*/
func Encrypt(toEncrypt string, key string) string {
	var encrypted string = ""
	toEncrypt = strings.ToUpper(toEncrypt)
	key = strings.ToUpper(key)

	for i, c := range toEncrypt {
		fmt.Println(c - 65)
		encrypted = encrypted + string((((c-65)+(rune(key[i%len(key)])-65))%26)+65)
	}

	return (encrypted)
}

/*
M_i = (C_i - K_i) mod 26

Text is placed into the same case then brought into range 0-26 for math then brought back to uppercase ASCII range
*/
func Decrypt(cipherText string, key string) string {
	var decrypted string = ""
	cipherText = strings.ToUpper(cipherText)
	key = strings.ToUpper(key)

	for i, c := range cipherText {
		decrypted = decrypted + string((((c-65)-(rune(key[i%len(key)])-65)+26)%26)+65)
	}

	return (decrypted)
}
