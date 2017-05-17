// Copyright (c) 2017 spinlock@github.com
// All Rights Reserved.
//
// Licensed under the MIT (MIT-LICENSE.txt) license.

package unsafe2

import (
	"testing"
)

func TestCastString(t *testing.T) {
	var b = []byte("hello")
	var s = CastString(b)
	b[0] = 'w'
	assertTrue(s == "wello")
}
