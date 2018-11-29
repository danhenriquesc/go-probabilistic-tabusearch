package app

import (
	"fmt"
	"time"
	"math/rand"
	"github.com/danhenriquesc/go-probabilistic-tabusearch/pkg/constants"
	"github.com/danhenriquesc/go-probabilistic-tabusearch/pkg/reading"
	"github.com/danhenriquesc/go-probabilistic-tabusearch/pkg/constructive"
	"github.com/danhenriquesc/go-probabilistic-tabusearch/pkg/types"
	"github.com/danhenriquesc/go-probabilistic-tabusearch/pkg/local_search"
)

const maxTabuSize = 50
const iterations = 700
const pertubation = 3

func Fitness(s *types.Solution, distances *[constants.PROBLEM_SIZE + 1][constants.PROBLEM_SIZE + 1]int) int {
	var fitness int
	
	for i := 1; i < constants.PROBLEM_SIZE; i++ {
		fitness += distances[s[i]][s[i+1]]
	}

	fitness += distances[s[constants.PROBLEM_SIZE]][s[1]]

	return fitness
}

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

func Contains(needed string, tabuList *[]string) bool{
	for _, item := range *tabuList {
		if item == needed {
			return true
		}
	}
	return false
}

func Run() error {
	rand.Seed(time.Now().UTC().UnixNano())
	start := time.Now()

	var distances [constants.PROBLEM_SIZE + 1][constants.PROBLEM_SIZE + 1]int

	cities := reading.ReadProblem()
	reading.CalculateDistances(&distances, &cities)

	initialSolution := constructive.NewGreedyInitialSolution(&distances)
	// initialSolution := constructive.NewRandomInitialSolution()
	fitnessInitialSolution := Fitness(&initialSolution, &distances)
	fullInitialSolution := types.NewFullSolution(initialSolution, fitnessInitialSolution, 0, 0)

	fmt.Println(initialSolution)

	fullBestSolution := fullInitialSolution

	fmt.Println(fullBestSolution)
	
	fullBestCandidate := fullInitialSolution
	
	var tabuList []string
	// tabuList = append(tabuList, Tokenizertypes.FullSolution(&fullBestCandidate))

	x := 1
	for x < iterations {
		var neighborhood []types.FullSolution

		if x % pertubation == 0 {
			neighborhood = local_search.NeighborhoodBySwap(&fullBestCandidate, &distances)
		} else {
			neighborhood = local_search.NeighborhoodBy2opt(&fullBestCandidate, &distances)
		}

		first := true
		for _, candidate := range neighborhood {
			// notTabu := !Contains(Tokenizertypes.FullSolution(&candidate), &tabuList)
			notTabu := !Contains(TokenizerChange(&candidate), &tabuList)

			if notTabu && (first || types.FullSolutionFitness(&candidate) < types.FullSolutionFitness(&fullBestCandidate) ){
				fullBestCandidate = candidate
				first = false
			}
		}

		if types.FullSolutionFitness(&fullBestCandidate) < types.FullSolutionFitness(&fullBestSolution) {
			fullBestSolution = fullBestCandidate
		}

		tabuList = append(tabuList, TokenizerChange(&fullBestCandidate))
		if len(tabuList) > maxTabuSize {
			tabuList = tabuList[1:]
		}

		x += 1
		fmt.Println(types.FullSolutionFitness(&fullBestSolution), types.FullSolutionFitness(&fullBestCandidate))
	}

	fmt.Println(fullBestSolution)

	fmt.Println(time.Since(start))

	return nil
}
