package factory

import "github.com/DallasWmk/security_suite/product"

type RSAFactory struct{}

func (b *RSAFactory) CreateEncryptor() product.Encryption {
	return &product.RSAEncryptor{}
}

func (b *RSAFactory) GenerateKey() ([]byte, error) {
	// TODO
	return []byte{}, nil
}
