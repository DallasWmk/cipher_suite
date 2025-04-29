package product

type RSAEncryptor struct{}

func (e *RSAEncryptor) Encrypt(key []byte, data string) ([]byte, error) {
	// TODO
	return []byte{}, nil
}

func (e *RSAEncryptor) Decrypt(key []byte, encryptedData []byte) ([]byte, error) {
	// TODO
	return []byte{}, nil
}
