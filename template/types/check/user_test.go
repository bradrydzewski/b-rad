// Copyright 2019 Brad Rydzewski. All rights reserved.
// Use of this source code is governed by the Polyform License
// that can be found in the LICENSE.md file.

package check

import (
	"testing"

	"github.com/{{toLower repo}}/types"
)

func TestUser(t *testing.T) {
	tests := []struct {
		email string
		error error
		valid bool
	}{
		{
			email: "jane@gmail.com",
			valid: true,
		},
		{
			email: "jane@mailinater.com",
			error: ErrEmailThrowaway,
			valid: false,
		},
	}
	for _, test := range tests {
		user := &types.User{Email: test.email}
		ok, err := User(user)
		if got, want := ok, test.valid; got != want {
			t.Errorf("Want user %s is valid %v, got %v", test.email, want, got)
		}
		if got, want := err, test.error; got != want {
			t.Errorf("Want user %s error %v, got %v", test.email, want, got)
		}
	}
}
