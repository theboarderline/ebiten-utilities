package snake

import (
	"github.com/rs/zerolog/log"
	"strings"
)

func makeSnakes(response string) map[string]*Snake {
	snakes := make(map[string]*Snake)

	for _, s := range parseSnakeResponse(response) {
		snakes[s.Name] = s
	}

	return snakes
}

func parseSnakeResponse(response string) []*Snake {
	snakes := make([]*Snake, 0)

	for _, item := range strings.Split(response, "\n") {
		if item == "" {
			continue
		}

		s, err := NewSnakeFromResponse(item)
		if err != nil {
			log.Error().Err(err).Msg("Error parsing snake response")
		}

		snakes = append(snakes, s)
	}

	return nil
}
