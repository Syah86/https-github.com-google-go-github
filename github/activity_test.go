// Copyright 2013 The go-github AUTHORS. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package github

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestActivityService_ListStarred_specifiedUser(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/users/u/starred", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, `[{"id":1}]`)
	})

	repos, err := client.Activity.ListStarred("u", nil)
	if err != nil {
		t.Errorf("Activity.ListStarred returned error: %v", err)
	}

	want := []Repository{Repository{ID: 1}}
	if !reflect.DeepEqual(repos, want) {
		t.Errorf("Repositories.List returned %+v, want %+v", repos, want)
	}
}

func TestActivityService_ListStarred_specifiedUser_optionalParameters(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/users/u/starred", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testFormValues(t, r, values{
			"sort":      "created",
			"direction": "asc",
			"page":      "2",
		})
		fmt.Fprint(w, `[{"id":2}]`)
	})

	opt := &ActivityListOptions{"created", "asc", 2}
	repos, err := client.Activity.ListStarred("u", opt)
	if err != nil {
		t.Errorf("Activity.ListStarred returned error: %v", err)
	}

	want := []Repository{Repository{ID: 2}}
	if !reflect.DeepEqual(repos, want) {
		t.Errorf("Activity.ListStarred returned %+v, want %+v", repos, want)
	}
}
