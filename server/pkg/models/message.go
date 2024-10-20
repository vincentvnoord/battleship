package models

type Message struct {
	Type    string      `json:"type"`
	Payload interface{} `json:"payload"`
}

type GameMessage struct {
	NewState string   `json:"new_state"`
	Players  []Player `json:"players"`
}

type Player struct {
	PlayerID string   `json:"player_id"`
	Name     string   `json:"name"`
	Board    [][]Cell `json:"board"`
}

type Cell struct {
	X   int  `json:"x"`
	Y   int  `json:"y"`
	Hit bool `json:"hit"`
}

type Ship struct {
	X           int    `json:"x"`
	Y           int    `json:"y"`
	Orientation string `json:"orientation"`
	Size        int    `json:"size"`
	Placed      bool   `json:"placed"`
}
