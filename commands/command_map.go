package commands

import (
	"fmt"

	"pokedexcli/requests"
)

func Map(cfg *Config) error {
	url := "https://pokeapi.co/api/v2/location-area/"
	if cfg != nil && cfg.Next != "" {
		url = cfg.Next
	}

	value, err := requests.GetLocationAreas(url)
	if err != nil {
		return err
	}

	if cfg != nil {
		cfg.Next = value.Next
		cfg.Previous = value.Previous
	}

	for _, location := range value.Results {
		fmt.Println(location.Name)
	}
	if cfg != nil && cfg.Previous == "" {
		fmt.Println("You are on the first page of results.")
	}
	return nil
}

func MapBack(cfg *Config) error {
	if cfg == nil || cfg.Previous == "" {
		fmt.Println("No previous page available.")
		return nil
	}

	value, err := requests.GetLocationAreas(cfg.Previous)
	if err != nil {
		return err
	}

	if cfg != nil {
		cfg.Next = value.Next
		cfg.Previous = value.Previous
	}

	for _, location := range value.Results {
		fmt.Println(location.Name)
	}
	if cfg != nil && cfg.Previous == "" {
		fmt.Println("You are on the first page of results.")
	}
	return nil
}
