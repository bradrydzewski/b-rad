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

// Package server implements an http server.
package server

import (
	"context"
	"crypto/tls"
	"net/http"

	"golang.org/x/crypto/acme/autocert"
	"golang.org/x/sync/errgroup"
)

// A Server defines parameters for running an HTTP server.
type Server struct {
	Acme    bool
	Addr    string
	Cert    string
	Key     string
	Host    string
	Handler http.Handler
}

// ListenAndServe initializes a server to respond to HTTP network requests.
func (s *Server) ListenAndServe(ctx context.Context) error {
	if s.Acme {
		return s.listenAndServeAcme(ctx)
	} else if s.Key != "" {
		return s.listenAndServeTLS(ctx)
	}
	return s.listenAndServe(ctx)
}

func (s *Server) listenAndServe(ctx context.Context) error {
	var g errgroup.Group
	s1 := &http.Server{
		Addr:    s.Addr,
		Handler: s.Handler,
	}
	g.Go(func() error {
		select {
		case <-ctx.Done():
			return s1.Shutdown(ctx)
		}
	})
	g.Go(func() error {
		return s1.ListenAndServe()
	})
	return g.Wait()
}

func (s *Server) listenAndServeTLS(ctx context.Context) error {
	var g errgroup.Group
	s1 := &http.Server{
		Addr:    ":http",
		Handler: http.HandlerFunc(redirect),
	}
	s2 := &http.Server{
		Addr:    ":https",
		Handler: s.Handler,
	}
	g.Go(func() error {
		return s1.ListenAndServe()
	})
	g.Go(func() error {
		return s2.ListenAndServeTLS(
			s.Cert,
			s.Key,
		)
	})
	g.Go(func() error {
		select {
		case <-ctx.Done():
			s1.Shutdown(ctx)
			s2.Shutdown(ctx)
			return nil
		}
	})
	return g.Wait()
}

func (s Server) listenAndServeAcme(ctx context.Context) error {
	var g errgroup.Group
	m := &autocert.Manager{
		Cache:      autocert.DirCache(".cache"),
		Prompt:     autocert.AcceptTOS,
		HostPolicy: autocert.HostWhitelist(s.Host),
	}
	s1 := &http.Server{
		Addr:    ":http",
		Handler: m.HTTPHandler(nil),
	}
	s2 := &http.Server{
		Addr:    ":https",
		Handler: s.Handler,
		TLSConfig: &tls.Config{
			GetCertificate: m.GetCertificate,
			NextProtos:     []string{"h2", "http/1.1"},
		},
	}
	g.Go(func() error {
		return s1.ListenAndServe()
	})
	g.Go(func() error {
		return s2.ListenAndServeTLS("", "")
	})
	g.Go(func() error {
		select {
		case <-ctx.Done():
			s1.Shutdown(ctx)
			s2.Shutdown(ctx)
			return nil
		}
	})
	return g.Wait()
}

func redirect(w http.ResponseWriter, req *http.Request) {
	target := "https://" + req.Host + req.URL.Path
	http.Redirect(w, req, target, http.StatusTemporaryRedirect)
}
