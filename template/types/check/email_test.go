// Copyright 2019 Brad Rydzewski. All rights reserved.
// Use of this source code is governed by the Polyform License
// that can be found in the LICENSE.md file.

package check

import "testing"

func TestIsThrowaway(t *testing.T) {
	tests := []struct {
		email string
		value bool
	}{
		{
			email: "jane@gmail.com",
			value: false,
		},
		{
			email: "jane@mailinater.com",
			value: true,
		},
	}
	for _, test := range tests {
		got, want := IsThrowaway(test.email), test.value
		if got != want {
			t.Errorf("Want %s is throwaway %v, got %v", test.email, want, got)
		}
	}
}

func TestIsPublic(t *testing.T) {
	tests := []struct {
		email string
		value bool
	}{
		{
			email: "jane@gmail.com",
			value: true,
		},
		{
			email: "jane@hotmail.com",
			value: true,
		},
		{
			email: "jane@outlook.com",
			value: true,
		},
		{
			email: "jane@acme.com",
			value: false,
		},
	}
	for _, test := range tests {
		got, want := IsPublic(test.email), test.value
		if got != want {
			t.Errorf("Want %s is public %v, got %v", test.email, want, got)
		}
	}
}
