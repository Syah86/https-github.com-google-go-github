// Copyright 2014 The go-github AUTHORS. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package github

import "fmt"

// ContributorStats represents a contributor to a repository and their
// weekly contributions to a given repo.
type ContributorStats struct {
	Author *Contributor `json:"author,omitempty"`
	Total  *int         `json:"total,omitempty"`
	Weeks  []WeeklyHash `json:"weeks,omitempty"`
}

// WeeklyHash represents the number of additions, deletions and commits
// a Contributor made in a given week.
type WeeklyHash struct {
	Week      *int `json:"w,omitempty"`
	Additions *int `json:"a,omitempty"`
	Deletions *int `json:"d,omitempty"`
	Commits   *int `json:"c,omitempty"`
}

// ListContributorsStats gets a repo's contributor list with additions, deletions and commit counts.
// If this is the first time these statistics are requested for the given repository, this method
// will return a non-nil error and a status code of 202. This is because this is the status that github
// returns to signify that it is now computing the requested statistics. A follow up request, after
// a delay of a second or so, should result in a successful request.
//
// GitHub API docs: https://developer.github.com/v3/repos/statistics/#contributors
func (s *RepositoriesService) ListContributorsStats(owner, repo string) (*[]ContributorStats, *Response, error) {
	u := fmt.Sprintf("repos/%v/%v/stats/contributors", owner, repo)
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	contributorStats := new([]ContributorStats)
	resp, err := s.client.Do(req, contributorStats)
	if err != nil {
		return nil, resp, err
	}

	return contributorStats, resp, err
}

// WeeklyCommitActivity represents the weekly commit activity for a repository.
// The days array is a group of commits per day, starting on Sunday.
type WeeklyCommitActivity struct {
	Days  []int `json:"days,omitempty"`
	Total *int  `json:"total,omitempty"`
	Week  *int  `json:"week,omitempty"`
}

// ListCommitActivity returns the last year of commit activity
// grouped by week. The days array is a group of commits per day,
// starting on Sunday. If this is the first time these statistics are
// requested for the given repository, this method will return a
// non-nil error and a status code of 202. This is because this is the
// status that github returns to signify that it is now computing the
// requested statistics. A follow up request, after a delay of a second
// or so, should result in a successful request.
//
// GitHub API docs: https://developer.github.com/v3/repos/statistics/#commit-activity
func (s *RepositoriesService) ListCommitActivity(owner, repo string) (*[]WeeklyCommitActivity, *Response, error) {
	u := fmt.Sprintf("repos/%v/%v/stats/commit_activity", owner, repo)
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	weeklyCommitActivity := new([]WeeklyCommitActivity)
	resp, err := s.client.Do(req, weeklyCommitActivity)
	if err != nil {
		return nil, resp, err
	}

	return weeklyCommitActivity, resp, err
}

// RepositoryParticipation is the number of commits by everyone
// who has contributed to the repository (including the owner)
// as well as the number of commits by the owner themself.
type RepositoryParticipation struct {
	All   []int `json:"all,omitempty"`
	Owner []int `json:"owner,omitempty"`
}

// ListParticipation returns the total commit counts for the 'owner'
// and total commit counts in 'all'. 'all' is everyone combined,
// including the 'owner' in the last 52 weeks. If you’d like to get
// the commit counts for non-owners, you can subtract 'all' from 'owner'.
//
// The array order is oldest week (index 0) to most recent week.
//
// If this is the first time these statistics are requested for the given
// repository, this method will return a non-nil error and a status code
// of 202. This is because this is the status that github returns to
// signify that it is now computing the requested statistics. A follow
// up request, after a delay of a second or so, should result in a
// successful request.
//
// GitHub API Docs: https://developer.github.com/v3/repos/statistics/#participation
func (s *RepositoriesService) ListParticipation(owner, repo string) (*RepositoryParticipation, *Response, error) {
	u := fmt.Sprintf("repos/%v/%v/stats/participation", owner, repo)
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	participation := new(RepositoryParticipation)
	resp, err := s.client.Do(req, participation)
	if err != nil {
		return nil, resp, err
	}

	return participation, resp, err
}
