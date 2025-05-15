package product

import (
	"bytes"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/gob"
	"fmt"
)

type RSAEncryptor struct{}

func (e *RSAEncryptor) Encrypt(key []byte, data string) ([]byte, error) {
	privKey, err := decodePrivKey(key)
	if err != nil {
		return []byte{}, err
	}

	label := []byte("message")

	rng := rand.Reader
	msgBytes := []byte(data)

	ciphertext, err := rsa.EncryptOAEP(sha256.New(), rng, &privKey.PublicKey, msgBytes, label)
	if err != nil {
		return []byte{}, fmt.Errorf("error from encryption: %w", err)
	}

	return ciphertext, nil
}

func (e *RSAEncryptor) Decrypt(key []byte, encryptedData []byte) ([]byte, error) {
	privKey, err := decodePrivKey(key)
	if err != nil {
		return []byte{}, err
	}
	label := []byte("message")

	plaintext, err := rsa.DecryptOAEP(sha256.New(), nil, privKey, encryptedData, label)
	if err != nil {
		return []byte{}, fmt.Errorf("error from decryption: %w", err)
	}
	return plaintext, nil
}

func decodePrivKey(key []byte) (*rsa.PrivateKey, error) {
	privKey := rsa.PrivateKey{}
	buffer := bytes.NewBuffer(key)
	decoder := gob.NewDecoder(buffer)
	err := decoder.Decode(&privKey)
	if err != nil {
		return nil, fmt.Errorf("decoding error: %w", err)
	}
	return &privKey, nil
}
