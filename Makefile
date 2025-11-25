.PHONY: bench bench-short

go-bench:
	go test -bench=. -benchmem ./gobench

go-bench-short:
	go test -bench=. -benchmem -benchtime=2x ./gobench

js-bench:
	node --expose-gc src/bench.js	
