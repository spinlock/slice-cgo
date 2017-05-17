// Copyright (c) 2017 spinlock@github.com
// All Rights Reserved.
//
// Licensed under the MIT (MIT-LICENSE.txt) license.

package unsafe2

import (
	"reflect"
	"runtime"
	"sync/atomic"
	"unsafe"
)

var allocOffheapBytes int64

func OffheapBytes() int64 {
	return allocOffheapBytes
}

type cgoSlice struct {
	ptr unsafe.Pointer
	buf []byte
}

func newCGoSlice(n int, force bool) Slice {
	after := atomic.AddInt64(&allocOffheapBytes, int64(n))
	if !force && after > MaxOffheapBytes() {
		atomic.AddInt64(&allocOffheapBytes, -int64(n))
		return nil
	}
	p := cgo_malloc(n)
	if p == nil {
		panic("cgo_malloc returns a nullptr")
	}
	s := &cgoSlice{
		ptr: p,
		buf: *(*[]byte)(unsafe.Pointer(&reflect.SliceHeader{
			Data: uintptr(p), Len: n, Cap: n,
		})),
	}
	runtime.SetFinalizer(s, (*cgoSlice).reclaim)
	return s
}

func (s *cgoSlice) Type() SliceType {
	return CGoSliceType
}

func (s *cgoSlice) Buffer() []byte {
	return s.buf
}

func (s *cgoSlice) reclaim() {
	if s.ptr == nil {
		return
	}
	cgo_free(s.ptr)
	atomic.AddInt64(&allocOffheapBytes, -int64(len(s.buf)))
	s.ptr = nil
	s.buf = nil
	runtime.SetFinalizer(s, nil)
}

func (s *cgoSlice) Slice2(beg, end int) Slice {
	return newGoSliceFrom(s, s.Buffer()[beg:end])
}

func (s *cgoSlice) Slice3(beg, end, cap int) Slice {
	return newGoSliceFrom(s, s.Buffer()[beg:end:cap])
}

func (s *cgoSlice) Parent() Slice {
	return nil
}
