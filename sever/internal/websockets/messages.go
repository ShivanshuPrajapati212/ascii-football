package websockets

const InitGameType string = "init_game"

type GameMessage struct {
	Type    string `json:"type"`
	Payload string `json:"payload"`
}
