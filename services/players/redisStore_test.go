package players_test

import (
	"testing"

	"github.com/tehsis/rabbitscore/services/players"

	redis "gopkg.in/redis.v5"
)

func TestIsValid(t *testing.T) {
	redisClient := redis.NewClient(&redis.Options{
		Addr: "127.0.0.1:6379",
	})

	store := players.NewRedisStore(redisClient)

	player := players.NewFromFacebook("tehsis", "123456")
	otherPlayer := players.NewFromFacebook("tehsis", "6543221")

	playerIsValid, err := store.IsValid(player)

	if !playerIsValid {
		t.Error("Player should be valid but is not valid")
	}

	otherPlayerIsValid, err := store.IsValid(otherPlayer)

	if otherPlayerIsValid {
		t.Error("otherPlayer should not be valid but is valid")
	}

	if err != nil {
		t.Error("Error should be nil")
	}
}
