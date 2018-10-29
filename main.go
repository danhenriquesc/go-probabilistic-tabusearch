package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
	"math"
	"time"
)

const problemFile = "berlin52.tsp.txt"
const problemSize = 52
const architectureBits = 64
const maxTabuSize = 10

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

func CalculateDistance(a, b City, ch chan float64) {
	ch <- math.Sqrt(math.Pow(a.x - b.x, 2) + math.Pow(a.y - b.y, 2))
}

func CalculateDistances(distances *[problemSize + 1][problemSize + 1]float64, cities *[problemSize + 1]City) {
	ch := make(chan float64)
	for i := 1; i <= problemSize; i++ {
		for j := i + 1; j <= problemSize; j++ {
			go CalculateDistance(cities[i], cities[j], ch)

			distance := <- ch
			distances[i][j] = distance
			distances[j][i] = distance
		}
	}	
}

func NewInitialSolution() Solution {
	initialSolution := Solution{}
	for i := 1; i <= problemSize; i++ {
		initialSolution[i] = i
	}
	return initialSolution
}

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

func Fitness(s *Solution, distances *[problemSize + 1][problemSize + 1]float64) float64 {
	var fitness float64

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
	start := time.Now()

	var cities [problemSize + 1]City
	var distances [problemSize + 1][problemSize + 1]float64

	ReadProblem(&cities)
	// PrintCities(&cities)

	CalculateDistances(&distances, &cities)

	initialSolution := NewInitialSolution()
	bestSolution := initialSolution
	fitnessBestSolution := Fitness(&bestSolution, &distances)
	bestCandidate := initialSolution
	fitnessBestCandidate := fitnessBestSolution
	
	var tabuList []string
	tabuList = append(tabuList, TokenizerSolution(&initialSolution))
	// fmt.Println(tabuList)

	x := 1
	for x < 1000 {
		neighborhood := GetNeighborhood(&bestCandidate)

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
