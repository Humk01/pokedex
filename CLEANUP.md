# Cleanup Report

The following files contain code that works but is worth cleaning up.

## commands/command_catch.go
- Uses a package-level `pokemon` map for captured state.
  - This is fine for the REPL, but it is global mutable state and not thread-safe.
  - It also makes the caught state hard to persist or reset cleanly.
- Normalizes names in one place, but `Inspect` duplicates normalization logic.
- The command name parsing is fine, but the command state could be structured more clearly.

## commands/command_inspect.go
- Works, but depends on the global `pokemon` map from `command_catch.go`.
- Duplicate URL/template logic with `command_catch.go`.
- Could be cleaned by extracting a shared Pokemon lookup / client helper.

## commands/command_map.go
- `Map` and `MapBack` contain very similar pagination logic.
- Duplicated printing and config updates could be extracted into a shared helper.
- The "first page" message is duplicated.

## commands/command_benchmark.go
- Contains a lot of repeated `fmt.Println()` calls for output formatting.
- The benchmark command is fine, but the output formatting could be factored.
- Could use a helper for `usage` / argument validation.

## commands/command_help.go
- Iterates `Commands()` map directly, which will print help entries in random order.
- If order matters, sort command names before printing.

## commands/command_exit.go
- `return nil` is unreachable after `os.Exit(0)`.
- This is harmless, but it is dead code and should be removed.

## requests/responseStruct.go
- Formatting is inconsistent around nested structs and indentation.
- `Stat` and `Type` contain anonymous nested structs, making fields harder to use.
- Consider extracting named types for the inner `Stat` and `Type` objects.

## requests/stats.go
- `requestStatsCollector` mixes logging, CSV write, and runtime stats all in one file.
- `loadNextRequestID()` parses the CSV line-by-line using a new `csv.Reader` for each line.
  - This works, but it is more complex than needed.
- `printTerminalLog()` is in the stats layer and is used for output formatting; if you want separation of concerns, this could be refactored.

## requests/request.go
- Works well, but the request layer is still responsible for some CLI formatting via logging.
- `fetchResponseBody()` returns raw bytes and HTTP status; good separation, but the CSV writer and terminal printing are tightly coupled in the same package.

## benchmark/benchmark.go
- The helper functions are good, but the benchmark package has no tests.
- The percentile implementation is generic and fine, but the benchmark logic assumes the first request is always a network request.
  - That is true given current usage, but it is an assumption worth documenting.

## General observations
- There are no obvious `TODO` or `FIXME` comments in the codebase.
- The code works, but several files carry duplicated logic or could use better separation of concerns.
- No unit tests were found for the command or benchmark code.

This report can be used as a starting point for refactoring and cleanup.
