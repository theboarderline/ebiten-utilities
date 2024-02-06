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
		Name:      s.Name,
		Alive:     s.Alive,
		Speed:     s.Speed,
		UnitHead:  s.UnitHead,
		UnitTail:  s.UnitTail,
		FoodEaten: s.FoodEaten,
	}
	return json.Marshal(minimalSnake)
}
