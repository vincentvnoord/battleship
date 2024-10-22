package simulation_handler

import (
	"github.com/vincentvnoord/battleship/internal/game"
	"github.com/vincentvnoord/battleship/internal/handlers"
	"github.com/vincentvnoord/battleship/pkg/models"
)

func PlaceShips(playerID string, currentGame *game.Game) {
	connection, exists := connManager.GetConnection(playerID)
	if !exists {
		return
	}

	player := currentGame.PlayersMap[playerID]
	if player == nil {
		handlers.SendMessage(connection, models.Message{
			Type:    "action_error",
			Payload: "Player not found",
		})
		return
	}

	if player.Board == nil {
		handlers.SendMessage(connection, models.Message{
			Type:    "action_error",
			Payload: "Player board not found",
		})
		return
	}

	if player.Ships == nil {
		handlers.SendMessage(connection, models.Message{
			Type:    "action_error",
			Payload: "Player ships not found",
		})
		return
	}

	for i, ship := range player.Ships {
		if !ship.Placed {
			err := currentGame.PlaceShip(playerID, ship, 0, i, game.Horizontal)
			if err != nil {
				handlers.SendMessage(connection, models.Message{
					Type:    "action_error",
					Payload: "Error placing ship",
				})
				return
			}

			handlers.SendMessage(connection, models.Message{
				Type: "action_success",
				Payload: models.PublicState{
					GameState: "placed_all_ships",
				},
			})
			return
		}
	}
}
