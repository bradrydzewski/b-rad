// Copyright 2019 Brad Rydzewski. All rights reserved.
// Use of this source code is governed by the Polyform License
// that can be found in the LICENSE.md file.

package enum

import "testing"

func TestParseOrder(t *testing.T) {
	tests := []struct {
		text string
		want Order
	}{
		{"asc", OrderAsc},
		{"Asc", OrderAsc},
		{"ASC", OrderAsc},
		{"ascending", OrderAsc},
		{"Ascending", OrderAsc},
		{"desc", OrderDesc},
		{"Desc", OrderDesc},
		{"DESC", OrderDesc},
		{"descending", OrderDesc},
		{"Descending", OrderDesc},
		{"", OrderDefault},
		{"invalid", OrderDefault},
	}

	for _, test := range tests {
		got, want := ParseOrder(test.text), test.want
		if got != want {
			t.Errorf("Want order %q parsed as %q, got %q", test.text, want, got)
		}
	}
}
