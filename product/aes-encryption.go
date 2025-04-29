package product

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"io"
)

type Encryption interface {
	Encrypt(key []byte, data string) ([]byte, error)
	Decrypt(key []byte, encryptedData []byte) ([]byte, error)
}

type AESEncryptor struct{}

func (e *AESEncryptor) Encrypt(key []byte, data string) ([]byte, error) {
	plaintext := []byte(data)
	if len(plaintext)%aes.BlockSize != 0 {
		plaintext = pkcs7Pad(plaintext, aes.BlockSize)
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return []byte{}, fmt.Errorf("failed creating new block: %w", err)
	}

	// The IV needs to be unique, but not secure. Therefore it's common to
	// include it at the beginning of the ciphertext.
	ciphertext := make([]byte, aes.BlockSize+len(plaintext))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return []byte{}, fmt.Errorf("failed reading in random IV: %w", err)
	}

	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(ciphertext[aes.BlockSize:], plaintext)

	return ciphertext, nil
}

func (e *AESEncryptor) Decrypt(key []byte, ciphertext []byte) ([]byte, error) {
	if len(ciphertext)%aes.BlockSize != 0 {
		return nil, fmt.Errorf("ciphertext is not a multiple of the block size")
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return []byte{}, fmt.Errorf("failed creating new block for decryption: %w", err)
	}

	// The IV needs to be unique, but not secure. Therefore it's common to
	// include it at the beginning of the ciphertext.
	if len(ciphertext) < aes.BlockSize {
		panic("ciphertext too short")
	}
	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]

	// CBC mode always works in whole blocks.
	if len(ciphertext)%aes.BlockSize != 0 {
		panic("ciphertext is not a multiple of the block size")
	}

	mode := cipher.NewCBCDecrypter(block, iv)

	// CryptBlocks can work in-place if the two arguments are the same.
	mode.CryptBlocks(ciphertext, ciphertext)
	return ciphertext, nil // Return the decrypted data (may need unpadding)
}

func pkcs7Pad(input []byte, blockSize int) []byte {
	paddingNeeded := blockSize - (len(input) % blockSize)
	padding := make([]byte, paddingNeeded)
	for i := range padding {
		padding[i] = byte(paddingNeeded)
	}
	return append(input, padding...)
}
