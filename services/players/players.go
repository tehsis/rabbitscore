package players

// Player is a user of Rabbit Wars
type Player struct {
	username   string
	externalID string
}

// Store is a store of players backed by an specify db
type Store interface {
	// IsValid validates that username is registerded with externalId.
	// If the user exists and matches externalId, it returns true.
	// If the user does not exists, it adds the mapping then returns true.
	// If the user exists but using a different externalId returns false.
	IsValid(player Player) bool
}

// NewFromFacebook returns a new player created with a fb id
func NewFromFacebook(username string, facebookID string) Player {
	return Player{
		username,
		"facebook|" + facebookID,
	}
}
