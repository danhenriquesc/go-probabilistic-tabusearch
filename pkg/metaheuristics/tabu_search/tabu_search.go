package tabu_search

import (
	"github.com/danhenriquesc/go-probabilistic-tabusearch/pkg/types"
	"github.com/danhenriquesc/go-probabilistic-tabusearch/pkg/local_search"
	"github.com/danhenriquesc/go-probabilistic-tabusearch/pkg/helpers"
)

const maxTabuSize = 50
const iterations = 700
const pertubation = 3

func Run(distances *types.Distances, fullInitialSolution types.FullSolution) (error, types.FullSolution){
	// fmt.Println(initialSolution)

	fullBestSolution := fullInitialSolution
	
	fullBestCandidate := fullInitialSolution
	
	var tabuList []string
	// tabuList = append(tabuList, Tokenizertypes.FullSolution(&fullBestCandidate))

	x := 1
	for x < iterations {
		var neighborhood []types.FullSolution

		if x % pertubation == 0 {
			neighborhood = local_search.NeighborhoodBySwap(&fullBestCandidate, distances)
		} else {
			neighborhood = local_search.NeighborhoodBy2opt(&fullBestCandidate, distances)
		}

		first := true
		for _, candidate := range neighborhood {
			// notTabu := !helpers.Contains(Tokenizertypes.FullSolution(&candidate), &tabuList)
			notTabu := !helpers.Contains(helpers.TokenizerChange(&candidate), &tabuList)

			if notTabu && (first || types.FullSolutionFitness(&candidate) < types.FullSolutionFitness(&fullBestCandidate) ){
				fullBestCandidate = candidate
				first = false
			}
		}

		if types.FullSolutionFitness(&fullBestCandidate) < types.FullSolutionFitness(&fullBestSolution) {
			fullBestSolution = fullBestCandidate
		}

		tabuList = append(tabuList, helpers.TokenizerChange(&fullBestCandidate))
		if len(tabuList) > maxTabuSize {
			tabuList = tabuList[1:]
		}

		x += 1
		// fmt.Println(types.FullSolutionFitness(fullBestSolution), types.FullSolutionFitness(fullBestCandidate))
	}

	return nil, fullBestSolution
}