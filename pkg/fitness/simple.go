package fitness

import (
	"github.com/danhenriquesc/go-probabilistic-tabusearch/pkg/constants"
	"github.com/danhenriquesc/go-probabilistic-tabusearch/pkg/types"
)

func Simple(s *types.Solution, distances *[constants.PROBLEM_SIZE + 1][constants.PROBLEM_SIZE + 1]int) int {
	var fitness int
	
	for i := 1; i < constants.PROBLEM_SIZE; i++ {
		fitness += distances[s[i]][s[i+1]]
	}

	fitness += distances[s[constants.PROBLEM_SIZE]][s[1]]

	return fitness
}