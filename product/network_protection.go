package product

type NetworkProtection interface {
	Protect() string
}

type HTTPSProtection struct{}

func (n *HTTPSProtection) Protect() string {
	return "Protected via HTTPS"
}

type VPNProtection struct{}

func (n *VPNProtection) Protect() string {
	return "Protected via VPN"
}
