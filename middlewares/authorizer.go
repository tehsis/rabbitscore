package middlewares

import (
	"context"
	"net/http"
	"strings"

	"github.com/tehsis/rabbitscore/services/authorizer"
)

type PlayerIdKey struct{}
type PlayerUsername struct{}

type Profile struct {
	ID   string
	Name string
}

const methodNotAllowed = "Method not allowed"
const playerNotFound = "Player not found"
const malformedAuthenticationHeader = "Malformed authentication header"

// Authorize authorizes the app
func Authorize(inner http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authorizationHeader := strings.Split(r.Header.Get("Authorization"), " ")

		if len(authorizationHeader) != 2 {
			http.Error(w, malformedAuthenticationHeader, http.StatusUnauthorized)
			return
		}

		authorizationMethod := authorizationHeader[0]
		authorizationToken := authorizationHeader[1]

		if authorizationMethod != "Bearer" {
			http.Error(w, methodNotAllowed, http.StatusUnauthorized)
			return
		}

		player, err := authorizer.GetPlayer(authorizationToken)

		if err != nil {
			http.Error(w, playerNotFound, http.StatusUnauthorized)
			return
		}

		r = r.WithContext(context.WithValue(r.Context(), PlayerIdKey{}, player.ID))
		r = r.WithContext(context.WithValue(r.Context(), PlayerUsername{}, player.Name))

		inner.ServeHTTP(w, r)
	})
}
