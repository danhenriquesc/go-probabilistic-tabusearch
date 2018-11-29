package reading

import (
	"os"
	"bufio"
	"strings"
	"strconv"
	"github.com/danhenriquesc/go-probabilistic-tabusearch/pkg/constants"
	"github.com/danhenriquesc/go-probabilistic-tabusearch/pkg/types"
)

func ReadProblem() [constants.PROBLEM_SIZE + 1]types.City {
	var cities [constants.PROBLEM_SIZE + 1]types.City

	file, err := os.Open(constants.PROBLEM_FILE)
	
	if err != nil {
		panic(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		words := strings.Fields(scanner.Text())
		
		index, err := strconv.ParseInt(words[0], 10, constants.ARCHITECTURE_BITS)
		if err != nil {
			panic(err)
		}
		x, err := strconv.ParseFloat(words[1], constants.ARCHITECTURE_BITS)
		if err != nil {
			panic(err)
		}
		y, err := strconv.ParseFloat(words[2], constants.ARCHITECTURE_BITS)
		if err != nil {
			panic(err)
		}
		
		cities[index] = types.NewCity(x, y)
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return cities
}