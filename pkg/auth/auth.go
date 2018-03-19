package auth

import (
	"net/http"

	"github.com/auth0-community/auth0"
	"gopkg.in/square/go-jose.v2"
)

// Handler is an HTTP handler of Auth0 provides an authentication middleware
type Handler struct {
	*auth0.JWTValidator
}

// NewHandler returns a new Auth handler with Auth0 JWT Validator
func NewHandler(domain string, audience []string, clientSecret string) Handler {
	return Handler{
		JWTValidator: auth0.NewValidator(
			auth0.NewConfiguration(
				auth0.NewKeyProvider([]byte(clientSecret)),
				audience,
				"https://"+domain+".auth0.com/",
				jose.RS256,
			),
			nil,
		),
	}
}

// Init is a function that used to register HTTP handlers to http.ServeMux object
func (h Handler) Init(mux *http.ServeMux) error {
	return nil
}
