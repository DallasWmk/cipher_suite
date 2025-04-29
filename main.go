package main

import (
	"fmt"

	"github.com/DallasWmk/security_suite/factory"
)

func useSecuritySuite(f factory.EncryptionFactory) {
	encryption := f.CreateEncryptor()

	key, err := f.GenerateKey()
	if err != nil {
		panic(fmt.Errorf("failed to generate key: %w", err))
	}
	encrypted, err := encryption.Encrypt(key, "This is a very secret message! Dont let anyone read me!!!")
	if err != nil {
		panic(fmt.Errorf("failed while encrypting: %w", err))
	}
	fmt.Printf("Encrypted: %x\n", encrypted)

	decrypted, err := encryption.Decrypt(key, encrypted)
	if err != nil {
		panic(fmt.Errorf("failed while decrypting: %w", err))
	}
	fmt.Printf("Decrypted: %s\n", decrypted)
}

func main() {
	fmt.Println("=== AES Encryption ===")
	useSecuritySuite(&factory.AESFactory{})

	fmt.Println("\n=== RSA Encrption ===")
	useSecuritySuite(&factory.RSAFactory{})
}
