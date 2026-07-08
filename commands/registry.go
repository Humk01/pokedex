package commands

type Config struct {
	Next     string
	Previous string
	Args     []string
}

type CliCommand struct {
	Name        string
	Description string
	Callback    func(*Config) error
}

type CommandMap map[string]CliCommand

func Commands() CommandMap {
	return CommandMap{
		"exit": {
			Name:        "exit",
			Description: "Exit the Pokedex",
			Callback:    Exit,
		},
		"help": {
			Name:        "help",
			Description: "Show this help message",
			Callback:    Help,
		},
		"map": {
			Name:        "map",
			Description: "Show the map of the Pokemon world",
			Callback:    Map,
		},
		"mapb": {
			Name:        "mapb",
			Description: "Show the previous page of the map",
			Callback:    MapBack,
		},
		"explore": {
			Name:        "explore",
			Description: "Explore a specific location in the Pokemon world",
			Callback:    callExplore,
		},
		"benchmark": {
			Name:        "benchmark",
			Description: "Benchmark cached and network request performance",
			Callback:    Benchmark,
		},
		"stats": {
			Name:        "stats",
			Description: "Show request performance statistics",
			Callback:    Stats,
		},
		"catch": {
			Name:        "catch",
			Description: "Catch a Pokemon by name",
			Callback:    callCatch,
		},
	}
}
