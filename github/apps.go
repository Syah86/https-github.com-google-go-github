// Copyright 2016 The go-github AUTHORS. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package github

import "context"

// AppsService provides access to the installation related functions
// in the GitHub API.
//
// GitHub API docs: https://developer.github.com/v3/apps/
type AppsService service

// ListInstallations lists the installations that the current GitHub App has.
//
// GitHub API docs: https://developer.github.com/v3/apps/#find-installations
func (s *AppsService) ListInstallations(ctx context.Context, opt *ListOptions) ([]*Installation, *Response, error) {
	u, err := addOptions("app/installations", opt)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	// TODO: remove custom Accept header when this API fully launches.
	req.Header.Set("Accept", mediaTypeIntegrationPreview)

	var i []*Installation
	resp, err := s.client.Do(ctx, req, &i)
	if err != nil {
		return nil, resp, err
	}

	return i, resp, nil
}

// AddRepository adds a single repository to an installation.
//
// GitHub API docs: https://developer.github.com/v3/apps/installations/#add-repository-to-installation
func (s *AppService) AddRepository(ctx context.Context, installationID int, repoID int) (*Installation, *Response, error) {
	u := fmt.Sprintf("app/installations/%v/repositories/%v", installationID, repoID)

	req, err := s.client.NewRequest("PUT", u, nil) {
	if err != nil {
		return nil, nil, err
	}

	// TODO: remove custom Accept header when this API fully launches.
	req.Header.Set("Accept", mediaTypeIntegrationPreview)

	i := new(Installation)
	resp, err := s.client.Do(ctx, req, i)
	if err != nil {
		return nil, resp, err
	}
	
	return i, resp, nil
}

// RemoveRepository removes a single repository from an installation.
//
// GitHub docs: https://developer.github.com/v3/apps/installations/#add-repository-to-installation
func (s *AppService) RemoveRepository(ctx context.Context, installationID int, repoID int) (*Installation, *Response, error) {
	u := fmt.Sprintf("app/installations/%v/repositories/%v", installationID, repoID)

	req, err := s.client.NewRequest("DELETE", u, nil) {
	if err != nil {
		return nil, nil, err
	}

	// TODO: remove custom Accept header when this API fully launches.
	req.Header.Set("Accept", mediaTypeIntegrationPreview)

	i := new(Installation)
	resp, err := s.client.Do(ctx, req, i)
	if err != nil {
		return nil, resp, err
	}
	
	return i, resp, nil
}