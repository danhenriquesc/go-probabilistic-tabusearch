package tabu_search

import (
	"fmt"
	"sort"
	"github.com/danhenriquesc/go-probabilistic-tabusearch/pkg/types"
	"github.com/danhenriquesc/go-probabilistic-tabusearch/pkg/local_search"
	"github.com/danhenriquesc/go-probabilistic-tabusearch/pkg/helpers"
	"github.com/danhenriquesc/go-probabilistic-tabusearch/pkg/constants"
	"github.com/danhenriquesc/go-probabilistic-tabusearch/pkg/fitness"
)

func Run(distances *types.Distances, fullInitialSolution types.FullSolution) (error, types.FullSolution){
	config := load_config(constants.PROBLEM_NAME)

	maxTabuSize := config["maxTabuSize"]
	iterations := config["iterations"]
	pertubation := config["pertubation"]
	optimal := config["optimal"]

	gap0 := float64(optimal) * 1.0
	gap0beat := false
	gap1 := float64(optimal) * 1.01
	gap1beat := false
	gap10 := float64(optimal) * 1.1
	gap10beat := false
	gap25 := float64(optimal) * 1.25
	gap25beat := false
	gap50 := float64(optimal) * 1.5
	gap50beat := false
	// fmt.Println(initialSolution)

	fullBestSolution := fullInitialSolution
	fullBestCandidate := fullInitialSolution
	
	var tabuList []string
	// tabuList = append(tabuList, Tokenizertypes.FullSolution(&fullBestCandidate))

	iteration := 0
	for iteration < iterations {
		var neighborhood []types.FullSolution

		if iteration % pertubation == 0 {
			neighborhood = local_search.NeighborhoodBySwap(&fullBestCandidate, distances)

			// TSP SINGLE IMPROVEMENT
			first := true
			for _, candidate := range neighborhood {
				// notTabu := !helpers.Contains(Tokenizertypes.FullSolution(&candidate), &tabuList)
				notTabu := !helpers.Contains(helpers.TokenizerChange(&candidate), &tabuList)

				if notTabu && (first || types.FullSolutionFitness(&candidate) < types.FullSolutionFitness(&fullBestCandidate) ){
					fullBestCandidate = candidate
					first = false
				} else if (!notTabu && types.FullSolutionFitness(&candidate) < types.FullSolutionFitness(&fullBestSolution) ) { // Aspiração
					fullBestCandidate = candidate
					break
				}
			}
		} else {
			neighborhood = local_search.NeighborhoodBy2opt(&fullBestCandidate, distances)

			// TSP MULTIIMPROVEMENT
			var goodCandidates []types.FullSolution
			var selecteds []types.FullSolution

			for _, candidate := range neighborhood {
				// notTabu := !helpers.Contains(Tokenizertypes.FullSolution(&candidate), &tabuList)
				notTabu := !helpers.Contains(helpers.TokenizerChange(&candidate), &tabuList)

				if notTabu && types.FullSolutionFitness(&candidate) < types.FullSolutionFitness(&fullBestCandidate) {
					goodCandidates = append(goodCandidates, candidate)
				} else if (!notTabu && types.FullSolutionFitness(&candidate) < types.FullSolutionFitness(&fullBestSolution) ) { // Aspiração
					goodCandidates = append(goodCandidates, candidate)
					break
				}
			}

			fmt.Println(len(goodCandidates))
			sort.Slice(goodCandidates, func(a, b int) bool {
				fitA := types.FullSolutionFitness(&goodCandidates[a])
				fitB := types.FullSolutionFitness(&goodCandidates[b])
				return fitA < fitB
			})

			fmt.Println("GOOD CANDIDATES:")
			for _, goodCandidate := range goodCandidates {
				improvement := types.FullSolutionFitness(&fullBestCandidate) - types.FullSolutionFitness(&goodCandidate)
				i, j := types.FullSolutionIndexes(&goodCandidate)
				fmt.Println(i, j, improvement)

				blocked := false

				for _, selected := range selecteds {
					s_i, s_j := types.FullSolutionIndexes(&selected)

					//conflict?
					if !((i < s_i && j < s_i) || (i > s_j && j > s_j)) {
						blocked = true
						break
					}
				}

				if blocked == false {
					selecteds = append(selecteds, goodCandidate)
				}
			}

			fmt.Println("SELECTEDS:")
			for _, selected := range selecteds {
				improvement := types.FullSolutionFitness(&fullBestCandidate) - types.FullSolutionFitness(&selected)
				i, j := types.FullSolutionIndexes(&selected)
				fmt.Println("I:", i, "J:", j, "DIFF:", improvement)
			}

			//apply swaps
			fmt.Println("PRE-SWAPS:")
			fmt.Println(fullBestCandidate)
			for _, selected := range selecteds {
				sn := types.FullSolutionSolution(&fullBestCandidate)
				i, j := types.FullSolutionIndexes(&selected)

				st := i
				end := j
				for st < end {
					sn[st], sn[end] = sn[end], sn[st]
					st++
					end--
				}

				s_before := types.FullSolutionSolution(&fullBestCandidate)
				fitness := fitness.Full2opt(&s_before, distances, types.FullSolutionFitness(&fullBestCandidate), i, j)
				fullBestCandidate = types.NewFullSolution(sn, fitness, i, j)
			}
			fmt.Println("POS-SWAPS:")
			fmt.Println(fullBestCandidate)
		}


		if types.FullSolutionFitness(&fullBestCandidate) < types.FullSolutionFitness(&fullBestSolution) {
			fullBestSolution = fullBestCandidate
		}

		tabuList = append(tabuList, helpers.TokenizerChange(&fullBestCandidate))
		if len(tabuList) > maxTabuSize {
			tabuList = tabuList[1:]
		}

		if !gap50beat && float64(types.FullSolutionFitness(&fullBestSolution)) <= gap50{
			fmt.Printf("GAP 50 in iteration %d\n", iteration)
			gap50beat = true
		}

		if !gap25beat && float64(types.FullSolutionFitness(&fullBestSolution)) <= gap25{
			fmt.Printf("GAP 25 in iteration %d\n", iteration)
			gap25beat = true
		}

		if !gap10beat && float64(types.FullSolutionFitness(&fullBestSolution)) <= gap10{
			fmt.Printf("GAP 10 in iteration %d\n", iteration)
			gap10beat = true
		}

		if !gap1beat && float64(types.FullSolutionFitness(&fullBestSolution)) <= gap1{
			fmt.Printf("GAP 1 in iteration %d\n", iteration)
			gap1beat = true
		}

		if !gap0beat && float64(types.FullSolutionFitness(&fullBestSolution)) <= gap0{
			fmt.Printf("OPTIMAL in iteration %d\n", iteration)
			gap0beat = true
		}


		iteration += 1
		// fmt.Println(types.FullSolutionFitness(fullBestSolution), types.FullSolutionFitness(fullBestCandidate))
	}

	return nil, fullBestSolution
}