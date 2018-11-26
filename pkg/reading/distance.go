package reading

import (
	"math"
	"github.com/danhenriquesc/go-probabilistic-tabusearch/pkg/constants"
)

func Nint(n float64) int {
	return int(n + 0.5)
}

func CalculateDistance(a, b City, ch chan int) {
	ch <- Nint(math.Sqrt(math.Pow(a.x - b.x, 2) + math.Pow(a.y - b.y, 2)))
}

func CalculateDistances(distances *[constants.PROBLEM_SIZE + 1][constants.PROBLEM_SIZE + 1]int, cities *[constants.PROBLEM_SIZE + 1]City) {
	ch := make(chan int)
	for i := 1; i <= constants.PROBLEM_SIZE; i++ {
		for j := i + 1; j <= constants.PROBLEM_SIZE; j++ {
			go CalculateDistance(cities[i], cities[j], ch)

			distance := <- ch
			distances[i][j] = distance
			distances[j][i] = distance
		}
	}	
}