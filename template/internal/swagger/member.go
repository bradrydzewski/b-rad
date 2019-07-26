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

package swagger

import (
	"net/http"

	"github.com/{{github}}/types"
)

// swagger:route GET /projects/{project}/members/{user} member getMember
//
// Get the project member with the matching email address.
//
//     Responses:
//       200: member
//
func memberFind(w http.ResponseWriter, r *http.Request) {}

// swagger:route GET  /projects/{project}/members member getMemberList
//
// Get the list of all project members.
//
//     Responses:
//       200: memberList
//
func memberList(w http.ResponseWriter, r *http.Request) {}

// swagger:route POST /projects/{project}/members member createMember
//
// Create a new project member.
//
//     Responses:
//       200: member
//
func memberCreate(w http.ResponseWriter, r *http.Request) {}

// swagger:route PATCH /projects/{project}/members/{user} member updateMember
//
// Update the project member.
//
//     Responses:
//       200: member
//
func memberUpdate(w http.ResponseWriter, r *http.Request) {}

// swagger:route DELETE /projects/{project}/members/{user} member deleteMember
//
// Delete the project member.
//
//     Responses:
//       204:
//
func memberDelete(w http.ResponseWriter, r *http.Request) {}

// swagger:parameters getMember deleteMember
type memberReq struct {
	// in: path
	Project int64 `json:"project"`

	// in: path
	Email string `json:"user"`
}

// swagger:parameters getMemberList
type memberListReq struct {
	// in: path
	Project int64 `json:"project"`
}

// swagger:parameters createMember
type memberCreateReq struct {
	// in: path
	Project int64 `json:"project"`

	// in: body
	Body types.MembershipInput
}

// swagger:parameters updateMember
type memberUpdateReq struct {
	// in: path
	Project int64 `json:"project"`

	// in: path
	Email string `json:"user"`

	// in: body
	Body types.MembershipInput
}

// swagger:response member
type memberResp struct {
	// in: body
	Body types.Member
}

// swagger:response memberList
type memberListResp struct {
	// in: body
	Body []types.Member
}
