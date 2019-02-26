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
		var selecteds []types.FullSolution

		if iteration % pertubation == 0 {
			neighborhood = local_search.NeighborhoodBySwap(&fullBestCandidate, distances)

			// TSP SINGLE IMPROVEMENT
			first := true
			for _, candidate := range neighborhood {
				// notTabu := !helpers.Contains(Tokenizertypes.FullSolution(&candidate), &tabuList)
				notTabu := !helpers.Contains(helpers.TokenizerChange(&candidate), &tabuList)

				if notTabu && (first || candidate.Fitness < fullBestCandidate.Fitness){
					fullBestCandidate = candidate
					first = false
				} else if (!notTabu && candidate.Fitness < fullBestSolution.Fitness ) { // Aspiração
					fullBestCandidate = candidate
					break
				}
			}
		} else {
			neighborhood = local_search.NeighborhoodBy2opt(&fullBestCandidate, distances)

			tmpBestCandidate := fullBestCandidate

			if constants.MULTI_IMPROVEMENT {
				// TSP MULTIIMPROVEMENT
				var goodCandidates []types.FullSolution

				first := true
				for _, candidate := range neighborhood {
					// notTabu := !helpers.Contains(Tokenizertypes.FullSolution(&candidate), &tabuList)
					notTabu := !helpers.Contains(helpers.TokenizerChange(&candidate), &tabuList)

					if notTabu && (first || candidate.Fitness < tmpBestCandidate.Fitness ){
						goodCandidates = append(goodCandidates, candidate)
						tmpBestCandidate = candidate
						first = false
					} else if (!notTabu && candidate.Fitness < fullBestSolution.Fitness ) { // Aspiração
						goodCandidates = append(goodCandidates, candidate)
						tmpBestCandidate = candidate
					}
				}

				if len(goodCandidates) > 0 {
					if constants.MULTI_IMPROVEMENT_DEBUG {
						fmt.Println(len(goodCandidates))
					}
					sort.Slice(goodCandidates, func(a, b int) bool {
						fitA := goodCandidates[a].Fitness
						fitB := goodCandidates[b].Fitness
						return fitA < fitB
					})
					improvement := fullBestCandidate.Fitness - goodCandidates[0].Fitness
					if improvement > 0 {
						tmpGoodCandidates := goodCandidates
						goodCandidates :=  goodCandidates[:0]

						for _, goodCandidate := range tmpGoodCandidates {
							improvement := fullBestCandidate.Fitness - goodCandidate.Fitness
							if improvement > 0 {
								goodCandidates = append(goodCandidates, goodCandidate)
							}
						}
					} else {
						if constants.MULTI_IMPROVEMENT_DEBUG {
							fmt.Println("NÃO MELHOROU CARALHO")
							fmt.Println(goodCandidates)
						}

						goodCandidates = goodCandidates[:1]

						if constants.MULTI_IMPROVEMENT_DEBUG {
							fmt.Println(goodCandidates)
						}
					}

					if constants.MULTI_IMPROVEMENT_DEBUG {
						fmt.Println("GOOD CANDIDATES:")
					}
					for _, goodCandidate := range goodCandidates {
						improvement := fullBestCandidate.Fitness- goodCandidate.Fitness
						i, j := goodCandidate.I, goodCandidate.J
						if constants.MULTI_IMPROVEMENT_DEBUG {
							fmt.Println(i, j, improvement)
						}

						blocked := false

						for _, selected := range selecteds {
							s_i, s_j := selected.I, selected.J

							//conflict?
							// if !((i < s_i && j < s_i) || (i > s_j && j > s_j)) {
							i_minus := i - 1
							if i_minus < 1 {
								i_minus = constants.PROBLEM_SIZE
							}
							j_plus := j + 1
							if j_plus > constants.PROBLEM_SIZE {
								j_plus = 1
							}
			

							if !((i < s_i && j < s_i) || (i > s_j && j > s_j)) {
								blocked = true
								break
							}

							if i_minus == s_j || j_plus == s_i {
								blocked = true
								break	
							} 
						}

						if blocked == false {
							selecteds = append(selecteds, goodCandidate)
						}
					}
				}

				if len(selecteds) > 1 {
					if constants.MULTI_IMPROVEMENT_DEBUG {
						fmt.Println("SELECTEDS:")
					}
					for _, selected := range selecteds {
						improvement := fullBestCandidate.Fitness - selected.Fitness
						i, j := selected.I, selected.J
						if constants.MULTI_IMPROVEMENT_DEBUG {
							fmt.Println("I:", i, "J:", j, "DIFF:", improvement)
						}
					}

					//apply swaps
					if constants.MULTI_IMPROVEMENT_DEBUG {
						fmt.Println("PRE-SWAPS:")
						fmt.Println(fullBestCandidate)
					}
				}

				for _, selected := range selecteds {
					sn := fullBestCandidate.Solution
					i, j := selected.I, selected.J

					st := i
					end := j
					for st < end {
						sn[st], sn[end] = sn[end], sn[st]
						st++
						end--
					}

					s_before := fullBestCandidate.Solution
					fitness := fitness.Full2opt(&s_before, distances, fullBestCandidate.Fitness, i, j)
					fullBestCandidate = types.FullSolution{sn, fitness, i, j}
				}

				if constants.MULTI_IMPROVEMENT_DEBUG {
					if len(selecteds) > 1 {
						fmt.Println("POS-SWAPS:")
						fmt.Println(fullBestCandidate)
					}
				}
			} else {
				first := true
				for _, candidate := range neighborhood {
					// notTabu := !helpers.Contains(Tokenizertypes.FullSolution(&candidate), &tabuList)
					notTabu := !helpers.Contains(helpers.TokenizerChange(&candidate), &tabuList)

					if notTabu && (first || candidate.Fitness < fullBestCandidate.Fitness){
						fullBestCandidate = candidate
						first = false
					} else if (!notTabu && candidate.Fitness < fullBestSolution.Fitness ) { // Aspiração
						fullBestCandidate = candidate
						break
					}
				}
			}
		}


		if fullBestCandidate.Fitness< fullBestSolution.Fitness {
			fullBestSolution = fullBestCandidate
		}

		if constants.MULTI_IMPROVEMENT {
			for _, selected := range selecteds {
				tabuList = append(tabuList, helpers.TokenizerChange(&selected))
				if len(tabuList) > maxTabuSize {
					tabuList = tabuList[1:]
				}
			}
		} else {
			tabuList = append(tabuList, helpers.TokenizerChange(&fullBestCandidate))
			if len(tabuList) > maxTabuSize {
				tabuList = tabuList[1:]
			}
		}

		// if constants.MULTI_IMPROVEMENT_DEBUG {
		// 	if len(tabuList) > 10 {
		// 		fmt.Println(tabuList[:5], "...", tabuList[len(tabuList)-5:], len(tabuList))
		// 	} else if len(tabuList) > 5 {
		// 		fmt.Println(tabuList[len(tabuList)-5:], len(tabuList))
		// 	} else {
		// 		fmt.Println(tabuList, len(tabuList))
		// 	}
		// }

		if !gap50beat && float64(fullBestSolution.Fitness) <= gap50{
			fmt.Printf("GAP 50 in iteration %d\n", iteration)
			gap50beat = true
		}

		if !gap25beat && float64(fullBestSolution.Fitness) <= gap25{
			fmt.Printf("GAP 25 in iteration %d\n", iteration)
			gap25beat = true
		}

		if !gap10beat && float64(fullBestSolution.Fitness) <= gap10{
			fmt.Printf("GAP 10 in iteration %d\n", iteration)
			gap10beat = true
		}

		if !gap1beat && float64(fullBestSolution.Fitness) <= gap1{
			fmt.Printf("GAP 1 in iteration %d\n", iteration)
			gap1beat = true
		}

		if !gap0beat && float64(fullBestSolution.Fitness) <= gap0{
			fmt.Printf("OPTIMAL in iteration %d\n", iteration)
			gap0beat = true
		}


		iteration += 1
		// fmt.Println(fullBestSolution.Fitness, fullBestCandidate.Fitness)
	}

	return nil, fullBestSolution
}