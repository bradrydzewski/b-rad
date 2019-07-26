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

package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/{{github}}/types"
)

// ensure HTTPClient implements Client interface.
var _ Client = (*HTTPClient)(nil)

// HTTPClient provides an HTTP client for interacting
// with the remote API.
type HTTPClient struct {
	client *http.Client
	base   string
	token  string
	debug  bool
}

// New returns a client at the specified url.
func New(uri string) *HTTPClient {
	return NewToken(uri, "")
}

// NewToken returns a client at the specified url that
// authenticates all outbound requests with the given token.
func NewToken(uri, token string) *HTTPClient {
	return &HTTPClient{http.DefaultClient, uri, token, false}
}

// SetClient sets the default http client. This can be
// used in conjunction with golang.org/x/oauth2 to
// authenticate requests to the server.
func (c *HTTPClient) SetClient(client *http.Client) {
	c.client = client
}

// SetDebug sets the debug flag. When the debug flag is
// true, the http.Resposne body to stdout which can be
// helpful when debugging.
func (c *HTTPClient) SetDebug(debug bool) {
	c.debug = debug
}

// Login authenticates the user and returns a JWT token.
func (c *HTTPClient) Login(username, password string) (*types.Token, error) {
	form := &url.Values{}
	form.Add("username", username)
	form.Add("password", password)
	out := new(types.UserToken)
	uri := fmt.Sprintf("%s/api/v1/login", c.base)
	err := c.post(uri, form, out)
	return out.Token, err
}

// Register registers a new  user and returns a JWT token.
func (c *HTTPClient) Register(username, password string) (*types.Token, error) {
	form := &url.Values{}
	form.Add("username", username)
	form.Add("password", password)
	out := new(types.UserToken)
	uri := fmt.Sprintf("%s/api/v1/register", c.base)
	err := c.post(uri, form, out)
	return out.Token, err
}

//
// User Endpoints
//

// Self returns the currently authenticated user.
func (c *HTTPClient) Self() (*types.User, error) {
	out := new(types.User)
	uri := fmt.Sprintf("%s/api/v1/user", c.base)
	err := c.get(uri, out)
	return out, err
}

// Token returns an oauth2 bearer token for the currently
// authenticated user.
func (c *HTTPClient) Token() (*types.Token, error) {
	out := new(types.Token)
	uri := fmt.Sprintf("%s/api/v1/user/token", c.base)
	err := c.post(uri, nil, out)
	return out, err
}

// User returns a user by ID or email.
func (c *HTTPClient) User(key string) (*types.User, error) {
	out := new(types.User)
	uri := fmt.Sprintf("%s/api/v1/users/%s", c.base, key)
	err := c.get(uri, out)
	return out, err
}

// UserList returns a list of all registered users.
func (c *HTTPClient) UserList() ([]*types.User, error) {
	out := []*types.User{}
	uri := fmt.Sprintf("%s/api/v1/users", c.base)
	err := c.get(uri, &out)
	return out, err
}

// UserCreate creates a new user account.
func (c *HTTPClient) UserCreate(user *types.User) (*types.User, error) {
	out := new(types.User)
	uri := fmt.Sprintf("%s/api/v1/users", c.base)
	err := c.post(uri, user, out)
	return out, err
}

// UserUpdate updates a user account by ID or email.
func (c *HTTPClient) UserUpdate(key string, user *types.UserInput) (*types.User, error) {
	out := new(types.User)
	uri := fmt.Sprintf("%s/api/v1/users/%s", c.base, key)
	err := c.patch(uri, user, out)
	return out, err
}

// UserDelete deletes a user account by ID or email.
func (c *HTTPClient) UserDelete(key string) error {
	uri := fmt.Sprintf("%s/api/v1/users/%s", c.base, key)
	err := c.delete(uri)
	return err
}

//
// Project endpoints
//

// Project returns a project by ID.
func (c *HTTPClient) Project(id int64) (*types.Project, error) {
	out := new(types.Project)
	uri := fmt.Sprintf("%s/api/v1/projects/%d", c.base, id)
	err := c.get(uri, out)
	return out, err
}

// ProjectList returns a list of all projects.
func (c *HTTPClient) ProjectList() ([]*types.Project, error) {
	out := []*types.Project{}
	uri := fmt.Sprintf("%s/api/v1/user/projects", c.base)
	err := c.get(uri, &out)
	return out, err
}

// ProjectCreate creates a new project.
func (c *HTTPClient) ProjectCreate(project *types.Project) (*types.Project, error) {
	out := new(types.Project)
	uri := fmt.Sprintf("%s/api/v1/projects", c.base)
	err := c.post(uri, project, out)
	return out, err
}

// ProjectUpdate updates a project.
func (c *HTTPClient) ProjectUpdate(id int64, user *types.ProjectInput) (*types.Project, error) {
	out := new(types.Project)
	uri := fmt.Sprintf("%s/api/v1/projects", c.base)
	err := c.patch(uri, user, out)
	return out, err
}

// ProjectDelete deletes a project.
func (c *HTTPClient) ProjectDelete(id int64) error {
	uri := fmt.Sprintf("%s/api/v1/projects/%d", c.base, id)
	err := c.delete(uri)
	return err
}

//
// Membership endpoints
//

// Member returns a membrer by ID.
func (c *HTTPClient) Member(project int64, user string) (*types.Member, error) {
	out := new(types.Member)
	uri := fmt.Sprintf("%s/api/v1/projects/%d/members/%s", c.base, project, user)
	err := c.get(uri, out)
	return out, err
}

// MemberList returns a list of all project members.
func (c *HTTPClient) MemberList(project int64) ([]*types.Member, error) {
	out := []*types.Member{}
	uri := fmt.Sprintf("%s/api/v1/projects/%d/members", c.base, project)
	err := c.get(uri, &out)
	return out, err
}

// MemberCreate creates a new project member.
func (c *HTTPClient) MemberCreate(member *types.MembershipInput) (*types.Member, error) {
	out := new(types.Member)
	uri := fmt.Sprintf("%s/api/v1/projects/%d/members/%s", c.base, member.Project, member.User)
	err := c.post(uri, member, out)
	return out, err
}

// MemberUpdate updates a project member.
func (c *HTTPClient) MemberUpdate(member *types.MembershipInput) (*types.Member, error) {
	out := new(types.Member)
	uri := fmt.Sprintf("%s/api/v1/projects/%d/members/%s", c.base, member.Project, member.User)
	err := c.patch(uri, member, out)
	return out, err
}

// MemberDelete deletes a project member.
func (c *HTTPClient) MemberDelete(project int64, user string) error {
	uri := fmt.Sprintf("%s/api/v1/projects/%d/members/%s", c.base, project, user)
	err := c.delete(uri)
	return err
}

//
// {{parent}} endpoints
//

// {{parent}} returns a {{toLower parent}} by ID.
func (c *HTTPClient) {{parent}}(project, id int64) (*types.{{parent}}, error) {
	out := new(types.{{parent}})
	uri := fmt.Sprintf("%s/api/v1/projects/%d/{{toLower parent}}s/%d", c.base, project, id)
	err := c.get(uri, out)
	return out, err
}

// {{parent}}List returns a list of all {{toLower parent}}s by project id.
func (c *HTTPClient) {{parent}}List(project int64) ([]*types.{{parent}}, error) {
	out := []*types.{{parent}}{}
	uri := fmt.Sprintf("%s/api/v1/projects/%d/{{toLower parent}}s", c.base, project)
	err := c.get(uri, &out)
	return out, err
}

// {{parent}}Create creates a new {{toLower parent}}.
func (c *HTTPClient) {{parent}}Create(project int64, {{toLower parent}} *types.{{parent}}) (*types.{{parent}}, error) {
	out := new(types.{{parent}})
	uri := fmt.Sprintf("%s/api/v1/projects/%d/{{toLower parent}}s", c.base, project)
	err := c.post(uri, {{toLower parent}}, out)
	return out, err
}

// {{parent}}Update updates a {{toLower parent}}.
func (c *HTTPClient) {{parent}}Update(project, id int64, {{toLower parent}} *types.{{parent}}Input) (*types.{{parent}}, error) {
	out := new(types.{{parent}})
	uri := fmt.Sprintf("%s/api/v1/projects/%d/{{toLower parent}}s/%d", c.base, project, id)
	err := c.patch(uri, {{toLower parent}}, out)
	return out, err
}

// {{parent}}Delete deletes a {{toLower parent}}.
func (c *HTTPClient) {{parent}}Delete(project, id int64) error {
	uri := fmt.Sprintf("%s/api/v1/projects/%d/{{toLower parent}}s/%d", c.base, project, id)
	err := c.delete(uri)
	return err
}

//
// {{child}} endpoints
//

// {{child}} returns a {{toLower child}} by ID.
func (c *HTTPClient) {{child}}(project, {{toLower parent}}, {{toLower child}} int64) (*types.{{child}}, error) {
	out := new(types.{{child}})
	uri := fmt.Sprintf("%s/api/v1/projects/%d/{{toLower parent}}s/%d/{{toLower child}}s/%d", c.base, project, {{toLower parent}}, {{toLower child}})
	err := c.get(uri, out)
	return out, err
}

// {{child}}List returns a list of all {{toLower child}}s by project id.
func (c *HTTPClient) {{child}}List(project, {{toLower parent}} int64) ([]*types.{{child}}, error) {
	out := []*types.{{child}}{}
	uri := fmt.Sprintf("%s/api/v1/projects/%d/{{toLower parent}}s/%d/{{toLower child}}s", c.base, project, {{toLower parent}})
	err := c.get(uri, &out)
	return out, err
}

// {{child}}Create creates a new {{toLower child}}.
func (c *HTTPClient) {{child}}Create(project, {{toLower parent}} int64, input *types.{{child}}) (*types.{{child}}, error) {
	out := new(types.{{child}})
	uri := fmt.Sprintf("%s/api/v1/projects/%d/{{toLower parent}}s/%d/{{toLower child}}s", c.base, project, {{toLower parent}})
	err := c.post(uri, input, out)
	return out, err
}

// {{child}}Update updates a {{toLower child}}.
func (c *HTTPClient) {{child}}Update(project, {{toLower parent}}, {{toLower child}} int64, input *types.{{child}}Input) (*types.{{child}}, error) {
	out := new(types.{{child}})
	uri := fmt.Sprintf("%s/api/v1/projects/%d/{{toLower parent}}s/%d/{{toLower child}}s/%d", c.base, project, {{toLower parent}}, {{toLower child}})
	err := c.patch(uri, input, out)
	return out, err
}

// {{child}}Delete deletes a {{toLower child}}.
func (c *HTTPClient) {{child}}Delete(project, {{toLower parent}}, {{toLower child}} int64) error {
	uri := fmt.Sprintf("%s/api/v1/projects/%d/{{toLower parent}}s/%d/{{toLower child}}s/%d", c.base, project, {{toLower parent}}, {{toLower child}})
	err := c.delete(uri)
	return err
}

//
// http request helper functions
//

// helper function for making an http GET request.
func (c *HTTPClient) get(rawurl string, out interface{}) error {
	return c.do(rawurl, "GET", nil, out)
}

// helper function for making an http POST request.
func (c *HTTPClient) post(rawurl string, in, out interface{}) error {
	return c.do(rawurl, "POST", in, out)
}

// helper function for making an http PUT request.
func (c *HTTPClient) put(rawurl string, in, out interface{}) error {
	return c.do(rawurl, "PUT", in, out)
}

// helper function for making an http PATCH request.
func (c *HTTPClient) patch(rawurl string, in, out interface{}) error {
	return c.do(rawurl, "PATCH", in, out)
}

// helper function for making an http DELETE request.
func (c *HTTPClient) delete(rawurl string) error {
	return c.do(rawurl, "DELETE", nil, nil)
}

// helper function to make an http request
func (c *HTTPClient) do(rawurl, method string, in, out interface{}) error {
	// executes the http request and returns the body as
	// and io.ReadCloser
	body, err := c.stream(rawurl, method, in, out)
	if body != nil {
		defer body.Close()
	}
	if err != nil {
		return err
	}

	// if a json response is expected, parse and return
	// the json response.
	if out != nil {
		return json.NewDecoder(body).Decode(out)
	}
	return nil
}

// helper function to stream an http request
func (c *HTTPClient) stream(rawurl, method string, in, out interface{}) (io.ReadCloser, error) {
	uri, err := url.Parse(rawurl)
	if err != nil {
		return nil, err
	}

	// if we are posting or putting data, we need to
	// write it to the body of the request.
	var buf io.ReadWriter
	if in != nil {
		buf = new(bytes.Buffer)
		// if posting form data, encode the form values.
		if form, ok := in.(*url.Values); ok {
			io.WriteString(buf, form.Encode())
		} else {
			if err := json.NewEncoder(buf).Encode(in); err != nil {
				return nil, err
			}
		}
	}

	// creates a new http request.
	req, err := http.NewRequest(method, uri.String(), buf)
	if err != nil {
		return nil, err
	}
	if in != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	if c.token != "" {
		req.Header.Set("Authorization", "Bearer "+c.token)
	}
	if _, ok := in.(*url.Values); ok {
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	}
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	if c.debug {
		dump, _ := httputil.DumpResponse(resp, true)
		fmt.Println(method, rawurl)
		fmt.Println(string(dump))
	}
	if resp.StatusCode > 299 {
		defer resp.Body.Close()
		err := new(remoteError)
		json.NewDecoder(resp.Body).Decode(err)
		return nil, err
	}
	return resp.Body, nil
}
