// Copyright 2013 The go-github AUTHORS. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package github

import (
	"fmt"
	"net/url"
	"time"
)

// PullRequestsService handles communication with the pull request related
// methods of the GitHub API.
//
// GitHub API docs: http://developer.github.com/v3/pulls/
type PullRequestsService struct {
	client *Client
}

// PullRequest represents a GitHub pull request on a repository.
type PullRequest struct {
	Number       int        `json:"number,omitempty"`
	State        string     `json:"state,omitempty"`
	Title        string     `json:"title,omitempty"`
	Body         string     `json:"body,omitempty"`
	CreatedAt    *time.Time `json:"created_at,omitempty"`
	UpdatedAt    *time.Time `json:"updated_at,omitempty"`
	ClosedAt     *time.Time `json:"closed_at,omitempty"`
	MergedAt     *time.Time `json:"merged_at,omitempty"`
	User         *User      `json:"user,omitempty"`
	Merged       bool       `json:"merged,omitempty"`
	Mergeable    bool       `json:"mergeable,omitempty"`
	MergedBy     *User      `json:"merged_by,omitempty"`
	Comments     int        `json:"comments,omitempty"`
	Commits      int        `json:"commits,omitempty"`
	Additions    int        `json:"additions,omitempty"`
	Deletions    int        `json:"deletions,omitempty"`
	ChangedFiles int        `json:"changed_files,omitempty"`

	// TODO(willnorris): add head and base once we have a Commit struct defined somewhere
}

// PullRequestListOptions specifies the optional parameters to the
// PullRequestsService.List method.
type PullRequestListOptions struct {
	// State filters pull requests based on their state.  Possible values are:
	// open, closed.  Default is "open".
	State string

	// Head filters pull requests by head user and branch name in the format of:
	// "user:ref-name".
	Head string

	// Base filters pull requests by base branch name.
	Base string
}

// List the pull requests for the specified repository.
//
// GitHub API docs: http://developer.github.com/v3/pulls/#list-pull-requests
func (s *PullRequestsService) List(owner string, repo string, opt *PullRequestListOptions) ([]PullRequest, error) {
	u := fmt.Sprintf("repos/%v/%v/pulls", owner, repo)
	if opt != nil {
		params := url.Values{
			"state": {opt.State},
			"head":  {opt.Head},
			"base":  {opt.Base},
		}
		u += "?" + params.Encode()
	}

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}

	pulls := new([]PullRequest)
	_, err = s.client.Do(req, pulls)
	return *pulls, err
}

// Get a single pull request.
//
// GitHub API docs: https://developer.github.com/v3/pulls/#get-a-single-pull-request
func (s *PullRequestsService) Get(owner string, repo string, number int) (*PullRequest, error) {
	u := fmt.Sprintf("repos/%v/%v/pulls/%d", owner, repo, number)
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}
	pull := new(PullRequest)
	_, err = s.client.Do(req, pull)
	return pull, err
}

// Create a new pull request on the specified repository.
//
// GitHub API docs: https://developer.github.com/v3/pulls/#create-a-pull-request
func (s *PullRequestsService) Create(owner string, repo string, pull *PullRequest) (*PullRequest, error) {
	u := fmt.Sprintf("repos/%v/%v/pulls", owner, repo)
	req, err := s.client.NewRequest("POST", u, pull)
	if err != nil {
		return nil, err
	}
	p := new(PullRequest)
	_, err = s.client.Do(req, p)
	return p, err
}

// Edit a pull request.
//
// GitHub API docs: https://developer.github.com/v3/pulls/#update-a-pull-request
func (s *PullRequestsService) Edit(owner string, repo string, number int, pull *PullRequest) (*PullRequest, error) {
	u := fmt.Sprintf("repos/%v/%v/pulls/%d", owner, repo, number)
	req, err := s.client.NewRequest("PATCH", u, pull)
	if err != nil {
		return nil, err
	}
	p := new(PullRequest)
	_, err = s.client.Do(req, p)
	return p, err
}
