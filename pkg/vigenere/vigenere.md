# Vigenere Cipher  

The Vigenere Cipher is an encryption method where each letter of the plain text is encoded with a different Caesar Cipher whose increment is determined by the corresponding letter of the key.

## Algebraic Description

Encryption:
C_i = (M_i + K_i) mod 26

Decryption:
M_i = (C_i - K_i) mod 26

Where:
M is the Message, K is the Key, and C is the Cipher-text

## Example

plaintext:  ATTACKATDAWN
key:        LEMON
ciphertext: LXFOPVEFRNHR

A + L -> 00 + 11 -> L
T + E -> 19 + 04 -> X
T + M -> 19 + 12 -> F
A + O -> 00 + 14 -> O
C + N -> 02 + 13 -> P
K + L -> 10 + 11 -> V
A + E -> 00 + 04 -> E
T + M -> 19 + 12 -> F
D + O -> 03 + 14 -> R
A + N -> 00 + 13 -> N
W + L -> 22 + 11 -> H
N + E -> 13 + 04 -> R

## Cryptanalysis

By chance, repeated words are sometimes encrypted using the same key letters, leading to repeated groups in the ciphertext. The distance between the repetitions can lead to the length of the key. For example, if the length between repetitions is 16 then the key could be 16, 8, 4, 2 or 1. Once the length of the key is known the ciphertext can be rewritten into that many columns where each column corresponds to a single letter of the key. Each column consists of text that has been encrypted by the same Caesar Cipher shift. This leaves each column vulnerable to frequency analysis.
