# toontest

`toontest` is a Go benchmarking project from [piraz](https://github.com/piraz)
designed to compare serialization and deserialization performance across
different formats: Toon, Gotoon, JSON, and Protobuf.

## What does it do?

This repository provides microbenchmarks for marshaling (serializing) and
unmarshaling (deserializing) a moderate-sized payload. It evaluates:
- [Toon](https://github.com/toon-format/toon-go)
- [Gotoon](https://github.com/someuser/gotoon) (if installed)
- `encoding/json` (standard Go library)
- [Protocol Buffers](https://developers.google.com/protocol-buffers)

The library benchmarks execution time, memory usage, and number of heap
allocations for each encoder/decoder.

## How to Run

First, clone this repository:

```sh
git clone https://github.com/piraz/toontest.git
cd toontest
```

Make sure you have your Go environment set up and dependencies installed.
You’ll also need `protoc` and its Go plugins if you want to regenerate protobuf
files.

### Benchmark

To run the full suite of benchmarks for all serialization formats:

```sh
make bench
```

This will output a table showing how many times each benchmark ran, the average
time per operation, memory used, and heap allocations.

### Quick (Short) Benchmark

For a faster, less statistically deep run (useful for development feedback):

```sh
make bench-short
```
This runs each benchmark only twice (see `-benchtime=2x` in the Makefile), so
results will be rough but quick.

## Adding New Benchmarks

Benchmarks are implemented in standard Go `*_test.go` files under `cmd/tut/`.  
To add another format, simply create a new `Benchmark...` function.

## Example Output

```
BenchmarkToonMarshal-12        8    142825750 ns/op   48315866 B/op    999951 allocs/op
BenchmarkGotoonMarshal-12      4    318810069 ns/op   72130438 B/op   1099957 allocs/op
BenchmarkJsonMarshal-12       49     24299468 ns/op    5299191 B/op         4 allocs/op
BenchmarkProtoMarshal-12     100     13222923 ns/op    2654362 B/op         2 allocs/op
```

## Requirements

- Go 1.20 or higher
- (For Protobuf) `protoc` and [protoc-gen-go](https://pkg.go.dev/mod/google.golang.org/protobuf/cmd/protoc-gen-go)

## License

MIT
