// Copyright (c) 2017 spinlock@github.com
// All Rights Reserved.
//
// Licensed under the MIT (MIT-LICENSE.txt) license.

package unsafe2

import (
	"testing"
)

func assertTrue(b bool) {
	if !b {
		panic("assertion failed")
	}
}

func TestMakeGoSlice(t *testing.T) {
	n := MinOffheapSlice - 1
	s := MakeSlice(n)
	assertTrue(s != nil)
	_, ok := s.(*goSlice)
	assertTrue(ok)

	c1 := s.Slice2(0, 1)
	c2 := s.Slice2(0, 2)
	assertTrue(c1.Parent() == nil)
	assertTrue(c2.Parent() == nil)
}

func TestMakeCGoSlice(t *testing.T) {
	n := MinOffheapSlice * 2
	SetMaxOffheapBytes(int64(n) * 2)

	s1 := MakeSlice(n)
	_, ok1 := s1.(*cgoSlice)
	assertTrue(ok1 && len(s1.Buffer()) == n)
	defer FreeSlice(s1)

	s2 := MakeSlice(n)
	_, ok2 := s2.(*cgoSlice)
	assertTrue(ok2 && len(s2.Buffer()) == n)
	defer FreeSlice(s2)

	assertTrue(OffheapBytes() == int64(n)*2)

	s3 := MakeSlice(n)
	_, ok3 := s3.(*goSlice)
	assertTrue(ok3 && len(s3.Buffer()) == n)
	defer FreeSlice(s3)

	assertTrue(OffheapBytes() == int64(n)*2)

	FreeSlice(s2)
	assertTrue(OffheapBytes() == int64(n))

	s4 := MakeSlice(n)
	_, ok4 := s4.(*cgoSlice)
	assertTrue(ok4 && len(s4.Buffer()) == n)
	defer FreeSlice(s4)

	assertTrue(OffheapBytes() == int64(n)*2)

	s5 := MakeOffheapSlice(n)
	assertTrue(s5 != nil && len(s5.Buffer()) == n)
	defer FreeSlice(s5)

	assertTrue(OffheapBytes() == int64(n)*3)

	c1 := s5.Slice2(0, 2)
	c2 := c1.Slice3(0, 1, 2)
	assertTrue(c1.Parent() == s5)
	assertTrue(c2.Parent() == s5)

	c1.Buffer()[0] = 'a'
	assertTrue(c2.Buffer()[0] == 'a')

	c1.Buffer()[0] = 'b'
	assertTrue(c2.Buffer()[0] == 'b')
}

func TestMake0CGoSlice(t *testing.T) {
	p := MakeOffheapSlice(0)
	_, ok := p.(*cgoSlice)
	assertTrue(ok && len(p.Buffer()) == 0)
	defer FreeSlice(p)
}
