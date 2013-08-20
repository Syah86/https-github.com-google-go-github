// Copyright 2013 The go-github AUTHORS. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package github

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"
	"time"
)

// Event represents a GitHub event.
type Event struct {
	Type       string          `json:"type,omitempty"`
	Public     bool            `json:"public"`
	RawPayload json.RawMessage `json:"payload,omitempty"`
	Repo       *Repository     `json:"repo,omitempty"`
	Actor      *User           `json:"actor,omitempty"`
	Org        *Organization   `json:"org,omitempty"`
	CreatedAt  *time.Time      `json:"created_at,omitempty"`
	ID         string          `json:"id,omitempty"`
}

// Payload returns the parsed event payload. For recognized event types
// (PushEvent), a value of the corresponding struct type will be returned.
func (e *Event) Payload() (payload interface{}) {
	switch e.Type {
	case "PushEvent":
		payload = &PushEvent{}
	}
	if err := json.Unmarshal(e.RawPayload, &payload); err != nil {
		panic(err.Error())
	}
	return payload
}

// PushEvent represents a git push to a GitHub repository.
//
// GitHub API docs: http://developer.github.com/v3/activity/events/types/#pushevent
type PushEvent struct {
	PushID  int               `json:"push_id,omitempty"`
	Head    string            `json:"head,omitempty"`
	Ref     string            `json:"ref,omitempty"`
	Size    int               `json:"ref,omitempty"`
	Commits []PushEventCommit `json:"commits,omitempty"`
}

// PushEventCommit represents a git commit in a GitHub PushEvent.
type PushEventCommit struct {
	SHA      string        `json:"sha,omitempty"`
	Message  string        `json:"message,omitempty"`
	Author   *CommitAuthor `json:"author,omitempty"`
	URL      string        `json:"url,omitempty"`
	Distinct bool          `json:"distinct"`
}

// List public events.
//
// GitHub API docs: http://developer.github.com/v3/activity/events/#list-public-events
func (s *ActivityService) ListPublicEvents(opt *ListOptions) ([]Event, *Response, error) {
	u := "events"
	if opt != nil {
		params := url.Values{
			"page": []string{strconv.Itoa(opt.Page)},
		}
		u += "?" + params.Encode()
	}

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	events := new([]Event)
	resp, err := s.client.Do(req, events)
	return *events, resp, err
}

// List repository events.
//
// GitHub API docs: http://developer.github.com/v3/activity/events/#list-repository-events
func (s *ActivityService) ListRepositoryEvents(owner, repo string, opt *ListOptions) ([]Event, *Response, error) {
	u := fmt.Sprintf("repos/%v/%v/events", owner, repo)
	if opt != nil {
		params := url.Values{
			"page": []string{strconv.Itoa(opt.Page)},
		}
		u += "?" + params.Encode()
	}

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	events := new([]Event)
	resp, err := s.client.Do(req, events)
	return *events, resp, err
}

// List issue events for a repository.
//
// GitHub API docs: http://developer.github.com/v3/activity/events/#list-issue-events-for-a-repository
func (s *ActivityService) ListIssueEventsForRepository(owner, repo string, opt *ListOptions) ([]Event, *Response, error) {
	u := fmt.Sprintf("repos/%v/%v/issues/events", owner, repo)
	if opt != nil {
		params := url.Values{
			"page": []string{strconv.Itoa(opt.Page)},
		}
		u += "?" + params.Encode()
	}

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	events := new([]Event)
	resp, err := s.client.Do(req, events)
	return *events, resp, err
}

// List public events for a network of repositories
//
// GitHub API docs: http://developer.github.com/v3/activity/events/#list-public-events-for-a-network-of-repositories
func (s *ActivityService) ListEventsForRepoNetwork(owner, repo string, opt *ListOptions) ([]Event, *Response, error) {
	u := fmt.Sprintf("networks/%v/%v/events", owner, repo)
	if opt != nil {
		params := url.Values{
			"page": []string{strconv.Itoa(opt.Page)},
		}
		u += "?" + params.Encode()
	}

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	events := new([]Event)
	resp, err := s.client.Do(req, events)
	return *events, resp, err
}

// List public events for an organization
//
// GitHub API docs: http://developer.github.com/v3/activity/events/#list-public-events-for-an-organization
func (s *ActivityService) ListEventsForOrganization(org string, opt *ListOptions) ([]Event, *Response, error) {
	u := fmt.Sprintf("orgs/%v/events", org)
	if opt != nil {
		params := url.Values{
			"page": []string{strconv.Itoa(opt.Page)},
		}
		u += "?" + params.Encode()
	}

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	events := new([]Event)
	resp, err := s.client.Do(req, events)
	return *events, resp, err
}

// ListEventsPerformedByUser lists the events performed by a user. If publicOnly is
// true, only public events will be returned.
//
// GitHub API docs: http://developer.github.com/v3/activity/events/#list-events-performed-by-a-user
func (s *ActivityService) ListEventsPerformedByUser(user string, publicOnly bool, opt *ListOptions) ([]Event, *Response, error) {
	var u string
	if publicOnly {
		u = fmt.Sprintf("users/%v/events/public", user)
	} else {
		u = fmt.Sprintf("users/%v/events", user)
	}

	if opt != nil {
		params := url.Values{
			"page": []string{strconv.Itoa(opt.Page)},
		}
		u += "?" + params.Encode()
	}

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	events := new([]Event)
	resp, err := s.client.Do(req, events)
	return *events, resp, err
}

// ListEventsRecievedByUser lists the events recieved by a user. If publicOnly is
// true, only public events will be returned.
//
// GitHub API docs: http://developer.github.com/v3/activity/events/#list-events-that-a-user-has-received
func (s *ActivityService) ListEventsRecievedByUser(user string, publicOnly bool, opt *ListOptions) ([]Event, *Response, error) {
	var u string
	if publicOnly {
		u = fmt.Sprintf("users/%v/received_events/public", user)
	} else {
		u = fmt.Sprintf("users/%v/received_events", user)
	}

	if opt != nil {
		params := url.Values{
			"page": []string{strconv.Itoa(opt.Page)},
		}
		u += "?" + params.Encode()
	}

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	events := new([]Event)
	resp, err := s.client.Do(req, events)
	return *events, resp, err
}

// ListEventsForOrganization provides the user’s organization dashboard. You
// must be authenticated as the user to view this.
//
// GitHub API docs: http://developer.github.com/v3/activity/events/#list-events-for-an-organization
func (s *ActivityService) ListEventsForOrganization(org, user string, opt *ListOptions) ([]Event, *Response, error) {
	u := fmt.Sprintf("users/%v/events/orgs/%v", user, org)
	if opt != nil {
		params := url.Values{
			"page": []string{strconv.Itoa(opt.Page)},
		}
		u += "?" + params.Encode()
	}
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	events := new([]Event)
	resp, err := s.client.Do(req, events)
	return *events, resp, err
}
