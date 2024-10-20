package game_service

import (
	"fmt"
	"sync"

	"github.com/vincentvnoord/battleship/internal/game"
)

type GameService struct {
	games map[string]*game.Game
	mu    sync.Mutex
}

func NewGameService() *GameService {
	return &GameService{
		games: make(map[string]*game.Game),
	}
}

func (s *GameService) CreateGame(playerNames []string, gameCode *string) (error, *string) {
	s.mu.Lock()
	defer s.mu.Unlock()

	gameId := ""
	if gameCode != nil {
		gameId = *gameCode
		if _, exists := s.games[gameId]; exists {
			return fmt.Errorf("Game with id %s already exists", gameId), nil
		}
	} else {
		gameId = fmt.Sprintf("game-%d", len(s.games)+1)
	}

	if _, exists := s.games[gameId]; exists {
		return fmt.Errorf("Game with id %s already exists", gameId), nil
	}

	// Create set of defaultShips to be copied for each player
	defaultShips := []*game.Ship{
		{Size: 5},
		{Size: 3},
		{Size: 2},
		{Size: 1},
	}

	players := make([]*game.Player, len(playerNames))
	for i, playerName := range playerNames {
		// Create a new player with random id
		player := game.NewPlayer(fmt.Sprintf("player-%d", i+1), playerName, game.NewBoard(10), defaultShips)
		players[i] = player
	}

	s.games[gameId] = game.NewGame(players)

	return nil, &gameId
}

func (s *GameService) GetGame(gameId string) (*game.Game, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	game, exists := s.games[gameId]
	if !exists {
		return nil, fmt.Errorf("Game with id %s does not exist", gameId)
	}

	return game, nil
}

func (s *GameService) GetOpenPlayerSlot(game *game.Game) (string, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	for _, playerId := range game.PlayerList {
		player := game.PlayersMap[*playerId]
		if !player.Connected {
			player.Connected = true
			return *playerId, nil
		}
	}

	return "", fmt.Errorf("No open player slots available")
}
