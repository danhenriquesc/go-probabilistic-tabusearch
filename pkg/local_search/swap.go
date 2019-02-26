package local_search

import (
	"github.com/danhenriquesc/go-probabilistic-tabusearch/pkg/constants"
	"github.com/danhenriquesc/go-probabilistic-tabusearch/pkg/types"
	"github.com/danhenriquesc/go-probabilistic-tabusearch/pkg/fitness"
)

func NeighborhoodBySwap(s *types.FullSolution, distances *types.Distances) []types.FullSolution {
	var neighbors []types.FullSolution

	for i := 1; i <= constants.PROBLEM_SIZE; i++ {
		for j := i + 1; j <= constants.PROBLEM_SIZE; j++ {
			sn := s.Solution
			sn[i], sn[j] = sn[j], sn[i]

			s_before :=  s.Solution
			fitness := fitness.Full(&s_before, distances, s.Fitness, i, j)
			fs := types.FullSolution{sn, fitness, i, j}

			neighbors = append(neighbors, fs)
		}
	}

	return neighbors
}