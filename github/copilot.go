// Copyright 2023 The go-github AUTHORS. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package github

import (
	"context"
	"encoding/json"
	"fmt"
)

// CopilotService provides access to the Copilot-related functions
// in the GitHub API.
//
// GitHub API docs: https://docs.github.com/en/rest/copilot/
type CopilotService service

// CopilotOrganizationDetails represents the details of an organization's Copilot for Business subscription.
type CopilotOrganizationDetails struct {
	SeatBreakdown         *CopilotSeatBreakdown `json:"seat_breakdown"`
	PublicCodeSuggestions string                `json:"public_code_suggestions"`
	CopilotChat           string                `json:"copilot_chat"`
	SeatManagementSetting string                `json:"seat_management_setting"`
}

// CopilotSeatBreakdown represents the breakdown of Copilot for Business seats for the organization.
type CopilotSeatBreakdown struct {
	Total               int `json:"total"`
	AddedThisCycle      int `json:"added_this_cycle"`
	PendingCancellation int `json:"pending_cancellation"`
	PendingInvitation   int `json:"pending_invitation"`
	ActiveThisCycle     int `json:"active_this_cycle"`
	InactiveThisCycle   int `json:"inactive_this_cycle"`
}

// ListCopilotSeatsResponse represents the Copilot for Business seat assignments for an organization.
type ListCopilotSeatsResponse struct {
	TotalSeats int64                 `json:"total_seats"`
	Seats      []*CopilotSeatDetails `json:"seats"`
}

// CopilotSeatDetails represents the details of a Copilot for Business seat.
type CopilotSeatDetails struct {
	// Assignee can either be a User, Team, or Organization.
	Assignee                interface{} `json:"assignee"`
	AssigningTeam           *Team       `json:"assigning_team,omitempty"`
	PendingCancellationDate *string     `json:"pending_cancellation_date,omitempty"`
	LastActivityAt          *string     `json:"last_activity_at,omitempty"`
	LastActivityEditor      *string     `json:"last_activity_editor,omitempty"`
	CreatedAt               string      `json:"created_at"`
	UpdatedAt               string      `json:"updated_at,omitempty"`
}

// SeatAssignments represents the number of seats assigned.
type SeatAssignments struct {
	SeatsCreated int `json:"seats_created"`
}

// SeatCancellations represents the number of seats cancelled.
type SeatCancellations struct {
	SeatsCancelled int `json:"seats_cancelled"`
}

func (cp *CopilotSeatDetails) UnmarshalJSON(data []byte) error {
	type alias CopilotSeatDetails
	var seatDetail alias

	if err := json.Unmarshal(data, &seatDetail); err != nil {
		return err
	}

	cp.AssigningTeam = seatDetail.AssigningTeam
	cp.PendingCancellationDate = seatDetail.PendingCancellationDate
	cp.LastActivityAt = seatDetail.LastActivityAt
	cp.LastActivityEditor = seatDetail.LastActivityEditor
	cp.CreatedAt = seatDetail.CreatedAt
	cp.UpdatedAt = seatDetail.UpdatedAt

	switch v := seatDetail.Assignee.(type) {
	case map[string]interface{}:
		jsonData, err := json.Marshal(seatDetail.Assignee)
		if err != nil {
			return err
		}

		if v["type"] == nil {
			return fmt.Errorf("assignee type field is not set")
		}

		if v["type"].(string) == "User" {
			user := &User{}
			if err := json.Unmarshal(jsonData, user); err != nil {
				return err
			}
			cp.Assignee = user
		} else if v["type"].(string) == "Team" {
			team := &Team{}
			if err := json.Unmarshal(jsonData, team); err != nil {
				return err
			}
			cp.Assignee = team
		} else if v["type"].(string) == "Organization" {
			organization := &Organization{}
			if err := json.Unmarshal(jsonData, organization); err != nil {
				return err
			}
			cp.Assignee = organization
		} else {
			return fmt.Errorf("unsupported assignee type %s", v["type"].(string))
		}
	default:
		return fmt.Errorf("unsupported assignee type %T", v)
	}

	return nil
}

// GetUser gets the User from the CopilotSeatDetails if the assignee is a user.
func (cp *CopilotSeatDetails) GetUser() (*User, bool) {
	if user, ok := cp.Assignee.(*User); ok {
		return user, true
	}
	return nil, false
}

// GetTeam gets the Team from the CopilotSeatDetails if the assignee is a team.
func (cp *CopilotSeatDetails) GetTeam() (*Team, bool) {
	if team, ok := cp.Assignee.(*Team); ok {
		return team, true
	}
	return nil, false
}

// GetOrganization gets the Organization from the CopilotSeatDetails if the assignee is an organization.
func (cp *CopilotSeatDetails) GetOrganization() (*Organization, bool) {
	if organization, ok := cp.Assignee.(*Organization); ok {
		return organization, true
	}
	return nil, false
}

// GetCopilotBilling gets Copilot for Business billing information and settings for an organization.
//
// GitHub API docs: https://docs.github.com/en/rest/copilot/copilot-for-business#get-copilot-for-business-seat-information-and-settings-for-an-organization
func (s *CopilotService) GetCopilotBilling(ctx context.Context, org string) (*CopilotOrganizationDetails, *Response, error) {
	u := fmt.Sprintf("orgs/%v/copilot/billing", org)

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	var copilotDetails *CopilotOrganizationDetails
	resp, err := s.client.Do(ctx, req, &copilotDetails)
	if err != nil {
		return nil, resp, err
	}

	return copilotDetails, resp, nil
}

// ListCopilotSeats lists Copilot for Business seat assignments for an organization.
//
// To paginate through all seats, populate 'Page' with the number of the last page.
//
// GitHub API docs: https://docs.github.com/en/rest/copilot/copilot-for-business#list-all-copilot-for-business-seat-assignments-for-an-organization
func (s *CopilotService) ListCopilotSeats(ctx context.Context, org string, opts *ListOptions) (*ListCopilotSeatsResponse, *Response, error) {
	u := fmt.Sprintf("orgs/%v/copilot/billing/seats", org)

	req, err := s.client.NewRequest("GET", u, opts)
	if err != nil {
		return nil, nil, err
	}

	var copilotSeats *ListCopilotSeatsResponse
	resp, err := s.client.Do(ctx, req, &copilotSeats)
	if err != nil {
		return nil, resp, err
	}

	return copilotSeats, resp, nil
}

// AddCopilotTeams adds teams to the Copilot for Business subscription for an organization.
//
// GitHub API docs: https://docs.github.com/en/rest/copilot/copilot-for-business#add-teams-to-the-copilot-for-business-subscription-for-an-organization
func (s *CopilotService) AddCopilotTeams(ctx context.Context, org string, teamNames []string) (*SeatAssignments, *Response, error) {
	u := fmt.Sprintf("orgs/%v/copilot/billing/selected_teams", org)

	req, err := s.client.NewRequest("POST", u, teamNames)
	if err != nil {
		return nil, nil, err
	}

	var seatAssignments *SeatAssignments
	resp, err := s.client.Do(ctx, req, &seatAssignments)
	if err != nil {
		return nil, resp, err
	}

	return seatAssignments, resp, nil
}

// RemoveCopilotTeams removes teams from the Copilot for Business subscription for an organization.
//
// GitHub API docs: https://docs.github.com/en/rest/copilot/copilot-for-business#remove-teams-from-the-copilot-for-business-subscription-for-an-organization
func (s *CopilotService) RemoveCopilotTeams(ctx context.Context, org string, teamNames []string) (*SeatCancellations, *Response, error) {
	u := fmt.Sprintf("orgs/%v/copilot/billing/selected_teams", org)

	req, err := s.client.NewRequest("DELETE", u, teamNames)
	if err != nil {
		return nil, nil, err
	}

	var seatCancellations *SeatCancellations
	resp, err := s.client.Do(ctx, req, &seatCancellations)
	if err != nil {
		return nil, resp, err
	}

	return seatCancellations, resp, nil
}

// AddCopilotUsers adds users to the Copilot for Business subscription for an organization
//
// GitHub API docs: https://docs.github.com/en/rest/copilot/copilot-for-business#add-users-to-the-copilot-for-business-subscription-for-an-organization
func (s *CopilotService) AddCopilotUsers(ctx context.Context, org string, userNames []string) (*SeatAssignments, *Response, error) {
	u := fmt.Sprintf("orgs/%v/copilot/billing/selected_users", org)

	req, err := s.client.NewRequest("POST", u, userNames)
	if err != nil {
		return nil, nil, err
	}

	var seatAssignments *SeatAssignments
	resp, err := s.client.Do(ctx, req, &seatAssignments)
	if err != nil {
		return nil, resp, err
	}

	return seatAssignments, resp, nil
}

// RemoveCopilotUsers removes users from the Copilot for Business subscription for an organization.
//
// GitHub API docs: https://docs.github.com/en/rest/copilot/copilot-for-business#remove-users-from-the-copilot-for-business-subscription-for-an-organization
func (s *CopilotService) RemoveCopilotUsers(ctx context.Context, org string, userNames []string) (*SeatCancellations, *Response, error) {
	u := fmt.Sprintf("orgs/%v/copilot/billing/selected_users", org)

	req, err := s.client.NewRequest("DELETE", u, userNames)
	if err != nil {
		return nil, nil, err
	}

	var seatCancellations *SeatCancellations
	resp, err := s.client.Do(ctx, req, &seatCancellations)
	if err != nil {
		return nil, resp, err
	}

	return seatCancellations, resp, nil
}

// GetSeatDetails gets Copilot for Business seat assignment details for a user.
//
// GitHub API docs: https://docs.github.com/en/rest/copilot/copilot-for-business#get-copilot-for-business-seat-assignment-details-for-a-user
func (s *CopilotService) GetSeatDetails(ctx context.Context, org, user string) (*CopilotSeatDetails, *Response, error) {
	u := fmt.Sprintf("orgs/%v/members/%v/copilot", org, user)

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	var seatDetails *CopilotSeatDetails
	resp, err := s.client.Do(ctx, req, &seatDetails)
	if err != nil {
		return nil, resp, err
	}

	return seatDetails, resp, nil
}
