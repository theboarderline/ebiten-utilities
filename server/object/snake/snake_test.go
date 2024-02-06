package snake_test

import (
	"fmt"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/theboarderline/ebiten-utilities/server/object/snake"
	"github.com/theboarderline/ebiten-utilities/server/param"
)

var _ = Describe("Snake", func() {

	var (
		name = "Snake"
		s    *snake.Snake
	)

	BeforeEach(func() {
		s = snake.NewSnakeRandDirLoc(name, param.SnakeLength, param.SnakeSpeedInitial, &param.ColorSnake1)
	})

	It("can marshal and unmarshal a snake", func() {
		snakeJSON, err := s.Marshal()
		Expect(err).To(BeNil())
		Expect(snakeJSON).NotTo(BeNil())
		Expect(snakeJSON).NotTo(BeEmpty())
		Expect(snakeJSON).To(ContainSubstring(name))
		Expect(snakeJSON).To(ContainSubstring(fmt.Sprint(param.SnakeSpeedInitial)))
		Expect(snakeJSON).To(ContainSubstring(fmt.Sprint(param.SnakeLength)))

		s2, err := snake.UnmarshalJSON(snakeJSON)
		Expect(err).To(BeNil())
		Expect(s2).NotTo(BeNil())
		Expect(s2.Name).To(Equal(name))
		Expect(s2.Speed).To(BeNumerically("==", param.SnakeSpeedInitial))
		Expect(s2.UnitHead).NotTo(BeNil())
		Expect(s2.UnitHead.HeadCenter).NotTo(BeNil())
		Expect(s2.UnitHead.HeadCenter.X).NotTo(BeZero())
		Expect(s2.UnitHead.HeadCenter.Y).NotTo(BeZero())
		Expect(s2.FoodEaten).To(BeNumerically("==", 0))

	})

})
