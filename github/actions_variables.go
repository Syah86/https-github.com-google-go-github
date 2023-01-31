// Copyright 2023 The go-github AUTHORS. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package github

import (
	"context"
	"fmt"
)

// Variable represents a repository action variable.
type Variable struct {
	Name       string    `json:"name"`
	Value      string    `json:"value"`
	CreatedAt  Timestamp `json:"created_at"`
	UpdatedAt  Timestamp `json:"updated_at"`
	Visibility string    `json:"visibility,omitempty"`
	// Used by ListOrgVariables and GetOrgVariables
	SelectedRepositoriesURL string `json:"selected_repositories_url,omitempty"`
	// Used by UpdateOrgVariable and CreateOrgVariable
	SelectedRepositoryIDs SelectedRepoIDs `json:"selected_repository_ids,omitempty"`
}

// Variables represents one item from the ListVariables response.
type Variables struct {
	TotalCount int         `json:"total_count"`
	Variables  []*Variable `json:"variables"`
}

func (s *ActionsService) listVariables(ctx context.Context, url string, opts *ListOptions) (*Variables, *Response, error) {
	u, err := addOptions(url, opts)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	variables := new(Variables)
	resp, err := s.client.Do(ctx, req, &variables)
	if err != nil {
		return nil, resp, err
	}

	return variables, resp, nil
}

// ListRepoVariables lists all variables available in a repository.
//
// GitHub API docs: https://docs.github.com/en/rest/actions/variables#list-repository-variables
func (s *ActionsService) ListRepoVariables(ctx context.Context, owner, repo string, opts *ListOptions) (*Variables, *Response, error) {
	url := fmt.Sprintf("repos/%v/%v/actions/variables", owner, repo)
	return s.listVariables(ctx, url, opts)
}

// ListOrgVariables lists all variables available in an organization.
//
// GitHub API docs: https://docs.github.com/en/rest/actions/variables#list-organization-variables
func (s *ActionsService) ListOrgVariables(ctx context.Context, org string, opts *ListOptions) (*Variables, *Response, error) {
	url := fmt.Sprintf("orgs/%v/actions/variables", org)
	return s.listVariables(ctx, url, opts)
}

// ListEnvVariables lists all variables available in an environment.
//
// GitHub API docs: https://docs.github.com/en/rest/actions/variables#list-environment-variables
func (s *ActionsService) ListEnvVariables(ctx context.Context, repoID int, env string, opts *ListOptions) (*Variables, *Response, error) {
	url := fmt.Sprintf("repositories/%v/environments/%v/variables", repoID, env)
	return s.listVariables(ctx, url, opts)
}

func (s *ActionsService) getVariable(ctx context.Context, url string) (*Variable, *Response, error) {
	req, err := s.client.NewRequest("GET", url, nil)
	if err != nil {
		return nil, nil, err
	}

	variable := new(Variable)
	resp, err := s.client.Do(ctx, req, variable)
	if err != nil {
		return nil, resp, err
	}

	return variable, resp, nil
}

// GetRepoVariable gets a single repository variable.
//
// GitHub API docs: https://docs.github.com/en/rest/actions/variables#get-a-repository-variable
func (s *ActionsService) GetRepoVariable(ctx context.Context, owner, repo, name string) (*Variable, *Response, error) {
	url := fmt.Sprintf("repos/%v/%v/actions/variables/%v", owner, repo, name)
	return s.getVariable(ctx, url)
}

// GetOrgVariable gets a single organization variable.
//
// GitHub API docs: https://docs.github.com/en/rest/actions/variables#get-an-organization-variable
func (s *ActionsService) GetOrgVariable(ctx context.Context, org, name string) (*Variable, *Response, error) {
	url := fmt.Sprintf("orgs/%v/actions/variables/%v", org, name)
	return s.getVariable(ctx, url)
}

// GetEnvVariable gets a single environment variable.
//
// GitHub API docs: https://docs.github.com/en/rest/actions/variables#get-an-environment-variable
func (s *ActionsService) GetEnvVariable(ctx context.Context, repoID int, env, variableName string) (*Variable, *Response, error) {
	url := fmt.Sprintf("repositories/%v/environments/%v/variables/%v", repoID, env, variableName)
	return s.getVariable(ctx, url)
}

func (s *ActionsService) postVariable(ctx context.Context, url string, variable *Variable) (*Response, error) {
	req, err := s.client.NewRequest("POST", url, variable)
	if err != nil {
		return nil, err
	}
	return s.client.Do(ctx, req, nil)
}

// CreateRepoVariable creates a repository variable.
//
// GitHub API docs: https://docs.github.com/en/rest/actions/variables#create-a-repository-variable
func (s *ActionsService) CreateRepoVariable(ctx context.Context, owner, repo string, variable *Variable) (*Response, error) {
	url := fmt.Sprintf("repos/%v/%v/actions/variables", owner, repo)
	return s.postVariable(ctx, url, variable)
}

// CreateOrgVariable creates an organization variable.
//
// GitHub API docs: https://docs.github.com/en/rest/actions/variables#create-an-organization-variable
func (s *ActionsService) CreateOrgVariable(ctx context.Context, org string, variable *Variable) (*Response, error) {
	url := fmt.Sprintf("orgs/%v/actions/variables", org)
	return s.postVariable(ctx, url, variable)
}

// CreateEnvVariable creates an environment variable.
//
// GitHub API docs: https://docs.github.com/en/rest/actions/variables#create-an-environment-variable
func (s *ActionsService) CreateEnvVariable(ctx context.Context, repoID int, env string, variable *Variable) (*Response, error) {
	url := fmt.Sprintf("repositories/%v/environments/%v/variables", repoID, env)
	return s.postVariable(ctx, url, variable)
}

func (s *ActionsService) patchVariable(ctx context.Context, url string, variable *Variable) (*Response, error) {
	req, err := s.client.NewRequest("PATCH", url, variable)
	if err != nil {
		return nil, err
	}
	return s.client.Do(ctx, req, nil)
}

// UpdateRepoVariable updates a repository variable.
//
// GitHub API docs: https://docs.github.com/en/rest/actions/variables#update-a-repository-variable
func (s *ActionsService) UpdateRepoVariable(ctx context.Context, owner, repo string, variable *Variable) (*Response, error) {
	url := fmt.Sprintf("repos/%v/%v/actions/variables/%v", owner, repo, variable.Name)
	return s.patchVariable(ctx, url, variable)
}

// UpdateOrgVariable updates an organization variable.
//
// GitHub API docs: https://docs.github.com/en/rest/actions/variables#update-an-organization-variable
func (s *ActionsService) UpdateOrgVariable(ctx context.Context, org string, variable *Variable) (*Response, error) {
	url := fmt.Sprintf("orgs/%v/actions/variables/%v", org, variable.Name)
	return s.patchVariable(ctx, url, variable)
}

// UpdateEnvVariable updates an environment variable.
//
// GitHub API docs: https://docs.github.com/en/rest/actions/variables#create-an-environment-variable
func (s *ActionsService) UpdateEnvVariable(ctx context.Context, repoID int, env string, variable *Variable) (*Response, error) {
	url := fmt.Sprintf("repositories/%v/environments/%v/variables/%v", repoID, env, variable.Name)
	return s.patchVariable(ctx, url, variable)
}

func (s *ActionsService) deleteVariable(ctx context.Context, url string) (*Response, error) {
	req, err := s.client.NewRequest("DELETE", url, nil)
	if err != nil {
		return nil, err
	}

	return s.client.Do(ctx, req, nil)
}

// DeleteRepoVariable deletes a variable in a repository.
//
// GitHub API docs: https://docs.github.com/en/rest/actions/variables#delete-a-repository-variable
func (s *ActionsService) DeleteRepoVariable(ctx context.Context, owner, repo, name string) (*Response, error) {
	url := fmt.Sprintf("repos/%v/%v/actions/variables/%v", owner, repo, name)
	return s.deleteVariable(ctx, url)
}

// DeleteOrgVariable deletes a variable in an organization.
//
// GitHub API docs: https://docs.github.com/en/rest/actions/variables#delete-an-organization-variable
func (s *ActionsService) DeleteOrgVariable(ctx context.Context, org, name string) (*Response, error) {
	url := fmt.Sprintf("orgs/%v/actions/variables/%v", org, name)
	return s.deleteVariable(ctx, url)
}

// DeleteEnvVariable deletes a variable in an environment.
//
// GitHub API docs: https://docs.github.com/en/rest/actions/variables#delete-an-environment-variable
func (s *ActionsService) DeleteEnvVariable(ctx context.Context, repoID int, env, variableName string) (*Response, error) {
	url := fmt.Sprintf("repositories/%v/environments/%v/variables/%v", repoID, env, variableName)
	return s.deleteVariable(ctx, url)
}

func (s *ActionsService) listSelectedReposForVariable(ctx context.Context, url string, opts *ListOptions) (*SelectedReposList, *Response, error) {
	u, err := addOptions(url, opts)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	result := new(SelectedReposList)
	resp, err := s.client.Do(ctx, req, result)
	if err != nil {
		return nil, resp, err
	}

	return result, resp, nil
}

// ListSelectedReposForOrgVariable lists all repositories that have access to a variable.
//
// GitHub API docs: https://docs.github.com/en/rest/actions/variables#list-selected-repositories-for-an-organization-variable
func (s *ActionsService) ListSelectedReposForOrgVariable(ctx context.Context, org, name string, opts *ListOptions) (*SelectedReposList, *Response, error) {
	url := fmt.Sprintf("orgs/%v/actions/variables/%v/repositories", org, name)
	return s.listSelectedReposForVariable(ctx, url, opts)
}

func (s *ActionsService) setSelectedReposForVariable(ctx context.Context, url string, ids SelectedRepoIDs) (*Response, error) {
	type repoIDs struct {
		SelectedIDs SelectedRepoIDs `json:"selected_repository_ids"`
	}

	req, err := s.client.NewRequest("PUT", url, repoIDs{SelectedIDs: ids})
	if err != nil {
		return nil, err
	}

	return s.client.Do(ctx, req, nil)
}

// SetSelectedReposForOrgVariable sets the repositories that have access to a variable.
//
// GitHub API docs: https://docs.github.com/en/rest/actions/variables#set-selected-repositories-for-an-organization-variable
func (s *ActionsService) SetSelectedReposForOrgVariable(ctx context.Context, org, name string, ids SelectedRepoIDs) (*Response, error) {
	url := fmt.Sprintf("orgs/%v/actions/variables/%v/repositories", org, name)
	return s.setSelectedReposForVariable(ctx, url, ids)
}

func (s *ActionsService) addSelectedRepoToVariable(ctx context.Context, url string) (*Response, error) {
	req, err := s.client.NewRequest("PUT", url, nil)
	if err != nil {
		return nil, err
	}

	return s.client.Do(ctx, req, nil)
}

// AddSelectedRepoToOrgVariable adds a repository to an organization variable.
//
// GitHub API docs: https://docs.github.com/en/rest/actions/variables#add-selected-repository-to-an-organization-variable
func (s *ActionsService) AddSelectedRepoToOrgVariable(ctx context.Context, org, name string, repo *Repository) (*Response, error) {
	url := fmt.Sprintf("orgs/%v/actions/variables/%v/repositories/%v", org, name, *repo.ID)
	return s.addSelectedRepoToVariable(ctx, url)
}

func (s *ActionsService) removeSelectedRepoFromVariable(ctx context.Context, url string) (*Response, error) {
	req, err := s.client.NewRequest("DELETE", url, nil)
	if err != nil {
		return nil, err
	}

	return s.client.Do(ctx, req, nil)
}

// RemoveSelectedRepoFromOrgVariable removes a repository from an organization variable.
//
// GitHub API docs: https://docs.github.com/en/rest/actions/variables#remove-selected-repository-from-an-organization-variable
func (s *ActionsService) RemoveSelectedRepoFromOrgVariable(ctx context.Context, org, name string, repo *Repository) (*Response, error) {
	url := fmt.Sprintf("orgs/%v/actions/variables/%v/repositories/%v", org, name, *repo.ID)
	return s.removeSelectedRepoFromVariable(ctx, url)
}
