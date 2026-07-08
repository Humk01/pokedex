package commands

import (
	"fmt"
	"math/rand/v2"
	"pokedexcli/requests"
	"strings"
)

var pokemon = make(map[string]struct{})

func Catch(name string, cfg *Config) error {
	normalized := strings.ToLower(name)
	fmt.Printf("Throwing a Pokeball at %s...\n", normalized)
	url := fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%s/", normalized)

	response, err := requests.GetPokemon(url)
	if err != nil {
		return err
	}

	chance := 100.0 / float64(response.BaseExperience)
	randomValue := rand.Float64()

	if randomValue <= chance {
		fmt.Printf("%s was caught!\n", normalized)
		pokemon[normalized] = struct{}{}
	} else {
		fmt.Printf("%s escaped!\n", normalized)
	}

	return nil
}

func callCatch(cfg *Config) error {
	if len(cfg.Args) == 0 {
		return fmt.Errorf("Please provide a Pokemon name to catch.")
	}
	return Catch(cfg.Args[0], cfg)
}

func Pokedex(cfg *Config) error {
	if len(pokemon) == 0 {
		fmt.Println("No Pokemon caught yet.")
		return nil
	}

	fmt.Println("Caught Pokemon:")
	for name := range pokemon {
		fmt.Println(name)
	}

	return nil
}

func callPokedex(cfg *Config) error {
	return Pokedex(cfg)
}
