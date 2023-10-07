package events

type Event struct {
	PlayerName     string     `json:"playerName,omitempty"`
	Type           string     `json:"type,omitempty"`
	Message        string     `json:"message,omitempty"`
	InputDirection DirectionT `json:"direction,omitempty"`
}
