package main

import (
	"fmt"
	"net/http"

	simulation_handler "github.com/vincentvnoord/battleship/internal/handlers/simulation"
	game_service "github.com/vincentvnoord/battleship/internal/services/game"
)

func main() {
	gameService := game_service.NewGameService() // Ensure this is correctly initialized

	playerNames := []string{"Alice", "Bob"}                          // This is a valid slice of strings
	err, _ := gameService.CreateGame(playerNames, StringPtr("test")) // Ensure CreateGame can handle []string
	if err != nil {
		panic(err)
	}

	http.HandleFunc("/game/", func(w http.ResponseWriter, r *http.Request) {
		simulation_handler.HandleGameConnection(w, r, gameService)
	})

	fmt.Printf("Starting test simulation on :8080\n")
	http.ListenAndServe(":8080", nil) // This will not return an error; remove the second call
}

func StringPtr(s string) *string {
	return &s
}
