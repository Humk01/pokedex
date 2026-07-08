package requests

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"net/url"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"sync"
	"time"
)

type RequestStatistics struct {
	TotalRequests           int
	CacheHits               int
	CacheMisses             int
	HitRatio                float64
	AverageNetworkRequestMs float64
	AverageCacheHitMs       float64
	FastestMs               float64
	SlowestMs               float64
}

type requestStatsCollector struct {
	mu                   sync.Mutex
	totalRequests        int
	cacheHits            int
	cacheMisses          int
	totalNetworkDuration time.Duration
	totalCacheDuration   time.Duration
	fastest              time.Duration
	slowest              time.Duration
	hasDuration          bool
	nextRequestID        int
	csvPath              string
}

var requestStats = newRequestStatsCollector()

func newRequestStatsCollector() *requestStatsCollector {
	collector := &requestStatsCollector{
		csvPath:       filepath.Join("logs", "requests.csv"),
		nextRequestID: 1,
	}
	collector.loadNextRequestID()
	return collector
}

func (s *requestStatsCollector) loadNextRequestID() {
	file, err := os.Open(s.csvPath)
	if err != nil {
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	firstLine := true
	lastID := 0
	for scanner.Scan() {
		if firstLine {
			firstLine = false
			continue
		}
		line := scanner.Text()
		if line == "" {
			continue
		}
		recordReader := csv.NewReader(strings.NewReader(line))
		record, err := recordReader.Read()
		if err != nil || len(record) == 0 {
			continue
		}
		id, err := strconv.Atoi(record[0])
		if err != nil {
			continue
		}
		lastID = id
	}
	if lastID > 0 {
		s.nextRequestID = lastID + 1
	}
}

func (s *requestStatsCollector) LogRequest(requestURL string, cacheHit bool, statusCode int, duration time.Duration, responseSize int) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if cacheHit {
		s.cacheHits++
		s.totalCacheDuration += duration
	} else {
		s.cacheMisses++
		s.totalNetworkDuration += duration
	}

	s.totalRequests++

	if !s.hasDuration || duration < s.fastest {
		s.fastest = duration
	}
	if !s.hasDuration || duration > s.slowest {
		s.slowest = duration
	}
	s.hasDuration = true

	requestID := s.nextRequestID
	s.nextRequestID++

	endpoint := parseEndpoint(requestURL)
	timestamp := time.Now().UTC().Format(time.RFC3339)
	durationMs := fmt.Sprintf("%.3f", float64(duration.Microseconds())/1000.0)
	record := []string{
		strconv.Itoa(requestID),
		timestamp,
		"GET",
		endpoint,
		requestURL,
		fmt.Sprintf("%t", cacheHit),
		strconv.Itoa(statusCode),
		durationMs,
		strconv.Itoa(responseSize),
	}

	if err := s.appendLogRecord(record); err != nil {
		fmt.Printf("warning: could not write request log: %v\n", err)
	}

	printTerminalLog(cacheHit, endpoint, statusCode, durationMs, responseSize)
}

func parseEndpoint(requestURL string) string {
	parsed, err := url.Parse(requestURL)
	if err != nil {
		return requestURL
	}
	return parsed.Path
}

func (s *requestStatsCollector) appendLogRecord(record []string) error {
	if err := os.MkdirAll(filepath.Dir(s.csvPath), 0o755); err != nil {
		return err
	}

	file, err := os.OpenFile(s.csvPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0o644)
	if err != nil {
		return err
	}
	defer file.Close()

	info, err := file.Stat()
	if err != nil {
		return err
	}

	writer := csv.NewWriter(file)
	if info.Size() == 0 {
		header := []string{"request_id", "timestamp", "method", "endpoint", "url", "cache_hit", "status_code", "duration_ms", "response_size_bytes"}
		if err := writer.Write(header); err != nil {
			return err
		}
	}

	if err := writer.Write(record); err != nil {
		return err
	}
	writer.Flush()
	return writer.Error()
}

func printTerminalLog(cacheHit bool, endpoint string, statusCode int, durationMs string, responseSize int) {
	if cacheHit {
		fmt.Printf("[CACHE HIT] GET %s\n", endpoint)
	} else {
		fmt.Printf("[CACHE MISS] GET %s\n", endpoint)
		fmt.Printf("Status: %d\n", statusCode)
	}

	fmt.Printf("Duration: %s ms\n", durationMs)
	fmt.Printf("Size: %d bytes\n\n", responseSize)
}

func GetStatistics() RequestStatistics {
	requestStats.mu.Lock()
	defer requestStats.mu.Unlock()

	stats := RequestStatistics{
		TotalRequests: requestStats.totalRequests,
		CacheHits:     requestStats.cacheHits,
		CacheMisses:   requestStats.cacheMisses,
	}

	if stats.TotalRequests > 0 {
		stats.HitRatio = float64(stats.CacheHits) / float64(stats.TotalRequests) * 100.0
	}
	if requestStats.cacheMisses > 0 {
		stats.AverageNetworkRequestMs = float64(requestStats.totalNetworkDuration.Microseconds()) / 1000.0 / float64(requestStats.cacheMisses)
	}
	if requestStats.cacheHits > 0 {
		stats.AverageCacheHitMs = float64(requestStats.totalCacheDuration.Microseconds()) / 1000.0 / float64(requestStats.cacheHits)
	}
	if requestStats.hasDuration {
		stats.FastestMs = float64(requestStats.fastest.Microseconds()) / 1000.0
		stats.SlowestMs = float64(requestStats.slowest.Microseconds()) / 1000.0
	}

	return stats
}
