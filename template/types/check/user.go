// Copyright 2019 Brad Rydzewski. All rights reserved.
// Use of this source code is governed by the Polyform License
// that can be found in the LICENSE.md file.

package check

import (
	"errors"

	"github.com/{{toLower repo}}/types"
)

var (
	// ErrEmailInvalid is returned when the email address
	// does not match a valid email patthern.
	ErrEmailInvalid = errors.New("Invalid email address")

	// ErrEmailThrowaway is returned when the email address
	// matches a throwaway provider.
	ErrEmailThrowaway = errors.New("Invalid email provider")

	// ErrEmailLen  is returned when the email address
	// exceeds the maximum number of characters.
	ErrEmailLen = errors.New("Email address cannot exceed 250 characters")
)

// User returns true if the User if valid.
func User(user *types.User) (bool, error) {
	if IsThrowaway(user.Email) {
		return false, ErrEmailThrowaway
	}
	if len(user.Email) > 250 {
		return false, Err{{title project}}IdentifierLen
	}
	return true, nil
}
