// Copyright 2019 Brad Rydzewski. All rights reserved.
// Use of this source code is governed by the Polyform License
// that can be found in the LICENSE.md file.

package enum

import "strings"

// UserField defines user attributes that can be
// used for sorting and filtering.
type UserAttr int

// Order enumeration.
const (
	UserAttrNone UserAttr = iota
	UserAttrId
	UserAttrName
	UserAttrEmail
	UserAttrAdmin
	UserAttrCreated
	UserAttrUpdated
)

// ParseUserAttr parses the user attribute string
// and returns the equivalent enumeration.
func ParseUserAttr(s string) UserAttr {
	switch strings.ToLower(s) {
	case "id":
		return UserAttrId
	case "name":
		return UserAttrName
	case "email":
		return UserAttrEmail
	case "admin":
		return UserAttrAdmin
	case "created", "created_at":
		return UserAttrCreated
	case "updated", "updated_at":
		return UserAttrUpdated
	default:
		return UserAttrNone
	}
}
