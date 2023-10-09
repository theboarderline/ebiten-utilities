package snake

import (
	"encoding/json"
	"github.com/rs/zerolog/log"
)

func UnmarshalJSON(b []byte) (*Snake, error) {
	var s Snake
	if err := json.Unmarshal(b, &s); err != nil {
		log.Error().Err(err).Msg("Error unmarshaling snake")
		return nil, err
	}

	return &s, nil
}

func (s *Snake) Marshal() ([]byte, error) {
	minimalSnake := Snake{
		Name:         s.Name,
		Dead:         s.Dead,
		Speed:        s.Speed,
		UnitHead:     s.UnitHead,
		unitTail:     s.unitTail,
		color:        s.color,
		FoodEaten:    s.FoodEaten,
		drawOptsHead: s.drawOptsHead,
	}
	return json.Marshal(minimalSnake)
}
