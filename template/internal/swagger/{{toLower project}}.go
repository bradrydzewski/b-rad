// Copyright 2019 Brad Rydzewski. All rights reserved.
// Use of this source code is governed by the Polyform License
// that can be found in the LICENSE.md file.

package swagger

import (
	"net/http"

	"github.com/{{toLower repo}}/types"
)

// swagger:route GET /{{toLower project}}s/{{`{`}}{{toLower project}}{{`}`}} {{toLower project}} find{{title project}}
//
// Get the {{toLower project}} with the matching {{toLower project}} id.
//
//     Responses:
//       200: {{toLower project}}
//
func {{toLower project}}Find(w http.ResponseWriter, r *http.Request) {}

// swagger:route POST /{{toLower project}}s {{toLower project}} create{{title project}}
//
// Create a new {{toLower project}}.
//
//     Responses:
//       200: {{toLower project}}
//
func {{toLower project}}Create(w http.ResponseWriter, r *http.Request) {}

// swagger:route PATCH /{{toLower project}}s/{{`{`}}{{toLower project}}{{`}`}} {{toLower project}} update{{title project}}
//
// Update the {{toLower project}} with the matching {{toLower project}} id.
//
//     Responses:
//       200: {{toLower project}}
//
func {{toLower project}}Update(w http.ResponseWriter, r *http.Request) {}

// swagger:route DELETE /{{toLower project}}s/{{`{`}}{{toLower project}}{{`}`}} {{toLower project}} delete{{title project}}
//
// Delete the {{toLower project}} with the matching {{toLower project}} id.
//
//     Responses:
//       204:
//
func {{toLower project}}Delete(w http.ResponseWriter, r *http.Request) {}

// swagger:parameters find{{title project}} delete{{title project}}
type {{toLower project}}Req struct {
	// in: path
	ID int64 `json:"{{toLower project}}"`
}

// swagger:parameters create{{title project}}
type {{toLower project}}CreateReq struct {
	// in: body
	Body types.{{title project}}Input
}

// swagger:parameters update{{title project}}
type {{toLower project}}UpdateReq struct {
	// in: path
	ID int64 `json:"{{toLower project}}"`

	// in: body
	Body types.{{title project}}Input
}

// swagger:parameters {{toLower project}}Delete
type {{toLower project}}DeleteInput struct {
	// in: path
	ID int64 `json:"{{toLower project}}"`
}

// swagger:response {{toLower project}}
type {{toLower project}}Resp struct {
	// in: body
	Body types.{{title project}}
}

// swagger:response {{toLower project}}List
type {{toLower project}}ListResp struct {
	// in: body
	Body []types.{{title project}}
}
