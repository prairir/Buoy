package translate

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"io"
)

// These use CBC AES and PKCS7 padding

// translate.encrypt: pad plainText with PKCS7 and AES blocksize and then
// encrypt with AES CBC mode and return
func encrypt(plainText, password []byte) ([]byte, error) {
	plainTextPad := pkcs7Pad(plainText, aes.BlockSize)

	cipherBlock, err := aes.NewCipher(password)
	if err != nil {
		return []byte{}, fmt.Errorf("translate.encrypt: %w", err)
	}

	// make a cipherText buffer
	cipherText := make([]byte, aes.BlockSize+len(plainTextPad))

	// fill IV with random bytes
	if _, err := io.ReadFull(rand.Reader, cipherText[:aes.BlockSize]); err != nil {
		return []byte{}, fmt.Errorf("translate.encrypt: %w", err)
	}

	cbc := cipher.NewCBCEncrypter(cipherBlock, cipherText[:aes.BlockSize])

	cbc.CryptBlocks(cipherText[aes.BlockSize:], plainTextPad)

	return cipherText, nil
}

// translate.decrypt: decrypts cipherText with AES CBC mode and then unpads using PKCS7
func decrypt(cipherText, password []byte) ([]byte, error) {
	cipherBlock, err := aes.NewCipher(password)
	if err != nil {
		return []byte{}, fmt.Errorf("translate.decrypt: %w", err)
	}

	iv := cipherText[:aes.BlockSize]

	cbc := cipher.NewCBCDecrypter(cipherBlock, iv)

	plainTextPad := make([]byte, len(cipherText[aes.BlockSize:]))

	cbc.CryptBlocks(plainTextPad, cipherText[aes.BlockSize:])

	plainText := pkcs7Unpad(plainTextPad)

	return plainText, nil
}

// translate.pkcs7Pad: Pads cipherText to be a multiple of blockSize
// It is an implementation of PKCS7 padding
// PKCS7 adds an extra block of padding if it fits the padding size
// Ryan straight up stole this from the internet
func pkcs7Pad(src []byte, blockSize int) []byte {
	padding := (blockSize - len(src)%blockSize)
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(src, padtext...)
}

// translate.pkcs7Unpad: This unpads a byte slice
// It is an implementation of PKCS7 padding
func pkcs7Unpad(src []byte) []byte {
	length := len(src)
	unpadding := int(src[length-1])
	return src[:(length - unpadding)]
}
