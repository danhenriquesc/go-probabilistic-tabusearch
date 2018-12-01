package tabu_search

func load_config(instance string) map[string]int {
	var config map[string]int
	config = make(map[string]int)

	switch instance {
	case "berlin52.tsp.txt":
		config["maxTabuSize"] = 50
		config["iterations"] = 700
		config["pertubation"] = 3
	}

	return config
}