package main

import (
	"fmt"
	"os"
	"github.com/danhenriquesc/go-probabilistic-tabusearch/pkg/app"
)

func main() {
	if err := app.Run("TABU_SEARCH"); err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
		os.Exit(1)
	}
}
