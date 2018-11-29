package local_search

import (
	"github.com/danhenriquesc/go-probabilistic-tabusearch/pkg/constants"
	"github.com/danhenriquesc/go-probabilistic-tabusearch/pkg/types"
	"github.com/danhenriquesc/go-probabilistic-tabusearch/pkg/fitness"
)

func NeighborhoodBySwap(s *types.FullSolution, distances *[constants.PROBLEM_SIZE + 1][constants.PROBLEM_SIZE + 1]int) []types.FullSolution {
	var neighbors []types.FullSolution

	for i := 1; i <= constants.PROBLEM_SIZE; i++ {
		for j := i + 1; j <= constants.PROBLEM_SIZE; j++ {
			sn := types.FullSolutionSolution(s)
			sn[i], sn[j] = sn[j], sn[i]

			s_before :=  types.FullSolutionSolution(s)
			fitness := fitness.Full(&s_before, distances, types.FullSolutionFitness(s), i, j)
			fs := types.NewFullSolution(sn, fitness, i, j)

			neighbors = append(neighbors, fs)
		}
	}

	return neighbors
}