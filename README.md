# Pokedex CLI

A small CLI app for exploring Pokémon location data with request instrumentation and caching.

## Overview

- `main.go` starts a REPL loop.
- `commands/` contains CLI command implementations.
- `requests/` handles HTTP requests, response parsing, caching, and CSV logging.
- `internal/pokecache/` stores request responses in memory.
- `benchmark/` contains reusable benchmark helpers.

## Core commands

- `help` - show available commands.
- `exit` - quit the REPL.
- `map` - list the current page of location areas.
- `mapb` - go back to the previous page of location areas.
- `explore <location-area>` - fetch encounters for a location area.
- `stats` - print runtime request performance statistics.
- `benchmark <iterations> <location-area>` - compare cold network request performance with cache hit performance.

## Instrumentation

- All request timing and logging is centralized in `requests/`.
- Requests are logged to `logs/requests.csv`.
- The CSV includes request metadata and timing for later analysis.
- Cache hits and misses are tracked separately.

## Notes

- `requests/request.go` is the main observability layer.
- `requests/stats.go` maintains in-memory statistics during runtime.
- `benchmark/benchmark.go` provides generic helpers for average, median, percentile, min, and max.

## Getting started

1. Run `go run .`
2. Use the REPL to execute commands.
3. Inspect `logs/requests.csv` for request timing data.
   
