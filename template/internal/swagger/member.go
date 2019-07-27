// Copyright 2019 Brad Rydzewski. All rights reserved.
// Use of this source code is governed by the Polyform License
// that can be found in the LICENSE.md file.

package swagger

import (
	"net/http"

	"github.com/{{toLower repo}}/types"
)

// swagger:route GET /{{toLower project}}s/{{`{`}}{{toLower project}}{{`}`}}/members/{user} member getMember
//
// Get the {{toLower project}} member with the matching email address.
//
//     Responses:
//       200: member
//
func memberFind(w http.ResponseWriter, r *http.Request) {}

// swagger:route GET  /{{toLower project}}s/{{`{`}}{{toLower project}}{{`}`}}/members member getMemberList
//
// Get the list of all {{toLower project}} members.
//
//     Responses:
//       200: memberList
//
func memberList(w http.ResponseWriter, r *http.Request) {}

// swagger:route POST /{{toLower project}}s/{{`{`}}{{toLower project}}{{`}`}}/members member createMember
//
// Create a new {{toLower project}} member.
//
//     Responses:
//       200: member
//
func memberCreate(w http.ResponseWriter, r *http.Request) {}

// swagger:route PATCH /{{toLower project}}s/{{`{`}}{{toLower project}}{{`}`}}/members/{user} member updateMember
//
// Update the {{toLower project}} member.
//
//     Responses:
//       200: member
//
func memberUpdate(w http.ResponseWriter, r *http.Request) {}

// swagger:route DELETE /{{toLower project}}s/{{`{`}}{{toLower project}}{{`}`}}/members/{user} member deleteMember
//
// Delete the {{toLower project}} member.
//
//     Responses:
//       204:
//
func memberDelete(w http.ResponseWriter, r *http.Request) {}

// swagger:parameters getMember deleteMember
type memberReq struct {
	// in: path
	{{title project}} int64 `json:"{{toLower project}}"`

	// in: path
	Email string `json:"user"`
}

// swagger:parameters getMemberList
type memberListReq struct {
	// in: path
	{{title project}} int64 `json:"{{toLower project}}"`
}

// swagger:parameters createMember
type memberCreateReq struct {
	// in: path
	{{title project}} int64 `json:"{{toLower project}}"`

	// in: body
	Body types.MembershipInput
}

// swagger:parameters updateMember
type memberUpdateReq struct {
	// in: path
	{{title project}} int64 `json:"{{toLower project}}"`

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
