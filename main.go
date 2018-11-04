package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
	"math"
	"time"
	"math/rand"
)

const problemFile = "berlin52.tsp.txt"
const problemSize = 52
const architectureBits = 64
const maxTabuSize = 50
const iterations = 2000
const pertubation = 1

type City struct {
	x, y float64
}

type Solution [problemSize + 1]int

func Check(e error) {
	if e != nil {
		panic(e)
	}
}

func ReadProblem(cities *[problemSize + 1]City) {
	file, err := os.Open(problemFile)
	Check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		words := strings.Fields(scanner.Text())
		
		index, err := strconv.ParseInt(words[0], 10, architectureBits)
		Check(err)
		x, err := strconv.ParseFloat(words[1], architectureBits)
		Check(err)
		y, err := strconv.ParseFloat(words[2], architectureBits)
		Check(err)
		
		cities[index] = City{x, y}
	}

	Check(scanner.Err())
}

func PrintCities(cities *[problemSize + 1]City) {
	for i := 1; i <= problemSize; i++ {
		fmt.Printf("City %d: %.2f x %.2f\n", i, cities[i].x, cities[i].y)
	}
}

func Nint(n float64) int {
	return int(n + 0.5)
}

func CalculateDistance(a, b City, ch chan int) {
	ch <- Nint(math.Sqrt(math.Pow(a.x - b.x, 2) + math.Pow(a.y - b.y, 2)))
}

func CalculateDistances(distances *[problemSize + 1][problemSize + 1]int, cities *[problemSize + 1]City) {
	ch := make(chan int)
	for i := 1; i <= problemSize; i++ {
		for j := i + 1; j <= problemSize; j++ {
			go CalculateDistance(cities[i], cities[j], ch)

			distance := <- ch
			distances[i][j] = distance
			distances[j][i] = distance
		}
	}	
}

/* ORDERED INITIAL */
func NewInitialSolution() Solution {
	initialSolution := Solution{}
	for i := 1; i <= problemSize; i++ {
		initialSolution[i] = i
	}
	return initialSolution
}

/* RAND INITIAL */
func NewRandomInitialSolution() Solution {
	initialSolution := Solution{}
	rands := rand.Perm(problemSize)

	for i := 1; i <= problemSize; i++ {
		initialSolution[i] = rands[i-1] + 1
	}

	return initialSolution
}

/* GREEDY INITIAL */
func newGreedyInitialSolution(distances *[problemSize + 1][problemSize + 1]int) Solution {
	var visited [problemSize + 1]bool
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
		for i := 1; i <= problemSize; i++ {
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
		for i := 1; i <= problemSize; i++ {
			all_visited = all_visited && visited[i]
		}
	}

	return initialSolution
}

/* ALL */
func GetNeighborhood(s *Solution) []Solution {
	var neighbors []Solution

	for i := 1; i <= problemSize; i++ {
		for j := i + 1; j <= problemSize; j++ {
			sn := *s
			sn[i], sn[j] = sn[j], sn[i]
			neighbors = append(neighbors, sn)
		}
	}

	return neighbors
}

/* 2-OPT */
func GetNeighborhood2opt(s *Solution) []Solution {
	var neighbors []Solution

	for i := 1; i <= problemSize; i++ {
		for j := i + 1; j <= problemSize; j++ {
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

func GetNeighborhoodInit() []Solution {
	var neighbors []Solution

	return neighbors
}

func Fitness(s *Solution, distances *[problemSize + 1][problemSize + 1]int) int {
	var fitness int
	
	for i := 1; i < problemSize; i++ {
		fitness += distances[s[i]][s[i+1]]
	}

	fitness += distances[s[problemSize]][s[1]]

	return fitness
}

func TokenizerSolution(s *Solution) string {
	return fmt.Sprint(*s)
}

func Contains(needed string, tabuList *[]string) bool{
	for _, item := range *tabuList {
		if item == needed {
			return true
		}
	}
	return false
}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	start := time.Now()

	var cities [problemSize + 1]City
	var distances [problemSize + 1][problemSize + 1]int

	ReadProblem(&cities)
	// PrintCities(&cities)

	CalculateDistances(&distances, &cities)

	initialSolution := newGreedyInitialSolution(&distances)
	fmt.Println(initialSolution)

	bestSolution := initialSolution
	fitnessBestSolution := Fitness(&bestSolution, &distances)
	fmt.Println(fitnessBestSolution)
	bestCandidate := initialSolution
	fitnessBestCandidate := fitnessBestSolution
	
	var tabuList []string
	tabuList = append(tabuList, TokenizerSolution(&initialSolution))
	// fmt.Println(tabuList)

	x := 1
	for x < iterations {
		neighborhood := GetNeighborhoodInit()

		if x % pertubation == 0 {
			neighborhood = GetNeighborhood(&bestCandidate)
		} else {
			neighborhood = GetNeighborhood2opt(&bestCandidate)
		}

		for _, candidate := range neighborhood {
			fitnessCandidate := Fitness(&candidate, &distances)
			notTabu := !Contains(TokenizerSolution(&candidate), &tabuList)

			if notTabu && fitnessCandidate < fitnessBestCandidate {
				bestCandidate = candidate
				fitnessBestCandidate = fitnessCandidate
			}
		}

		if fitnessBestCandidate < fitnessBestSolution {
			bestSolution = bestCandidate
			fitnessBestSolution = fitnessBestCandidate
		}

		tabuList = append(tabuList, TokenizerSolution(&bestCandidate))
		if len(tabuList) > maxTabuSize {
			tabuList = tabuList[1:]
		}

		// fmt.Println(fitnessBestSolution)

		x += 1
	}

	fmt.Println(bestSolution)
	fmt.Println(fitnessBestSolution)

	fmt.Println(time.Since(start))
}
