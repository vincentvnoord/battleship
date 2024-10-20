package lobby

type Lobby struct {
	ID          string
	PlayerNames []*string
}

func NewLobby(id string) *Lobby {
	return &Lobby{ID: id}
}

func (l *Lobby) AddPlayer(playerName string) {
	l.PlayerNames = append(l.PlayerNames, &playerName)
}
