package game_handler

import (
	"log"
	"net/http"

	"github.com/vincentvnoord/battleship/internal/game"
	"github.com/vincentvnoord/battleship/internal/handlers"
	game_service "github.com/vincentvnoord/battleship/internal/services/game"
	"github.com/vincentvnoord/battleship/pkg/models"

	"github.com/gorilla/websocket"
)

// HandleGameConnection manages the WebSocket connection for a game.
func HandleGameConnection(w http.ResponseWriter, r *http.Request, gameService *game_service.GameService) {
	gameId := r.URL.Path[len("/game/"):]
	if gameId == "" {
		http.Error(w, "gameId is required", http.StatusBadRequest)
		return
	}

	// Get the game instance from the game service
	currentGame, err := gameService.GetGame(gameId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	// Upgrade HTTP connection to WebSocket
	conn, err := handlers.Upgrader.Upgrade(w, r, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Handle the connection
	go handleWebSocketConnection(conn, currentGame, gameService)
}

// handleWebSocketConnection manages incoming messages and responses for the WebSocket connection.
func handleWebSocketConnection(conn *websocket.Conn, currentGame *game.Game, gameService *game_service.GameService) {
	defer conn.Close()

	// Assign a player ID and notify the client
	playerID, err := gameService.GetOpenPlayerSlot(currentGame)
	if err != nil {
		log.Println("Error assigning player ID:", err)
		return
	}

	if playerID == "" {
		log.Println("Failed to assign player ID")
		return
	}

	initialMessage := models.Message{
		Type:    "player_assigned",
		Payload: playerID,
	}

	// Respond with the assigned player ID
	err = handlers.SendMessage(conn, initialMessage)
	if err != nil {
		log.Println("Error writing message:", err)
		return
	}

	// Optional: Handle further messages if needed (for now, we will just close after sending the ID)
	// You can keep this loop if you want to manage further interactions
	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			log.Println("Error reading message:", err)
			break
		}

		// Here, you can choose to respond or handle further messages as needed
		log.Printf("Received message: %s\n", message)
	}
}
