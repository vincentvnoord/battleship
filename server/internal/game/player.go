package game

type Player struct {
	Connected bool
	Id        string
	Name      string
	Board     *Board
	Ships     []*Ship
}

func (p *Player) IsDefeated() bool {
	for _, ship := range p.Ships {
		if !ship.IsDestroyed() {
			return false
		}
	}
	return true
}

func NewPlayer(id, name string, board *Board, ships []*Ship) *Player {
	startingShips := make([]*Ship, len(ships))
	copy(startingShips, ships)

	return &Player{Id: id, Name: name, Board: board, Ships: startingShips}
}
