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

// helper function that constructs the openapi specification
// for user account resources.
func buildUser(reflector *openapi3.Reflector) {

	opFind := openapi3.Operation{}
	opFind.WithTags("user")
	reflector.SetRequest(&opFind, nil, http.MethodGet)
	reflector.SetJSONResponse(&opFind, new(types.User), http.StatusOK)
	reflector.SetJSONResponse(&opFind, new(render.Error), http.StatusInternalServerError)
	reflector.Spec.AddOperation(http.MethodGet, "/api/v1/user", opFind)

	opUpdate := openapi3.Operation{}
	opUpdate.WithTags("user")
	reflector.SetRequest(&opUpdate, new(types.UserInput), http.MethodPatch)
	reflector.SetJSONResponse(&opUpdate, new(types.User), http.StatusOK)
	reflector.SetJSONResponse(&opUpdate, new(render.Error), http.StatusInternalServerError)
	reflector.Spec.AddOperation(http.MethodPatch, "/api/v1/user", opUpdate)

	opToken := openapi3.Operation{}
	opToken.WithTags("user")
	reflector.SetRequest(&opToken, new(types.Token), http.MethodPost)
	reflector.SetJSONResponse(&opToken, new(types.User), http.StatusOK)
	reflector.SetJSONResponse(&opToken, new(render.Error), http.StatusInternalServerError)
	reflector.Spec.AddOperation(http.MethodPost, "/api/v1/user/token", opToken)

	on{{title project}}s := openapi3.Operation{}
	on{{title project}}s.WithTags("user")
	reflector.SetRequest(&on{{title project}}s, new(paginationRequest), http.MethodGet)
	reflector.SetJSONResponse(&on{{title project}}s, new([]*types.{{title project}}), http.StatusOK)
	reflector.SetJSONResponse(&on{{title project}}s, new(render.Error), http.StatusInternalServerError)
	reflector.Spec.AddOperation(http.MethodGet, "/api/v1/user/{{toLower project}}s", on{{title project}}s)
}
