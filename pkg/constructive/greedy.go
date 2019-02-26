package constructive

import (
	"github.com/danhenriquesc/go-probabilistic-tabusearch/pkg/constants"
	"github.com/danhenriquesc/go-probabilistic-tabusearch/pkg/types"
)

type GreedyConstructive struct {
	Distances *types.Distances	
}

func (gc GreedyConstructive) generateInitialSolution() types.Solution {
	var visited [constants.PROBLEM_SIZE + 1]bool
	initialSolution := types.Solution{}

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
			if gc.Distances[current_city][i] < min_distance && visited[i] == false {
				min_distance = gc.Distances[current_city][i]
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