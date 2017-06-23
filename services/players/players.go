package players

type SocialPlayer struct {
	ID       string
	Provider string
}

// Player is a user of Rabbit Wars
type Player struct {
	ID       string
	Name     string
	SocialID SocialPlayer
}

// Store is a store of players backed by an specify db
type Store interface {
	// GetID returns the player ID. saving it if it's not already exists
	GetID(profile Player) string
	// IsValid validates that username is registerded with externalId.
	// If the user exists and matches externalId, it returns true.
	// If the user does not exists, it adds the mapping then returns true.
	// If the user exists but using a different externalId returns false.
	IsValid(player Player) bool
}
