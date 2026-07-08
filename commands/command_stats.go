package commands

import (
	"fmt"
	"pokedexcli/requests"
)

func Stats(cfg *Config) error {
	stats := requests.GetStatistics()

	fmt.Println("Performance Statistics")
	fmt.Println()
	fmt.Printf("Total Requests: %d\n", stats.TotalRequests)
	fmt.Printf("Cache Hits: %d\n", stats.CacheHits)
	fmt.Printf("Cache Misses: %d\n", stats.CacheMisses)
	fmt.Printf("Hit Ratio: %.2f%%\n", stats.HitRatio)
	fmt.Println()
	fmt.Printf("Average Network Request: %.3f ms\n", stats.AverageNetworkRequestMs)
	fmt.Printf("Average Cache Hit: %.3f ms\n", stats.AverageCacheHitMs)
	fmt.Println()
	fmt.Printf("Fastest Request: %.3f ms\n", stats.FastestMs)
	fmt.Printf("Slowest Request: %.3f ms\n", stats.SlowestMs)

	return nil
}
