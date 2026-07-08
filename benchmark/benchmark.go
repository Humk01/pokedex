package benchmark

import (
	"fmt"
	"sort"
	"time"
)

type BenchmarkResult struct {
	Iterations              int
	NetworkRequests         int
	CacheHits               int
	CacheMisses             int
	AverageRequestMs        float64
	AverageNetworkRequestMs float64
	AverageCacheHitMs       float64
	FastestMs               float64
	SlowestMs               float64
	MedianMs                float64
	P95Ms                   float64
	Speedup                 float64
}

type RequestFunc func() error

func RunBenchmark(iterations int, request RequestFunc) (BenchmarkResult, error) {
	if iterations <= 0 {
		return BenchmarkResult{}, fmt.Errorf("iterations must be greater than zero")
	}

	durations := make([]float64, 0, iterations)
	networkDurations := make([]float64, 0, 1)
	cacheDurations := make([]float64, 0, iterations-1)

	for i := 0; i < iterations; i++ {
		start := time.Now()
		if err := request(); err != nil {
			return BenchmarkResult{}, err
		}
		duration := time.Since(start)
		dms := float64(duration.Microseconds()) / 1000.0
		durations = append(durations, dms)

		if i == 0 {
			networkDurations = append(networkDurations, dms)
			continue
		}
		cacheDurations = append(cacheDurations, dms)
	}

	result := BenchmarkResult{
		Iterations:              iterations,
		NetworkRequests:         len(networkDurations),
		CacheHits:               len(cacheDurations),
		CacheMisses:             len(networkDurations),
		AverageRequestMs:        Average(durations),
		AverageNetworkRequestMs: Average(networkDurations),
		AverageCacheHitMs:       Average(cacheDurations),
		FastestMs:               Min(durations),
		SlowestMs:               Max(durations),
		MedianMs:                Median(durations),
		P95Ms:                   Percentile(durations, 95),
	}

	if result.AverageCacheHitMs > 0 {
		result.Speedup = result.AverageNetworkRequestMs / result.AverageCacheHitMs
	}

	return result, nil
}

func Average(values []float64) float64 {
	if len(values) == 0 {
		return 0
	}
	var sum float64
	for _, value := range values {
		sum += value
	}
	return sum / float64(len(values))
}

func Median(values []float64) float64 {
	if len(values) == 0 {
		return 0
	}
	sorted := make([]float64, len(values))
	copy(sorted, values)
	sort.Float64s(sorted)
	mid := len(sorted) / 2
	if len(sorted)%2 == 0 {
		return (sorted[mid-1] + sorted[mid]) / 2
	}
	return sorted[mid]
}

func Percentile(values []float64, percentile float64) float64 {
	if len(values) == 0 {
		return 0
	}
	sorted := make([]float64, len(values))
	copy(sorted, values)
	sort.Float64s(sorted)

	if percentile <= 0 {
		return sorted[0]
	}
	if percentile >= 100 {
		return sorted[len(sorted)-1]
	}

	rank := percentile / 100 * float64(len(sorted)-1)
	lower := int(rank)
	upper := lower + 1
	if upper >= len(sorted) {
		return sorted[lower]
	}
	weight := rank - float64(lower)
	return sorted[lower] + weight*(sorted[upper]-sorted[lower])
}

func Min(values []float64) float64 {
	if len(values) == 0 {
		return 0
	}
	min := values[0]
	for _, value := range values[1:] {
		if value < min {
			min = value
		}
	}
	return min
}

func Max(values []float64) float64 {
	if len(values) == 0 {
		return 0
	}
	max := values[0]
	for _, value := range values[1:] {
		if value > max {
			max = value
		}
	}
	return max
}
