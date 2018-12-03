package reading

import (
	"math"
	"github.com/danhenriquesc/go-probabilistic-tabusearch/pkg/constants"
	"github.com/danhenriquesc/go-probabilistic-tabusearch/pkg/types"
)

func Nint(n float64) int {
	return int(n + 0.5)
}

func CalculateDistanceParallel(a, b types.City, ch chan int) {
	a_x, a_y := types.CityXY(a)
	b_x, b_y := types.CityXY(b)
	ch <- Nint(math.Sqrt(math.Pow(a_x - b_x, 2) + math.Pow(a_y - b_y, 2)))
}

func CalculateDistance(a, b types.City) int {
	a_x, a_y := types.CityXY(a)
	b_x, b_y := types.CityXY(b)
	return Nint(math.Sqrt(math.Pow(a_x - b_x, 2) + math.Pow(a_y - b_y, 2)))
}

func CalculateDistances(distances *types.Distances, cities *types.Cities) {
	ch := make(chan int)
	for i := 1; i <= constants.PROBLEM_SIZE; i++ {
		for j := i + 1; j <= constants.PROBLEM_SIZE; j++ {
			go CalculateDistanceParallel(cities[i], cities[j], ch)

			distance := <- ch
			distances[i][j] = distance
			distances[j][i] = distance
		}
	}	
}