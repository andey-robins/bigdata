# MapReduce

An exploration of MapReduce data processing

## Quick Reference

Easiest way to install go (that I've found): `brew install go`

Run the wordcount example: `go run main.go -wc`

An example of generating the rainbow table for the 10_000 most common passwords over md5, sha1, sha2, and sha3 is run using the command below.
```bash
$ go build
$ time ./mapreduce -rbow
./mapreduce -rbow  0.56s user 0.24s system 214% cpu 0.371 total
```

The corresponding discussion is in `writeup.pdf`.

## Dependencies

`go version go1.19.1 darwin/arm64`

This also depends on Glow, a golang MapReduce library. It should install automatically during the build process; however, if it doesn't:

```bash
$ go get github.com/chrislusf/glow
$ go get github.com/chrislusf/glow/flow
```

## Running

`go run main.go --help`

### Arguments

Arguments are presented and described in the help text.

Run `go run main.go --help`

### Example

`go run main.go -wc`

```
Words = 20459
```

## Building

`go build main.go`