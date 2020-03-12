package services

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/nevillejrbrown/tictacgo/backend/domain"
	"github.com/nevillejrbrown/tictacgo/backend/storage"
)

func CreateGame(w http.ResponseWriter, r *http.Request) {
	// unmarshall the JSON
	game := domain.Game{}
	decodeError := json.NewDecoder(r.Body).Decode(&game)
	if decodeError != nil {
		http.Error(w, decodeError.Error(), http.StatusBadRequest)
		return
	}
	gameCreated, createErr := storage.GameStore{}.CreateGame(game)

	if createErr != nil {
		fmt.Printf("\nerror!" + createErr.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	data, _ := json.Marshal(gameCreated)
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Content-Length", strconv.Itoa(len(data)))
	w.Header().Set("Location", "games/"+string(gameCreated.Id))

	w.WriteHeader(http.StatusCreated)
	w.Write(data)
}
