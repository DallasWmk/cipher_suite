package product

type Authentication interface {
	Authenticate(user string) string
}

type OAuth2Auth struct{}

func (a *OAuth2Auth) Authenticate(user string) string {
	return "Authenticated " + user + " via OAuth2"
}

type JWTAuth struct{}

func (a *JWTAuth) Authenticate(user string) string {
	return "Authenticated " + user + " via JWT"
}
