package auth

import (
	"context"
	"net/http"

	"github.com/auth0-community/auth0"
	"github.com/sirupsen/logrus"
	"gopkg.in/square/go-jose.v2"
)

const (
	_ Key = iota
	// Name (String) is an End-User's full name in displayable form including all name parts,
	// possibly including titles and suffixes, ordered according to the End-User's locale and preferences.
	Name
	// Email (String) is an End-User's preferred e-mail address.
	// Its value MUST conform to the RFC 5322 [RFC5322] addr-spec syntax.
	// The RP MUST NOT rely upon this value being unique.
	Email
	// EmailVerified (Boolean) ensure that this e-mail address was controlled by the End-User at the time the verification was performed.
	// The means by which an e-mail address is verified is context-specific,
	// and dependent upon the trust framework or contractual agreements within which the parties are operating.
	EmailVerified
)

// Key used to define a context.Context's key instead of basic type
type Key int

// NewMiddleware returns a new authentication middleware which validates token on request header
func NewMiddleware(jwksURI string, audience []string, issuer string) func(http.Handler) http.Handler {
	validator := auth0.NewValidator(
		auth0.NewConfiguration(
			auth0.NewJWKClient(
				auth0.JWKClientOptions{
					URI: jwksURI,
				},
				nil,
			),
			audience,
			issuer,
			jose.RS256,
		),
		nil,
	)

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := context.Background()

			token, err := validator.ValidateRequest(r)
			if err != nil {
				logrus.Warnf("[%s] %s: %v", r.Method, r.URL.String(), err)
			} else {
				claims := make(map[string]interface{})
				validator.Claims(r, token, &claims)
				ctx = context.WithValue(ctx, Name, claims["name"])
				ctx = context.WithValue(ctx, Email, claims["email"])
				ctx = context.WithValue(ctx, EmailVerified, claims["email_verified"])
			}

			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
