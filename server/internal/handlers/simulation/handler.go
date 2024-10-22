package simulation_handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/vincentvnoord/battleship/internal/game"
	"github.com/vincentvnoord/battleship/internal/handlers"
	"github.com/vincentvnoord/battleship/internal/services"
	game_service "github.com/vincentvnoord/battleship/internal/services/game"
	"github.com/vincentvnoord/battleship/internal/transformers"
	"github.com/vincentvnoord/battleship/pkg/models"
)

var connManager = services.NewConnectionManager()

func HandleGameConnection(w http.ResponseWriter, r *http.Request, gameService *game_service.GameService) {
	gameId := r.URL.Path[len("/game/"):]
	if gameId == "" {
		http.Error(w, "gameId is required", http.StatusBadRequest)
		return
	}

	currentGame, err := gameService.GetGame(gameId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	conn, err := handlers.Upgrader.Upgrade(w, r, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	go handleGameConnection(conn, currentGame, gameService)
}

func handleGameConnection(conn *websocket.Conn, currentGame *game.Game, gameService *game_service.GameService) {
	defer conn.Close()

	playerID, err := gameService.GetOpenPlayerSlot(currentGame)
	if err != nil {
		log.Println("Error assigning player ID:", err)
		return
	}

	connManager.AddConnection(playerID, conn)

	initialMessage := models.Message{
		Type:    "player_assigned",
		Payload: playerID,
	}

	err = handlers.SendMessage(conn, initialMessage)
	if err != nil {
		log.Println("Error writing message:", err)
		return
	}

	playerLength := len(currentGame.PlayerList)
	connectedLength := 0
	if playerLength >= 2 {
		for _, player := range currentGame.PlayerList {
			if currentGame.PlayersMap[*player].Connected {
				connectedLength++
			}
		}
	}

	if connectedLength == playerLength {
		OnStartGame(currentGame)
	}

	for {
		_, messageBytes, err := conn.ReadMessage()
		if err != nil {
			log.Println("Error reading message:", err)
			break
		}

		var message models.Message
		if err := json.Unmarshal(messageBytes, &message); err != nil {
			log.Println("Error unmarshalling message:", err)
			continue
		}

		switch message.Type {
		case "place_ships":
			PlaceShips(playerID, currentGame)
		}

		log.Printf("Received message: %s\n", message)
	}

	connManager.RemoveConnection(playerID)
}

func OnStartGame(currentGame *game.Game) {
	gameState := transformers.TransformPublicState(currentGame, "placing_ships")

	for _, id := range currentGame.PlayerList {
		conn, exists := connManager.GetConnection(*id)
		if !exists {
			log.Println("Connection not found for player:", *id)
			continue
		}

		handlers.SendMessage(conn, models.Message{
			Type:    "game_start",
			Payload: gameState,
		})
	}
}
