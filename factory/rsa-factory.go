package factory

import (
	"bytes"
	"crypto/rand"
	"crypto/rsa"
	"encoding/gob"
	"fmt"

	"github.com/DallasWmk/security_suite/product"
)

type RSAFactory struct{}

func (b *RSAFactory) CreateEncryptor() product.Encryption {
	return &product.RSAEncryptor{}
}

func (b *RSAFactory) GenerateKey() ([]byte, error) {
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return []byte{}, fmt.Errorf("error generating RSA key: %s", err)
	}

	var buffer bytes.Buffer
	encoder := gob.NewEncoder(&buffer)

	err = encoder.Encode(privateKey)
	if err != nil {
		return []byte{}, fmt.Errorf("encoding error: %w", err)
	}

	bytes := buffer.Bytes()
	return bytes, nil
}
