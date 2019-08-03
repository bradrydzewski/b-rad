// Copyright 2019 Brad Rydzewski. All rights reserved.
// Use of this source code is governed by the Polyform License
// that can be found in the LICENSE.md file.

// Package router provides http handlers for serving the
// web applicationa and API endpoints.
package router

import (
	"context"
	"net/http"

	"github.com/{{toLower repo}}/internal/api/handler/account"
	"github.com/{{toLower repo}}/internal/api/handler/{{toLower parent}}s"
	"github.com/{{toLower repo}}/internal/api/handler/members"
	"github.com/{{toLower repo}}/internal/api/handler/{{toLower project}}s"
	"github.com/{{toLower repo}}/internal/api/handler/{{toLower child}}s"
	"github.com/{{toLower repo}}/internal/api/handler/system"
	"github.com/{{toLower repo}}/internal/api/handler/user"
	"github.com/{{toLower repo}}/internal/api/handler/users"
	"github.com/{{toLower repo}}/internal/api/middleware/access"
	"github.com/{{toLower repo}}/internal/api/middleware/address"
	"github.com/{{toLower repo}}/internal/api/middleware/token"
	"github.com/{{toLower repo}}/internal/api/openapi"
	"github.com/{{toLower repo}}/internal/store"
	"github.com/{{toLower repo}}/web/dist"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"github.com/rs/zerolog/hlog"
	"github.com/rs/zerolog/log"
	"github.com/swaggest/swgui/v3emb"
	"github.com/unrolled/secure"
)

// empty context
var nocontext = context.Background()

// New returns a new http.Handler that routes traffic
// to the appropriate http.Handlers.
func New(
	{{toLower child}}Store store.{{title child}}Store,
	{{toLower parent}}Store store.{{title parent}}Store,
	memberStore store.MemberStore,
	{{toLower project}}Store store.{{title project}}Store,
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

		// configure middleware to help ascertain the true
		// server address from the incoming http.Request
		r.Use(
			address.Handler(
				config.Server.Proto,
				config.Server.Host,
			),
		)

		// configure logging middleware.
		r.Use(hlog.NewHandler(log.Logger))
		r.Use(hlog.URLHandler("path"))
		r.Use(hlog.MethodHandler("method"))
		r.Use(hlog.RequestIDHandler("request", "Request-Id"))

		// configure cors middleware
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

		// {{toLower project}} endpoints
		r.Route("/{{toLower project}}s", func(r chi.Router) {
			r.Use(auth)
			r.Post("/", {{toLower project}}s.HandleCreate({{toLower project}}Store, memberStore))

			// {{toLower project}} endpoints
			r.Route("/{{`{`}}{{toLower project}}{{`}`}}", func(r chi.Router) {
				r.Use(access.{{title project}}Access({{toLower project}}Store, memberStore))

				r.Get("/", {{toLower project}}s.HandleFind({{toLower project}}Store, memberStore))
				r.Patch("/", {{toLower project}}s.HandleUpdate({{toLower project}}Store))
				r.Delete("/", {{toLower project}}s.HandleDelete({{toLower project}}Store))

				// {{toLower parent}} endpoints
				r.Route("/{{toLower parent}}s", func(r chi.Router) {
					r.Get("/", {{toLower parent}}s.HandleList({{toLower project}}Store, {{toLower parent}}Store))
					r.Post("/", {{toLower parent}}s.HandleCreate({{toLower project}}Store, {{toLower parent}}Store))
					r.Get("/{{`{`}}{{toLower parent}}{{`}`}}", {{toLower parent}}s.HandleFind({{toLower project}}Store, {{toLower parent}}Store))
					r.Patch("/{{`{`}}{{toLower parent}}{{`}`}}", {{toLower parent}}s.HandleUpdate({{toLower project}}Store, {{toLower parent}}Store))
					r.With(
						access.{{title project}}Admin(memberStore),
					).Delete("/{{`{`}}{{toLower parent}}{{`}`}}", {{toLower parent}}s.HandleDelete({{toLower project}}Store, {{toLower parent}}Store))

					// {{toLower child}} endpoints
					r.Route("/{{`{`}}{{toLower parent}}{{`}`}}/{{toLower child}}s", func(r chi.Router) {
						r.Get("/", {{toLower child}}s.HandleList({{toLower project}}Store, {{toLower parent}}Store, {{toLower child}}Store))
						r.Post("/", {{toLower child}}s.HandleCreate({{toLower project}}Store, {{toLower parent}}Store, {{toLower child}}Store))
						r.Get("/{{`{`}}{{toLower child}}{{`}`}}", {{toLower child}}s.HandleFind({{toLower project}}Store, {{toLower parent}}Store, {{toLower child}}Store))
						r.Patch("/{{`{`}}{{toLower child}}{{`}`}}", {{toLower child}}s.HandleUpdate({{toLower project}}Store, {{toLower parent}}Store, {{toLower child}}Store))
						r.With(
							access.{{title project}}Admin(memberStore),
						).Delete("/{{`{`}}{{toLower child}}{{`}`}}", {{toLower child}}s.HandleDelete({{toLower project}}Store, {{toLower parent}}Store, {{toLower child}}Store))
					})
				})

				// {{toLower project}} member endpoints
				r.Route("/members", func(r chi.Router) {
					r.Use(access.{{title project}}Admin(memberStore))

					r.Get("/", members.HandleList({{toLower project}}Store, memberStore))
					r.Get("/{user}", members.HandleFind(userStore, {{toLower project}}Store, memberStore))
					r.Post("/{user}", members.HandleCreate(userStore, {{toLower project}}Store, memberStore))
					r.Patch("/{user}", members.HandleUpdate(userStore, {{toLower project}}Store, memberStore))
					r.Delete("/{user}", members.HandleDelete(userStore, {{toLower project}}Store, memberStore))
				})
			})
		})

		// authenticated user endpoints
		r.Route("/user", func(r chi.Router) {
			r.Use(auth)

			r.Get("/", user.HandleFind())
			r.Patch("/", user.HandleUpdate(userStore))
			r.Get("/{{toLower project}}s", user.Handle{{title project}}s({{toLower project}}Store))
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
		r.Post("/login", account.HandleLogin(userStore, systemStore))

		// user registration endpoint
		r.Post("/register", account.HandleRegister(userStore, systemStore))

		// openapi specification endpoints
		swagger := openapi.Handler()
		r.Handle("/swagger.json", swagger)
		r.Handle("/swagger.yaml", swagger)
	})

	// openapi playground endpoints
	swagger := v3emb.NewHandler("API Definition", "/api/v1/swagger.yaml", "/swagger")
	r.Handle("/swagger", swagger)
	r.Handle("/swagger/*", swagger)

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
	r.With(sec.Handler).NotFound(
		dist.Handler(),
	)

	return r
}
