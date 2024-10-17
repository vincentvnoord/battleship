package game

import (
	"fmt"
)

type Cell struct {
	X, Y       int
	OccupiedBy *Ship
	Hit        bool
}

type Board struct {
	Cells [][]Cell
}

func NewBoard(size int) *Board {
	b := &Board{}
	b.Cells = make([][]Cell, size)
	for i := range b.Cells {
		b.Cells[i] = make([]Cell, size)
	}
	return b
}

func (board *Board) getSize() (int, int) {
	boardWidth := len(board.Cells)
	boardHeight := len(board.Cells[0])
	return boardWidth, boardHeight
}

func (board *Board) PlaceShip(ship *Ship, x, y int, orientation ShipOrientation) error {
	if x < 0 || y < 0 {
		return fmt.Errorf("Ship out of bounds at x: %d, y: %d", x, y)
	}

	width, height := board.getSize()

	if orientation == Horizontal {
		for i := 0; i < ship.Size; i++ {
			x := x + i
			if x >= width {
				return fmt.Errorf("Ship out of bounds at x: %d", x)
			}

			return board.placeShipInCell(x, y, ship)
		}
	} else {
		for i := 0; i < ship.Size; i++ {
			y := y + i
			if y >= height {
				return fmt.Errorf("Ship out of bounds at y: %d", y)
			}

			return board.placeShipInCell(x, y, ship)
		}
	}

	return nil
}

func (board *Board) placeShipInCell(x, y int, ship *Ship) error {
	cell := &board.Cells[x][y]
	if cell.OccupiedBy != nil {
		return fmt.Errorf("Ship already placed at x: %d, y: %d", x, y)
	}

	cell.OccupiedBy = ship
	return nil
}

func (board *Board) Attack(x, y int) error {
	if x < 0 || y < 0 {
		return fmt.Errorf("Attack out of bounds at x: %d, y: %d", x, y)
	}

	width, height := board.getSize()

	if x >= width || y >= height {
		return fmt.Errorf("Attack out of bounds at x: %d, y: %d", x, y)
	}

	cell := &board.Cells[x][y]
	if cell.Hit {
		return fmt.Errorf("Cell already hit at x: %d, y: %d", x, y)
	}

	cell.Hit = true
	if cell.OccupiedBy != nil {
		cell.OccupiedBy.Hits++
	}

	return nil
}
