package requests

type Response struct {
	Count    int            `json:"count"`
	Next     string         `json:"next"`
	Previous string         `json:"previous"`
	Results  []LocationArea `json:"results"`
}

type LocationArea struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type Pokemon struct {
    Name           string `json:"name"`
    BaseExperience int    `json:"base_experience"`
}
