package constructive

import (
	"github.com/danhenriquesc/go-probabilistic-tabusearch/pkg/constants"
	"github.com/danhenriquesc/go-probabilistic-tabusearch/pkg/types"
)

func NewInitialSolution() types.Solution {
	initialSolution := types.Solution{}
	for i := 1; i <= constants.PROBLEM_SIZE; i++ {
		initialSolution[i] = i
	}
	return initialSolution
}