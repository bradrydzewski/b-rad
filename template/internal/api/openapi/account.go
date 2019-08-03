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

// request to login to an account.
type loginRequest struct {
	Username string `formData:"username"`
	Password string `formData:"password"`
}

// request to register an account.
type registerRequest struct {
	Username string `formData:"username"`
	Password string `formData:"password"`
}

// helper function that constructs the openapi specification
// for the account registration and login endpoints.
func buildAccount(reflector *openapi3.Reflector) {

	onLogin := openapi3.Operation{}
	onLogin.WithTags("account")
	reflector.SetRequest(&onLogin, new(loginRequest), http.MethodPost)
	reflector.SetJSONResponse(&onLogin, new(types.UserToken), http.StatusOK)
	reflector.SetJSONResponse(&onLogin, new(render.Error), http.StatusBadRequest)
	reflector.SetJSONResponse(&onLogin, new(render.Error), http.StatusInternalServerError)
	reflector.SetJSONResponse(&onLogin, new(render.Error), http.StatusNotFound)
	reflector.Spec.AddOperation(http.MethodPost, "/api/v1/login", onLogin)

	onRegister := openapi3.Operation{}
	onRegister.WithTags("account")
	reflector.SetRequest(&onRegister, new(registerRequest), http.MethodPost)
	reflector.SetJSONResponse(&onRegister, new(types.UserToken), http.StatusOK)
	reflector.SetJSONResponse(&onRegister, new(render.Error), http.StatusInternalServerError)
	reflector.SetJSONResponse(&onRegister, new(render.Error), http.StatusBadRequest)
	reflector.Spec.AddOperation(http.MethodPost, "/api/v1/register", onRegister)
}
