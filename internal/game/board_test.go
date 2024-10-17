package game

import (
	"testing"
)

func TestCreateBoard(t *testing.T) {
	b := NewBoard(10)
	w, h := b.getSize()
	if w != 10 || h != 10 {
		t.Errorf("Expected board size to be 10, got %d", w)
	}
}

func TestPlaceShip(t *testing.T) {
	b := NewBoard(10)
	ship := &Ship{Size: 5}
	err := b.PlaceShip(ship, 0, 0, Horizontal)
	if err != nil {
		t.Errorf("Expected ship to be placed, got error: %s", err)
	}

	_, height := b.getSize()
	err = b.PlaceShip(ship, 0, height, Vertical)
	if err == nil {
		t.Errorf("Expected ship to not be placed, got ship placed outside of bounds")
	}

	err = b.PlaceShip(ship, -1, -1, Vertical)
	if err == nil {
		t.Errorf("Expected ship to not be placed, got ship placed outside of bounds (negative coordinates)")
	}
}

func TestPlaceShipOccupied(t *testing.T) {
	b := NewBoard(10)
	ship1 := &Ship{Size: 5}
	ship2 := &Ship{Size: 5}
	err := b.PlaceShip(ship1, 0, 0, Horizontal)
	if err != nil {
		t.Errorf("Expected ship to be placed, got error: %s", err)
	}

	err = b.PlaceShip(ship2, 0, 0, Horizontal)
	if err == nil {
		t.Errorf("Expected ship to not be placed, got ship placed on top of another ship")
	}
}

func TestAttackHit(t *testing.T) {
	b := NewBoard(10)
	ship := &Ship{Size: 5}
	b.PlaceShip(ship, 0, 0, Horizontal)
	b.Attack(0, 0)
	if ship.Hits != 1 {
		t.Errorf("Expected ship to be hit, got %d hits", ship.Hits)
	}
}

func TestAttackMiss(t *testing.T) {
	b := NewBoard(10)
	ship := &Ship{Size: 5}
	b.PlaceShip(ship, 0, 0, Horizontal)
	b.Attack(0, 1)
	if ship.Hits != 0 {
		t.Errorf("Expected ship to not be hit, got %d hits", ship.Hits)
	}
}

func TestAttackOutOfBounds(t *testing.T) {
	b := NewBoard(10)
	err := b.Attack(-1, -1)
	if err == nil {
		t.Errorf("Expected attack to be out of bounds, got no error")
	}

	err = b.Attack(10, 10)
	if err == nil {
		t.Errorf("Expected attack to be out of bounds, got no error")
	}
}
