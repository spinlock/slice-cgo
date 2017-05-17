// Copyright (c) 2017 spinlock@github.com
// All Rights Reserved.
//
// Licensed under the MIT (MIT-LICENSE.txt) license.

package unsafe2

import (
	"fmt"
	"sync/atomic"
)

type Slice interface {
	Type() SliceType

	Buffer() []byte
	reclaim()

	Slice2(beg, end int) Slice
	Slice3(beg, end, cap int) Slice
	Parent() Slice
}

var maxOffheapBytes int64

func MaxOffheapBytes() int64 {
	return maxOffheapBytes
}

func SetMaxOffheapBytes(n int64) {
	if n < 0 {
		panic(fmt.Sprintf("set maxOffheapBytes=%d", n))
	}
	atomic.StoreInt64(&maxOffheapBytes, n)
}

const MinOffheapSlice = 1024 * 16

func MakeSlice(n int) Slice {
	if n >= MinOffheapSlice {
		if s := newCGoSlice(n, false); s != nil {
			return s
		}
	}
	return newGoSlice(n)
}

func MakeOffheapSlice(n int) Slice {
	if n >= 0 {
		return newCGoSlice(n, true)
	}
	panic("make slice with negative size")
}

func FreeSlice(s Slice) {
	if s != nil {
		s.reclaim()
	}
}
