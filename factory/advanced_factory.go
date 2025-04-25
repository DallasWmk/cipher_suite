package factory

import "github.com/DallasWmk/security_suite/product"

type AdvancedSecurityFactory struct{}

func (a *AdvancedSecurityFactory) CreateEncryption() product.Encryption {
	return &product.RSAEncryption{}
}

func (a *AdvancedSecurityFactory) CreateAuthentication() product.Authentication {
	return &product.JWTAuth{}
}

func (a *AdvancedSecurityFactory) CreateNetworkProtection() product.NetworkProtection {
	return &product.VPNProtection{}
}
