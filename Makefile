.DEFAULT_GOAL := all

export GO15VENDOREXPERIMENT=1

all:

# build_tags=
build_tags=-tags "cgo_jemalloc"

godeps:
	@make --no-print-directory -C vendor/github.com/spinlock/jemalloc-go/

install: godeps
	go install ${build_tags} ./unsafe2

gotest: godeps
	go test ${build_tags} ./unsafe2

clean:

distclean: clean
	@make --no-print-directory --quiet -C vendor/github.com/spinlock/jemalloc-go/ distclean

