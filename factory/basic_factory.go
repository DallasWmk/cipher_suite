package factory

import "github.com/DallasWmk/security_suite/product"

type BasicSecurityFactory struct{}

func (b *BasicSecurityFactory) CreateEncryption() product.Encryption {
	return &product.AESEncryption{}
}

func (b *BasicSecurityFactory) CreateAuthentication() product.Authentication {
	return &product.OAuth2Auth{}
}

func (b *BasicSecurityFactory) CreateNetworkProtection() product.NetworkProtection {
	return &product.HTTPSProtection{}
}
