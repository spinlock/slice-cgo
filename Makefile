.DEFAULT_GOAL := build-all

export GO15VENDOREXPERIMENT=1

build-all: godeps

godeps:
	@make --no-print-directory -C vendor/github.com/spinlock/jemalloc-go/

install: godeps
	go install -tags "cgo_jemalloc" ./unsafe2

clean:

distclean: clean
	@make --no-print-directory --quiet -C vendor/github.com/spinlock/jemalloc-go/ distclean

gotest: godeps
	go test -tags "cgo_jemalloc" ./unsafe2
