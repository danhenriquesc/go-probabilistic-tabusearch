package helpers

import (
	"fmt"
)

type gapCounter struct {
	optimal 	int
	gaps 		[]float64
	gapsStatus  map[float64]bool
}

func (gc gapCounter) CheckBeats(fitness float64, iteration int) {
	for _, gap := range gc.gaps {
		if !gc.gapsStatus[gap] && fitness <= float64(gc.optimal) * gap {
			fmt.Printf("GAP %.0f in iteration %d\n", gap * 100 - 100 ,iteration)
			gc.gapsStatus[gap] = true
		}
	}
}

func NewGapCounter(optimal int, gaps []float64) gapCounter {
	return gapCounter{optimal, gaps, map[float64]bool{}}
}
