package caesargo

// EncryptText function to encrypt text with caesar chiper algorithm.
func EncryptText(text string, shift int) string {
	runes := []rune(text)
	result := make([]rune, len(runes))

	for i, char := range runes {
		if char >= 'A' && char <= 'Z' {
			result[i] = rune((int(char-'A')+shift)%26 + 'A')
		} else if char >= 'a' && char <= 'z' {
			result[i] = rune((int(char-'a')+shift)%26 + 'a')
		} else {
			result[i] = char
		}
	}

	return string(result)
}

// DecryptText function to decrypt text with caesar chiper algorithm.
func DecryptText(text string, shift int) string {
	runes := []rune(text)
	result := make([]rune, len(runes))

	for i, char := range runes {
		if char >= 'A' && char <= 'Z' {
			result[i] = rune((int(char-'A')-shift+26)%26 + 'A')
		} else if char >= 'a' && char <= 'z' {
			result[i] = rune((int(char-'a')-shift+26)%26 + 'a')
		} else {
			result[i] = char
		}
	}

	return string(result)
}
