// Copyright 2019 Brad Rydzewski. All rights reserved.
// Use of this source code is governed by the Polyform License
// that can be found in the LICENSE.md file.

package openapi

import (
	"bytes"
	"net/http"
	"strings"

	"github.com/{{toLower repo}}/version"
	"github.com/swaggest/openapi-go/openapi3"
)

type (
	// base request for pagination
	paginationRequest struct {
		Page int `query:"page"     default:"1"`
		Size int `query:"per_page" default:"100"`
	}

	// base response for pagination
	paginationResponse struct {
		Total   int      `header:"x-total"`
		Pagelen int      `header:"x-total-pages"`
		Page    int      `header:"x-page"`
		Size    int      `header:"x-per-page"`
		Next    int      `header:"x-next"`
		Prev    int      `header:"x-prev"`
		Link    []string `header:"Link"`
	}
)

// Handler returns an http.HandlerFunc that writes the openapi v3
// specification file to the http.Response body.
func Handler() http.HandlerFunc {
	spec := build()
	yaml, _ := spec.MarshalYAML()
	json, _ := spec.MarshalJSON()

	yaml = normalize(yaml)
	json = normalize(json)

	return func(w http.ResponseWriter, r *http.Request) {
		switch {
		case strings.HasSuffix(r.URL.Path, ".json"):
			w.Write(json)
		default:
			w.Write(yaml)
		}
	}
}

// build is a helper function that constructs the
// openapi specification object, which can be marshaled
// to json or yaml, as needed.
func build() *openapi3.Spec {
	reflector := openapi3.Reflector{}
	reflector.Spec = &openapi3.Spec{Openapi: "3.0.0"}
	reflector.Spec.Info.
		WithTitle("API Specification").
		WithVersion(version.Version.String())

	//
	// register endpoints
	//

	buildAccount(&reflector)
	build{{title project}}(&reflector)
	build{{title parent}}(&reflector)
	build{{title child}}(&reflector)
	buildUser(&reflector)
	buildUsers(&reflector)

	//
	// define security scheme
	//

	scheme := openapi3.SecuritySchemeOrRef{
		SecurityScheme: &openapi3.SecurityScheme{
			HTTPSecurityScheme: &openapi3.HTTPSecurityScheme{
				Scheme: "bearerAuth",
				Bearer: &openapi3.Bearer{},
			},
		},
	}
	security := openapi3.ComponentsSecuritySchemes{}
	security.WithMapOfSecuritySchemeOrRefValuesItem("bearerAuth", scheme)
	reflector.Spec.Components.WithSecuritySchemes(security)

	//
	// enforce security scheme globally
	//

	reflector.Spec.WithSecurity(map[string][]string{
		"bearerAuth": {},
	})

	return reflector.Spec
}

// helper function normalizes the output to ensure
// automatically-generated names are more user friendly.
func normalize(data []byte) []byte {
	data = bytes.ReplaceAll(data, []byte("Types"), []byte(""))
	data = bytes.ReplaceAll(data, []byte("Openapi"), []byte(""))
	data = bytes.ReplaceAll(data, []byte("FormData"), []byte(""))
	data = bytes.ReplaceAll(data, []byte("RenderError"), []byte("Error"))
	return data
}
