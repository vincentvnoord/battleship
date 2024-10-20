package game

import (
	"testing"
)

func attackHelper(t *testing.T, game *Game, player string, x, y int) {
	err := game.Attack(player, x, y)
	if err != nil {
		t.Fatalf("Expected no error during attack but got %s", err.Error())
	}
	game.NextTurn()
}

func TestNewGame(t *testing.T) {
	players := []*Player{
		{Id: "1", Name: "Player 1"},
		{Id: "2", Name: "Player 2"},
	}

	game := NewGame(players)
	game.AddPlayer(players[0])
	game.AddPlayer(players[1])

	if game.Turn != 0 {
		t.Errorf("Expected turn to be 0 but got %d", game.Turn)
	}

	if game.PlayersMap[players[0].Id] != players[0] {
		t.Errorf("Expected player 1 to be in the game but it wasn't")
	}

	if game.PlayersMap[players[1].Id] != players[1] {
		t.Errorf("Expected player 2 to be in the game but it wasn't")
	}
}

func TestCurrentPlayer(t *testing.T) {
	players := []*Player{
		{Id: "1", Name: "Player 1"},
		{Id: "2", Name: "Player 2"},
	}

	game := NewGame(players)

	currentPlayer, err := game.CurrentPlayer()

	if err != nil {
		t.Errorf("Expected no error but got %s", err.Error())
	}

	if currentPlayer != players[0] {
		t.Errorf("Expected current player to be player 1 but got %s", currentPlayer.Name)
	}

	game.RemovePlayer("1")
	game.RemovePlayer("2")

	_, err = game.CurrentPlayer()
	if err == nil {
		t.Errorf("Expected an error but got none")
	}
}

func TestGetPlayer(t *testing.T) {
	players := []*Player{
		{Id: "1", Name: "Player 1"},
		{Id: "2", Name: "Player 2"},
	}

	game := NewGame(players)

	player := game.getPlayerById("1")
	if player != players[0] {
		t.Errorf("Expected player 1 but got %s", player.Name)
	}

	player = game.getPlayerById("3")
	if player != nil {
		t.Errorf("Expected no player but got %s", player.Name)
	}
}

func TestGamePlaceShip(t *testing.T) {
	players := []*Player{
		{Id: "1", Name: "Player 1", Board: NewBoard(10)},
		{Id: "2", Name: "Player 2", Board: NewBoard(10)},
	}

	game := NewGame(players)

	ship := &Ship{Size: 3}

	err := game.PlaceShip("1", ship, 0, 0, Horizontal)
	if err != nil {
		t.Errorf("Expected no error but got %s", err.Error())
	}

	err = game.PlaceShip("3", ship, 0, 0, Horizontal)
	if err == nil {
		t.Errorf("Expected an error but got none")
	}
}

func TestGameAttack(t *testing.T) {
	players := []*Player{
		{Id: "1", Name: "Player 1", Board: NewBoard(10)},
		{Id: "2", Name: "Player 2", Board: NewBoard(10)},
	}

	game := NewGame(players)

	err := game.Attack("2", 0, 0)
	if err != nil {
		t.Errorf("Expected no error but got %s", err.Error())
	}
	game.NextTurn()

	err = game.Attack("1", 0, 0)
	if err != nil {
		t.Errorf("Expected no error but got %s", err.Error())
	}

	err = game.Attack("3", 0, 0)
	if err == nil {
		t.Errorf("Expected an error but got none")
	}
}

func TestGameFinish(t *testing.T) {
	players := []*Player{
		{Id: "1", Name: "Player 1", Board: NewBoard(10)},
		{Id: "2", Name: "Player 2", Board: NewBoard(10)},
	}

	game := NewGame(players)

	game.PlayersMap["1"].Board.PlaceShip(&Ship{Size: 4}, 0, 0, Horizontal)

	game.PlayersMap["2"].Board.PlaceShip(&Ship{Size: 3}, 0, 0, Horizontal)

	attackHelper(t, game, "2", 0, 0)
	attackHelper(t, game, "1", 0, 0)
	attackHelper(t, game, "2", 1, 0)
	attackHelper(t, game, "1", 1, 0)
	attackHelper(t, game, "2", 2, 0)
	attackHelper(t, game, "1", 2, 0)

	if !game.Finished() {
		t.Errorf("Expected game to be finished but it wasn't")
	}
}
