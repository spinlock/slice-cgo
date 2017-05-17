.DEFAULT_GOAL := build-all

export GO15VENDOREXPERIMENT=1

build-all: godeps

godeps:
	@make --no-print-directory -C vendor/github.com/spinlock/jemalloc-go/

clean:

distclean: clean
	@make --no-print-directory --quiet -C vendor/github.com/spinlock/jemalloc-go/ distclean

gotest: codis-deps
	go test ./...

gobench: codis-deps
	go test -gcflags -l -bench=. -v ./...
