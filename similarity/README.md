# Sentence Similarity

## Quick Reference

Easiest way to install go (that I've found): `brew install go`

Run the case with the 100K sentence file: `go run main.go -in 100K.txt -k 0`

## Dependencies

`go version go1.19.3 darwin/arm64`

## Running

`go run main.go --help`

### Arguments

Arguments are presented and described in the help text.

Run `go run main.go --help`

## Data

|File|Distance 0 Time|Distance 0 Count|Distance 1 Time|Distance 1 Count| -size |
|----|---------------|----------------|---------------|----------------|-------|
|tiny.txt|81.042µs|2|723.167µs|0|100|
|small.txt|101.916µs|1|8.733083ms|0|100|
|100.txt|251.791µs|2|24.666458ms|0|100|
|1k.txt|2.074583ms|34|137.512375ms|0|1000|
|10k.txt|20.184375ms|488|1.405221042s|25|10_000|
|100k.txt|175.108875ms|7124|21.876438208s|1186|100_000|
|1M.txt|1.83640875s|79902|9m25.496216417s|38098|1_000_000|
|5M.txt|9.21365425s|516778|1h57m52.556996041s|353343|5_000_000|
|25M.txt|46.974499541s|4432935|12h53m51.515688709s|1976565|25_000_000|
