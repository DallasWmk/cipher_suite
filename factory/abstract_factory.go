package factory

import "github.com/DallasWmk/security_suite/product"

type EncryptionFactory interface {
	CreateEncryptor() product.Encryption
	GenerateKey() ([]byte, error)
}
