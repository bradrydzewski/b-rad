// Copyright 2019 Brad Rydzewski. All rights reserved.
// Use of this source code is governed by the Polyform License
// that can be found in the LICENSE.md file.

package {{toLower parent}}s

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/{{toLower repo}}/internal/api/render"
	"github.com/{{toLower repo}}/internal/logger"
	"github.com/{{toLower repo}}/internal/store"
	"github.com/{{toLower repo}}/types"
	"github.com/go-chi/chi"
)

// HandleCreate returns an http.HandlerFunc that creates
// the object and persists to the datastore.
func HandleCreate({{toLower parent}}s store.{{title parent}}Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		{{toLower project}}, err := strconv.ParseInt(chi.URLParam(r, "{{toLower project}}"), 10, 64)
		if err != nil {
			render.BadRequest(w, err)
			logger.FromRequest(r).
				WithError(err).
				Debugln("cannot parse {{toLower project}} id")
			return
		}

		in := new(types.{{title parent}}Input)
		err = json.NewDecoder(r.Body).Decode(in)
		if err != nil {
			render.BadRequest(w, err)
			logger.FromRequest(r).
				WithError(err).
				WithField("{{toLower project}}", {{toLower project}}).
				Debugln("cannot unmarshal json request")
			return
		}

		{{toLower parent}} := &types.{{title parent}}{
			{{title project}}: {{toLower project}},
			Name:    in.Name.String,
			Desc:    in.Desc.String,
			Created: time.Now().Unix(),
			Updated: time.Now().Unix(),
		}

		err = {{toLower parent}}s.Create(r.Context(), {{toLower parent}})
		if err != nil {
			render.InternalError(w, err)
			logger.FromRequest(r).
				WithError(err).
				WithField("name", {{toLower parent}}.Name).
				WithField("{{toLower project}}", {{toLower project}}).
				Errorln("cannot create {{toLower parent}}")
		} else {
			render.JSON(w, {{toLower parent}}, 200)
		}
	}
}
