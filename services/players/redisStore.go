package players

import (
	"fmt"

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
	redisPlayerCmd := store.client.Get(player.username)

	socialID, _ := redisPlayerCmd.Result()

	if socialID == "" || redisPlayerCmd.Err() != nil {
		err := store.client.Set(player.username, player.externalID, 0).Err()

		if err != nil {
			fmt.Printf("EERRROR  or here %v", err)
			return false, err
		}

		return true, nil
	}

	if socialID == player.externalID {
		return true, nil
	}

	return false, nil
}
