package handlers

import (
	"encoding/json"
	"errors"
	"net/http"

	"strconv"

	"github.com/tehsis/rabbitscore/rabbitContext"
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
	username := r.FormValue("username")
	score := r.FormValue("score")

	if username == "" {
		ResponseError(w, http.StatusBadRequest, "username is required")
		return
	}

	if score == "" {
		ResponseError(w, http.StatusBadRequest, "score is required")
		return
	}

	if FbID := r.Context().Value(rabbitContext.Context.Auth); FbID != nil {
		player := players.NewFromFacebook(username, FbID.(string))

		store := players.GetStore()

		playerIsValid, err := store.IsValid(player)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		if !playerIsValid {
			ResponseError(w, http.StatusUnauthorized, username+" does not belongs to you")
			return
		}

		scoreInt, err := strconv.Atoi(score)

		if err != nil {
			ResponseError(w, http.StatusInternalServerError, "Internal Error 42")
		}

		position := leaderboard.AddScore(username, uint(scoreInt))

		ResponsePosition(w, username, position)
	}

}
