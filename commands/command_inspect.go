package commands

import (
	"fmt"
	"pokedexcli/requests"
	"strings"
)

func Inspect(name string, cfg *Config) error {
	normalized := strings.ToLower(name)
	if _, ok := pokemon[normalized]; !ok {
		return fmt.Errorf("%s has not been caught yet", normalized)
	}

	url := fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%s/", normalized)
	response, err := requests.GetPokemon(url)
	if err != nil {
		return err
	}

	fmt.Printf("Pokemon: %s\n", response.Name)
	fmt.Printf("Base Experience: %d\n", response.BaseExperience)
	fmt.Printf("Height: %d\n", response.Height)
	fmt.Printf("Weight: %d\n", response.Weight)
	fmt.Println("Types:")
	for _, t := range response.Types {
		fmt.Printf("- %s\n", t.Type.Name)
	}
	fmt.Println("Stats:")
	for _, s := range response.Stats {
		fmt.Printf("- %s: %d\n", s.Stat.Name, s.BaseStat)
	}

	return nil
}

func callInspect(cfg *Config) error {
	if len(cfg.Args) == 0 {
		return fmt.Errorf("Please provide a Pokemon name to inspect.")
	}
	return Inspect(cfg.Args[0], cfg)
}
