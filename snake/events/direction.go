package events

import "math/rand"

type DirectionT uint8

const (
	DirectionUp    DirectionT = 'u'
	DirectionDown             = 'd'
	DirectionLeft             = 'l'
	DirectionRight            = 'r'
)

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

func (d DirectionT) IsVertical() bool {
	return d == DirectionUp || d == DirectionDown
}

func (d DirectionT) IsHorizontal() bool {
	return d == DirectionLeft || d == DirectionRight
}

func (d DirectionT) IsValid() bool {
	return d.IsVertical() || d.IsHorizontal()
}
