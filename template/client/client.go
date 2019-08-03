// Copyright 2019 Brad Rydzewski. All rights reserved.
// Use of this source code is governed by the Polyform License
// that can be found in the LICENSE.md file.

package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/{{toLower repo}}/types"
	"github.com/{{toLower repo}}/version"
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
func (c *HTTPClient) UserList(params types.Params) ([]*types.User, error) {
	out := []*types.User{}
	uri := fmt.Sprintf("%s/api/v1/users?page=%d&per_page=%d", c.base, params.Page, params.Size)
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
// {{title project}} endpoints
//

//
// {{title project}} endpoints
//

// {{title project}} returns a {{toLower project}} by slug.
func (c *HTTPClient) {{title project}}(slug string) (*types.{{title project}}, error) {
	out := new(types.{{title project}})
	uri := fmt.Sprintf("%s/api/v1/{{toLower project}}s/%s", c.base, slug)
	err := c.get(uri, out)
	return out, err
}

// {{title project}}List returns a list of all {{toLower project}}s.
func (c *HTTPClient) {{title project}}List(params types.Params) ([]*types.{{title project}}, error) {
	out := []*types.{{title project}}{}
	uri := fmt.Sprintf("%s/api/v1/user/{{toLower project}}s?page=%dper_page=%d", c.base, params.Page, params.Size)
	err := c.get(uri, &out)
	return out, err
}

// {{title project}}Create creates a new {{toLower project}}.
func (c *HTTPClient) {{title project}}Create({{toLower project}} *types.{{title project}}) (*types.{{title project}}, error) {
	out := new(types.{{title project}})
	uri := fmt.Sprintf("%s/api/v1/{{toLower project}}s", c.base)
	err := c.post(uri, {{toLower project}}, out)
	return out, err
}

// {{title project}}Update updates a {{toLower project}}.
func (c *HTTPClient) {{title project}}Update(key string, user *types.{{title project}}Input) (*types.{{title project}}, error) {
	out := new(types.{{title project}})
	uri := fmt.Sprintf("%s/api/v1/{{toLower project}}s/%s", c.base, key)
	err := c.patch(uri, user, out)
	return out, err
}

// {{title project}}Delete deletes a {{toLower project}}.
func (c *HTTPClient) {{title project}}Delete(key string) error {
	uri := fmt.Sprintf("%s/api/v1/{{toLower project}}s/%s", c.base, key)
	err := c.delete(uri)
	return err
}

//
// Membership endpoints
//

// Member returns a membrer.
func (c *HTTPClient) Member({{toLower project}}, user string) (*types.Member, error) {
	out := new(types.Member)
	uri := fmt.Sprintf("%s/api/v1/{{toLower project}}s/%s/members/%s", c.base, {{toLower project}}, user)
	err := c.get(uri, out)
	return out, err
}

// MemberList returns a list of all {{toLower project}} members.
func (c *HTTPClient) MemberList({{toLower project}} string, params types.Params) ([]*types.Member, error) {
	out := []*types.Member{}
	uri := fmt.Sprintf("%s/api/v1/{{toLower project}}s/%s/members?page=%d&per_page=%d", c.base, {{toLower project}}, params.Page, params.Size)
	err := c.get(uri, &out)
	return out, err
}

// MemberCreate creates a new {{toLower project}} member.
func (c *HTTPClient) MemberCreate(member *types.MembershipInput) (*types.Member, error) {
	out := new(types.Member)
	uri := fmt.Sprintf("%s/api/v1/{{toLower project}}s/%s/members/%s", c.base, member.{{title project}}, member.User)
	err := c.post(uri, member, out)
	return out, err
}

// MemberUpdate updates a {{toLower project}} member.
func (c *HTTPClient) MemberUpdate(member *types.MembershipInput) (*types.Member, error) {
	out := new(types.Member)
	uri := fmt.Sprintf("%s/api/v1/{{toLower project}}s/%s/members/%s", c.base, member.{{title project}}, member.User)
	err := c.patch(uri, member, out)
	return out, err
}

// MemberDelete deletes a {{toLower project}} member.
func (c *HTTPClient) MemberDelete({{toLower project}}, user string) error {
	uri := fmt.Sprintf("%s/api/v1/{{toLower project}}s/%s/members/%s", c.base, {{toLower project}}, user)
	err := c.delete(uri)
	return err
}

//
// {{title parent}} endpoints
//

// {{title parent}} returns a {{toLower parent}} by ID.
func (c *HTTPClient) {{title parent}}({{toLower project}}, slug string) (*types.{{title parent}}, error) {
	out := new(types.{{title parent}})
	uri := fmt.Sprintf("%s/api/v1/{{toLower project}}s/%s/{{toLower parent}}s/%s", c.base, {{toLower project}}, slug)
	err := c.get(uri, out)
	return out, err
}

// {{title parent}}List returns a list of all {{toLower parent}}s by {{toLower project}} id.
func (c *HTTPClient) {{title parent}}List({{toLower project}} string, params types.Params) ([]*types.{{title parent}}, error) {
	out := []*types.{{title parent}}{}
	uri := fmt.Sprintf("%s/api/v1/{{toLower project}}s/%s/{{toLower parent}}s?page=%dper_page=%d", c.base, {{toLower project}}, params.Page, params.Size)
	err := c.get(uri, &out)
	return out, err
}

// {{title parent}}Create creates a new {{toLower parent}}.
func (c *HTTPClient) {{title parent}}Create({{toLower project}} string, {{toLower parent}} *types.{{title parent}}) (*types.{{title parent}}, error) {
	out := new(types.{{title parent}})
	uri := fmt.Sprintf("%s/api/v1/{{toLower project}}s/%s/{{toLower parent}}s", c.base, {{toLower project}})
	err := c.post(uri, {{toLower parent}}, out)
	return out, err
}

// {{title parent}}Update updates a {{toLower parent}}.
func (c *HTTPClient) {{title parent}}Update({{toLower project}}, slug string, {{toLower parent}} *types.{{title parent}}Input) (*types.{{title parent}}, error) {
	out := new(types.{{title parent}})
	uri := fmt.Sprintf("%s/api/v1/{{toLower project}}s/%s/{{toLower parent}}s/%s", c.base, {{toLower project}}, slug)
	err := c.patch(uri, {{toLower parent}}, out)
	return out, err
}

// {{title parent}}Delete deletes a {{toLower parent}}.
func (c *HTTPClient) {{title parent}}Delete({{toLower project}}, slug string) error {
	uri := fmt.Sprintf("%s/api/v1/{{toLower project}}s/%s/{{toLower parent}}s/%s", c.base, {{toLower project}}, slug)
	err := c.delete(uri)
	return err
}

//
// {{title child}} endpoints
//

// {{title child}} returns a {{toLower child}} by ID.
func (c *HTTPClient) {{title child}}({{toLower project}}, {{toLower parent}}, {{toLower child}} string) (*types.{{title child}}, error) {
	out := new(types.{{title child}})
	uri := fmt.Sprintf("%s/api/v1/{{toLower project}}s/%s/{{toLower parent}}s/%s/{{toLower child}}s/%s", c.base, {{toLower project}}, {{toLower parent}}, {{toLower child}})
	err := c.get(uri, out)
	return out, err
}

// {{title child}}List returns a list of all {{toLower child}}s by {{toLower project}} id.
func (c *HTTPClient) {{title child}}List({{toLower project}}, {{toLower parent}} string, params types.Params) ([]*types.{{title child}}, error) {
	out := []*types.{{title child}}{}
	uri := fmt.Sprintf("%s/api/v1/{{toLower project}}s/%s/{{toLower parent}}s/%s/{{toLower child}}s?page=%dper_page=%d", c.base, {{toLower project}}, {{toLower parent}}, params.Page, params.Size)
	err := c.get(uri, &out)
	return out, err
}

// {{title child}}Create creates a new {{toLower child}}.
func (c *HTTPClient) {{title child}}Create({{toLower project}}, {{toLower parent}} string, input *types.{{title child}}) (*types.{{title child}}, error) {
	out := new(types.{{title child}})
	uri := fmt.Sprintf("%s/api/v1/{{toLower project}}s/%s/{{toLower parent}}s/%s/{{toLower child}}s", c.base, {{toLower project}}, {{toLower parent}})
	err := c.post(uri, input, out)
	return out, err
}

// {{title child}}Update updates a {{toLower child}}.
func (c *HTTPClient) {{title child}}Update({{toLower project}}, {{toLower parent}}, {{toLower child}} string, input *types.{{title child}}Input) (*types.{{title child}}, error) {
	out := new(types.{{title child}})
	uri := fmt.Sprintf("%s/api/v1/{{toLower project}}s/%s/{{toLower parent}}s/%s/{{toLower child}}s/%s", c.base, {{toLower project}}, {{toLower parent}}, {{toLower child}})
	err := c.patch(uri, input, out)
	return out, err
}

// {{title child}}Delete deletes a {{toLower child}}.
func (c *HTTPClient) {{title child}}Delete({{toLower project}}, {{toLower parent}}, {{toLower child}} string) error {
	uri := fmt.Sprintf("%s/api/v1/{{toLower project}}s/%s/{{toLower parent}}s/%s/{{toLower child}}s/%s", c.base, {{toLower project}}, {{toLower parent}}, {{toLower child}})
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

	// include the client version information in the
	// http accept header for debugging purposes.
	req.Header.Set("Accept", "application/json;version="+version.Version.String())

	// send the http request.
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
