package app

import (
	"fmt"
	"math"
	"time"
	"math/rand"
	"github.com/danhenriquesc/go-probabilistic-tabusearch/pkg/constants"
	"github.com/danhenriquesc/go-probabilistic-tabusearch/pkg/reading"
	"github.com/danhenriquesc/go-probabilistic-tabusearch/pkg/constructive"
	"github.com/danhenriquesc/go-probabilistic-tabusearch/pkg/types"
)

const maxTabuSize = 50
const iterations = 700
const pertubation = 3

/* ALL */
func GetNeighborhood(s *types.Solution) []types.Solution {
	var neighbors []types.Solution

	for i := 1; i <= constants.PROBLEM_SIZE; i++ {
		for j := i + 1; j <= constants.PROBLEM_SIZE; j++ {
			sn := *s
			sn[i], sn[j] = sn[j], sn[i]
			neighbors = append(neighbors, sn)
		}
	}

	return neighbors
}

func GetFullNeighborhood(s *types.FullSolution, distances *[constants.PROBLEM_SIZE + 1][constants.PROBLEM_SIZE + 1]int) []types.FullSolution {
	var neighbors []types.FullSolution

	for i := 1; i <= constants.PROBLEM_SIZE; i++ {
		for j := i + 1; j <= constants.PROBLEM_SIZE; j++ {
			sn := types.FullSolutionSolution(s)
			sn[i], sn[j] = sn[j], sn[i]

			s_before :=  types.FullSolutionSolution(s)
			fitness := FullFitness(&s_before, distances, types.FullSolutionFitness(s), i, j)
			fs := types.NewFullSolution(sn, fitness, i, j)

			neighbors = append(neighbors, fs)
		}
	}

	return neighbors
}


/* 2-OPT */
func GetNeighborhood2opt(s *types.Solution) []types.Solution {
	var neighbors []types.Solution

	for i := 1; i <= constants.PROBLEM_SIZE; i++ {
		for j := i + 1; j <= constants.PROBLEM_SIZE; j++ {
			sn := *s

			st := i
			end := j
			for st < end {
				sn[st], sn[end] = sn[end], sn[st]
				st++
				end--
			}

			neighbors = append(neighbors, sn)
		}
	}

	return neighbors
}

func GetFullNeighborhood2opt(s *types.FullSolution, distances *[constants.PROBLEM_SIZE + 1][constants.PROBLEM_SIZE + 1]int) []types.FullSolution {
	var neighbors []types.FullSolution

	for i := 1; i <= constants.PROBLEM_SIZE; i++ {
		for j := i + 2; j <= constants.PROBLEM_SIZE; j++ {
			sn := types.FullSolutionSolution(s)

			st := i
			end := j
			for st < end {
				sn[st], sn[end] = sn[end], sn[st]
				st++
				end--
			}

			s_before := types.FullSolutionSolution(s)
			fitness := FullFitness2opt(&s_before, distances, types.FullSolutionFitness(s), i, j)
			fs := types.NewFullSolution(sn, fitness, i, j)

			neighbors = append(neighbors, fs)
		}
	}

	return neighbors
}

/* PERTURB */
func PerturbNeighborhood(s *types.Solution) {
	r1 := (rand.Int() % constants.PROBLEM_SIZE) + 1
	r2 := (rand.Int() % constants.PROBLEM_SIZE) + 1

	st := int(math.Min(float64(r1), float64(r2)))
	end := int(math.Max(float64(r1), float64(r2)))

	for st < end {
		s[st], s[end] = s[end], s[st]
		st++
		end--
	}
}

func GetNeighborhoodInit() []types.Solution {
	var neighbors []types.Solution

	return neighbors
}

func GetNeighborhoodFullInit() []types.FullSolution {
	var neighbors []types.FullSolution

	return neighbors
}

func Fitness(s *types.Solution, distances *[constants.PROBLEM_SIZE + 1][constants.PROBLEM_SIZE + 1]int) int {
	var fitness int
	
	for i := 1; i < constants.PROBLEM_SIZE; i++ {
		fitness += distances[s[i]][s[i+1]]
	}

	fitness += distances[s[constants.PROBLEM_SIZE]][s[1]]

	return fitness
}

func FullFitness(s *types.Solution, distances *[constants.PROBLEM_SIZE + 1][constants.PROBLEM_SIZE + 1]int, current_fitness, i, j int) int {
	fitness := current_fitness
	
	if j - i != constants.PROBLEM_SIZE - 1 {
		if i > 1 {
			fitness -= distances[s[i-1]][s[i]]
			fitness += distances[s[i-1]][s[j]]
		} else {
			fitness -= distances[s[constants.PROBLEM_SIZE]][s[i]]
			fitness += distances[s[constants.PROBLEM_SIZE]][s[j]]
		}

		if j - i > 1 {
			fitness -= distances[s[i]][s[i+1]]
			fitness += distances[s[j]][s[i+1]]
			
			fitness -= distances[s[j-1]][s[j]]
			fitness += distances[s[j-1]][s[i]]
		}

		if j < constants.PROBLEM_SIZE {
			fitness -= distances[s[j]][s[j+1]]
			fitness += distances[s[i]][s[j+1]]
		} else {
			fitness -= distances[s[j]][s[1]]
			fitness += distances[s[i]][s[1]]
		}
	} else {
		fitness -= distances[s[j-1]][s[j]]
		fitness += distances[s[j-1]][s[1]]

		fitness -= distances[s[1]][s[2]]
		fitness += distances[s[j]][s[2]]
	}

	return fitness
}

func FullFitness2opt(s *types.Solution, distances *[constants.PROBLEM_SIZE + 1][constants.PROBLEM_SIZE + 1]int, current_fitness, i, j int) int {
	fitness := current_fitness
	
	if j - i != constants.PROBLEM_SIZE - 1 {
		if i > 1 {
			fitness -= distances[s[i-1]][s[i]]
			fitness += distances[s[i-1]][s[j]]
		} else {
			fitness -= distances[s[constants.PROBLEM_SIZE]][s[i]]
			fitness += distances[s[constants.PROBLEM_SIZE]][s[j]]
		}

		if j < constants.PROBLEM_SIZE {
			fitness -= distances[s[j]][s[j+1]]
			fitness += distances[s[i]][s[j+1]]
		} else {
			fitness -= distances[s[j]][s[1]]
			fitness += distances[s[i]][s[1]]
		}
	}
	
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
	// initialSolution := NewRandomInitialSolution()
	fitnessInitialSolution := Fitness(&initialSolution, &distances)
	fullInitialSolution := types.NewFullSolution(initialSolution, fitnessInitialSolution, 0, 0)

	fmt.Println(initialSolution)

	// bestSolution := initialSolution
	// fitnessBestSolution := Fitness(&bestSolution, &distances)
	fullBestSolution := fullInitialSolution

	fmt.Println(fullBestSolution)
	// bestCandidate := initialSolution
	// fitnessBestCandidate := fitnessBestSolution
	fullBestCandidate := fullInitialSolution
	
	var tabuList []string
	// tabuList = append(tabuList, Tokenizertypes.FullSolution(&fullBestCandidate))

	x := 1
	for x < iterations {
		neighborhood := GetNeighborhoodFullInit()

		// if x % 300 == 0 {
		// 	PerturbNeighborhood(&bestCandidate)
		// }
		if x % pertubation == 0 {
			neighborhood = GetFullNeighborhood(&fullBestCandidate, &distances)
		} else {
			neighborhood = GetFullNeighborhood2opt(&fullBestCandidate, &distances)
		}

		first := true
		for _, candidate := range neighborhood {
			// fitnessCandidate := Fitness(&candidate, &distances)
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

		// fmt.Println(fitnessBestSolution)

		x += 1
		fmt.Println(types.FullSolutionFitness(&fullBestSolution), types.FullSolutionFitness(&fullBestCandidate))
	}

	fmt.Println(fullBestSolution)
	// fmt.Println(fitnessBestSolution)

	fmt.Println(time.Since(start))

	return nil
}
