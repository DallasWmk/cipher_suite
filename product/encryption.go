package product

type Encryption interface {
	Encrypt(data string) string
}

type AESEncryption struct{}

func (e *AESEncryption) Encrypt(data string) string {
	return "AES Encrypted: " + data
}

type RSAEncryption struct{}

func (e *RSAEncryption) Encrypt(data string) string {
	return "RSA Encrypted: " + data
}
