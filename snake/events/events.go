package events

type Event struct {
	PlayerName string `json:"playerName"`
	Type       string `json:"type"`
	Message    string `json:"message"`
}
