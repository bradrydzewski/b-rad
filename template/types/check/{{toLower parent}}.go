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
	// Err{{title parent}}Identifier is returned when the {{toLower parent}}
	// slug is an invalid format.
	Err{{title parent}}Identifier = errors.New("Invalid {{toLower parent}} identifier")

	// Err{{title parent}}IdentifierLen is returned when the {{toLower parent}}
	// name exceeds the maximum number of characters.
	Err{{title parent}}IdentifierLen = errors.New("{{title parent}} identifier cannot exceed 250 characters")

	// Err{{title parent}}NameLen is returned when the {{toLower parent}} name
	// exceeds the maximum number of characters.
	Err{{title parent}}NameLen = errors.New("{{title parent}} name cannot exceed 250 characters")

	// Err{{title parent}}DescLen is returned when the {{toLower parent}} desc
	// exceeds the maximum number of characters.
	Err{{title parent}}DescLen = errors.New("{{title parent}} description cannot exceed 250 characters")
)

// {{title parent}} returns true if the {{title parent}} if valid.
func {{title parent}}({{toLower parent}} *types.{{title parent}}) (bool, error) {
	if !slug.IsSlug({{toLower parent}}.Slug) {
		return false, Err{{title parent}}Identifier
	}
	if len({{toLower parent}}.Slug) > 250 {
		return false, Err{{title parent}}IdentifierLen
	}
	if len({{toLower parent}}.Name) > 250 {
		return false, Err{{title parent}}NameLen
	}
	if len({{toLower parent}}.Desc) > 500 {
		return false, Err{{title parent}}DescLen
	}
	return true, nil
}
