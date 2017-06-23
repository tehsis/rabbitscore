package middlewares

import (
	"net/http"
	"strings"

	"github.com/tehsis/rabbitscore/services/authorizer"
)

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

		if authorizationMethod == "Bearer" {
			playerID, err := authorizer.GetPlayerID(authorizationToken)
			if err != nil {
				http.Error(w, playerNotFound, http.StatusUnauthorized)
				return
			}

			http.Error(w, playerID, http.StatusOK)
			return
		}

		http.Error(w, methodNotAllowed, http.StatusUnauthorized)
		return

		//inner.ServeHTTP(w, r)
	})
}
