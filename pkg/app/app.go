package app

import (
	"fmt"
	"time"
	"math/rand"
	"github.com/danhenriquesc/go-probabilistic-tabusearch/pkg/metaheuristics/tabu_search"
)

func Run() error {
	rand.Seed(time.Now().UTC().UnixNano())
	start := time.Now()

	err := tabu_search.Run()

	fmt.Println(time.Since(start))

	return err
}
