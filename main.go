package main

import (
	"fmt"

	"github.com/DallasWmk/security_suite/factory"
)

func useSecuritySuite(f factory.SecuritySuiteFactory) {
	encryption := f.CreateEncryption()
	authentication := f.CreateAuthentication()
	protection := f.CreateNetworkProtection()

	fmt.Println(encryption.Encrypt("SensitiveData"))
	fmt.Println(authentication.Authenticate("Alice"))
	fmt.Println(protection.Protect())
}

func main() {
	fmt.Println("=== Basic Security Suite ===")
	useSecuritySuite(&factory.BasicSecurityFactory{})

	fmt.Println("\n=== Advanced Security Suite ===")
	useSecuritySuite(&factory.AdvancedSecurityFactory{})
}
