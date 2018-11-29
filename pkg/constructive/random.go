package constructive

import (
	"math/rand"
	"github.com/danhenriquesc/go-probabilistic-tabusearch/pkg/constants"
	"github.com/danhenriquesc/go-probabilistic-tabusearch/pkg/types"
)

func NewRandomInitialSolution() types.Solution {
	initialSolution := types.Solution{}
	rands := rand.Perm(constants.PROBLEM_SIZE)

	for i := 1; i <= constants.PROBLEM_SIZE; i++ {
		initialSolution[i] = rands[i-1] + 1
	}

	return initialSolution
}
