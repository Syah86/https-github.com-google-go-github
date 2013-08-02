// Copyright 2013 The go-github AUTHORS. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package github

import (
	"fmt"
	"net/url"
	"strconv"
)

// ActivityListStarredOptions specifies the optional parameters to the
// ActivityService.ListStarred method.
type ActivityListStarredOptions struct {
	// How to sort the repository list.  Possible values are: created, updated,
	// pushed, full_name.  Default is "full_name".
	Sort string

	// Direction in which to sort repositories.  Possible values are: asc, desc.
	// Default is "asc" when sort is "full_name", otherwise default is "desc".
	Direction string

	// For paginated result sets, page of results to retrieve.
	Page int
}

// ListStarred lists all the repos starred by a user.  Passing the empty string
// will list the starred repositories for the authenticated user.
//
// GitHub API docs: http://developer.github.com/v3/activity/starring/#list-repositories-being-starred
func (s *ActivityService) ListStarred(user string, opt *ActivityListStarredOptions) ([]Repository, *Response, error) {
	var u string
	if user != "" {
		u = fmt.Sprintf("users/%v/starred", user)
	} else {
		u = "user/starred"
	}
	if opt != nil {
		params := url.Values{
			"sort":      []string{opt.Sort},
			"direction": []string{opt.Direction},
			"page":      []string{strconv.Itoa(opt.Page)},
		}
		u += "?" + params.Encode()
	}
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	repos := new([]Repository)
	resp, err := s.client.Do(req, repos)
	return *repos, resp, err
}
