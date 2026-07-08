package requests

type LocationAreaDetail struct {
	Name              string             `json:"name"`
	PokemonEncounters []PokemonEncounter `json:"pokemon_encounters"`
}

type PokemonEncounter struct {
	Pokemon NamedAPIResource `json:"pokemon"`
}

type NamedAPIResource struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}
