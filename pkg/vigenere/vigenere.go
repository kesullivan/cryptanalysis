package vigenere

import (
	"strings"
)

func Encrypt(toEncrypt string, key string) string {
	keyArr := []byte(strings.ToUpper(key))
	textArr := []byte(strings.ToUpper(toEncrypt))
	var encrypted string = ""

	for i := 0; i < len(textArr); i++ {
		encrypted = encrypted + string((((textArr[i]-65)+(keyArr[i%len(key)]-65))%26)+65)
	}

	return (encrypted)
}

func Decrypt(cipherText string, key string) string {
	keyArr := []byte(strings.ToUpper(key))
	textArr := []byte(strings.ToUpper(cipherText))
	var decrypted string = ""

	for i := 0; i < len(textArr); i++ {
		decrypted = decrypted + string((((textArr[i]-65)-(keyArr[i%len(key)]-65)+26)%26)+65)
	}

	return (decrypted)
}
