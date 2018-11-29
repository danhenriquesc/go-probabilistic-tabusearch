package helpers

import (
	"fmt"
	"github.com/danhenriquesc/go-probabilistic-tabusearch/pkg/types"
)

func TokenizerSolution(s *types.Solution) string {
	return fmt.Sprint(*s)
}

func TokenizerFullSolution(fs *types.FullSolution) string {
	return fmt.Sprint(*fs)
}

func TokenizerChange(fs *types.FullSolution) string {
	i, j := types.FullSolutionIndexes(fs)
	return fmt.Sprint(i, "|", j)
}
