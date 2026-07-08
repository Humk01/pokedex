package commands

import (
	"fmt"
	"pokedexcli/requests"
)

func callExplore(cfg *Config) error {
	if len(cfg.Args) == 0 {
		return fmt.Errorf("Please provide a location to explore.")
	}
	return Explore(cfg.Args[0], cfg)
}

func Explore(location string, cfg *Config) error {
	fmt.Println("Exploring the Pokemon world...")
	url := fmt.Sprintf("https://pokeapi.co/api/v2/location-area/%s/", location)

	response, err := requests.GetLocationArea(url)
	if err != nil {
		return err
	}

	fmt.Println("Pokemon Encounters:")
	for _, encounter := range response.PokemonEncounters {
		fmt.Printf("- %s\n", encounter.Pokemon.Name)
	}

	return nil
}
