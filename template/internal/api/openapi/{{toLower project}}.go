// Copyright 2019 Brad Rydzewski. All rights reserved.
// Use of this source code is governed by the Polyform License
// that can be found in the LICENSE.md file.

package openapi

import (
	"net/http"

	"github.com/{{toLower repo}}/internal/api/render"
	"github.com/{{toLower repo}}/types"

	"github.com/swaggest/openapi-go/openapi3"
)

type (
	// request to find or delete a {{toLower project}}.
	{{toLower project}}Request struct {
		Param string `path:"{{toLower project}}"`
	}

	// request to update a {{toLower project}}.
	{{toLower project}}PatchRequest struct {
		Param string `path:"{{toLower project}}"`

		// include request body input.
		types.{{title project}}Input
	}
)

// helper function that constructs the openapi specification
// for {{toLower project}} resources.
func build{{title project}}(reflector *openapi3.Reflector) {

	opFind := openapi3.Operation{}
	opFind.WithTags("{{toLower project}}")
	reflector.SetRequest(&opFind, new({{toLower project}}Request), http.MethodGet)
	reflector.SetJSONResponse(&opFind, new(types.{{title project}}), http.StatusOK)
	reflector.SetJSONResponse(&opFind, new(render.Error), http.StatusInternalServerError)
	reflector.SetJSONResponse(&opFind, new(render.Error), http.StatusNotFound)
	reflector.Spec.AddOperation(http.MethodGet, "/api/v1/{{toLower project}}s/{{`{`}}{{toLower project}}{{`}`}}", opFind)

	opCreate := openapi3.Operation{}
	opCreate.WithTags("{{toLower project}}")
	reflector.SetRequest(&opCreate, new(types.{{title project}}Input), http.MethodPost)
	reflector.SetJSONResponse(&opCreate, new(types.{{title project}}), http.StatusOK)
	reflector.SetJSONResponse(&opCreate, new(render.Error), http.StatusBadRequest)
	reflector.SetJSONResponse(&opCreate, new(render.Error), http.StatusInternalServerError)
	reflector.SetJSONResponse(&opCreate, new(render.Error), http.StatusNotFound)
	reflector.Spec.AddOperation(http.MethodPost, "/api/v1/{{toLower project}}s", opCreate)

	opUpdate := openapi3.Operation{}
	opUpdate.WithTags("{{toLower project}}")
	reflector.SetRequest(&opUpdate, new({{toLower project}}PatchRequest), http.MethodPatch)
	reflector.SetJSONResponse(&opUpdate, new(types.{{title project}}), http.StatusOK)
	reflector.SetJSONResponse(&opUpdate, new(render.Error), http.StatusBadRequest)
	reflector.SetJSONResponse(&opUpdate, new(render.Error), http.StatusInternalServerError)
	reflector.SetJSONResponse(&opUpdate, new(render.Error), http.StatusNotFound)
	reflector.Spec.AddOperation(http.MethodPatch, "/api/v1/{{toLower project}}s/{{`{`}}{{toLower project}}{{`}`}}", opUpdate)

	opDelete := openapi3.Operation{}
	opDelete.WithTags("{{toLower project}}")
	reflector.SetRequest(&opDelete, new({{toLower project}}Request), http.MethodDelete)
	reflector.SetJSONResponse(&opDelete, nil, http.StatusNoContent)
	reflector.SetJSONResponse(&opDelete, new(render.Error), http.StatusInternalServerError)
	reflector.SetJSONResponse(&opDelete, new(render.Error), http.StatusNotFound)
	reflector.Spec.AddOperation(http.MethodDelete, "/api/v1/{{toLower project}}s/{{`{`}}{{toLower project}}{{`}`}}", opDelete)
}
