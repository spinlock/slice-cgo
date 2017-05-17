// Copyright (c) 2017 spinlock@github.com
// All Rights Reserved.
//
// Licensed under the MIT (MIT-LICENSE.txt) license.

// +build linux

package unsafe2

import (
	"unsafe"

	jemalloc "github.com/spinlock/jemalloc-go"
)

func cgo_malloc(n int) unsafe.Pointer {
	return jemalloc.Malloc(n)
}

func cgo_free(ptr unsafe.Pointer) {
	jemalloc.Free(ptr)
}
