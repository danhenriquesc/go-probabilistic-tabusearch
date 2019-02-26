package types

import (
	"github.com/danhenriquesc/go-probabilistic-tabusearch/pkg/constants"
)

type City struct {
	X, Y float64
}

type Cities [constants.PROBLEM_SIZE + 1]City

type Distances [constants.PROBLEM_SIZE + 1][constants.PROBLEM_SIZE + 1]int

type Solution [constants.PROBLEM_SIZE + 1]int

type FullSolution struct {
	Solution Solution
	Fitness int
	I, J int
}
