package authenticator

import (
	"errors"

	"github.com/tehsis/rabbitscore/services/authenticator/providers/facebook"
)

// Authenticator interface allows to authenticate agaisnt different providers
type Authenticator interface {
	Authenticate(string)
}

const (
	// ProviderFacebook is the fb provider identifier
	ProviderFacebook = "Facebook"
)

// SocialProfile is the normalized RabbitScores profile
type SocialProfile struct {
	ID        string
	Provider  string
	FirstName string
	LastName  string
}

type facebookCredentials struct {
	accessToken string
}

// Authenticate is a method that allows authentication by multiple providers
func Authenticate(method string, credentials string) (SocialProfile, error) {
	var profile SocialProfile
	var err error

	if method == ProviderFacebook {
		var fbProfile facebook.FacebookProfile
		fbProfile, err = facebook.Authenticate(credentials)
		profile = SocialProfile{
			fbProfile.ID,
			ProviderFacebook,
			fbProfile.FirstName,
			fbProfile.LastName,
		}
	} else {
		err = errors.New("method_not_available")
	}

	return profile, err
}
