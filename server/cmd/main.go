package main

import (
	"fmt"

	"github.com/vincentvnoord/battleship/internal/game"
)

func main() {
	fmt.Println("Hello, World!")

	ship1 := &game.Ship{Size: 3}
	ship2 := &game.Ship{Size: 3}

	player1 := &game.Player{Id: "1", Name: "Player 1", Board: game.NewBoard(5), Ships: []*game.Ship{ship1}}
	ship1.OwnedBy = player1

	player2 := &game.Player{Id: "2", Name: "Player 2", Board: game.NewBoard(5), Ships: []*game.Ship{ship2}}
	ship2.OwnedBy = player2

	currentGame := game.NewGame([]*game.Player{player1, player2})

	for {
		currentPlayer, err := currentGame.CurrentPlayer()
		if err != nil {
			fmt.Println(err)
			break
		}

		fmt.Printf("%s's turn\n", currentPlayer.Name)
		fmt.Println("Place your ship")
		var position string
		fmt.Scan(&position)

		err = currentGame.PlaceShip(currentPlayer.Id, currentPlayer.Ships[0], 0, 0, game.Horizontal)
		if err != nil {
			fmt.Println(err)
			continue
		}

		currentGame.Turn++
		if currentGame.Turn == 2 {
			break
		}
	}
}
