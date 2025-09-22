# Herman - Modular Blockchain (Prototype)

Herman is a compact, learning-oriented prototype of a modular blockchain. It demonstrates how to decouple a blockchain into four layers and iterate on each layer independently:

- Execution
- Settlement
- Consensus
- Data Availability

The codebase is intentionally concise and evolves through small, traceable increments to mimic real-world iterative development.

## Features

- Pluggable architecture across four core layers
- Multiple mock implementations to illustrate extension points
- Minimal CLI entrypoint for smoke testing
- Unit tests covering execution behavior samples

## Architecture

- Execution: defines `ExecutionEngine` interfaces and sample engines (EVM-style mock). Intended to host EVM/WASM execution paths in real systems.
- Settlement: defines finalization contracts and a mock prover to validate state roots or finality signals.
- Consensus: defines consensus interfaces and a pseudo-random leader elector example.
- Data Availability: defines DA interface and a simple in-memory store with sampling utilities.

## Repository Structure

```
cmd/herman/                # CLI entrypoint
internal/execution/        # Execution layer interfaces and examples
internal/settlement/       # Settlement layer interfaces and examples
internal/consensus/        # Consensus layer interfaces and examples
internal/dataavailability/ # Data availability layer interfaces and examples
scripts/                   # Automation and utility scripts
go.mod                     # Go module definition
```

## Quick Start

Prerequisites:

- Go 1.21+

Build and run CLI:

```
go build -o bin/herman ./cmd/herman
./bin/herman
```

Run tests:

```
go test ./...
```

## Development Guidelines

- Language: Go (prototype-level code)
- Style: prefer small, incremental, readable commits
- Commit messages: use conventional commits (e.g., `feat: ...`, `fix: ...`, `test: ...`, `docs: ...`)
- Modularity: add new implementations behind interfaces and keep cross-layer coupling minimal

## Extending the Prototype

- Execution
  - Add new engines implementing `ExecutionEngine` (e.g., a WASM mock)
  - Enrich transaction models and state transitions
- Settlement
  - Extend provers and finality checks
  - Integrate multi-signature verification flows (mocked)
- Consensus
  - Add algorithm variants (PoS/PoA flavors)
  - Simulate validator sets and block proposal logic
- Data Availability
  - Swap in alternative storage backends (memory, file, network mocks)
  - Experiment with sampling strategies

## Roadmap (Example)

1. Expand execution mocks to include WASM-style engine
2. Add basic cross-layer wiring in a minimal node loop
3. Introduce more tests for settlement/DA error paths
4. Provide simple config and metrics stubs
5. Document integration patterns and benchmarking guidelines

## Security

This is not production software. It contains mocked components and simplified logic, and it omits security hardening, networking, and persistence.

## License

MIT License. See `LICENSE` for details.

