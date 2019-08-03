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
	// request to find or delete a {{toLower parent}}.
	{{toLower parent}}Request struct {
		{{title project}} string `path:"{{toLower project}}"`
		{{title parent}}   string `path:"{{toLower parent}}"`
	}

	// request to list all {{toLower parent}}s
	{{toLower parent}}ListRequest struct {
		{{title project}} string `path:"{{toLower project}}"`

		// include pagination
		paginationRequest
	}

	// request to create a {{toLower parent}}.
	{{toLower parent}}PostRequest struct {
		{{title project}} string `path:"{{toLower project}}"`

		// include request body input.
		types.{{title parent}}Input
	}

	// request to update a {{toLower parent}}.
	{{toLower parent}}PatchRequest struct {
		{{title project}} string `path:"{{toLower project}}"`
		{{title parent}}   string `path:"{{toLower parent}}"`

		// include request body input.
		types.{{title parent}}Input
	}
)

// helper function that constructs the openapi specification
// for {{toLower parent}} resources.
func build{{title parent}}(reflector *openapi3.Reflector) {

	opFind := openapi3.Operation{}
	opFind.WithTags("{{toLower parent}}")
	reflector.SetRequest(&opFind, new({{toLower parent}}Request), http.MethodGet)
	reflector.SetJSONResponse(&opFind, new(types.{{title parent}}), http.StatusOK)
	reflector.SetJSONResponse(&opFind, new(render.Error), http.StatusInternalServerError)
	reflector.SetJSONResponse(&opFind, new(render.Error), http.StatusNotFound)
	reflector.Spec.AddOperation(http.MethodGet, "/api/v1/{{toLower project}}s/{{`{`}}{{toLower project}}{{`}`}}/{{toLower parent}}s/{{`{`}}{{toLower parent}}{{`}`}}", opFind)

	opList := openapi3.Operation{}
	opList.WithTags("{{toLower parent}}")
	reflector.SetRequest(&opList, new({{toLower parent}}ListRequest), http.MethodGet)
	reflector.SetJSONResponse(&opList, new([]*types.{{title parent}}), http.StatusOK)
	reflector.SetJSONResponse(&opList, new(render.Error), http.StatusInternalServerError)
	reflector.SetJSONResponse(&opList, new(render.Error), http.StatusNotFound)
	reflector.Spec.AddOperation(http.MethodGet, "/api/v1/{{toLower project}}s/{{`{`}}{{toLower project}}{{`}`}}/{{toLower parent}}s", opList)

	opCreate := openapi3.Operation{}
	opCreate.WithTags("{{toLower parent}}")
	reflector.SetRequest(&opCreate, new({{toLower parent}}PostRequest), http.MethodPost)
	reflector.SetJSONResponse(&opCreate, new(types.{{title parent}}), http.StatusOK)
	reflector.SetJSONResponse(&opCreate, new(render.Error), http.StatusBadRequest)
	reflector.SetJSONResponse(&opCreate, new(render.Error), http.StatusInternalServerError)
	reflector.SetJSONResponse(&opCreate, new(render.Error), http.StatusNotFound)
	reflector.Spec.AddOperation(http.MethodPost, "/api/v1/{{toLower project}}s/{{`{`}}{{toLower project}}{{`}`}}/{{toLower parent}}s", opCreate)

	opUpdate := openapi3.Operation{}
	opUpdate.WithTags("{{toLower parent}}")
	reflector.SetRequest(&opUpdate, new({{toLower parent}}PatchRequest), http.MethodPatch)
	reflector.SetJSONResponse(&opUpdate, new(types.{{title parent}}), http.StatusOK)
	reflector.SetJSONResponse(&opUpdate, new(render.Error), http.StatusBadRequest)
	reflector.SetJSONResponse(&opUpdate, new(render.Error), http.StatusInternalServerError)
	reflector.SetJSONResponse(&opUpdate, new(render.Error), http.StatusNotFound)
	reflector.Spec.AddOperation(http.MethodPatch, "/api/v1/{{toLower project}}s/{{`{`}}{{toLower project}}{{`}`}}/{{toLower parent}}s/{{`{`}}{{toLower parent}}{{`}`}}", opUpdate)

	opDelete := openapi3.Operation{}
	opDelete.WithTags("{{toLower parent}}")
	reflector.SetRequest(&opDelete, new({{toLower parent}}Request), http.MethodDelete)
	reflector.SetJSONResponse(&opDelete, nil, http.StatusNoContent)
	reflector.SetJSONResponse(&opDelete, new(render.Error), http.StatusInternalServerError)
	reflector.SetJSONResponse(&opDelete, new(render.Error), http.StatusNotFound)
	reflector.Spec.AddOperation(http.MethodDelete, "/api/v1/{{toLower project}}s/{{`{`}}{{toLower project}}{{`}`}}/{{toLower parent}}s/{{`{`}}{{toLower parent}}{{`}`}}", opDelete)
}
