package facebook

import (
	gograph "github.com/tehsis/gographbook"
)

// FacebookProfile represents a Facebook profile
type FacebookProfile struct {
	ID        string
	FirstName string
	LastName  string
}

// Authenticate with Facebook
func Authenticate(accessToken string) (FacebookProfile, error) {
	fbAPI := gograph.New(accessToken)

	profile, err := fbAPI.Me([]string{"first_name", "last_name"})

	if err != nil {
		return FacebookProfile{}, err
	}

	return FacebookProfile{
		profile.ID,
		profile.FirstName,
		profile.LastName,
	}, err
}
