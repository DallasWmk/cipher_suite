package factory

import "github.com/DallasWmk/security_suite/product"

type SecuritySuiteFactory interface {
	CreateEncryption() product.Encryption
	CreateAuthentication() product.Authentication
	CreateNetworkProtection() product.NetworkProtection
}
