// Copyright 2019 Brad Rydzewski. All rights reserved.
// Use of this source code is governed by the Polyform License
// that can be found in the LICENSE.md file.

package render

// pagelen calculates to total number of pages given the
// page size and total count of all paginated items.
func pagelen(size, total int) int {
	quotient, remainder := total/size, total%size
	switch {
	case quotient == 0:
		return 1
	case remainder == 0:
		return quotient
	default:
		return quotient + 1
	}
}

// max returns the larger of x or y.
func max(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}

// max returns the smaller of x or y.
func min(x, y int) int {
	if y == 0 {
		return x
	} else if x < y {
		return x
	} else {
		return y
	}
}
