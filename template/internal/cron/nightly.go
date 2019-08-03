// Copyright 2019 Brad Rydzewski. All rights reserved.
// Use of this source code is governed by the Polyform License
// that can be found in the LICENSE.md file.

package cron

import (
	"context"
	"time"

	"github.com/rs/zerolog/log"
)

// helper function returns the current time.
var now = time.Now

// Nightly is a sub-routine that periodically purges historical data.
type Nightly struct {
	// Inject required stores here
}

// NewNightly returns a new Nightly sub-routine.
func NewNightly() *Nightly {
	return &Nightly{}
}

// Run runs the purge sub-routine.
func (n *Nightly) Run(ctx context.Context) {
	ticker := time.NewTicker(time.Hour * 24)
	logger := log.Ctx(ctx)
	for {
		select {
		case <-ctx.Done():
			return // break
		case <-ticker.C:
			// TODO replace this with your nightly
			// cron tasks.
			logger.Trace().Msg("cron job executed")
		}
	}
}
