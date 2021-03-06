package tabu_search

func loadConfig(instance string) map[string]int {
	var config map[string]int
	config = make(map[string]int)

	switch instance {
	case "berlin52.tsp.txt": // OPTIMAL: 7542 | BEAT: 7542
		config["maxTabuSize"] = 75
		config["iterations"] = 5000
		config["pertubation"] = 3
		config["optimal"] = 7542
	case "eil51.tsp.txt": // OPTIMAL: 426 | BEAT: 427
		config["maxTabuSize"] = 50
		config["iterations"] = 500
		config["pertubation"] = 3
		config["optimal"] = 426
	case "bays29.tsp.txt": // OPTIMAL: 2020 | BEAT: 9073
		config["maxTabuSize"] = 10
		config["iterations"] = 500
		config["pertubation"] = 3
		config["optimal"] = 2020
	case "att48.tsp.txt": // OPTIMAL: 33522 | BEAT: 33522
		config["maxTabuSize"] = 30
		config["iterations"] = 5000
		config["pertubation"] = 5
		config["optimal"] = 33522
	case "ch150.tsp.txt": // OPTIMAL: 6528 | BEAT: 6561
		config["maxTabuSize"] = 150
		config["iterations"] = 2500
		config["pertubation"] = 3
		config["optimal"] = 6528
	case "fl417.tsp.txt": // OPTIMAL: 11861 | BEAT: 12304
		config["maxTabuSize"] = 150
		config["iterations"] = 200
		config["pertubation"] = 3
		config["optimal"] = 11861		
	default:
		config["maxTabuSize"] = 50
		config["iterations"] = 700
		config["pertubation"] = 3
		config["optimal"] = 0
	}

	return config
}
