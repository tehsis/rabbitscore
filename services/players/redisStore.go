package players

import (
	"fmt"

	uuid "github.com/satori/go.uuid"
	"github.com/tehsis/rabbitscore/services/logger"
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

func (store *RedisStore) GetPlayerName(id string) (string, error) {
	cmd := store.client.Get("user:" + id)

	name, err := cmd.Result()

	return name, err
}

// GetID gets the id of the given player
func (store *RedisStore) GetID(profile Player) (Player, error) {
	if profile.ID != "" {
		logger.Log().Debug("User ID Provided", profile)
		return profile, nil
	}

	providerString := fmt.Sprintf("%s|%s", profile.SocialID.Provider, profile.SocialID.ID)

	cmd := store.client.Get(providerString)
	profile.ID = cmd.Val()

	if profile.ID == "" {
		profile.ID = uuid.NewV4().String()
		store.client.Set(providerString, profile.ID, 0)
	}

	store.client.Set("user:"+profile.ID, profile.Name, 0)

	return profile, nil
}
