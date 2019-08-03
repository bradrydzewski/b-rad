// Copyright 2019 Brad Rydzewski. All rights reserved.
// Use of this source code is governed by the Polyform License
// that can be found in the LICENSE.md file.

package database

// default query range limit.
const defaultLimit = 100

// limit returns the page size to a sql limit.
func limit(size int) int {
	if size == 0 {
		size = defaultLimit
	}
	return size
}

// offset converts the page to a sql offset.
func offset(page, size int) int {
	if page == 0 {
		page = 1
	}
	if size == 0 {
		size = defaultLimit
	}
	page = page - 1
	return page * size
}
