package core

// Rectangle compatible with float32 type fields of ebiten.Vertex struct.
type RectF32 struct {
	Pos       Vec32
	Size      Vec32
	PosInUnit Vec32
}

func NewRect(pos, size Vec32) *RectF32 {
	return &RectF32{pos, size, Vec32{0, 0}}
}
