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

package logger

import (
	"context"
	"net/http"
	"testing"

	"github.com/sirupsen/logrus"
)

func TestContext(t *testing.T) {
	entry := logrus.NewEntry(logrus.StandardLogger())

	ctx := WithContext(context.Background(), entry)
	got := FromContext(ctx)

	if got != entry {
		t.Errorf("Expected Logger from context")
	}
}

func TestEmptyContext(t *testing.T) {
	got := FromContext(context.Background())
	if got != L {
		t.Errorf("Expected default Logger from context")
	}
}

func TestRequest(t *testing.T) {
	entry := logrus.NewEntry(logrus.StandardLogger())

	ctx := WithContext(context.Background(), entry)
	req := new(http.Request)
	req = req.WithContext(ctx)

	got := FromRequest(req)

	if got != entry {
		t.Errorf("Expected Logger from http.Request")
	}
}
