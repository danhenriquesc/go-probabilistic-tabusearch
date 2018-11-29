package local_search

import (
	"github.com/danhenriquesc/go-probabilistic-tabusearch/pkg/constants"
	"github.com/danhenriquesc/go-probabilistic-tabusearch/pkg/types"
)

func NeighborhoodBy2opt(s *types.FullSolution, distances *[constants.PROBLEM_SIZE + 1][constants.PROBLEM_SIZE + 1]int) []types.FullSolution {
	var neighbors []types.FullSolution

	for i := 1; i <= constants.PROBLEM_SIZE; i++ {
		for j := i + 2; j <= constants.PROBLEM_SIZE; j++ {
			sn := types.FullSolutionSolution(s)

			st := i
			end := j
			for st < end {
				sn[st], sn[end] = sn[end], sn[st]
				st++
				end--
			}

			s_before := types.FullSolutionSolution(s)
			fitness := FullFitness2opt(&s_before, distances, types.FullSolutionFitness(s), i, j)
			fs := types.NewFullSolution(sn, fitness, i, j)

			neighbors = append(neighbors, fs)
		}
	}

	return neighbors
}


func Fitness(s *types.Solution, distances *[constants.PROBLEM_SIZE + 1][constants.PROBLEM_SIZE + 1]int) int {
	var fitness int
	
	for i := 1; i < constants.PROBLEM_SIZE; i++ {
		fitness += distances[s[i]][s[i+1]]
	}

	fitness += distances[s[constants.PROBLEM_SIZE]][s[1]]

	return fitness
}

func FullFitness(s *types.Solution, distances *[constants.PROBLEM_SIZE + 1][constants.PROBLEM_SIZE + 1]int, current_fitness, i, j int) int {
	fitness := current_fitness
	
	if j - i != constants.PROBLEM_SIZE - 1 {
		if i > 1 {
			fitness -= distances[s[i-1]][s[i]]
			fitness += distances[s[i-1]][s[j]]
		} else {
			fitness -= distances[s[constants.PROBLEM_SIZE]][s[i]]
			fitness += distances[s[constants.PROBLEM_SIZE]][s[j]]
		}

		if j - i > 1 {
			fitness -= distances[s[i]][s[i+1]]
			fitness += distances[s[j]][s[i+1]]
			
			fitness -= distances[s[j-1]][s[j]]
			fitness += distances[s[j-1]][s[i]]
		}

		if j < constants.PROBLEM_SIZE {
			fitness -= distances[s[j]][s[j+1]]
			fitness += distances[s[i]][s[j+1]]
		} else {
			fitness -= distances[s[j]][s[1]]
			fitness += distances[s[i]][s[1]]
		}
	} else {
		fitness -= distances[s[j-1]][s[j]]
		fitness += distances[s[j-1]][s[1]]

		fitness -= distances[s[1]][s[2]]
		fitness += distances[s[j]][s[2]]
	}

	return fitness
}

func FullFitness2opt(s *types.Solution, distances *[constants.PROBLEM_SIZE + 1][constants.PROBLEM_SIZE + 1]int, current_fitness, i, j int) int {
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