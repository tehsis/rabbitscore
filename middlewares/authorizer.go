package middlewares

import (
	"context"
	"net/http"

	"github.com/tehsis/rabbitscore/authenticator"
	"github.com/tehsis/rabbitscore/rabbitContext"
)

// Authorize authorizes the app
func Authorize(inner http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		authtoken := r.Header.Get("Authorization")

		if authtoken == "" {
			inner.ServeHTTP(w, r)
			return
		}

		fbID, err := authenticator.Facebook.Authenticate(authtoken)

		if err != nil {
			inner.ServeHTTP(w, r)
			return
		}

		ctx := context.WithValue(r.Context(), rabbitContext.Context.Auth, fbID)

		inner.ServeHTTP(w, r.WithContext(ctx))
	})
}
