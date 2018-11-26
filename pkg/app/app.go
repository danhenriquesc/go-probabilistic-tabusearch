package app

import (
	"fmt"
	"math"
	"time"
	"math/rand"
	"github.com/danhenriquesc/go-probabilistic-tabusearch/pkg/constants"
	"github.com/danhenriquesc/go-probabilistic-tabusearch/pkg/reading"
)

const maxTabuSize = 50
const iterations = 700
const pertubation = 3

type City struct {
	x, y float64
}

type Solution [constants.PROBLEM_SIZE + 1]int

type FullSolution struct {
	solution Solution
	fitness int
	i, j int
}

func PrintCities(cities *[constants.PROBLEM_SIZE + 1]City) {
	for i := 1; i <= constants.PROBLEM_SIZE; i++ {
		fmt.Printf("City %d: %.2f x %.2f\n", i, cities[i].x, cities[i].y)
	}
}




/* ORDERED INITIAL */
func NewInitialSolution() Solution {
	initialSolution := Solution{}
	for i := 1; i <= constants.PROBLEM_SIZE; i++ {
		initialSolution[i] = i
	}
	return initialSolution
}

/* RAND INITIAL */
func NewRandomInitialSolution() Solution {
	initialSolution := Solution{}
	rands := rand.Perm(constants.PROBLEM_SIZE)

	for i := 1; i <= constants.PROBLEM_SIZE; i++ {
		initialSolution[i] = rands[i-1] + 1
	}

	return initialSolution
}

/* GREEDY INITIAL */
func NewGreedyInitialSolution(distances *[constants.PROBLEM_SIZE + 1][constants.PROBLEM_SIZE + 1]int) Solution {
	var visited [constants.PROBLEM_SIZE + 1]bool
	initialSolution := Solution{}

	current_city := 1
	current_index := 1
	initialSolution[current_index] = current_city
	visited[current_city] = true
	current_index++

	all_visited := false

	for all_visited != true {
		min_distance := 1000000000
		min_distance_city := -1
		for i := 1; i <= constants.PROBLEM_SIZE; i++ {
			if distances[current_city][i] < min_distance && visited[i] == false {
				min_distance = distances[current_city][i]
				min_distance_city = i
			}
		}

		current_city = min_distance_city
		initialSolution[current_index] = current_city
		visited[current_city] = true
		current_index++

		all_visited = true
		for i := 1; i <= constants.PROBLEM_SIZE; i++ {
			all_visited = all_visited && visited[i]
		}
	}

	return initialSolution
}

/* ALL */
func GetNeighborhood(s *Solution) []Solution {
	var neighbors []Solution

	for i := 1; i <= constants.PROBLEM_SIZE; i++ {
		for j := i + 1; j <= constants.PROBLEM_SIZE; j++ {
			sn := *s
			sn[i], sn[j] = sn[j], sn[i]
			neighbors = append(neighbors, sn)
		}
	}

	return neighbors
}

func GetFullNeighborhood(s *FullSolution, distances *[constants.PROBLEM_SIZE + 1][constants.PROBLEM_SIZE + 1]int) []FullSolution {
	var neighbors []FullSolution

	for i := 1; i <= constants.PROBLEM_SIZE; i++ {
		for j := i + 1; j <= constants.PROBLEM_SIZE; j++ {
			sn := s.solution
			sn[i], sn[j] = sn[j], sn[i]
			fs := FullSolution{sn, FullFitness(&s.solution, distances, s.fitness, i, j), i, j}
			// fs := FullSolution{sn, Fitness(&sn, distances)}
			neighbors = append(neighbors, fs)
		}
	}

	return neighbors
}


/* 2-OPT */
func GetNeighborhood2opt(s *Solution) []Solution {
	var neighbors []Solution

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

func GetFullNeighborhood2opt(s *FullSolution, distances *[constants.PROBLEM_SIZE + 1][constants.PROBLEM_SIZE + 1]int) []FullSolution {
	var neighbors []FullSolution

	for i := 1; i <= constants.PROBLEM_SIZE; i++ {
		for j := i + 2; j <= constants.PROBLEM_SIZE; j++ {
			sn := s.solution

			st := i
			end := j
			for st < end {
				sn[st], sn[end] = sn[end], sn[st]
				st++
				end--
			}

			fs := FullSolution{sn, FullFitness2opt(&s.solution, distances, s.fitness, i, j), i, j}
			// fs := FullSolution{sn, Fitness(&sn, distances)}
			neighbors = append(neighbors, fs)
		}
	}

	return neighbors
}

/* PERTURB */
func PerturbNeighborhood(s *Solution) {
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

func GetNeighborhoodInit() []Solution {
	var neighbors []Solution

	return neighbors
}

func GetNeighborhoodFullInit() []FullSolution {
	var neighbors []FullSolution

	return neighbors
}

func Fitness(s *Solution, distances *[constants.PROBLEM_SIZE + 1][constants.PROBLEM_SIZE + 1]int) int {
	var fitness int
	
	for i := 1; i < constants.PROBLEM_SIZE; i++ {
		fitness += distances[s[i]][s[i+1]]
	}

	fitness += distances[s[constants.PROBLEM_SIZE]][s[1]]

	return fitness
}

func FullFitness(s *Solution, distances *[constants.PROBLEM_SIZE + 1][constants.PROBLEM_SIZE + 1]int, current_fitness, i, j int) int {
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

func FullFitness2opt(s *Solution, distances *[constants.PROBLEM_SIZE + 1][constants.PROBLEM_SIZE + 1]int, current_fitness, i, j int) int {
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

func TokenizerSolution(s *Solution) string {
	return fmt.Sprint(*s)
}

func TokenizerFullSolution(fs *FullSolution) string {
	return fmt.Sprint(*fs)
}

func TokenizerChange(fs *FullSolution) string {
	return fmt.Sprint(fs.i,"|",fs.j)
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

	initialSolution := NewGreedyInitialSolution(&distances)
	// initialSolution := NewRandomInitialSolution()
	fitnessInitialSolution := Fitness(&initialSolution, &distances)
	fullInitialSolution := FullSolution{initialSolution, fitnessInitialSolution, 0, 0}

	fmt.Println(initialSolution)

	// bestSolution := initialSolution
	// fitnessBestSolution := Fitness(&bestSolution, &distances)
	fullBestSolution := fullInitialSolution

	fmt.Println(fullBestSolution)
	// bestCandidate := initialSolution
	// fitnessBestCandidate := fitnessBestSolution
	fullBestCandidate := fullInitialSolution
	
	var tabuList []string
	// tabuList = append(tabuList, TokenizerFullSolution(&fullBestCandidate))

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
			// notTabu := !Contains(TokenizerFullSolution(&candidate), &tabuList)
			notTabu := !Contains(TokenizerChange(&candidate), &tabuList)

			if notTabu && (first || candidate.fitness < fullBestCandidate.fitness) {
				fullBestCandidate = candidate
				first = false
			}
		}

		if fullBestCandidate.fitness < fullBestSolution.fitness {
			fullBestSolution = fullBestCandidate
		}

		tabuList = append(tabuList, TokenizerChange(&fullBestCandidate))
		if len(tabuList) > maxTabuSize {
			tabuList = tabuList[1:]
		}

		// fmt.Println(fitnessBestSolution)

		x += 1
		fmt.Println(fullBestSolution.fitness, fullBestCandidate.fitness)
	}

	fmt.Println(fullBestSolution)
	// fmt.Println(fitnessBestSolution)

	fmt.Println(time.Since(start))

	return nil
}
