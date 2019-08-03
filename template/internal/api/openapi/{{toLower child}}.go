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
	// request to find or delete a {{toLower child}}.
	{{toLower child}}Request struct {
		{{title project}} string `path:"{{toLower project}}"`
		{{title parent}}   string `path:"{{toLower parent}}"`
		{{title child}} string `path:"{{toLower child}}"`
	}

	// request to list all {{toLower child}}s
	{{toLower child}}ListRequest struct {
		{{title project}} string `path:"{{toLower project}}"`
		{{title parent}}   string `path:"{{toLower parent}}"`

		// include pagination
		paginationRequest
	}

	// request to create a {{toLower parent}}.
	{{toLower child}}PostRequest struct {
		{{title project}} string `path:"{{toLower project}}"`
		{{title parent}}   string `path:"{{toLower parent}}"`

		// include request body input.
		types.{{title child}}Input
	}

	// request to update a {{toLower parent}}.
	{{toLower child}}PatchRequest struct {
		{{title project}} string `path:"{{toLower project}}"`
		{{title parent}}   string `path:"{{toLower parent}}"`
		{{title child}} string `path:"{{toLower child}}"`

		// include request body input.
		types.{{title child}}Input
	}
)

// helper function that constructs the openapi specification
// for {{toLower parent}} resources.
func build{{title child}}(reflector *openapi3.Reflector) {

	opFind := openapi3.Operation{}
	opFind.WithTags("{{toLower child}}")
	reflector.SetRequest(&opFind, new({{toLower child}}Request), http.MethodGet)
	reflector.SetJSONResponse(&opFind, new(types.{{title child}}), http.StatusOK)
	reflector.SetJSONResponse(&opFind, new(render.Error), http.StatusInternalServerError)
	reflector.SetJSONResponse(&opFind, new(render.Error), http.StatusNotFound)
	reflector.Spec.AddOperation(http.MethodGet, "/api/v1/{{toLower project}}s/{{`{`}}{{toLower project}}{{`}`}}/{{toLower parent}}s/{{`{`}}{{toLower parent}}{{`}`}}/{{toLower child}}s/{{`{`}}{{toLower child}}{{`}`}}", opFind)

	opList := openapi3.Operation{}
	opList.WithTags("{{toLower child}}")
	reflector.SetRequest(&opList, new({{toLower child}}ListRequest), http.MethodGet)
	reflector.SetJSONResponse(&opList, new([]*types.{{title child}}), http.StatusOK)
	reflector.SetJSONResponse(&opList, new(render.Error), http.StatusInternalServerError)
	reflector.SetJSONResponse(&opList, new(render.Error), http.StatusNotFound)
	reflector.Spec.AddOperation(http.MethodGet, "/api/v1/{{toLower project}}s/{{`{`}}{{toLower project}}{{`}`}}/{{toLower parent}}s/{{`{`}}{{toLower parent}}{{`}`}}/{{toLower child}}s", opList)

	opCreate := openapi3.Operation{}
	opCreate.WithTags("{{toLower child}}")
	reflector.SetRequest(&opCreate, new({{toLower child}}PostRequest), http.MethodPost)
	reflector.SetJSONResponse(&opCreate, new(types.{{title child}}), http.StatusOK)
	reflector.SetJSONResponse(&opCreate, new(render.Error), http.StatusBadRequest)
	reflector.SetJSONResponse(&opCreate, new(render.Error), http.StatusInternalServerError)
	reflector.SetJSONResponse(&opCreate, new(render.Error), http.StatusNotFound)
	reflector.Spec.AddOperation(http.MethodPost, "/api/v1/{{toLower project}}s/{{`{`}}{{toLower project}}{{`}`}}/{{toLower parent}}s/{{`{`}}{{toLower parent}}{{`}`}}/{{toLower child}}s", opCreate)

	opUpdate := openapi3.Operation{}
	opUpdate.WithTags("{{toLower child}}")
	reflector.SetRequest(&opUpdate, new({{toLower child}}PatchRequest), http.MethodPatch)
	reflector.SetJSONResponse(&opUpdate, new(types.{{title child}}), http.StatusOK)
	reflector.SetJSONResponse(&opUpdate, new(render.Error), http.StatusBadRequest)
	reflector.SetJSONResponse(&opUpdate, new(render.Error), http.StatusInternalServerError)
	reflector.SetJSONResponse(&opUpdate, new(render.Error), http.StatusNotFound)
	reflector.Spec.AddOperation(http.MethodPatch, "/api/v1/{{toLower project}}s/{{`{`}}{{toLower project}}{{`}`}}/{{toLower parent}}s/{{`{`}}{{toLower parent}}{{`}`}}/{{toLower child}}s/{{`{`}}{{toLower child}}{{`}`}}", opUpdate)

	opDelete := openapi3.Operation{}
	opDelete.WithTags("{{toLower child}}")
	reflector.SetRequest(&opDelete, new({{toLower child}}Request), http.MethodDelete)
	reflector.SetJSONResponse(&opDelete, nil, http.StatusNoContent)
	reflector.SetJSONResponse(&opDelete, new(render.Error), http.StatusInternalServerError)
	reflector.SetJSONResponse(&opDelete, new(render.Error), http.StatusNotFound)
	reflector.Spec.AddOperation(http.MethodDelete, "/api/v1/{{toLower project}}s/{{`{`}}{{toLower project}}{{`}`}}/{{toLower parent}}s/{{`{`}}{{toLower parent}}{{`}`}}/{{toLower child}}s/{{`{`}}{{toLower child}}{{`}`}}", opDelete)
}
