// Copyright 2019 Brad Rydzewski. All rights reserved.
// Use of this source code is governed by the Polyform License
// that can be found in the LICENSE.md file.

// Package mocks provides mock interfaces.
package mocks

//go:generate mockgen -package=mocks -destination=mock_store.go github.com/{{toLower repo}}/internal/store {{title child}}Store,{{title parent}}Store,MemberStore,{{title project}}Store,SystemStore,UserStore
//go:generate mockgen -package=mocks -destination=mock_client.go github.com/{{toLower repo}}/client Client
