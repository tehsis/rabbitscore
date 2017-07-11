package handlers

import (
	"encoding/json"
	"errors"
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

	scores := leaderboard.GetTopTen()

	if err := json.NewEncoder(w).Encode(scores); err != nil {
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

	name, err := players.GetStore().GetPlayerName(userID)

	if uint(scoreInt) > currentScore {
		position = leaderboard.AddScore(name, uint(scoreInt))
	} else {
		position = leaderboard.GetScore(name)
	}

	ResponsePosition(w, username, position)
}
