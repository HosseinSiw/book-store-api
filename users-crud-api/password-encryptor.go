package users_crud_api

// caesarEncrypt encrypts the given plaintext using the Caesar Cipher with the specified shift.

func CaesarEncrypt(plaintext string, shift int) string {
	ciphertext := ""
	for _, char := range plaintext {
		if char >= 'A' && char <= 'Z' {
			ciphertext += string((char-'A'+rune(shift))%26 + 'A')
		} else if char >= 'a' && char <= 'z' {
			ciphertext += string((char-'a'+rune(shift))%26 + 'a')
		} else {
			ciphertext += string(char)
		}
	}
	return ciphertext
}

func CaesarDecrypt(ciphertext string, shift int) string {
	return CaesarEncrypt(ciphertext, 26-shift)
}
