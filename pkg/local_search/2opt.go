package local_search

import (
	"github.com/danhenriquesc/go-probabilistic-tabusearch/pkg/constants"
	"github.com/danhenriquesc/go-probabilistic-tabusearch/pkg/types"
	"github.com/danhenriquesc/go-probabilistic-tabusearch/pkg/fitness"
)

func NeighborhoodBy2opt(s *types.FullSolution, distances *types.Distances) []types.FullSolution {
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
			fitness := fitness.Full2opt(&s_before, distances, types.FullSolutionFitness(s), i, j)
			fs := types.NewFullSolution(sn, fitness, i, j)

			neighbors = append(neighbors, fs)
		}
	}

	return neighbors
}
