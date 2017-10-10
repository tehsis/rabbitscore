package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/tehsis/rabbitscore/middlewares"
	"github.com/tehsis/rabbitscore/services/leaderboard"
	"github.com/tehsis/rabbitscore/services/players"
)

type score struct {
	Username string `json:"username"`
	Points   uint   `json:"points"`
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

	scoreResponse := make([]score, len(scores))

	for index, score := range scores {
		username, err := store.GetPlayerName(score.Username)

		if err != nil {
			username = "unknown"
		}

		scoreResponse[index].Username = username
		scoreResponse[index].Points = score.Points
	}

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

	// this user id should be properly formated (eg. facebook|id, twitter|id)

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

	fmt.Printf("Score: %v\n", scoreInt)
	fmt.Printf("Current: %v\n", currentScore)

	if uint(scoreInt) > uint(currentScore) {
		position = leaderboard.AddScore(userID, uint(scoreInt))
	} else {
		position = leaderboard.GetScore(userID)
	}

	ResponsePosition(w, username, position)
}
