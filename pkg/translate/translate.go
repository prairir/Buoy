package translate

import (
	"fmt"

	"github.com/prairir/Buoy/pkg/config"
)

// translate.Translate: compresses and encrypts plainText
func Translate(plainText []byte) ([]byte, error) {
	compressed, err := compress(plainText)
	if err != nil {
		return []byte{}, fmt.Errorf("translate.Translate: %w", err)
	}

	encrypted, err := encrypt(compressed, config.Config.Password)
	if err != nil {
		return []byte{}, fmt.Errorf("translate.Translate: %w", err)
	}

	return encrypted, nil
}

// translate.Untranslate: decrypts and decompresses already "translated" cipherText.
func Untranslate(cipherText []byte) ([]byte, error) {
	decrypted, err := decrypt(cipherText, config.Config.Password)
	if err != nil {
		return []byte{}, fmt.Errorf("translate.Untranslate: %w", err)
	}

	decompressed, err := decompress(decrypted)
	if err != nil {
		return []byte{}, fmt.Errorf("translate.Untranslate: %w", err)
	}

	return decompressed, nil
}
