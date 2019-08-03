// Copyright 2019 Brad Rydzewski. All rights reserved.
// Use of this source code is governed by the Polyform License
// that can be found in the LICENSE.md file.

package check

import (
	"errors"

	"github.com/{{toLower repo}}/types"

	"github.com/gosimple/slug"
)

var (
	// Err{{title child}}Identifier is returned when the {{toLower child}}
	// slug is an invalid format.
	Err{{title child}}Identifier = errors.New("Invalid {{toLower child}} identifier")

	// Err{{title child}}IdentifierLen is returned when the {{toLower child}}
	// name exceeds the maximum number of characters.
	Err{{title child}}IdentifierLen = errors.New("{{title child}} identifier cannot exceed 250 characters")

	// Err{{title child}}NameLen is returned when the {{toLower child}} name
	// exceeds the maximum number of characters.
	Err{{title child}}NameLen = errors.New("{{title child}} name cannot exceed 250 characters")

	// Err{{title child}}DescLen is returned when the {{toLower child}} desc
	// exceeds the maximum number of characters.
	Err{{title child}}DescLen = errors.New("{{title child}} description cannot exceed 250 characters")
)

// {{title child}} returns true if the {{title child}} if valid.
func {{title child}}({{toLower child}} *types.{{title child}}) (bool, error) {
	if !slug.IsSlug({{toLower child}}.Slug) {
		return false, Err{{title child}}Identifier
	}
	if len({{toLower child}}.Slug) > 250 {
		return false, Err{{title child}}IdentifierLen
	}
	if len({{toLower child}}.Name) > 250 {
		return false, Err{{title child}}NameLen
	}
	if len({{toLower child}}.Desc) > 500 {
		return false, Err{{title child}}DescLen
	}
	return true, nil
}
