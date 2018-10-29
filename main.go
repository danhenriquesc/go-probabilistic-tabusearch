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

type City struct {
	x, y float64
}

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

func main() {
	start := time.Now()

	var cities [problemSize + 1]City
	var distances [problemSize + 1][problemSize + 1]float64

	ReadProblem(&cities)

	for i := 1; i <= problemSize; i++ {
		fmt.Println(cities[i])
	}

	CalculateDistances(&distances, &cities)

	// fmt.Println(distances)

	fmt.Println(time.Since(start))
}
