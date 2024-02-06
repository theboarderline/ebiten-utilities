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

package object

import (
	c "github.com/theboarderline/ebiten-utilities/server/core"
	"github.com/theboarderline/ebiten-utilities/server/events"
	"github.com/theboarderline/ebiten-utilities/server/param"
)

const (
	decrementAlpha = 8.0
	scoreAnimSpeed = 25
)

var (
	scoreAnimShiftY float32
)

type ScoreAnim struct {
	pos       c.Vec32
	alpha     uint8
	direction events.DirectionT
}

func NewScoreAnim(pos c.Vec32) *ScoreAnim {
	newAnim := &ScoreAnim{
		pos: c.Vec32{
			X: pos.X,
			Y: pos.Y - scoreAnimShiftY,
		},
		alpha:     param.ColorScore.A,
		direction: events.DirectionUp,
	}

	return newAnim
}

// Update Returns true when the animation is finished
func (s *ScoreAnim) Update() bool {
	// Move animation
	s.pos.Y -= scoreAnimSpeed * param.DeltaTime

	// Decrease alpha
	if s.alpha < decrementAlpha {
		return true
	}

	s.alpha -= decrementAlpha
	return false
}
