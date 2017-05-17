// Copyright (c) 2017 spinlock@github.com
// All Rights Reserved.
//
// Licensed under the MIT (MIT-LICENSE.txt) license.

package unsafe2

type SliceType int

const (
	GoSliceType SliceType = iota
	CGoSliceType
)

func (t SliceType) String() string {
	switch t {
	case GoSliceType:
		return "go_slice"
	case CGoSliceType:
		return "cgo_slice"
	}
	panic("unknown type")
}
