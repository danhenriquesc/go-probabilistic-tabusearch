package app

import (
	"fmt"
	"github.com/danhenriquesc/go-probabilistic-tabusearch/pkg/reading"
	"github.com/danhenriquesc/go-probabilistic-tabusearch/pkg/constructive"
	"github.com/danhenriquesc/go-probabilistic-tabusearch/pkg/fitness"
	"github.com/danhenriquesc/go-probabilistic-tabusearch/pkg/types"
	"github.com/danhenriquesc/go-probabilistic-tabusearch/pkg/metaheuristics/tabu_search"
)

func Run() error {
	var distances types.Distances

	cities := reading.ReadProblem()
	reading.CalculateDistances(&distances, &cities)

	initialSolution := constructive.NewGreedyInitialSolution(&distances)
	// initialSolution := constructive.NewRandomInitialSolution()

	fitnessInitialSolution := fitness.Simple(&initialSolution, &distances)
	fullInitialSolution := types.NewFullSolution(initialSolution, fitnessInitialSolution, 0, 0)
	fmt.Println(fullInitialSolution)

	err, fullBestSolution := tabu_search.Run(&distances, fullInitialSolution)

	fmt.Println(fullBestSolution)

	return err
}
