package game

type Ship struct {
	OwnedBy *Player
	Size    int
	Hits    int
	Placed  bool
}

type ShipOrientation int

const (
	Horizontal ShipOrientation = iota
	Vertical
)
