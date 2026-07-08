package requests

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	cache "pokedexcli/internal/pokecache"
)

var cacheInstance = cache.NewCache(5 * time.Minute) // 5 minutes in milliseconds

func GetLocationAreas(url string) (Response, error) {
	body, _, err := fetchResponseBody(url)
	if err != nil {
		return Response{}, err
	}

	var response Response
	err = json.Unmarshal(body, &response)
	if err != nil {
		return Response{}, fmt.Errorf("invalid JSON response from %s: %w", url, err)
	}

	return response, nil
}

func GetLocationArea(url string) (LocationAreaDetail, error) {
	body, _, err := fetchResponseBody(url)
	if err != nil {
		return LocationAreaDetail{}, err
	}

	var response LocationAreaDetail
	err = json.Unmarshal(body, &response)
	if err != nil {
		return LocationAreaDetail{}, fmt.Errorf("invalid JSON response from %s: %w", url, err)
	}

	return response, nil
}

func GetPokemon(url string) (Pokemon, error) {
	body, _, err := fetchResponseBody(url)
	if err != nil {
		return Pokemon{}, err
	}

	var response Pokemon
	err = json.Unmarshal(body, &response)
	if err != nil {
		return Pokemon{}, fmt.Errorf("invalid JSON response from %s: %w", url, err)
	}

	return response, nil
}

func ClearCache() {
	cacheInstance.Clear()
}

func fetchResponseBody(requestURL string) ([]byte, int, error) {
	start := time.Now()
	body, ok := cacheInstance.Get(requestURL)
	duration := time.Since(start)
	if ok {
		requestStats.LogRequest(requestURL, true, http.StatusOK, duration, len(body))
		return body, http.StatusOK, nil
	}

	start = time.Now()
	resp, err := http.Get(requestURL)
	if err != nil {
		duration = time.Since(start)
		requestStats.LogRequest(requestURL, false, 0, duration, 0)
		return nil, 0, err
	}
	defer resp.Body.Close()

	body, err = io.ReadAll(resp.Body)
	duration = time.Since(start)
	if err != nil {
		requestStats.LogRequest(requestURL, false, resp.StatusCode, duration, 0)
		return nil, resp.StatusCode, err
	}

	if resp.StatusCode != http.StatusOK {
		requestStats.LogRequest(requestURL, false, resp.StatusCode, duration, len(body))
		return nil, resp.StatusCode, fmt.Errorf("request failed: %s: %s", resp.Status, strings.TrimSpace(string(body)))
	}

	requestStats.LogRequest(requestURL, false, resp.StatusCode, duration, len(body))
	cacheInstance.Add(requestURL, body)
	return body, resp.StatusCode, nil
}
