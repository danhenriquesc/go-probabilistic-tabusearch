package constructive

import (
	"github.com/danhenriquesc/go-probabilistic-tabusearch/pkg/types"
)

type constructor interface {
	generateInitialSolution() types.Solution
}

func NewInitialSolution(c constructor) types.Solution{
	return c.generateInitialSolution()
}
