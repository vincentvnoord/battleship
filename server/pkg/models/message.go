package models

type Message struct {
	Type    string      `json:"type"`
	Payload interface{} `json:"payload"`
}

type PublicState struct {
	GameID                string   `json:"game_id"`
	GameState             string   `json:"game_state"`
	Players               []Player `json:"players"`
	CurrentAttackedPlayer string   `json:"current_attacked_player"`
}

type Player struct {
	Name  string   `json:"name"`
	Board [][]Cell `json:"board"`
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
