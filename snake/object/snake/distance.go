package snake

import (
	c "github.com/theboarderline/ebiten-utilities/snake/core"
	"github.com/theboarderline/ebiten-utilities/snake/events"
	"github.com/theboarderline/ebiten-utilities/snake/object"
	"github.com/theboarderline/ebiten-utilities/snake/param"
	"math"
)

func (s *Snake) CheckSnakeIntersection() {
	curUnit := s.UnitHead.Next
	if curUnit == nil {
		return
	}

	var tolerance float32 = param.ToleranceDefault
	if len(curUnit.CompCollision.Rects) > 1 { // If second unit is on an edge
		tolerance = param.ToleranceScreenEdge // To avoid false collisions on screen edges
	}

	for curUnit != nil {
		if object.Collides(s.UnitHead, curUnit, tolerance) {
			return
		}
		curUnit = curUnit.Next
	}
}

func (s *Snake) CalcSnakeFoodDist(foodLoc c.Vec64) float32 {
	headLoc := s.UnitHead.HeadCenter

	minDist := c.Distance(headLoc, foodLoc)

	if headLoc.X < param.HalfScreenWidth { // Left projection distance
		virtualFood := c.Vec64{X: foodLoc.X - param.ScreenWidth, Y: foodLoc.Y}
		minDist = math.Min(minDist, c.Distance(headLoc, virtualFood))
	} else if headLoc.X >= param.HalfScreenWidth { // Right projection distance
		virtualFood := c.Vec64{X: foodLoc.X + param.ScreenWidth, Y: foodLoc.Y}
		minDist = math.Min(minDist, c.Distance(headLoc, virtualFood))
	}

	if headLoc.Y < param.HalfScreenHeight { // Upper projection distance
		virtualFood := c.Vec64{X: foodLoc.X, Y: foodLoc.Y - param.ScreenHeight}
		minDist = math.Min(minDist, c.Distance(headLoc, virtualFood))
	} else if headLoc.Y >= param.HalfScreenHeight { // Bottom projection distance
		virtualFood := c.Vec64{X: foodLoc.X, Y: foodLoc.Y + param.ScreenHeight}
		minDist = math.Min(minDist, c.Distance(headLoc, virtualFood))
	}

	return float32(minDist)
}

func (s *Snake) HandleDirectionInput(direction events.DirectionT) {
	pressedLeft := direction == events.DirectionLeft
	pressedRight := direction == events.DirectionRight
	pressedUp := direction == events.DirectionUp
	pressedDown := direction == events.DirectionDown

	if !pressedLeft && !pressedRight && !pressedUp && !pressedDown {
		return
	}

	dirCurrent := s.LastDirection()
	dirNew := dirCurrent

	if dirCurrent.IsVertical() {
		if pressedLeft {
			dirNew = events.DirectionLeft
		} else if pressedRight {
			dirNew = events.DirectionRight
		}
	} else {
		if pressedUp {
			dirNew = events.DirectionUp
		} else if pressedDown {
			dirNew = events.DirectionDown
		}
	}

	if dirNew == dirCurrent {
		return
	}

	newTurn := NewTurn(dirCurrent, dirNew)
	s.TurnTo(newTurn, false)
}

func (s *Snake) CheckFoodCollision(distToFood float32) bool {

	if distToFood <= param.RadiusEating {
		s.Grow()
		s.TriggerScoreAnim()
		return true
	}

	return false
}

func (s *Snake) TriggerScoreAnim() *object.ScoreAnim {
	corrCenter := s.UnitHead.HeadCenter

	switch s.UnitHead.Direction {
	case events.DirectionUp:
		corrCenter.Y -= param.RadiusSnake
	case events.DirectionDown:
		corrCenter.Y += param.RadiusSnake
	case events.DirectionRight:
		corrCenter.X += param.RadiusSnake
	case events.DirectionLeft:
		corrCenter.X -= param.RadiusSnake
	}

	return object.NewScoreAnim(corrCenter.To32())
}
