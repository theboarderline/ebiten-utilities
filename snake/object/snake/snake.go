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
	"encoding/json"
	"fmt"
	"github.com/theboarderline/ebiten-utilities/snake/events"
	"image/color"
	"log"
	"math"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	c "github.com/theboarderline/ebiten-utilities/snake/core"
	"github.com/theboarderline/ebiten-utilities/snake/param"
	"github.com/theboarderline/ebiten-utilities/snake/shader"
)

var (
	imageCircle    = ebiten.NewImage(param.SnakeWidth, param.SnakeWidth)
	shaderMouth    = shader.New(shader.PathCircleMouth)
	MouthEnabled   = false
	optTriangEmpty ebiten.DrawTrianglesOptions
)

func init() {
	rand.Seed(time.Now().UnixNano())

	// Prepare circle image whose radius is snake's half width
	imageCircle.DrawRectShader(param.SnakeWidth, param.SnakeWidth, &shader.Circle, &ebiten.DrawRectShaderOptions{
		Uniforms: map[string]interface{}{
			"Radius": float32(param.RadiusSnake),
		},
	})

}

type Snake struct {
	Name            string                            `json:"name"`
	Dead            bool                              `json:"alive"`
	Speed           float64                           `json:"speed"`
	UnitHead        *Unit                             `json:"head"`
	unitTail        *Unit                             `json:"tail"`
	turnPrev        *Turn                             `json:"turnPrev"`
	turnQueue       []*Turn                           `json:"turnQueue"`
	distAfterTurn   float64                           `json:"distAfterTurn"`
	growthRemaining float64                           `json:"growthRemaining"`
	growthTarget    float64                           `json:"growthTarget"`
	FoodEaten       uint8                             `json:"foodEaten"`
	color           *color.RGBA                       `json:"color"`
	drawOptsHead    ebiten.DrawTrianglesShaderOptions `json:"drawOptsHead"`
}

func NewSnake(name string, headCenter c.Vec64, initialLength uint16, speed float64, direction events.DirectionT, color *color.RGBA) *Snake {

	if name == "" {
		name = fmt.Sprintf("Snake %f-%f", headCenter.X, headCenter.Y)
	}

	if !direction.IsValid() {
		direction = events.DirectionRight
	}

	if headCenter.X > param.ScreenWidth {
		headCenter.X = param.ScreenWidth / 2
	}

	if headCenter.Y > param.ScreenHeight {
		headCenter.Y = param.ScreenHeight / 2
	}

	if isVertical := direction.IsVertical(); (isVertical && (initialLength > param.ScreenHeight)) ||
		(!isVertical && (initialLength > param.ScreenWidth)) {
		panic("Initial snake intersects itself.")
	}
	if color == nil {
		panic("Snake color cannot be nil")
	}

	initialUnit := NewUnit(headCenter, float64(initialLength), direction, color)

	snake := &Snake{
		Name:     name,
		Speed:    speed,
		UnitHead: initialUnit,
		unitTail: initialUnit,
		color:    color,
		drawOptsHead: ebiten.DrawTrianglesShaderOptions{
			Uniforms: map[string]interface{}{
				"Radius":      float32(param.RadiusSnake),
				"RadiusMouth": float32(param.RadiusMouth),
			},
		},
	}

	return snake
}

func NewSnakeFromResponse(response string) (*Snake, error) {
	var s Snake

	if err := json.Unmarshal([]byte(response), &s); err != nil {
		log.Print(err)
		return nil, err
	}

	return &s, nil
}

func NewSnakeRandDir(name string, headCenter c.Vec64, initialLength uint16, speed float64, color *color.RGBA) *Snake {
	direction := events.NewRandomDirection()
	return NewSnake(name, headCenter, initialLength, speed, direction, color)
}

func NewSnakeRandDirLoc(name string, initialLength uint16, speed float64, color *color.RGBA) *Snake {
	headCenter := c.Vec64{
		X: float64(rand.Intn(param.ScreenWidth)),
		Y: float64(rand.Intn(param.ScreenHeight)),
	}
	return NewSnakeRandDir(name, headCenter, initialLength, speed, color)
}

func NewRandSnake(name string) *Snake {
	if name == "" {
		randNumber := rand.Intn(100)
		name = fmt.Sprintf("Snake %d", randNumber)
	}
	return NewSnakeRandDirLoc(name, param.SnakeLength, param.SnakeSpeedInitial, &param.ColorSnake1)
}

func (s *Snake) Update(distToFood float32) {
	if s.Dead {
		return
	}

	moveDistance := s.Speed * param.DeltaTime

	// if the snake has moved a safe distance after the last turn, take the next turn in the queue.
	if (len(s.turnQueue) > 0) && (s.distAfterTurn+param.ToleranceDefault >= param.SnakeWidth) {
		var nextTurn *Turn
		nextTurn, s.turnQueue = s.turnQueue[0], s.turnQueue[1:] // Pop front
		s.TurnTo(nextTurn, true)
	}

	s.updateHead(moveDistance, distToFood)
	s.updateTail(moveDistance, distToFood)
}

func (s *Snake) updateHead(dist float64, distToFood float32) {
	// Increase head length
	s.UnitHead.length += dist

	// Move head
	switch s.UnitHead.Direction {
	case events.DirectionRight:
		s.UnitHead.moveRight(dist)
	case events.DirectionLeft:
		s.UnitHead.moveLeft(dist)
	case events.DirectionUp:
		s.UnitHead.moveUp(dist)
	case events.DirectionDown:
		s.UnitHead.moveDown(dist)
	}

	if s.UnitHead != s.unitTail { // Avoid unnecessary updates
		s.UnitHead.update(distToFood)
	}

	// Distance to food

	// Update draw options
	proxToFood := 1.0 - distToFood/param.MouthAnimStartDistance
	s.drawOptsHead.Uniforms["Direction"] = float32(s.UnitHead.Direction)
	s.drawOptsHead.Uniforms["ProximityToFood"] = proxToFood

	s.distAfterTurn += dist
}

func (s *Snake) updateTail(dist float64, distToFood float32) {
	decreaseAmount := dist
	if s.growthRemaining > 0 {
		// Calculate the tail reduction with the square function so that the growth doesn't look ugly.
		// f(x) = 0.8 + 3.0*(x-0.5)^2)/4.0 where  0 <= x <= 1
		growthCompletion := 1 - s.growthRemaining/s.growthTarget
		decreaseAmount *= (0.8 + 3.0*(growthCompletion-0.5)*(growthCompletion-0.5)/4.0)
		s.growthRemaining -= (dist - decreaseAmount)
	} else {
		s.growthTarget = s.growthRemaining
	}

	// Decrease tail length
	s.unitTail.length -= decreaseAmount

	// Delete tail if its length is less than width of the snake
	if (s.unitTail.prev != nil) && (s.unitTail.length <= param.SnakeWidth) {
		s.unitTail.prev.length += s.unitTail.length
		s.unitTail = s.unitTail.prev
		s.unitTail.Next = nil
	}

	s.unitTail.update(distToFood)
}

func (s *Snake) TurnTo(newTurn *Turn, isFromQueue bool) {
	if !isFromQueue {
		// Check if the new turn is dangerous (twice same turns rapidly).
		if (s.turnPrev != nil) &&
			(s.turnPrev.isTurningLeft == newTurn.isTurningLeft) &&
			(s.distAfterTurn+param.ToleranceDefault <= param.SnakeWidth) {
			// New turn cannot be taken now, push it into the queue
			s.turnQueue = append(s.turnQueue, newTurn)
			return
		}
		// If there are turns in the queue then add the new turn to the queue as well.
		if len(s.turnQueue) > 0 {
			s.turnQueue = append(s.turnQueue, newTurn)
			return
		}
	}
	s.distAfterTurn = 0

	oldHead := s.UnitHead

	// Decide on the color of the new head unit.
	newColor := s.color
	// if param.DebugUnits && (oldHead.color == s.color) {
	// 	newColor = &param.ColorSnake2
	// }

	// Create a new head unit.
	newHead := NewUnit(oldHead.HeadCenter, 0, newTurn.directionTo, newColor)

	// Add the new head unit to the beginning of the unit doubly linked list.
	newHead.Next = oldHead
	oldHead.prev = newHead
	s.UnitHead = newHead

	// Update prev turn
	s.turnPrev = newTurn
}

func (s *Snake) Grow() {
	// Compute the new growth and add to the remaining growth value.
	// f(x)=50+5*log2(x/10.0+1)
	newGrowth := 50.0 + 5.0*math.Log2(float64(s.FoodEaten)/10.0+1.0)
	s.growthRemaining += newGrowth
	s.growthTarget += newGrowth
	s.FoodEaten++

	// Update snake speed
	// f(x)=250+25/e^(0.0075x)
	s.Speed = param.SnakeSpeedFinal + (param.SnakeSpeedInitial-param.SnakeSpeedFinal)/math.Exp(0.0075*float64(s.FoodEaten))
}

func (s *Snake) LastDirection() events.DirectionT {
	// if the turn queue is not empty, return the direction of the last turn to be taken.
	if queueLength := len(s.turnQueue); queueLength > 0 {
		return s.turnQueue[queueLength-1].directionTo
	}

	// return current head direction
	return s.UnitHead.Direction
}

func (s *Snake) Draw(dst *ebiten.Image) {
	if s.Dead {
		return
	}

	for unit := s.UnitHead; unit != nil; unit = unit.Next {
		// Draw circle centered on unit's head center
		vertices, indices := unit.CompTriangHead.Triangles()
		if MouthEnabled && (unit == s.UnitHead) {
			dst.DrawTrianglesShader(vertices, indices, shaderMouth, &s.drawOptsHead)
		} else {
			dst.DrawTriangles(vertices, indices, imageCircle, &optTriangEmpty)
		}

		if unit.Next == nil {
			// Draw circle centered on unit's tail center
			vertices, indices = unit.CompTriangTail.Triangles()
			dst.DrawTriangles(vertices, indices, imageCircle, &optTriangEmpty)
		}

		// Draw rectangle starts from unit's head center to the tail head center
		unit.CompBody.Draw(dst)

		if param.DebugUnits {
			unit.DrawDebugInfo(dst)
		}
	}
}
