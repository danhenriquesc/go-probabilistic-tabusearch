package types

import (
	"github.com/danhenriquesc/go-probabilistic-tabusearch/pkg/constants"
)

type City struct {
	x, y float64
}

func NewCity(x, y float64) City {
	return City{x, y}
}

func CityXY(c City) (float64, float64) {
	return c.x, c.y
}

type Cities [constants.PROBLEM_SIZE + 1]City
type Distances [constants.PROBLEM_SIZE + 1][constants.PROBLEM_SIZE + 1]int