.DEFAULT_GOAL := build-all

export GO15VENDOREXPERIMENT=1

build-all: godeps

godeps:
	@make --no-print-directory -C vendor/github.com/spinlock/jemalloc-go/

clean:

distclean: clean
	@make --no-print-directory --quiet -C vendor/github.com/spinlock/jemalloc-go/ distclean

gotest: godeps
	go test ./...

gobench: godeps
	go test -gcflags -l -bench=. -v ./...
