package events

import "math/rand"

type Event struct {
	PlayerName     string     `json:"playerName,omitempty"`
	Type           string     `json:"type,omitempty"`
	Message        string     `json:"message,omitempty"`
	InputDirection DirectionT `json:"direction,omitempty"`
}

func NewRandomDirection() (direction DirectionT) {

	if rand.Float64() < 0.25 {
		direction = DirectionUp
	} else if rand.Float64() < 0.5 {
		direction = DirectionDown
	} else if rand.Float64() < 0.75 {
		direction = DirectionLeft
	} else {
		direction = DirectionRight
	}

	return direction
}
