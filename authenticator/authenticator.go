package authenticator

import "github.com/tehsis/gographbook"

// Authenticator interface allows to authenticate agaisnt different providers
type Authenticator interface {
	Authenticate(string, string)
}

// FacebookAuth is a struct that represents the needed fields to authorize to fb
type facebookAuth struct {
}

// Facebook object to authorize with facebook
var Facebook facebookAuth

// Authorize with Facebook
func (fb *facebookAuth) Authenticate(accessToken string) (string, error) {
	fbAPI := gograph.New(accessToken)

	profile, err := fbAPI.Me([]string{"email"})

	return profile.ID, err
}
