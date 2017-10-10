package leaderboard

import (
	"github.com/tehsis/leaderboard"
	"github.com/tehsis/rabbitscore/services/redis"
)

// Repo is a persistent storage
var repo *leaderboard.LeaderBoard

func getLeaderboard() *leaderboard.LeaderBoard {
	if repo == nil {
		c := redis.GetClient()
		l := leaderboard.NewRedisLeaderBoard(c)
		repo = &l
	}

	return repo
}

func AddScore(name string, points uint) uint {
	return getLeaderboard().Set(name, points)
}

func GetScore(name string) uint {
	_, score := getLeaderboard().Get(name)
	return score
}

func GetTopFive() []leaderboard.Score {
	return getLeaderboard().GetTop(5)
}
