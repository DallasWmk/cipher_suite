package main

import (
	"crypto/rand"
	"fmt"

	"github.com/DallasWmk/security_suite/factory"
)

func useSecuritySuite(f factory.SecuritySuiteFactory) {
	encryption := f.CreateEncryption()
	authentication := f.CreateAuthentication()
	protection := f.CreateNetworkProtection()

	key := make([]byte, 32)
	_, err := rand.Read(key)
	if err != nil {
		fmt.Println("here")
		panic(err)
	}
	output, err := encryption.Encrypt(key, "This is a very secret message! Dont let anyone read me!!!")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(output)
	fmt.Println(authentication.Authenticate("Alice"))
	fmt.Println(protection.Protect())
}

func main() {
	fmt.Println("=== Basic Security Suite ===")
	useSecuritySuite(&factory.BasicSecurityFactory{})

	fmt.Println("\n=== Advanced Security Suite ===")
	useSecuritySuite(&factory.AdvancedSecurityFactory{})
}
