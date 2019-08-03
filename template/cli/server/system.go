// Copyright 2019 Brad Rydzewski. All rights reserved.
// Use of this source code is governed by the Polyform License
// that can be found in the LICENSE.md file.

package server

import (
	"github.com/{{toLower repo}}/internal/cron"
	"github.com/{{toLower repo}}/internal/server"
)

// system stores high level system sub-routines.
type system struct {
	server  *server.Server
	nightly *cron.Nightly
}

// newSystem returns a new system structure.
func newSystem(server *server.Server, nightly *cron.Nightly) *system {
	return &system{
		server:  server,
		nightly: nightly,
	}
}
