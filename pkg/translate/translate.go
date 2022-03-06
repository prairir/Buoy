package translate

import (
	"fmt"
)

// translate.Translate: compresses and encrypts plainText
func Translate(plainText, password []byte) ([]byte, error) {
	compressed, err := compress(plainText)
	if err != nil {
		return []byte{}, fmt.Errorf("translate.Translate: %w", err)
	}

	encrypted, err := encrypt(compressed, password)
	if err != nil {
		return []byte{}, fmt.Errorf("translate.Translate: %w", err)
	}

	return encrypted, nil
}

// translate.Untranslate: decrypts and decompresses already "translated" cipherText.
func Untranslate(cipherText, password []byte) ([]byte, error) {
	decrypted, err := decrypt(cipherText, password)
	if err != nil {
		return []byte{}, fmt.Errorf("translate.Untranslate: %w", err)
	}

	decompressed, err := decompress(decrypted)
	if err != nil {
		return []byte{}, fmt.Errorf("translate.Untranslate: %w", err)
	}

	return decompressed, nil
}
