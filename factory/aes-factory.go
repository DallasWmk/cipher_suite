package factory

import (
	"crypto/rand"
	"fmt"

	"github.com/DallasWmk/security_suite/product"
)

type AESFactory struct{}

func (b *AESFactory) CreateEncryptor() product.Encryption {
	return &product.AESEncryptor{}
}

func (b *AESFactory) GenerateKey() ([]byte, error) {
	key := make([]byte, 32)
	_, err := rand.Read(key)
	if err != nil {
		return []byte{}, fmt.Errorf("failed generating key: %w", err)
	}
	return key, nil
}
