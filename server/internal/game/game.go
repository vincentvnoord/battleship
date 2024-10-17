package game

import "fmt"

type Player struct {
	Id    string
	Name  string
	Board *Board
	Ships []*Ship
}

type Game struct {
	Board   *Board
	Players []*Player
	Turn    int
}

func NewGame(board *Board, players []*Player) *Game {
	return &Game{board, players, 0}
}

func (g *Game) CurrentPlayer() *Player {
	return g.Players[g.Turn]
}

func (g *Game) PlaceShip(playerId string, ship *Ship, x, y int, orientation ShipOrientation) error {
	player := g.getPlayerById(playerId)
	if player == nil {
		return fmt.Errorf("Player with id %s not found", playerId)
	}

	return player.Board.PlaceShip(ship, x, y, orientation)
}

func (g *Game) getPlayerById(playerId string) *Player {
	for _, player := range g.Players {
		if player.Id == playerId {
			return player
		}
	}
	return nil
}
