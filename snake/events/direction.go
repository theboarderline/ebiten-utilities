package events

type DirectionT uint8

const (
	DirectionUp    DirectionT = 'u'
	DirectionDown             = 'd'
	DirectionLeft             = 'l'
	DirectionRight            = 'r'
)

func (d DirectionT) IsVertical() bool {
	return d == DirectionUp || d == DirectionDown
}

func (d DirectionT) IsHorizontal() bool {
	return d == DirectionLeft || d == DirectionRight
}

func (d DirectionT) IsValid() bool {
	return d.IsVertical() || d.IsHorizontal()
}
