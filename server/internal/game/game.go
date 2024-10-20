package game

import "fmt"

type GameState int

const (
	PlacingShips GameState = iota
	Attacking
	GameOver
)

type Game struct {
	PlayerList []*string
	PlayersMap map[string]*Player
	Turn       int
	State      GameState
}

func NewGame(players []*Player) *Game {
	playersMap := make(map[string]*Player)
	playerList := make([]*string, len(players))
	for i, player := range players {
		playersMap[player.Id] = player
		playerList[i] = &player.Id
	}

	return &Game{playerList, playersMap, 0, PlacingShips}
}

func (g *Game) Finished() bool {
	playersNotDefeated := 0

	for _, player := range g.PlayersMap {
		if !player.IsDefeated() {
			playersNotDefeated++
		}
	}

	return playersNotDefeated <= 1
}

func (g *Game) AddPlayer(player *Player) error {
	if _, exists := g.PlayersMap[player.Id]; exists {
		return fmt.Errorf("Player with id %s already exists", player.Id)
	}
	g.PlayersMap[player.Id] = player
	g.PlayerList = append(g.PlayerList, &player.Id)

	return nil
}

func (g *Game) RemovePlayer(playerId string) error {
	if _, exists := g.PlayersMap[playerId]; !exists {
		return fmt.Errorf("Player with id %s does not exist", playerId)
	}

	delete(g.PlayersMap, playerId)
	for i, id := range g.PlayerList {
		if *id == playerId {
			g.PlayerList = append(g.PlayerList[:i], g.PlayerList[i+1:]...)
			break
		}
	}

	return nil
}

func (g *Game) CurrentPlayer() (*Player, error) {
	if len(g.PlayersMap) == 0 {
		return nil, fmt.Errorf("No players in the game")
	}

	playerId := g.PlayerList[g.Turn]
	return g.PlayersMap[*playerId], nil
}

func (g *Game) getPlayerById(playerId string) *Player {
	return g.PlayersMap[playerId]
}

func (g *Game) PlaceShip(playerId string, ship *Ship, x, y int, orientation ShipOrientation) error {
	player := g.getPlayerById(playerId)
	if player == nil {
		return fmt.Errorf("Player with id %s not found", playerId)
	}

	return player.Board.PlaceShip(ship, x, y, orientation)
}

func (g *Game) Attack(playerId string, x, y int) error {
	player := g.PlayersMap[playerId]
	if player == nil {
		return fmt.Errorf("Player with id %s not found", playerId)
	}

	currentPlayer, err := g.CurrentPlayer()
	if err != nil {
		return err
	}

	if currentPlayer.Id == playerId {
		return fmt.Errorf("Player %s cannot attack themselves", player.Name)
	}

	return currentPlayer.Board.Attack(x, y)
}

func (g *Game) NextTurn() {
	nextTurn := g.Turn + 1
	if nextTurn >= len(g.PlayersMap) {
		nextTurn = 0
	}

	g.Turn = nextTurn
}
