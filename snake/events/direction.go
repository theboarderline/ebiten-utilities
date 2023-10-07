package events

type DirectionT uint8

const (
	DirectionUp DirectionT = iota
	DirectionDown
	DirectionLeft
	DirectionRight
	DirectionTotal
)

func (d DirectionT) IsVertical() bool {
	if d >= DirectionTotal {
		panic("wrong direction")
	}
	return (d == DirectionUp) || (d == DirectionDown)
}
