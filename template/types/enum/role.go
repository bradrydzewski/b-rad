// Copyright 2019 Brad Rydzewski. All rights reserved.
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

package enum

import "encoding/json"

// Role defines the member role.
type Role int

// Role enumeration.
const (
	RoleDeveloper Role = iota
	RoleAdmin
)

// String returns the Role as a string.
func (e Role) String() string {
	switch e {
	case RoleDeveloper:
		return "developer"
	case RoleAdmin:
		return "admin"
	default:
		return "developer"
	}
}

// MarshalJSON marshals the Type as a JSON string.
func (e Role) MarshalJSON() ([]byte, error) {
	return json.Marshal(e.String())
}

// UnmarshalJSON unmashals a quoted json string to the enum value.
func (e *Role) UnmarshalJSON(b []byte) error {
	var v string
	json.Unmarshal(b, &v)
	switch v {
	case "admin":
		*e = RoleAdmin
	case "developer":
		*e = RoleDeveloper
	default:
		*e = RoleDeveloper
	}
	return nil
}
