package param

// Snake parameters
const (
	Localhost              = "127.0.0.1"
	GameserverPort         = 7777
	ClientPort             = 8080
	SnakeSpeedInitial      = 275
	SnakeSpeedFinal        = 250
	SnakeLength            = 240
	SnakeWidth             = 30
	MouthAnimStartDistance = 120
	RadiusSnake            = SnakeWidth / 2.0
	RadiusMouth            = RadiusSnake * 0.625
)

// Snake collision tolerances must be an integer or false collisions will occur.
const (
	ToleranceDefault    = 2 //param.SnakeWidth / 16.0
	ToleranceScreenEdge = RadiusSnake
)
