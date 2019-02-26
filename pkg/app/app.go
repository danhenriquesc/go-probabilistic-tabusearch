package app

import (
	"fmt"
	"errors"
	"github.com/danhenriquesc/go-probabilistic-tabusearch/pkg/reading"
	"github.com/danhenriquesc/go-probabilistic-tabusearch/pkg/constructive"
	"github.com/danhenriquesc/go-probabilistic-tabusearch/pkg/fitness"
	"github.com/danhenriquesc/go-probabilistic-tabusearch/pkg/types"
	"github.com/danhenriquesc/go-probabilistic-tabusearch/pkg/metaheuristics/tabu_search"
)

func Run(metaheuristic string) error {
	var distances types.Distances

	cities := reading.ReadProblem()
	reading.CalculateDistances(&distances, &cities)

	constructiveMetaheuristic := constructive.GreedyConstructive{&distances}
	initialSolution := constructive.NewInitialSolution(constructiveMetaheuristic)

	fitnessInitialSolution := fitness.Simple(&initialSolution, &distances)
	fullInitialSolution := types.FullSolution{initialSolution, fitnessInitialSolution, 0, 0}
	fmt.Println(fullInitialSolution)

	switch metaheuristic {
		case "TABU_SEARCH":
			err, fullBestSolution := tabu_search.Run(&distances, fullInitialSolution)

			fmt.Println(fullBestSolution)
			return err
		default:
			return errors.New("INVALID METAHEURISTIC")
	}

	return nil
}
