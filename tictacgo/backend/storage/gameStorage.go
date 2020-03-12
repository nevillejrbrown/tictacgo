package storage

import (
	"github.com/nevillejrbrown/tictacgo/backend/domain"
)

type IGameStore interface {
	// Init()
	CreateGame(gameToStore domain.Game) (gameCreated domain.Game, err error)
}

type GameStore struct {
}

var instance *GameStore

func GameStoreInstance() *GameStore {
	if instance == nil {
		instance = &GameStore{}
	}
	return instance
}

var games = make(map[string]domain.Game) // map[string]domain.Game
var gamesAdded int

// func (gs GameStore) Init() {
// 	games = make(map[string]domain.Game)
// 	gamesAdded = 0
// }

func (gs GameStore) CreateGame(gameToStore domain.Game) (gameCreated domain.Game, err error) {
	gamesAdded++
	gameToStore.Id = gamesAdded
	games[string(gamesAdded)] = gameToStore
	return gameToStore, nil // won't
}
