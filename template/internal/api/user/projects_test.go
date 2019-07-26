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

package user

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/{{github}}/internal/api/render"
	"github.com/{{github}}/internal/api/request"
	"github.com/{{github}}/mocks"
	"github.com/{{github}}/types"

	"github.com/golang/mock/gomock"
	"github.com/google/go-cmp/cmp"
)

func TestProjectList(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	mockUser := &types.User{
		ID:    1,
		Email: "octocat@github.com",
	}

	mockList := []*types.Project{
		{
			Name: "test",
			Desc: "desc",
		},
	}

	projs := mocks.NewMockProjectStore(controller)
	projs.EXPECT().List(gomock.Any(), mockUser.ID, gomock.Any()).Return(mockList, nil)

	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/", nil)
	r = r.WithContext(
		request.WithUser(r.Context(), mockUser),
	)

	HandleProjects(projs)(w, r)
	if got, want := w.Code, http.StatusOK; want != got {
		t.Errorf("Want response code %d, got %d", want, got)
	}

	got, want := []*types.Project{}, mockList
	json.NewDecoder(w.Body).Decode(&got)
	if diff := cmp.Diff(got, want); len(diff) > 0 {
		t.Errorf(diff)
	}
}

func TestProjectListErr(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	mockUser := &types.User{
		ID:    1,
		Email: "octocat@github.com",
	}

	projs := mocks.NewMockProjectStore(controller)
	projs.EXPECT().List(gomock.Any(), mockUser.ID, gomock.Any()).Return(nil, render.ErrNotFound)

	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/", nil)
	r = r.WithContext(
		request.WithUser(r.Context(), mockUser),
	)

	HandleProjects(projs)(w, r)
	if got, want := w.Code, http.StatusInternalServerError; want != got {
		t.Errorf("Want response code %d, got %d", want, got)
	}

	got, want := &render.Error{}, render.ErrNotFound
	json.NewDecoder(w.Body).Decode(got)
	if diff := cmp.Diff(got, want); len(diff) > 0 {
		t.Errorf(diff)
	}
}
