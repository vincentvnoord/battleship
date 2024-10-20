package transformers

import (
	"github.com/vincentvnoord/battleship/internal/game"
	"github.com/vincentvnoord/battleship/pkg/models"
)

func TransformCell(cell game.Cell) models.Cell {
	return models.Cell{
		X:   cell.X,
		Y:   cell.Y,
		Hit: cell.Hit,
	}
}

func TransformBoard(board game.Board) [][]models.Cell {
	boardModel := [][]models.Cell{}
	for _, row := range board.Cells {
		cellRow := []models.Cell{}
		for _, cell := range row {
			cellModel := TransformCell(cell)
			cellRow = append(cellRow, cellModel)
		}
		boardModel = append(boardModel, cellRow)
	}
	return boardModel
}

func TransformPlayers(players []*string, g *game.Game) []models.Player {
	playersModel := []models.Player{}

	for _, player := range players {
		playerData := g.PlayersMap[*player]

		playerModel := models.Player{
			PlayerID: *player,
			Name:     playerData.Name,
			Board:    TransformBoard(*playerData.Board),
		}

		playersModel = append(playersModel, playerModel)
	}

	return playersModel
}

func TransformGame(g *game.Game, newState string) models.GameMessage {
	players := TransformPlayers(g.PlayerList, g)

	return models.GameMessage{
		NewState: newState,
		Players:  players,
	}
}
