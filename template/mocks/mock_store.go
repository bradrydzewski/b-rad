// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/{{toLower repo}}/internal/store (interfaces: {{title child}}Store,{{title parent}}Store,MemberStore,{{title project}}Store,SystemStore,UserStore)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	types "github.com/{{toLower repo}}/types"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// Mock{{title child}}Store is a mock of {{title child}}Store interface
type Mock{{title child}}Store struct {
	ctrl     *gomock.Controller
	recorder *Mock{{title child}}StoreMockRecorder
}

// Mock{{title child}}StoreMockRecorder is the mock recorder for Mock{{title child}}Store
type Mock{{title child}}StoreMockRecorder struct {
	mock *Mock{{title child}}Store
}

// NewMock{{title child}}Store creates a new mock instance
func NewMock{{title child}}Store(ctrl *gomock.Controller) *Mock{{title child}}Store {
	mock := &Mock{{title child}}Store{ctrl: ctrl}
	mock.recorder = &Mock{{title child}}StoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *Mock{{title child}}Store) EXPECT() *Mock{{title child}}StoreMockRecorder {
	return m.recorder
}

// Create mocks base method
func (m *Mock{{title child}}Store) Create(arg0 context.Context, arg1 *types.{{title child}}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create
func (mr *Mock{{title child}}StoreMockRecorder) Create(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*Mock{{title child}}Store)(nil).Create), arg0, arg1)
}

// Delete mocks base method
func (m *Mock{{title child}}Store) Delete(arg0 context.Context, arg1 *types.{{title child}}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete
func (mr *Mock{{title child}}StoreMockRecorder) Delete(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*Mock{{title child}}Store)(nil).Delete), arg0, arg1)
}

// Find mocks base method
func (m *Mock{{title child}}Store) Find(arg0 context.Context, arg1 int64) (*types.{{title child}}, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Find", arg0, arg1)
	ret0, _ := ret[0].(*types.{{title child}})
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Find indicates an expected call of Find
func (mr *Mock{{title child}}StoreMockRecorder) Find(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Find", reflect.TypeOf((*Mock{{title child}}Store)(nil).Find), arg0, arg1)
}

// FindSlug mocks base method
func (m *Mock{{title child}}Store) FindSlug(arg0 context.Context, arg1 int64, arg2 string) (*types.{{title child}}, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindSlug", arg0, arg1, arg2)
	ret0, _ := ret[0].(*types.{{title child}})
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindSlug indicates an expected call of FindSlug
func (mr *Mock{{title child}}StoreMockRecorder) FindSlug(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindSlug", reflect.TypeOf((*Mock{{title child}}Store)(nil).FindSlug), arg0, arg1, arg2)
}

// List mocks base method
func (m *Mock{{title child}}Store) List(arg0 context.Context, arg1 int64, arg2 types.Params) ([]*types.{{title child}}, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0, arg1, arg2)
	ret0, _ := ret[0].([]*types.{{title child}})
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List
func (mr *Mock{{title child}}StoreMockRecorder) List(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*Mock{{title child}}Store)(nil).List), arg0, arg1, arg2)
}

// Update mocks base method
func (m *Mock{{title child}}Store) Update(arg0 context.Context, arg1 *types.{{title child}}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update
func (mr *Mock{{title child}}StoreMockRecorder) Update(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*Mock{{title child}}Store)(nil).Update), arg0, arg1)
}

// Mock{{title parent}}Store is a mock of {{title parent}}Store interface
type Mock{{title parent}}Store struct {
	ctrl     *gomock.Controller
	recorder *Mock{{title parent}}StoreMockRecorder
}

// Mock{{title parent}}StoreMockRecorder is the mock recorder for Mock{{title parent}}Store
type Mock{{title parent}}StoreMockRecorder struct {
	mock *Mock{{title parent}}Store
}

// NewMock{{title parent}}Store creates a new mock instance
func NewMock{{title parent}}Store(ctrl *gomock.Controller) *Mock{{title parent}}Store {
	mock := &Mock{{title parent}}Store{ctrl: ctrl}
	mock.recorder = &Mock{{title parent}}StoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *Mock{{title parent}}Store) EXPECT() *Mock{{title parent}}StoreMockRecorder {
	return m.recorder
}

// Create mocks base method
func (m *Mock{{title parent}}Store) Create(arg0 context.Context, arg1 *types.{{title parent}}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create
func (mr *Mock{{title parent}}StoreMockRecorder) Create(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*Mock{{title parent}}Store)(nil).Create), arg0, arg1)
}

// Delete mocks base method
func (m *Mock{{title parent}}Store) Delete(arg0 context.Context, arg1 *types.{{title parent}}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete
func (mr *Mock{{title parent}}StoreMockRecorder) Delete(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*Mock{{title parent}}Store)(nil).Delete), arg0, arg1)
}

// Find mocks base method
func (m *Mock{{title parent}}Store) Find(arg0 context.Context, arg1 int64) (*types.{{title parent}}, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Find", arg0, arg1)
	ret0, _ := ret[0].(*types.{{title parent}})
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Find indicates an expected call of Find
func (mr *Mock{{title parent}}StoreMockRecorder) Find(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Find", reflect.TypeOf((*Mock{{title parent}}Store)(nil).Find), arg0, arg1)
}

// FindSlug mocks base method
func (m *Mock{{title parent}}Store) FindSlug(arg0 context.Context, arg1 int64, arg2 string) (*types.{{title parent}}, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindSlug", arg0, arg1, arg2)
	ret0, _ := ret[0].(*types.{{title parent}})
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindSlug indicates an expected call of FindSlug
func (mr *Mock{{title parent}}StoreMockRecorder) FindSlug(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindSlug", reflect.TypeOf((*Mock{{title parent}}Store)(nil).FindSlug), arg0, arg1, arg2)
}

// List mocks base method
func (m *Mock{{title parent}}Store) List(arg0 context.Context, arg1 int64, arg2 types.Params) ([]*types.{{title parent}}, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0, arg1, arg2)
	ret0, _ := ret[0].([]*types.{{title parent}})
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List
func (mr *Mock{{title parent}}StoreMockRecorder) List(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*Mock{{title parent}}Store)(nil).List), arg0, arg1, arg2)
}

// Update mocks base method
func (m *Mock{{title parent}}Store) Update(arg0 context.Context, arg1 *types.{{title parent}}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update
func (mr *Mock{{title parent}}StoreMockRecorder) Update(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*Mock{{title parent}}Store)(nil).Update), arg0, arg1)
}

// MockMemberStore is a mock of MemberStore interface
type MockMemberStore struct {
	ctrl     *gomock.Controller
	recorder *MockMemberStoreMockRecorder
}

// MockMemberStoreMockRecorder is the mock recorder for MockMemberStore
type MockMemberStoreMockRecorder struct {
	mock *MockMemberStore
}

// NewMockMemberStore creates a new mock instance
func NewMockMemberStore(ctrl *gomock.Controller) *MockMemberStore {
	mock := &MockMemberStore{ctrl: ctrl}
	mock.recorder = &MockMemberStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockMemberStore) EXPECT() *MockMemberStoreMockRecorder {
	return m.recorder
}

// Create mocks base method
func (m *MockMemberStore) Create(arg0 context.Context, arg1 *types.Membership) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create
func (mr *MockMemberStoreMockRecorder) Create(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockMemberStore)(nil).Create), arg0, arg1)
}

// Delete mocks base method
func (m *MockMemberStore) Delete(arg0 context.Context, arg1, arg2 int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete
func (mr *MockMemberStoreMockRecorder) Delete(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockMemberStore)(nil).Delete), arg0, arg1, arg2)
}

// Find mocks base method
func (m *MockMemberStore) Find(arg0 context.Context, arg1, arg2 int64) (*types.Member, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Find", arg0, arg1, arg2)
	ret0, _ := ret[0].(*types.Member)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Find indicates an expected call of Find
func (mr *MockMemberStoreMockRecorder) Find(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Find", reflect.TypeOf((*MockMemberStore)(nil).Find), arg0, arg1, arg2)
}

// List mocks base method
func (m *MockMemberStore) List(arg0 context.Context, arg1 int64, arg2 types.Params) ([]*types.Member, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0, arg1, arg2)
	ret0, _ := ret[0].([]*types.Member)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List
func (mr *MockMemberStoreMockRecorder) List(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockMemberStore)(nil).List), arg0, arg1, arg2)
}

// Update mocks base method
func (m *MockMemberStore) Update(arg0 context.Context, arg1 *types.Membership) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update
func (mr *MockMemberStoreMockRecorder) Update(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockMemberStore)(nil).Update), arg0, arg1)
}

// Mock{{title project}}Store is a mock of {{title project}}Store interface
type Mock{{title project}}Store struct {
	ctrl     *gomock.Controller
	recorder *Mock{{title project}}StoreMockRecorder
}

// Mock{{title project}}StoreMockRecorder is the mock recorder for Mock{{title project}}Store
type Mock{{title project}}StoreMockRecorder struct {
	mock *Mock{{title project}}Store
}

// NewMock{{title project}}Store creates a new mock instance
func NewMock{{title project}}Store(ctrl *gomock.Controller) *Mock{{title project}}Store {
	mock := &Mock{{title project}}Store{ctrl: ctrl}
	mock.recorder = &Mock{{title project}}StoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *Mock{{title project}}Store) EXPECT() *Mock{{title project}}StoreMockRecorder {
	return m.recorder
}

// Create mocks base method
func (m *Mock{{title project}}Store) Create(arg0 context.Context, arg1 *types.{{title project}}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create
func (mr *Mock{{title project}}StoreMockRecorder) Create(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*Mock{{title project}}Store)(nil).Create), arg0, arg1)
}

// Delete mocks base method
func (m *Mock{{title project}}Store) Delete(arg0 context.Context, arg1 *types.{{title project}}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete
func (mr *Mock{{title project}}StoreMockRecorder) Delete(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*Mock{{title project}}Store)(nil).Delete), arg0, arg1)
}

// Find mocks base method
func (m *Mock{{title project}}Store) Find(arg0 context.Context, arg1 int64) (*types.{{title project}}, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Find", arg0, arg1)
	ret0, _ := ret[0].(*types.{{title project}})
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Find indicates an expected call of Find
func (mr *Mock{{title project}}StoreMockRecorder) Find(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Find", reflect.TypeOf((*Mock{{title project}}Store)(nil).Find), arg0, arg1)
}

// FindSlug mocks base method
func (m *Mock{{title project}}Store) FindSlug(arg0 context.Context, arg1 string) (*types.{{title project}}, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindSlug", arg0, arg1)
	ret0, _ := ret[0].(*types.{{title project}})
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindSlug indicates an expected call of FindSlug
func (mr *Mock{{title project}}StoreMockRecorder) FindSlug(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindSlug", reflect.TypeOf((*Mock{{title project}}Store)(nil).FindSlug), arg0, arg1)
}

// FindToken mocks base method
func (m *Mock{{title project}}Store) FindToken(arg0 context.Context, arg1 string) (*types.{{title project}}, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindToken", arg0, arg1)
	ret0, _ := ret[0].(*types.{{title project}})
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindToken indicates an expected call of FindToken
func (mr *Mock{{title project}}StoreMockRecorder) FindToken(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindToken", reflect.TypeOf((*Mock{{title project}}Store)(nil).FindToken), arg0, arg1)
}

// List mocks base method
func (m *Mock{{title project}}Store) List(arg0 context.Context, arg1 int64, arg2 types.Params) ([]*types.{{title project}}, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0, arg1, arg2)
	ret0, _ := ret[0].([]*types.{{title project}})
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List
func (mr *Mock{{title project}}StoreMockRecorder) List(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*Mock{{title project}}Store)(nil).List), arg0, arg1, arg2)
}

// Update mocks base method
func (m *Mock{{title project}}Store) Update(arg0 context.Context, arg1 *types.{{title project}}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update
func (mr *Mock{{title project}}StoreMockRecorder) Update(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*Mock{{title project}}Store)(nil).Update), arg0, arg1)
}

// MockSystemStore is a mock of SystemStore interface
type MockSystemStore struct {
	ctrl     *gomock.Controller
	recorder *MockSystemStoreMockRecorder
}

// MockSystemStoreMockRecorder is the mock recorder for MockSystemStore
type MockSystemStoreMockRecorder struct {
	mock *MockSystemStore
}

// NewMockSystemStore creates a new mock instance
func NewMockSystemStore(ctrl *gomock.Controller) *MockSystemStore {
	mock := &MockSystemStore{ctrl: ctrl}
	mock.recorder = &MockSystemStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockSystemStore) EXPECT() *MockSystemStoreMockRecorder {
	return m.recorder
}

// Config mocks base method
func (m *MockSystemStore) Config(arg0 context.Context) *types.Config {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Config", arg0)
	ret0, _ := ret[0].(*types.Config)
	return ret0
}

// Config indicates an expected call of Config
func (mr *MockSystemStoreMockRecorder) Config(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Config", reflect.TypeOf((*MockSystemStore)(nil).Config), arg0)
}

// MockUserStore is a mock of UserStore interface
type MockUserStore struct {
	ctrl     *gomock.Controller
	recorder *MockUserStoreMockRecorder
}

// MockUserStoreMockRecorder is the mock recorder for MockUserStore
type MockUserStoreMockRecorder struct {
	mock *MockUserStore
}

// NewMockUserStore creates a new mock instance
func NewMockUserStore(ctrl *gomock.Controller) *MockUserStore {
	mock := &MockUserStore{ctrl: ctrl}
	mock.recorder = &MockUserStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockUserStore) EXPECT() *MockUserStoreMockRecorder {
	return m.recorder
}

// Count mocks base method
func (m *MockUserStore) Count(arg0 context.Context) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Count", arg0)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Count indicates an expected call of Count
func (mr *MockUserStoreMockRecorder) Count(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Count", reflect.TypeOf((*MockUserStore)(nil).Count), arg0)
}

// Create mocks base method
func (m *MockUserStore) Create(arg0 context.Context, arg1 *types.User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create
func (mr *MockUserStoreMockRecorder) Create(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockUserStore)(nil).Create), arg0, arg1)
}

// Delete mocks base method
func (m *MockUserStore) Delete(arg0 context.Context, arg1 *types.User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete
func (mr *MockUserStoreMockRecorder) Delete(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockUserStore)(nil).Delete), arg0, arg1)
}

// Find mocks base method
func (m *MockUserStore) Find(arg0 context.Context, arg1 int64) (*types.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Find", arg0, arg1)
	ret0, _ := ret[0].(*types.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Find indicates an expected call of Find
func (mr *MockUserStoreMockRecorder) Find(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Find", reflect.TypeOf((*MockUserStore)(nil).Find), arg0, arg1)
}

// FindEmail mocks base method
func (m *MockUserStore) FindEmail(arg0 context.Context, arg1 string) (*types.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindEmail", arg0, arg1)
	ret0, _ := ret[0].(*types.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindEmail indicates an expected call of FindEmail
func (mr *MockUserStoreMockRecorder) FindEmail(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindEmail", reflect.TypeOf((*MockUserStore)(nil).FindEmail), arg0, arg1)
}

// FindKey mocks base method
func (m *MockUserStore) FindKey(arg0 context.Context, arg1 string) (*types.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindKey", arg0, arg1)
	ret0, _ := ret[0].(*types.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindKey indicates an expected call of FindKey
func (mr *MockUserStoreMockRecorder) FindKey(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindKey", reflect.TypeOf((*MockUserStore)(nil).FindKey), arg0, arg1)
}

// List mocks base method
func (m *MockUserStore) List(arg0 context.Context, arg1 types.UserFilter) ([]*types.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0, arg1)
	ret0, _ := ret[0].([]*types.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List
func (mr *MockUserStoreMockRecorder) List(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockUserStore)(nil).List), arg0, arg1)
}

// Update mocks base method
func (m *MockUserStore) Update(arg0 context.Context, arg1 *types.User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update
func (mr *MockUserStoreMockRecorder) Update(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockUserStore)(nil).Update), arg0, arg1)
}
