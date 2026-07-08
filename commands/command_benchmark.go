package commands

import (
	"fmt"
	"strconv"

	"pokedexcli/benchmark"
	"pokedexcli/requests"
)

func Benchmark(cfg *Config) error {
	if len(cfg.Args) < 2 {
		return fmt.Errorf("usage: benchmark <iterations> <location-area>")
	}

	iterations, err := strconv.Atoi(cfg.Args[0])
	if err != nil || iterations <= 0 {
		return fmt.Errorf("iterations must be a positive integer")
	}

	location := cfg.Args[1]
	if location == "" {
		return fmt.Errorf("location-area is required")
	}

	requests.ClearCache()

	requestURL := fmt.Sprintf("https://pokeapi.co/api/v2/location-area/%s/", location)
	result, err := benchmark.RunBenchmark(iterations, func() error {
		_, err := requests.GetLocationArea(requestURL)
		return err
	})
	if err != nil {
		return err
	}

	fmt.Println("Benchmark Results")
	fmt.Println("=================")
	fmt.Println()
	fmt.Printf("Location Area: %s\n", location)
	fmt.Printf("Iterations: %d\n", iterations)
	fmt.Println()
	fmt.Printf("Network Requests: %d\n", result.NetworkRequests)
	fmt.Printf("Cache Hits: %d\n", result.CacheHits)
	fmt.Printf("Cache Misses: %d\n", result.CacheMisses)
	fmt.Println()
	fmt.Printf("Average Network Request: %.3f ms\n", result.AverageNetworkRequestMs)
	fmt.Printf("Average Cache Hit: %.3f ms\n", result.AverageCacheHitMs)
	fmt.Println()
	fmt.Printf("Fastest Request: %.3f ms\n", result.FastestMs)
	fmt.Printf("Slowest Request: %.3f ms\n", result.SlowestMs)
	fmt.Println()
	fmt.Printf("Median: %.3f ms\n", result.MedianMs)
	fmt.Printf("P95: %.3f ms\n", result.P95Ms)
	fmt.Println()
	fmt.Printf("Estimated Speedup: %.0fx\n", result.Speedup)

	return nil
}
