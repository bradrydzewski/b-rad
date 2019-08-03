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
	// Err{{title project}}Identifier is returned when the {{toLower project}}
	// slug is an invalid format.
	Err{{title project}}Identifier = errors.New("Invalid {{toLower project}} identifier")

	// Err{{title project}}IdentifierLen is returned when the {{toLower project}}
	// name exceeds the maximum number of characters.
	Err{{title project}}IdentifierLen = errors.New("{{title project}} identifier cannot exceed 250 characters")

	// Err{{title project}}NameLen is returned when the {{toLower project}} name
	// exceeds the maximum number of characters.
	Err{{title project}}NameLen = errors.New("{{title project}} name cannot exceed 250 characters")

	// Err{{title project}}DescLen is returned when the {{toLower project}} desc
	// exceeds the maximum number of characters.
	Err{{title project}}DescLen = errors.New("{{title project}} description cannot exceed 250 characters")
)

// {{title project}} returns true if the {{title project}} if valid.
func {{title project}}({{toLower project}} *types.{{title project}}) (bool, error) {
	if !slug.IsSlug({{toLower project}}.Slug) {
		return false, Err{{title project}}Identifier
	}
	if len({{toLower project}}.Slug) > 250 {
		return false, Err{{title project}}IdentifierLen
	}
	if len({{toLower project}}.Name) > 250 {
		return false, Err{{title project}}NameLen
	}
	if len({{toLower project}}.Desc) > 500 {
		return false, Err{{title project}}DescLen
	}
	return true, nil
}
