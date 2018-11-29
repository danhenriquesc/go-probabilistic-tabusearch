package fitness

import (
	"github.com/danhenriquesc/go-probabilistic-tabusearch/pkg/constants"
	"github.com/danhenriquesc/go-probabilistic-tabusearch/pkg/types"
)

func Full2opt(s *types.Solution, distances *types.Distances, current_fitness, i, j int) int {
	fitness := current_fitness
	
	if j - i != constants.PROBLEM_SIZE - 1 {
		if i > 1 {
			fitness -= distances[s[i-1]][s[i]]
			fitness += distances[s[i-1]][s[j]]
		} else {
			fitness -= distances[s[constants.PROBLEM_SIZE]][s[i]]
			fitness += distances[s[constants.PROBLEM_SIZE]][s[j]]
		}

		if j < constants.PROBLEM_SIZE {
			fitness -= distances[s[j]][s[j+1]]
			fitness += distances[s[i]][s[j+1]]
		} else {
			fitness -= distances[s[j]][s[1]]
			fitness += distances[s[i]][s[1]]
		}
	}
	
	return fitness
}