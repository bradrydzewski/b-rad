// Copyright 2019 Brad Rydzewski. All rights reserved.
// Use of this source code is governed by the Polyform License
// that can be found in the LICENSE.md file.

package enum

import "testing"

func TestParseUserAttr(t *testing.T) {
	tests := []struct {
		text string
		want UserAttr
	}{
		{"id", UserAttrId},
		{"name", UserAttrName},
		{"email", UserAttrEmail},
		{"created", UserAttrCreated},
		{"updated", UserAttrUpdated},
		{"admin", UserAttrAdmin},
		{"", UserAttrNone},
		{"invalid", UserAttrNone},
	}

	for _, test := range tests {
		got, want := ParseUserAttr(test.text), test.want
		if got != want {
			t.Errorf("Want user attribute %q parsed as %q, got %q", test.text, want, got)
		}
	}
}
