package players

import (
	"fmt"

	uuid "github.com/satori/go.uuid"
	redis "gopkg.in/redis.v5"
)

// RedisStore is a Player store backed by Redis
type RedisStore struct {
	client *redis.Client
}

// NewRedisStore returns a new RedisStore
func NewRedisStore(client *redis.Client) RedisStore {
	return RedisStore{
		client,
	}
}

// IsValid validates a player against a RedisDB
func (store *RedisStore) IsValid(player Player) (bool, error) {
	return true, nil
}

// GetID gets the id of the given player
func (store *RedisStore) GetID(profile Player) (Player, error) {
	if profile.ID != "" {
		return profile, nil
	}

	m := make(map[string]string)

	m["name"] = profile.Name
	providerString := "provider:" + profile.SocialID.Provider

	cmd := store.client.HGet(providerString, profile.SocialID.ID)
	profile.ID = cmd.Val()

	if profile.ID == "" {
		profile.ID = uuid.NewV4().String()
		store.client.HMSet("user:"+profile.ID, m)
	}

	store.client.HSet(providerString, profile.SocialID.ID, profile.ID)
	fmt.Printf("profile: %v", profile)
	return profile, nil
}
