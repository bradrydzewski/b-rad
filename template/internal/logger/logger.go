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

// Package logger provides logrus helper functions.
package logger

import (
	"context"
	"net/http"

	"github.com/sirupsen/logrus"
)

type loggerKey struct{}

// L is an alias for the the standard logger.
var L = logrus.NewEntry(logrus.StandardLogger())

// WithContext returns a new context with the provided logger. Use in
// combination with logger.WithField(s) for great effect.
func WithContext(ctx context.Context, logger *logrus.Entry) context.Context {
	return context.WithValue(ctx, loggerKey{}, logger)
}

// FromContext retrieves the current logger from the context. If no
// logger is available, the default logger is returned.
func FromContext(ctx context.Context) *logrus.Entry {
	logger := ctx.Value(loggerKey{})
	if logger == nil {
		return L
	}
	return logger.(*logrus.Entry)
}

// FromRequest retrieves the current logger from the request. If no
// logger is available, the default logger is returned.
func FromRequest(r *http.Request) *logrus.Entry {
	return FromContext(r.Context())
}
