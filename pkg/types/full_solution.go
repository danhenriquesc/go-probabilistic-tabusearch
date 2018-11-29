package types

type FullSolution struct {
	solution Solution
	fitness int
	i, j int
}

func FullSolutionSolution(fs *FullSolution) Solution {
	return fs.solution
}

func FullSolutionFitness(fs *FullSolution) int {
	return fs.fitness
}

func FullSolutionIndexes(fs *FullSolution) (int, int) {
	return fs.i, fs.j
}

func NewFullSolution(s Solution, fitness, i, j int) FullSolution {
	return FullSolution{s, fitness, i, j}
}