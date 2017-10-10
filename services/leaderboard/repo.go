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
		l := leaderboard.NewRedisLeaderBoard("rabbits-leaderboard", c)
		repo = &l
	}

	return repo
}

func AddScore(name string, points uint) uint {
	position, err := getLeaderboard().Set(name, points)
	if err != nil {
		panic(err)
	}

	return position
}

func GetScore(name string) uint {
	score, _, _ := getLeaderboard().Get(name)

	return score
}

func GetTopFive() []leaderboard.Score {
	top5, err := getLeaderboard().GetTop(5)
	if err != nil {
		panic(err)
	}
	return top5
}
