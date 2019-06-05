// Copyright 2016 The go-github AUTHORS. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package github

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"testing"
	"time"
)

func TestAppsService_Get_authenticatedApp(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/app", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Accept", mediaTypeIntegrationPreview)
		fmt.Fprint(w, `{"id":1}`)
	})

	app, _, err := client.Apps.Get(context.Background(), "")
	if err != nil {
		t.Errorf("Apps.Get returned error: %v", err)
	}

	want := &App{ID: Int64(1)}
	if !reflect.DeepEqual(app, want) {
		t.Errorf("Apps.Get returned %+v, want %+v", app, want)
	}
}

func TestAppsService_Get_specifiedApp(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/apps/a", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Accept", mediaTypeIntegrationPreview)
		fmt.Fprint(w, `{"html_url":"https://github.com/apps/a"}`)
	})

	app, _, err := client.Apps.Get(context.Background(), "a")
	if err != nil {
		t.Errorf("Apps.Get returned error: %v", err)
	}

	want := &App{HTMLURL: String("https://github.com/apps/a")}
	if !reflect.DeepEqual(app, want) {
		t.Errorf("Apps.Get returned %+v, want %+v", *app.HTMLURL, *want.HTMLURL)
	}
}

func TestAppsService_ListInstallations(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/app/installations", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Accept", mediaTypeIntegrationPreview)
		testFormValues(t, r, values{
			"page":     "1",
			"per_page": "2",
		})
		fmt.Fprint(w, `[{
                                   "id":1,
                                   "app_id":1,
                                   "target_id":1,
                                   "target_type": "Organization",
                                   "permissions": {
                                       "metadata": "read",
                                       "contents": "read",
                                       "issues": "write",
                                       "single_file": "write"
                                   },
                                  "events": [
                                      "push",
                                      "pull_request"
                                  ],
                                 "single_file_name": "config.yml",
                                 "repository_selection": "selected",
                                 "created_at": "2018-01-01T00:00:00Z",
                                 "updated_at": "2018-01-01T00:00:00Z"}]`,
		)
	})

	opt := &ListOptions{Page: 1, PerPage: 2}
	installations, _, err := client.Apps.ListInstallations(context.Background(), opt)
	if err != nil {
		t.Errorf("Apps.ListInstallations returned error: %v", err)
	}

	date := Timestamp{Time: time.Date(2018, time.January, 1, 0, 0, 0, 0, time.UTC)}
	want := []*Installation{{
		ID:                  Int64(1),
		AppID:               Int64(1),
		TargetID:            Int64(1),
		TargetType:          String("Organization"),
		SingleFileName:      String("config.yml"),
		RepositorySelection: String("selected"),
		Permissions: &InstallationPermissions{
			Metadata:   String("read"),
			Contents:   String("read"),
			Issues:     String("write"),
			SingleFile: String("write")},
		Events:    []string{"push", "pull_request"},
		CreatedAt: &date,
		UpdatedAt: &date,
	}}
	if !reflect.DeepEqual(installations, want) {
		t.Errorf("Apps.ListInstallations returned %+v, want %+v", installations, want)
	}
}

func TestAppsService_GetInstallation(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/app/installations/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Accept", mediaTypeIntegrationPreview)
		fmt.Fprint(w, `{"id":1, "app_id":1, "target_id":1, "target_type": "Organization"}`)
	})

	installation, _, err := client.Apps.GetInstallation(context.Background(), 1)
	if err != nil {
		t.Errorf("Apps.GetInstallation returned error: %v", err)
	}

	want := &Installation{ID: Int64(1), AppID: Int64(1), TargetID: Int64(1), TargetType: String("Organization")}
	if !reflect.DeepEqual(installation, want) {
		t.Errorf("Apps.GetInstallation returned %+v, want %+v", installation, want)
	}
}

func TestAppsService_ListUserInstallations(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/user/installations", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Accept", mediaTypeIntegrationPreview)
		testFormValues(t, r, values{
			"page":     "1",
			"per_page": "2",
		})
		fmt.Fprint(w, `{"installations":[{"id":1, "app_id":1, "target_id":1, "target_type": "Organization"}]}`)
	})

	opt := &ListOptions{Page: 1, PerPage: 2}
	installations, _, err := client.Apps.ListUserInstallations(context.Background(), opt)
	if err != nil {
		t.Errorf("Apps.ListUserInstallations returned error: %v", err)
	}

	want := []*Installation{{ID: Int64(1), AppID: Int64(1), TargetID: Int64(1), TargetType: String("Organization")}}
	if !reflect.DeepEqual(installations, want) {
		t.Errorf("Apps.ListUserInstallations returned %+v, want %+v", installations, want)
	}
}

func TestAppsService_CreateInstallationToken(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/app/installations/1/access_tokens", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		testHeader(t, r, "Accept", mediaTypeIntegrationPreview)
		fmt.Fprint(w, `{"token":"t"}`)
	})

	token, _, err := client.Apps.CreateInstallationToken(context.Background(), 1)
	if err != nil {
		t.Errorf("Apps.CreateInstallationToken returned error: %v", err)
	}

	want := &InstallationToken{Token: String("t")}
	if !reflect.DeepEqual(token, want) {
		t.Errorf("Apps.CreateInstallationToken returned %+v, want %+v", token, want)
	}
}

func TestAppsService_CreateRepositoryScopedInstallationToken(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	input := &ScopedInstallationTokenRequest{RepositoryIds: &[]int64{1}}

	mux.HandleFunc("/app/installations/1/access_tokens", func(w http.ResponseWriter, r *http.Request) {
		v := new(ScopedInstallationTokenRequest)
		json.NewDecoder(r.Body).Decode(v)

		testMethod(t, r, "POST")
		testHeader(t, r, "Accept", mediaTypeIntegrationPreview)
		if !reflect.DeepEqual(v, input) {
			t.Errorf("Request body = %+v, want %+v", v, input)
		}

		fmt.Fprint(w, `{"token":"t", "repositories": [{"id": 1}]}`)
	})

	token, _, err := client.Apps.CreateScopedInstallationToken(context.Background(), 1, input)
	if err != nil {
		t.Errorf("Apps.CreateScopedInstallationToken returned error: %v", err)
	}

	want := &InstallationToken{
		Token: String("t"),
		Repositories: []*Repository{
			{ID: Int64(1)},
		},
	}

	if !reflect.DeepEqual(token, want) {
		t.Errorf("Apps.CreateScopedInstallationToken returned %+v, want %+v", token, want)
	}
}

func TestAppsService_CreateRepositoryAndPermissionScopedInstallationToken(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	input := &ScopedInstallationTokenRequest{
		RepositoryIds: &[]int64{1},
		Permissions: &InstallationPermissions{
			Metadata: String("read"),
			Contents: String("read"),
		},
	}

	mux.HandleFunc("/app/installations/1/access_tokens", func(w http.ResponseWriter, r *http.Request) {
		v := new(ScopedInstallationTokenRequest)
		json.NewDecoder(r.Body).Decode(v)

		testMethod(t, r, "POST")
		testHeader(t, r, "Accept", mediaTypeIntegrationPreview)
		if !reflect.DeepEqual(v, input) {
			t.Errorf("Request body = %+v, want %+v", v, input)
		}

		fmt.Fprint(w, `{"token":"t", "repositories": [{"id": 1}], "permissions": { "metadata": "read", "contents": "read" }}`)
	})

	token, _, err := client.Apps.CreateScopedInstallationToken(context.Background(), 1, input)
	if err != nil {
		t.Errorf("Apps.CreateScopedInstallationToken returned error: %v", err)
	}

	want := &InstallationToken{
		Token: String("t"),
		Repositories: []*Repository{
			{ID: Int64(1)},
		},
		Permissions: &InstallationPermissions{
			Metadata: String("read"),
			Contents: String("read"),
		},
	}

	if !reflect.DeepEqual(token, want) {
		t.Errorf("Apps.CreateScopedInstallationToken returned %+v, want %+v", token, want)
	}
}

func TestAppsService_CreateAttachement(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/content_references/11/attachments", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		testHeader(t, r, "Accept", mediaTypeReactionsPreview)

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"id":1,"title":"title1","body":"body1"}`))
	})

	got, _, err := client.Apps.CreateAttachment(context.Background(), 11, "title1", "body1")
	if err != nil {
		t.Errorf("CreateAttachment returned error: %v", err)
	}

	want := &Attachment{ID: Int64(1), Title: String("title1"), Body: String("body1")}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("CreateAttachment = %+v, want %+v", got, want)
	}
}
func TestAppsService_FindOrganizationInstallation(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/orgs/o/installation", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Accept", mediaTypeIntegrationPreview)
		fmt.Fprint(w, `{"id":1, "app_id":1, "target_id":1, "target_type": "Organization"}`)
	})

	installation, _, err := client.Apps.FindOrganizationInstallation(context.Background(), "o")
	if err != nil {
		t.Errorf("Apps.FindOrganizationInstallation returned error: %v", err)
	}

	want := &Installation{ID: Int64(1), AppID: Int64(1), TargetID: Int64(1), TargetType: String("Organization")}
	if !reflect.DeepEqual(installation, want) {
		t.Errorf("Apps.FindOrganizationInstallation returned %+v, want %+v", installation, want)
	}
}

func TestAppsService_FindRepositoryInstallation(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/repos/o/r/installation", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Accept", mediaTypeIntegrationPreview)
		fmt.Fprint(w, `{"id":1, "app_id":1, "target_id":1, "target_type": "Organization"}`)
	})

	installation, _, err := client.Apps.FindRepositoryInstallation(context.Background(), "o", "r")
	if err != nil {
		t.Errorf("Apps.FindRepositoryInstallation returned error: %v", err)
	}

	want := &Installation{ID: Int64(1), AppID: Int64(1), TargetID: Int64(1), TargetType: String("Organization")}
	if !reflect.DeepEqual(installation, want) {
		t.Errorf("Apps.FindRepositoryInstallation returned %+v, want %+v", installation, want)
	}
}

func TestAppsService_FindRepositoryInstallationByID(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/repositories/1/installation", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Accept", mediaTypeIntegrationPreview)
		fmt.Fprint(w, `{"id":1, "app_id":1, "target_id":1, "target_type": "Organization"}`)
	})

	installation, _, err := client.Apps.FindRepositoryInstallationByID(context.Background(), 1)
	if err != nil {
		t.Errorf("Apps.FindRepositoryInstallationByID returned error: %v", err)
	}

	want := &Installation{ID: Int64(1), AppID: Int64(1), TargetID: Int64(1), TargetType: String("Organization")}
	if !reflect.DeepEqual(installation, want) {
		t.Errorf("Apps.FindRepositoryInstallationByID returned %+v, want %+v", installation, want)
	}
}

func TestAppsService_FindUserInstallation(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/users/u/installation", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Accept", mediaTypeIntegrationPreview)
		fmt.Fprint(w, `{"id":1, "app_id":1, "target_id":1, "target_type": "User"}`)
	})

	installation, _, err := client.Apps.FindUserInstallation(context.Background(), "u")
	if err != nil {
		t.Errorf("Apps.FindUserInstallation returned error: %v", err)
	}

	want := &Installation{ID: Int64(1), AppID: Int64(1), TargetID: Int64(1), TargetType: String("User")}
	if !reflect.DeepEqual(installation, want) {
		t.Errorf("Apps.FindUserInstallation returned %+v, want %+v", installation, want)
	}
}
