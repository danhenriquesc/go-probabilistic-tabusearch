package fitness

import (
	"github.com/danhenriquesc/go-probabilistic-tabusearch/pkg/constants"
	"github.com/danhenriquesc/go-probabilistic-tabusearch/pkg/types"
)

func Simple(s *types.Solution, distances *types.Distances) int {
	var fitness int
	
	for i := 1; i < constants.PROBLEM_SIZE; i++ {
		fitness += distances[s[i]][s[i+1]]
	}

	fitness += distances[s[constants.PROBLEM_SIZE]][s[1]]

	return fitness
}