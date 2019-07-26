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

// swagger:route GET /projects/{project} project findProject
//
// Get the project with the matching project id.
//
//     Responses:
//       200: project
//
func projectFind(w http.ResponseWriter, r *http.Request) {}

// swagger:route POST /projects project createProject
//
// Create a new project.
//
//     Responses:
//       200: project
//
func projectCreate(w http.ResponseWriter, r *http.Request) {}

// swagger:route PATCH /projects/{project} project updateProject
//
// Update the project with the matching project id.
//
//     Responses:
//       200: project
//
func projectUpdate(w http.ResponseWriter, r *http.Request) {}

// swagger:route DELETE /projects/{project} project deleteProject
//
// Delete the project with the matching project id.
//
//     Responses:
//       204:
//
func projectDelete(w http.ResponseWriter, r *http.Request) {}

// swagger:parameters findProject deleteProject
type projectReq struct {
	// in: path
	ID int64 `json:"project"`
}

// swagger:parameters createProject
type projectCreateReq struct {
	// in: body
	Body types.ProjectInput
}

// swagger:parameters updateProject
type projectUpdateReq struct {
	// in: path
	ID int64 `json:"project"`

	// in: body
	Body types.ProjectInput
}

// swagger:parameters projectDelete
type projectDeleteInput struct {
	// in: path
	ID int64 `json:"project"`
}

// swagger:response project
type projectResp struct {
	// in: body
	Body types.Project
}

// swagger:response projectList
type projectListResp struct {
	// in: body
	Body []types.Project
}
