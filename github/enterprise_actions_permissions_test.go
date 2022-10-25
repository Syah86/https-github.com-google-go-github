// Copyright 2021 The go-github AUTHORS. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package github

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestEnterpriseService_GetActionsPermissions(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/enterprises/e/actions/permissions", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, `{"enabled_organizations": "all", "allowed_actions": "all"}`)
	})

	ctx := context.Background()
	ent, _, err := client.Enterprise.GetActionsPermissions(ctx, "e")
	if err != nil {
		t.Errorf("Enterprise.GetActionsPermissions returned error: %v", err)
	}
	want := &ActionsPermissionsEnterprise{EnabledOrganizations: String("all"), AllowedActions: String("all")}
	if !cmp.Equal(ent, want) {
		t.Errorf("Enterprise.GetActionsPermissions returned %+v, want %+v", ent, want)
	}

	const methodName = "GetActionsPermissions"
	testBadOptions(t, methodName, func() (err error) {
		_, _, err = client.Enterprise.GetActionsPermissions(ctx, "\n")
		return err
	})

	testNewRequestAndDoFailure(t, methodName, client, func() (*Response, error) {
		got, resp, err := client.Enterprise.GetActionsPermissions(ctx, "e")
		if got != nil {
			t.Errorf("testNewRequestAndDoFailure %v = %#v, want nil", methodName, got)
		}
		return resp, err
	})
}

func TestEnterpriseService_EditActionsPermissions(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	input := &ActionsPermissionsEnterprise{EnabledOrganizations: String("all"), AllowedActions: String("selected")}

	mux.HandleFunc("/enterprises/e/actions/permissions", func(w http.ResponseWriter, r *http.Request) {
		v := new(ActionsPermissionsEnterprise)
		json.NewDecoder(r.Body).Decode(v)

		testMethod(t, r, "PUT")
		if !cmp.Equal(v, input) {
			t.Errorf("Request body = %+v, want %+v", v, input)
		}

		fmt.Fprint(w, `{"enabled_organizations": "all", "allowed_actions": "selected"}`)
	})

	ctx := context.Background()
	ent, _, err := client.Enterprise.EditActionsPermissions(ctx, "e", *input)
	if err != nil {
		t.Errorf("Enterprise.EditActionsPermissions returned error: %v", err)
	}

	want := &ActionsPermissionsEnterprise{EnabledOrganizations: String("all"), AllowedActions: String("selected")}
	if !cmp.Equal(ent, want) {
		t.Errorf("Enterprise.EditActionsPermissions returned %+v, want %+v", ent, want)
	}

	const methodName = "EditActionsPermissions"
	testBadOptions(t, methodName, func() (err error) {
		_, _, err = client.Enterprise.EditActionsPermissions(ctx, "\n", *input)
		return err
	})

	testNewRequestAndDoFailure(t, methodName, client, func() (*Response, error) {
		got, resp, err := client.Enterprise.EditActionsPermissions(ctx, "e", *input)
		if got != nil {
			t.Errorf("testNewRequestAndDoFailure %v = %#v, want nil", methodName, got)
		}
		return resp, err
	})
}

func TestActionsService_ListEnabledOrgsInEnterprise(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/enterprises/e/actions/permissions/organizations", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testFormValues(t, r, values{
			"page": "1",
		})
		fmt.Fprint(w, `{"total_count":2,"organizations":[{"id":2}, {"id":3}]}`)
	})

	ctx := context.Background()
	opt := &ListOptions{
		Page: 1,
	}
	got, _, err := client.Actions.ListEnabledOrgsInEnterprise(ctx, "e", opt)
	if err != nil {
		t.Errorf("Actions.ListEnabledOrgsInEnterprise returned error: %v", err)
	}

	want := &ActionsEnabledOnEnterpriseOrgs{TotalCount: int(2), Organizations: []*Organization{
		{ID: Int64(2)},
		{ID: Int64(3)},
	}}
	if !cmp.Equal(got, want) {
		t.Errorf("Actions.ListEnabledOrgsInEnterprise returned %+v, want %+v", got, want)
	}

	const methodName = "ListEnabledOrgsInEnterprise"
	testBadOptions(t, methodName, func() (err error) {
		_, _, err = client.Actions.ListEnabledOrgsInEnterprise(ctx, "\n", opt)
		return err
	})

	testNewRequestAndDoFailure(t, methodName, client, func() (*Response, error) {
		got, resp, err := client.Actions.ListEnabledOrgsInEnterprise(ctx, "e", opt)
		if got != nil {
			t.Errorf("testNewRequestAndDoFailure %v = %#v, want nil", methodName, got)
		}
		return resp, err
	})
}

func TestActionsService_SetEnabledOrgsInEnterprise(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/enterprises/e/actions/permissions/organizations", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "PUT")
		testHeader(t, r, "Content-Type", "application/json")
		testBody(t, r, `{"selected_organization_ids":[123,1234]}`+"\n")
		w.WriteHeader(http.StatusNoContent)
	})

	ctx := context.Background()
	_, err := client.Actions.SetEnabledOrgsInEnterprise(ctx, "e", []int64{123, 1234})
	if err != nil {
		t.Errorf("Actions.SetEnabledOrgsInEnterprise returned error: %v", err)
	}

	const methodName = "SetEnabledOrgsInEnterprise"

	testBadOptions(t, methodName, func() (err error) {
		_, err = client.Actions.SetEnabledOrgsInEnterprise(ctx, "\n", []int64{123, 1234})
		return err
	})

	testNewRequestAndDoFailure(t, methodName, client, func() (*Response, error) {
		return client.Actions.SetEnabledOrgsInEnterprise(ctx, "e", []int64{123, 1234})
	})
}

func TestActionsService_AddEnabledOrgInEnterprise(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/enterprises/e/actions/permissions/organizations/123", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "PUT")
		w.WriteHeader(http.StatusNoContent)
	})

	ctx := context.Background()
	_, err := client.Actions.AddEnabledOrgInEnterprise(ctx, "e", 123)
	if err != nil {
		t.Errorf("Actions.AddEnabledOrgInEnterprise returned error: %v", err)
	}

	const methodName = "AddEnabledOrgInEnterprise"

	testBadOptions(t, methodName, func() (err error) {
		_, err = client.Actions.AddEnabledOrgInEnterprise(ctx, "\n", 123)
		return err
	})

	testNewRequestAndDoFailure(t, methodName, client, func() (*Response, error) {
		return client.Actions.AddEnabledOrgInEnterprise(ctx, "e", 123)
	})
}

func TestActionsService_RemoveEnabledOrgInEnterprise(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/enterprises/e/actions/permissions/organizations/123", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "DELETE")
		w.WriteHeader(http.StatusNoContent)
	})

	ctx := context.Background()
	_, err := client.Actions.RemoveEnabledOrgInEnterprise(ctx, "e", 123)
	if err != nil {
		t.Errorf("Actions.RemoveEnabledOrgInEnterprise returned error: %v", err)
	}

	const methodName = "RemoveEnabledOrgInEnterprise"

	testBadOptions(t, methodName, func() (err error) {
		_, err = client.Actions.RemoveEnabledOrgInEnterprise(ctx, "\n", 123)
		return err
	})

	testNewRequestAndDoFailure(t, methodName, client, func() (*Response, error) {
		return client.Actions.RemoveEnabledOrgInEnterprise(ctx, "e", 123)
	})
}
