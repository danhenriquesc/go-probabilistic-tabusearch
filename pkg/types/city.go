package types

type City struct {
	x, y float64
}

func NewCity(x, y float64) City {
	return City{x, y}
}

func CityXY(c City) (float64, float64) {
	return c.x, c.y
}