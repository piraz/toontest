.PHONY: bench bench-short

bench:
	go test -bench=. -benchmem ./cmd/tut

bench-short:
	go test -bench=. -benchmem -benchtime=2x ./cmd/tut
