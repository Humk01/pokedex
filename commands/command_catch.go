package commands

import (
	"fmt"
	"math/rand/v2"
	"pokedexcli/requests"
)

var pokemon = make(map[string]struct{})

func Catch(name string, cfg *Config) error {
	fmt.Printf("Throwing a Pokeball at %s.../n", name)
	url := fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%s/", name)

	response, err := requests.GetPokemon(url)
	if err != nil {
		return err
	}

	chance := 100.0 / float64(response.BaseExperience)
	randomValue := rand.Float64() * 100.0

	if randomValue <= chance {
		fmt.Printf("You caught %s!\n", name)
		pokemon[name] = struct{}{}
	} else {
		fmt.Printf("%s escaped!\n", name)
	}

	return nil
}

func callCatch(cfg *Config) error {
	if len(cfg.Args) == 0 {
		return fmt.Errorf("Please provide a Pokemon name to catch.")
	}
	return Catch(cfg.Args[0], cfg)
}
