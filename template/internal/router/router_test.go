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

package router

import "testing"

// this unit test ensures routes that require authorization
// return a 401 unauthorized if no token, or an invalid token
// is provided.
func TestTokenGate(t *testing.T) {
	t.Skip()
}

// this unit test ensures routes that require project access
// return a 403 forbidden if the user does not have acess
// to the project
func TestProjectGate(t *testing.T) {
	t.Skip()
}

// this unit test ensures routes that require system access
// return a 403 forbidden if the user does not have acess
// to the project
func TestSystemGate(t *testing.T) {
	t.Skip()
}
