package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"sync"

	tl "github.com/tehsis/leaderboard"
	"github.com/tehsis/rabbitscore/middlewares"
	"github.com/tehsis/rabbitscore/services/leaderboard"
	"github.com/tehsis/rabbitscore/services/players"
)

type score struct {
	Username string `json:"Username"`
	Points   uint   `json:"Points"`
}

func validateScore(currentScore score) error {
	if currentScore.Username == "" {
		return errors.New("Malformed JSON")
	}

	return nil
}

// LeaderBoardHandler is a list of scores
func LeaderBoardHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)

	store := players.GetStore()

	scores := leaderboard.GetTopFive()

	fmt.Printf("Top five %v", scores)

	scoreResponse := make([]score, len(scores))

	var wg sync.WaitGroup
	wg.Add(len(scores))
	for index, score := range scores {
		go func(score tl.Score, index int) {
			defer wg.Done()
			username, err := store.GetPlayerName(score.Username)

			if err != nil {
				username = "unknown"
			}

			scoreResponse[index].Username = username
			scoreResponse[index].Points = score.Points
		}(score, index)
	}

	wg.Wait()

	if err := json.NewEncoder(w).Encode(scoreResponse); err != nil {
		panic(err)
	}
}

// AddScore adds a new score
func AddScore(w http.ResponseWriter, r *http.Request) {
	score := r.FormValue("score")
	userID, ok := r.Context().Value(middlewares.PlayerIdKey{}).(string)
	username, ok := r.Context().Value(middlewares.PlayerUsername{}).(string)

	if !ok {
		ResponseError(w, http.StatusInternalServerError, "Unknown error, please try again later")
	}

	if score == "" {
		ResponseError(w, http.StatusBadRequest, "score is required")
		return
	}

	if username == "" || userID == "" {
		ResponseError(w, http.StatusInternalServerError, "Internal Error 40")
	}

	scoreInt, err := strconv.Atoi(score)

	if err != nil {
		ResponseError(w, http.StatusInternalServerError, "Internal Error 42")
	}

	currentScore := leaderboard.GetScore(userID)

	var position uint

	if uint(scoreInt) > uint(currentScore) {
		position = leaderboard.AddScore(userID, uint(scoreInt))
	} else {
		position = leaderboard.GetScore(userID)
	}

	ResponsePosition(w, username, position)
}
