/*
Copyright (C) 2022 Anıl Konaç

This file is part of snake-ebiten.

snake-ebiten is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

snake-ebiten is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with snake-ebiten. If not, see <https://www.gnu.org/licenses/>.
*/

package snake

import (
	c "github.com/theboarderline/ebiten-utilities/server/core"
	"github.com/theboarderline/ebiten-utilities/server/events"
	"github.com/theboarderline/ebiten-utilities/server/param"
	"image/color"
	"math"
)

type Unit struct {
	ProximityToFood float32           `json:"proxToFood"`
	HeadCenter      c.Vec64           `json:"headCenter"`
	Length          float64           `json:"Length"`
	Direction       events.DirectionT `json:"direction"`
	Next            *Unit             `json:"next"`
	Prev            *Unit             `json:"Prev"`
}

func NewUnit(headCenter c.Vec64, length float64, direction events.DirectionT, color *color.RGBA) *Unit {
	return &Unit{
		HeadCenter: headCenter,
		Length:     length,
		Direction:  direction,
	}
}

func (u *Unit) createRectCollision() (rectColl *c.RectF32) {
	length32 := float32(math.Floor(u.Length))
	flCenter := u.HeadCenter.Floor().To32()

	switch u.Direction {
	case events.DirectionRight:
		rectColl = c.NewRect(
			c.Vec32{
				X: flCenter.X - length32 + param.RadiusSnake,
				Y: flCenter.Y - param.RadiusSnake,
			},
			c.Vec32{X: length32, Y: param.SnakeWidth},
		)
	case events.DirectionLeft:
		rectColl = c.NewRect(
			c.Vec32{
				X: flCenter.X - param.RadiusSnake,
				Y: flCenter.Y - param.RadiusSnake,
			},
			c.Vec32{X: length32, Y: param.SnakeWidth},
		)
	case events.DirectionUp:
		rectColl = c.NewRect(
			c.Vec32{
				X: flCenter.X - param.RadiusSnake,
				Y: flCenter.Y - param.RadiusSnake,
			},
			c.Vec32{X: param.SnakeWidth, Y: length32},
		)
	case events.DirectionDown:
		rectColl = c.NewRect(
			c.Vec32{
				X: flCenter.X - param.RadiusSnake,
				Y: flCenter.Y - length32 + param.RadiusSnake,
			},
			c.Vec32{X: param.SnakeWidth, Y: length32})
	default:
		panic("Wrong unit direction!!")
	}

	return
}

func (u *Unit) createRectDraw(rectColl *c.RectF32) (rectDraw *c.RectF32) {
	if u.Next == nil {
		rectDraw = rectColl
		return
	}

	switch u.Direction {
	case events.DirectionRight:
		rectDraw = c.NewRect(c.Vec32{X: rectColl.Pos.X - param.SnakeWidth, Y: rectColl.Pos.Y}, c.Vec32{X: rectColl.Size.X + param.SnakeWidth, Y: rectColl.Size.Y})
	case events.DirectionLeft:
		rectDraw = c.NewRect(c.Vec32{X: rectColl.Pos.X, Y: rectColl.Pos.Y}, c.Vec32{X: rectColl.Size.X + param.SnakeWidth, Y: rectColl.Size.Y})
	case events.DirectionUp:
		rectDraw = c.NewRect(c.Vec32{X: rectColl.Pos.X, Y: rectColl.Pos.Y}, c.Vec32{X: rectColl.Size.X, Y: rectColl.Size.Y + param.SnakeWidth})
	case events.DirectionDown:
		rectDraw = c.NewRect(c.Vec32{X: rectColl.Pos.X, Y: rectColl.Pos.Y - param.SnakeWidth}, c.Vec32{X: rectColl.Size.X, Y: rectColl.Size.Y + param.SnakeWidth})
	default:
		panic("Wrong unit direction!!")
	}

	return
}

func (u *Unit) createRectHead() *c.RectF32 {
	headCenter32 := u.HeadCenter.Floor().To32()
	return c.NewRect(c.Vec32{X: headCenter32.X - param.RadiusSnake, Y: headCenter32.Y - param.RadiusSnake}, c.Vec32{X: param.SnakeWidth, Y: param.SnakeWidth})
}

func (u *Unit) createRectTail(rectHead *c.RectF32) (rectTail *c.RectF32) {
	size := c.Vec32{X: param.SnakeWidth, Y: param.SnakeWidth}
	switch u.Direction {
	case events.DirectionUp:
		rectTail = c.NewRect(c.Vec32{X: rectHead.Pos.X, Y: rectHead.Pos.Y + float32(u.Length) - param.SnakeWidth}, size)
	case events.DirectionDown:
		rectTail = c.NewRect(c.Vec32{X: rectHead.Pos.X, Y: rectHead.Pos.Y - float32(u.Length) + param.SnakeWidth}, size)
	case events.DirectionLeft:
		rectTail = c.NewRect(c.Vec32{X: rectHead.Pos.X + float32(u.Length) - param.SnakeWidth, Y: rectHead.Pos.Y}, size)
	case events.DirectionRight:
		rectTail = c.NewRect(c.Vec32{X: rectHead.Pos.X - float32(u.Length) + param.SnakeWidth, Y: rectHead.Pos.Y}, size)
	}
	return
}

func (u *Unit) createRectBody(rectColl *c.RectF32) (rectBody *c.RectF32) {

	switch u.Direction {
	case events.DirectionUp:
		rectBody = c.NewRect(
			c.Vec32{X: rectColl.Pos.X, Y: rectColl.Pos.Y + param.RadiusSnake},
			c.Vec32{X: rectColl.Size.X, Y: rectColl.Size.Y - param.SnakeWidth},
		)
		if u.Next != nil {
			rectBody.Size.Y += param.SnakeWidth
		}
	case events.DirectionDown:
		rectBody = c.NewRect(
			c.Vec32{X: rectColl.Pos.X, Y: rectColl.Pos.Y + param.RadiusSnake},
			c.Vec32{X: rectColl.Size.X, Y: rectColl.Size.Y - param.SnakeWidth},
		)
		if u.Next != nil {
			rectBody.Pos.Y -= param.SnakeWidth
			rectBody.Size.Y += param.SnakeWidth
		}
	case events.DirectionLeft:
		rectBody = c.NewRect(
			c.Vec32{X: rectColl.Pos.X + param.RadiusSnake, Y: rectColl.Pos.Y},
			c.Vec32{X: rectColl.Size.X - param.SnakeWidth, Y: rectColl.Size.Y},
		)
		if u.Next != nil {
			rectBody.Size.X += param.SnakeWidth
		}
	case events.DirectionRight:
		rectBody = c.NewRect(
			c.Vec32{X: rectColl.Pos.X + param.RadiusSnake, Y: rectColl.Pos.Y},
			c.Vec32{X: rectColl.Size.X - param.SnakeWidth, Y: rectColl.Size.Y},
		)
		if u.Next != nil {
			rectBody.Size.X += param.SnakeWidth
			rectBody.Pos.X -= param.SnakeWidth
		}
	}
	return
}

func (u *Unit) moveUp(dist float64) {
	u.HeadCenter.Y -= dist

	// teleport if head center is offscreen.
	if param.TeleportEnabled && (u.HeadCenter.Y < 0) {
		u.HeadCenter.Y += param.ScreenHeight
	}
}

func (u *Unit) moveDown(dist float64) {
	u.HeadCenter.Y += dist

	// teleport if head center is offscreen.
	if param.TeleportEnabled && (u.HeadCenter.Y > param.ScreenHeight) {
		u.HeadCenter.Y -= param.ScreenHeight
	}
}

func (u *Unit) moveRight(dist float64) {
	u.HeadCenter.X += dist

	// teleport if head center is offscreen.
	if param.TeleportEnabled && (u.HeadCenter.X > param.ScreenWidth) {
		u.HeadCenter.X -= param.ScreenWidth
	}
}

func (u *Unit) moveLeft(dist float64) {
	u.HeadCenter.X -= dist

	// teleport if head center is offscreen.
	if param.TeleportEnabled && (u.HeadCenter.X < 0) {
		u.HeadCenter.X += param.ScreenWidth
	}
}

func (u *Unit) markHeadCenters() {

	var offset float64 = 0
	if u.Next == nil {
		offset = param.SnakeWidth
	}

	backCenter := u.HeadCenter
	switch u.Direction {
	case events.DirectionUp:
		backCenter.Y = u.HeadCenter.Y + u.Length - offset
	case events.DirectionDown:
		backCenter.Y = u.HeadCenter.Y - u.Length + offset
	case events.DirectionRight:
		backCenter.X = u.HeadCenter.X - u.Length + offset
	case events.DirectionLeft:
		backCenter.X = u.HeadCenter.X + u.Length - offset
	}
}

// Implement collidable interface
// ------------------------------
func (u *Unit) CollEnabled() bool {
	return true
}

func (u *Unit) CollisionRects() []c.RectF32 {
	return []c.RectF32{}
}
