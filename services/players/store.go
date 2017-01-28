package players

import "github.com/tehsis/rabbitscore/services/redis"

// Players is a store of players
var playersStore *RedisStore

// GetStore allows getting playersStore singletone
func GetStore() *RedisStore {
	if playersStore == nil {
		p := NewRedisStore(redis.GetClient())
		playersStore = &p
	}

	return playersStore
}
