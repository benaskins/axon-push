@AGENTS.md

## Conventions
- EventBus[T] is generic; callers define the event type
- Slow subscribers get dropped events with a warning (backpressure)
- SSE helpers are standalone functions, not tied to EventBus

## Constraints
- Stdlib only; no external dependencies
- Do not add persistence or durability; this is in-memory only
- Do not import any axon-* packages

## Testing
- `go test ./...` runs all tests
- `go vet ./...` for lint
