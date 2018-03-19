package auth

import (
	"context"
	"net/http"

	"gopkg.in/square/go-jose.v2"

	"github.com/auth0-community/auth0"
	"github.com/sirupsen/logrus"
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

// NewRS256JSONWebTokenMiddleware uses to validate a JSON web token that signing in RS256 format
func NewRS256JSONWebTokenMiddleware(domain string, audience []string, secret interface{}) func(http.Handler) http.Handler {
	validator := auth0.NewValidator(
		auth0.NewConfiguration(
			auth0.NewKeyProvider(secret),
			audience,
			"https://"+domain+".auth0.com/",
			jose.RS256,
		),
		nil,
	)

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := context.Background()

			token, err := validator.ValidateRequest(r)
			if err != nil {
				logrus.Errorf("[%s] %s: %v", r.Method, r.URL.String(), err)
			} else {
				var name, email string
				var emailVerified bool

				token.Claims("name", &name)
				token.Claims("email", &email)
				token.Claims("email_verified", &emailVerified)

				ctx = context.WithValue(ctx, Name, name)
				ctx = context.WithValue(ctx, Email, email)
				ctx = context.WithValue(ctx, EmailVerified, emailVerified)
			}

			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
