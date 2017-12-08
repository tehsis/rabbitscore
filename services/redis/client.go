package redis

import (
	redis "gopkg.in/redis.v5"
)

var client *redis.Client

// GetClient returns a singleton for the redis client
func GetClient() *redis.Client {
	if client == nil {
		client = redis.NewClient(&redis.Options{
			Addr: "db:6379",
		})
	}

	return client
}
