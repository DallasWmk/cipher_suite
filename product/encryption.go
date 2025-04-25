package product

import (
	"crypto/aes"
	"fmt"
)

type Encryption interface {
	Encrypt(key []byte, data string) (string, error)
}

type AESEncryption struct{}

func (e *AESEncryption) Encrypt(key []byte, data string) (string, error) {
	dataBytes := []byte(data)

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", fmt.Errorf("failed creating new block: %w", err)
	}

	encryptedBytes := make([]byte, 16)
	dataBytesPadded := pkcs7Pad(dataBytes, 16)
	block.Encrypt(encryptedBytes, dataBytesPadded)

	return fmt.Sprintf("AES Encrypted: %x", encryptedBytes), nil
}

func pkcs7Pad(input []byte, blockSize int) []byte {
	paddingNeeded := blockSize - (len(input) % blockSize)
	padding := make([]byte, paddingNeeded)
	for i := range padding {
		padding[i] = byte(paddingNeeded)
	}
	return append(input, padding...)
}

type RSAEncryption struct{}

func (e *RSAEncryption) Encrypt(key []byte, data string) (string, error) {
	return "RSA Encrypted: " + data, nil
}
