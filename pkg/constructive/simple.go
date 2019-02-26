package constructive

import (
	"github.com/danhenriquesc/go-probabilistic-tabusearch/pkg/constants"
	"github.com/danhenriquesc/go-probabilistic-tabusearch/pkg/types"
)

type SimpleConstructive struct {}

func (sc SimpleConstructive) generateInitialSolution() types.Solution {
	initialSolution := types.Solution{}
	for i := 1; i <= constants.PROBLEM_SIZE; i++ {
		initialSolution[i] = i
	}
	return initialSolution
}