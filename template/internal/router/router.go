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

// Package router provides http handlers for serving the
// web applicationa and API endpoints.
package router

import (
	"context"
	"net/http"
	"path/filepath"

	"github.com/{{github}}/internal/api/access"
	"github.com/{{github}}/internal/api/{{toLower child}}s"
	"github.com/{{github}}/internal/api/{{toLower parent}}s"
	"github.com/{{github}}/internal/api/login"
	"github.com/{{github}}/internal/api/members"
	"github.com/{{github}}/internal/api/projects"
	"github.com/{{github}}/internal/api/register"
	"github.com/{{github}}/internal/api/system"
	"github.com/{{github}}/internal/api/token"
	"github.com/{{github}}/internal/api/user"
	"github.com/{{github}}/internal/api/users"
	"github.com/{{github}}/internal/logger"
	"github.com/{{github}}/internal/store"
	"github.com/{{github}}/internal/swagger"
	"github.com/{{github}}/web/dist"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"github.com/unrolled/secure"
)

// empty context
var nocontext = context.Background()

// New returns a new http.Handler that routes traffic
// to the appropriate http.Handlers.
func New(
	{{toLower child}}Store store.{{child}}Store,
	{{toLower parent}}Store store.{{parent}}Store,
	memberStore store.MemberStore,
	projectStore store.ProjectStore,
	userStore store.UserStore,
	systemStore store.SystemStore,
) http.Handler {

	// create the router with caching disabled
	// for API endpoints
	r := chi.NewRouter()

	// create the auth middleware.
	auth := token.Must(userStore)

	// retrieve system configuration in order to
	// retrieve security and cors configuration options.
	config := systemStore.Config(nocontext)

	r.Route("/api/v1", func(r chi.Router) {
		r.Use(middleware.NoCache)
		r.Use(middleware.Recoverer)
		r.Use(logger.Middleware)

		cors := cors.New(
			cors.Options{
				AllowedOrigins:   config.Cors.AllowedOrigins,
				AllowedMethods:   config.Cors.AllowedMethods,
				AllowedHeaders:   config.Cors.AllowedHeaders,
				ExposedHeaders:   config.Cors.ExposedHeaders,
				AllowCredentials: config.Cors.AllowCredentials,
				MaxAge:           config.Cors.MaxAge,
			},
		)
		r.Use(cors.Handler)

		// project endpoints
		r.Route("/projects", func(r chi.Router) {
			r.Use(auth)
			r.Post("/", projects.HandleCreate(projectStore, memberStore))

			// project endpoints
			r.Route("/{project}", func(r chi.Router) {
				r.Use(access.ProjectAccess(memberStore))

				r.Get("/", projects.HandleFind(projectStore))
				r.Patch("/", projects.HandleUpdate(projectStore))
				r.Delete("/", projects.HandleDelete(projectStore))

				// {{toLower parent}} endpoints
				r.Route("/{{toLower parent}}s", func(r chi.Router) {
					r.Get("/", {{toLower parent}}s.HandleList({{toLower parent}}Store))
					r.Post("/", {{toLower parent}}s.HandleCreate({{toLower parent}}Store))
					r.Get("/{{`{`}}{{toLower parent}}{{`}`}}", {{toLower parent}}s.HandleFind({{toLower parent}}Store))
					r.Patch("/{{`{`}}{{toLower parent}}{{`}`}}", {{toLower parent}}s.HandleUpdate({{toLower parent}}Store))
					r.With(
						access.ProjectAdmin(memberStore),
					).Delete("/{{`{`}}{{toLower parent}}{{`}`}}", {{toLower parent}}s.HandleDelete({{toLower parent}}Store))

					// {{toLower child}} endpoints
					r.Route("/{{`{`}}{{toLower parent}}{{`}`}}/{{toLower child}}s", func(r chi.Router) {
						r.Get("/", {{toLower child}}s.HandleList({{toLower parent}}Store, {{toLower child}}Store))
						r.Post("/", {{toLower child}}s.HandleCreate({{toLower parent}}Store, {{toLower child}}Store))
						r.Get("/{{`{`}}{{toLower child}}{{`}`}}", {{toLower child}}s.HandleFind({{toLower parent}}Store, {{toLower child}}Store))
						r.Patch("/{{`{`}}{{toLower child}}{{`}`}}", {{toLower child}}s.HandleUpdate({{toLower parent}}Store, {{toLower child}}Store))
						r.With(
							access.ProjectAdmin(memberStore),
						).Delete("/{{`{`}}{{toLower child}}{{`}`}}", {{toLower child}}s.HandleDelete({{toLower parent}}Store, {{toLower child}}Store))
					})
				})

				// project member endpoints
				r.Route("/members", func(r chi.Router) {
					r.Use(access.ProjectAdmin(memberStore))

					r.Get("/", members.HandleList(projectStore, memberStore))
					r.Get("/{user}", members.HandleFind(userStore, projectStore, memberStore))
					r.Post("/{user}", members.HandleCreate(userStore, projectStore, memberStore))
					r.Patch("/{user}", members.HandleUpdate(userStore, projectStore, memberStore))
					r.Delete("/{user}", members.HandleDelete(userStore, projectStore, memberStore))
				})
			})
		})

		// authenticated user endpoints
		r.Route("/user", func(r chi.Router) {
			r.Use(auth)

			r.Get("/", user.HandleFind())
			r.Get("/projects", user.HandleProjects(projectStore))
			r.Patch("/user", user.HandleUpdate(userStore))
			r.Post("/token", user.HandleToken(userStore))
		})

		// user management endpoints
		r.Route("/users", func(r chi.Router) {
			r.Use(auth)
			r.Use(access.SystemAdmin())

			r.Get("/", users.HandleList(userStore))
			r.Post("/", users.HandleCreate(userStore))
			r.Get("/{user}", users.HandleFind(userStore))
			r.Patch("/{user}", users.HandleUpdate(userStore))
			r.Delete("/{user}", users.HandleDelete(userStore))
		})

		// system management endpoints
		r.Route("/system", func(r chi.Router) {
			r.Get("/health", system.HandleHealth)
			r.Get("/version", system.HandleVersion)
		})

		// user login endpoint
		r.Post("/login", login.HandleLogin(userStore, systemStore))

		// user registration endpoint
		r.Post("/register", register.HandleRegister(userStore, systemStore))
	})

	// serve swagger for embedded filesystem.
	swaggerFS := http.FileServer(swagger.FileSystem())
	r.Handle("/swagger", http.RedirectHandler("/swagger/", http.StatusSeeOther))
	r.Handle("/swagger/*", http.StripPrefix("/swagger/", swaggerFS))

	// create middleware to enforce security best practices.
	sec := secure.New(
		secure.Options{
			AllowedHosts:          config.Secure.AllowedHosts,
			HostsProxyHeaders:     config.Secure.HostsProxyHeaders,
			SSLRedirect:           config.Secure.SSLRedirect,
			SSLTemporaryRedirect:  config.Secure.SSLTemporaryRedirect,
			SSLHost:               config.Secure.SSLHost,
			SSLProxyHeaders:       config.Secure.SSLProxyHeaders,
			STSSeconds:            config.Secure.STSSeconds,
			STSIncludeSubdomains:  config.Secure.STSIncludeSubdomains,
			STSPreload:            config.Secure.STSPreload,
			ForceSTSHeader:        config.Secure.ForceSTSHeader,
			FrameDeny:             config.Secure.FrameDeny,
			ContentTypeNosniff:    config.Secure.ContentTypeNosniff,
			BrowserXssFilter:      config.Secure.BrowserXSSFilter,
			ContentSecurityPolicy: config.Secure.ContentSecurityPolicy,
			ReferrerPolicy:        config.Secure.ReferrerPolicy,
		},
	)

	// serve all other routes from the embedded filesystem.
	fs := http.FileServer(dist.FileSystem())
	r.With(sec.Handler).NotFound(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// because this is a single page application,
			// we need to always load the index.html file
			// in the root of the project, unless the path
			// points to a file with an extension (css, js, etc)
			if filepath.Ext(r.URL.Path) == "" {
				// HACK: alter the path to point to the
				// root of the project.
				r.URL.Path = "/"
			}
			// and finally server the file.
			fs.ServeHTTP(w, r)
		}),
	)

	return r
}
