# Streaming

An exploration of streaming data processing

## Quick Reference

Easiest way to install go (that I've found): `brew install go`

Run the example from class: `go run main.go -class`

Run a big example: `ggo run main.go -seed 11 -moments 1000 -in data.csv `
    - Ommit the `-seed 11` for random results, keep and change the seed value for deterministic behavior

## Dependencies

`go version go1.19.1 darwin/arm64`

## Running

`go run main.go --help`

### Arguments

Arguments are presented and described in the help text.

Run `go run main.go --help`

### Example

`go run main.go -seed 11 -moments 1000 -in data.csv `

Provides an example output of:

```bash
Stream cardinality   = 350079
S cardinality        = 55607
True surprise number = 23737517
Estimated surprise   = 77838178
```

## RNG Information

All examples are generated using the golang rand package seeded either with Unix time or a specified seed value. Examples are computed on macOS Monterey 12.6, with arm architecture. Exact outputs may vary depending upon the system's handling of RNG.

## Building

`go build main.go`