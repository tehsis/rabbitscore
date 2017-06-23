package handlers

import (
	"net/http"

	"github.com/tehsis/rabbitscore/services/authenticator"
	"github.com/tehsis/rabbitscore/services/authorizer"
	"github.com/tehsis/rabbitscore/services/players"
)

// AuthenticationHandler should respond with a token if the user is succesfully authenticated
func AuthenticationHandler(w http.ResponseWriter, r *http.Request) {

	method := r.FormValue("method")
	credentials := r.FormValue("credentials")

	if method == "" {
		http.Error(w, errorMissingAuthenticationMethod, http.StatusBadRequest)
		return
	}

	if credentials == "" {
		http.Error(w, errorMissingCredentials, http.StatusBadRequest)
		return
	}

	profile, err := authenticator.Authenticate(method, credentials)

	if err != nil {
		if err.Error() == "method_not_available" {
			http.Error(w, errorBadMethod, http.StatusUnauthorized)
			return
		}

		http.Error(w, errorUnauthorized, http.StatusUnauthorized)
		return
	}

	player, _ := players.GetStore().GetID(players.Player{
		Name: profile.FirstName + " " + string(profile.LastName[0]),
		SocialID: players.SocialPlayer{
			Provider: method,
			ID:       profile.ID,
		},
	})

	token, _ := authorizer.GetAccessToken(player.ID)

	err = ResponseToken(w, string(token))

	if err != nil {
		panic(err)
	}
}
