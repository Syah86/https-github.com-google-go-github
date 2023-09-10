// Copyright 2018 The go-github AUTHORS. All rights reserved.
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
	"time"

	"github.com/google/go-cmp/cmp"
)

func TestTeamsService_ListDiscussionsByID(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/organizations/1/team/2/discussions", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testFormValues(t, r, values{
			"direction": "desc",
			"page":      "2",
		})
		fmt.Fprintf(w,
			`[
				{
					"author": {
						"login": "author",
						"id": 0,
						"avatar_url": "https://avatars1.githubusercontent.com/u/0?v=4",
						"gravatar_id": "",
						"url": "https://api.github.com/users/author",
						"html_url": "https://github.com/author",
						"followers_url": "https://api.github.com/users/author/followers",
						"following_url": "https://api.github.com/users/author/following{/other_user}",
						"gists_url": "https://api.github.com/users/author/gists{/gist_id}",
						"starred_url": "https://api.github.com/users/author/starred{/owner}{/repo}",
						"subscriptions_url": "https://api.github.com/users/author/subscriptions",
						"organizations_url": "https://api.github.com/users/author/orgs",
						"repos_url": "https://api.github.com/users/author/repos",
						"events_url": "https://api.github.com/users/author/events{/privacy}",
						"received_events_url": "https://api.github.com/users/author/received_events",
						"type": "User",
						"site_admin": false
					},
					"body": "test",
					"body_html": "<p>test</p>",
					"body_version": "version",
					"comments_count": 1,
					"comments_url": "https://api.github.com/teams/2/discussions/3/comments",
					"created_at": "2018-01-01T00:00:00Z",
					"last_edited_at": null,
					"html_url": "https://github.com/orgs/1/teams/2/discussions/3",
					"node_id": "node",
					"number": 3,
					"pinned": false,
					"private": false,
					"team_url": "https://api.github.com/teams/2",
					"title": "test",
					"updated_at": "2018-01-01T00:00:00Z",
					"url": "https://api.github.com/teams/2/discussions/3"
				}
			]`)
	})
	ctx := context.Background()
	discussions, _, err := client.Teams.ListDiscussionsByID(ctx, 1, 2, &DiscussionListOptions{"desc", ListOptions{Page: 2}})
	if err != nil {
		t.Errorf("Teams.ListDiscussionsByID returned error: %v", err)
	}

	want := []*TeamDiscussion{
		{
			Author: &User{
				Login:             String("author"),
				ID:                Int64(0),
				AvatarURL:         String("https://avatars1.githubusercontent.com/u/0?v=4"),
				GravatarID:        String(""),
				URL:               String("https://api.github.com/users/author"),
				HTMLURL:           String("https://github.com/author"),
				FollowersURL:      String("https://api.github.com/users/author/followers"),
				FollowingURL:      String("https://api.github.com/users/author/following{/other_user}"),
				GistsURL:          String("https://api.github.com/users/author/gists{/gist_id}"),
				StarredURL:        String("https://api.github.com/users/author/starred{/owner}{/repo}"),
				SubscriptionsURL:  String("https://api.github.com/users/author/subscriptions"),
				OrganizationsURL:  String("https://api.github.com/users/author/orgs"),
				ReposURL:          String("https://api.github.com/users/author/repos"),
				EventsURL:         String("https://api.github.com/users/author/events{/privacy}"),
				ReceivedEventsURL: String("https://api.github.com/users/author/received_events"),
				Type:              String("User"),
				SiteAdmin:         Bool(false),
			},
			Body:          String("test"),
			BodyHTML:      String("<p>test</p>"),
			BodyVersion:   String("version"),
			CommentsCount: Int(1),
			CommentsURL:   String("https://api.github.com/teams/2/discussions/3/comments"),
			CreatedAt:     &Timestamp{time.Date(2018, time.January, 1, 0, 0, 0, 0, time.UTC)},
			LastEditedAt:  nil,
			HTMLURL:       String("https://github.com/orgs/1/teams/2/discussions/3"),
			NodeID:        String("node"),
			Number:        Int(3),
			Pinned:        Bool(false),
			Private:       Bool(false),
			TeamURL:       String("https://api.github.com/teams/2"),
			Title:         String("test"),
			UpdatedAt:     &Timestamp{time.Date(2018, time.January, 1, 0, 0, 0, 0, time.UTC)},
			URL:           String("https://api.github.com/teams/2/discussions/3"),
		},
	}
	if !cmp.Equal(discussions, want) {
		t.Errorf("Teams.ListDiscussionsByID returned %+v, want %+v", discussions, want)
	}

	const methodName = "ListDiscussionsByID"
	testBadOptions(t, methodName, func() (err error) {
		_, _, err = client.Teams.ListDiscussionsByID(ctx, -1, -2, nil)
		return err
	})

	testNewRequestAndDoFailure(t, methodName, client, func() (*Response, error) {
		got, resp, err := client.Teams.ListDiscussionsByID(ctx, 1, 2, &DiscussionListOptions{"desc", ListOptions{Page: 2}})
		if got != nil {
			t.Errorf("testNewRequestAndDoFailure %v = %#v, want nil", methodName, got)
		}
		return resp, err
	})
}

func TestTeamsService_ListDiscussionsBySlug(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/orgs/o/teams/s/discussions", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testFormValues(t, r, values{
			"direction": "desc",
			"page":      "2",
		})
		fmt.Fprintf(w,
			`[
				{
					"author": {
						"login": "author",
						"id": 0,
						"avatar_url": "https://avatars1.githubusercontent.com/u/0?v=4",
						"gravatar_id": "",
						"url": "https://api.github.com/users/author",
						"html_url": "https://github.com/author",
						"followers_url": "https://api.github.com/users/author/followers",
						"following_url": "https://api.github.com/users/author/following{/other_user}",
						"gists_url": "https://api.github.com/users/author/gists{/gist_id}",
						"starred_url": "https://api.github.com/users/author/starred{/owner}{/repo}",
						"subscriptions_url": "https://api.github.com/users/author/subscriptions",
						"organizations_url": "https://api.github.com/users/author/orgs",
						"repos_url": "https://api.github.com/users/author/repos",
						"events_url": "https://api.github.com/users/author/events{/privacy}",
						"received_events_url": "https://api.github.com/users/author/received_events",
						"type": "User",
						"site_admin": false
					},
					"body": "test",
					"body_html": "<p>test</p>",
					"body_version": "version",
					"comments_count": 1,
					"comments_url": "https://api.github.com/teams/2/discussions/3/comments",
					"created_at": "2018-01-01T00:00:00Z",
					"last_edited_at": null,
					"html_url": "https://github.com/orgs/1/teams/2/discussions/3",
					"node_id": "node",
					"number": 3,
					"pinned": false,
					"private": false,
					"team_url": "https://api.github.com/teams/2",
					"title": "test",
					"updated_at": "2018-01-01T00:00:00Z",
					"url": "https://api.github.com/teams/2/discussions/3"
				}
			]`)
	})
	ctx := context.Background()
	discussions, _, err := client.Teams.ListDiscussionsBySlug(ctx, "o", "s", &DiscussionListOptions{"desc", ListOptions{Page: 2}})
	if err != nil {
		t.Errorf("Teams.ListDiscussionsBySlug returned error: %v", err)
	}

	want := []*TeamDiscussion{
		{
			Author: &User{
				Login:             String("author"),
				ID:                Int64(0),
				AvatarURL:         String("https://avatars1.githubusercontent.com/u/0?v=4"),
				GravatarID:        String(""),
				URL:               String("https://api.github.com/users/author"),
				HTMLURL:           String("https://github.com/author"),
				FollowersURL:      String("https://api.github.com/users/author/followers"),
				FollowingURL:      String("https://api.github.com/users/author/following{/other_user}"),
				GistsURL:          String("https://api.github.com/users/author/gists{/gist_id}"),
				StarredURL:        String("https://api.github.com/users/author/starred{/owner}{/repo}"),
				SubscriptionsURL:  String("https://api.github.com/users/author/subscriptions"),
				OrganizationsURL:  String("https://api.github.com/users/author/orgs"),
				ReposURL:          String("https://api.github.com/users/author/repos"),
				EventsURL:         String("https://api.github.com/users/author/events{/privacy}"),
				ReceivedEventsURL: String("https://api.github.com/users/author/received_events"),
				Type:              String("User"),
				SiteAdmin:         Bool(false),
			},
			Body:          String("test"),
			BodyHTML:      String("<p>test</p>"),
			BodyVersion:   String("version"),
			CommentsCount: Int(1),
			CommentsURL:   String("https://api.github.com/teams/2/discussions/3/comments"),
			CreatedAt:     &Timestamp{time.Date(2018, time.January, 1, 0, 0, 0, 0, time.UTC)},
			LastEditedAt:  nil,
			HTMLURL:       String("https://github.com/orgs/1/teams/2/discussions/3"),
			NodeID:        String("node"),
			Number:        Int(3),
			Pinned:        Bool(false),
			Private:       Bool(false),
			TeamURL:       String("https://api.github.com/teams/2"),
			Title:         String("test"),
			UpdatedAt:     &Timestamp{time.Date(2018, time.January, 1, 0, 0, 0, 0, time.UTC)},
			URL:           String("https://api.github.com/teams/2/discussions/3"),
		},
	}
	if !cmp.Equal(discussions, want) {
		t.Errorf("Teams.ListDiscussionsBySlug returned %+v, want %+v", discussions, want)
	}

	const methodName = "ListDiscussionsBySlug"
	testBadOptions(t, methodName, func() (err error) {
		_, _, err = client.Teams.ListDiscussionsBySlug(ctx, "o\no", "s\ns", nil)
		return err
	})

	testNewRequestAndDoFailure(t, methodName, client, func() (*Response, error) {
		got, resp, err := client.Teams.ListDiscussionsBySlug(ctx, "o", "s", &DiscussionListOptions{"desc", ListOptions{Page: 2}})
		if got != nil {
			t.Errorf("testNewRequestAndDoFailure %v = %#v, want nil", methodName, got)
		}
		return resp, err
	})
}

func TestTeamsService_GetDiscussionByID(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/organizations/1/team/2/discussions/3", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, `{"number":3}`)
	})

	ctx := context.Background()
	discussion, _, err := client.Teams.GetDiscussionByID(ctx, 1, 2, 3)
	if err != nil {
		t.Errorf("Teams.GetDiscussionByID returned error: %v", err)
	}

	want := &TeamDiscussion{Number: Int(3)}
	if !cmp.Equal(discussion, want) {
		t.Errorf("Teams.GetDiscussionByID returned %+v, want %+v", discussion, want)
	}

	const methodName = "GetDiscussionByID"
	testBadOptions(t, methodName, func() (err error) {
		_, _, err = client.Teams.GetDiscussionByID(ctx, -1, -2, -3)
		return err
	})

	testNewRequestAndDoFailure(t, methodName, client, func() (*Response, error) {
		got, resp, err := client.Teams.GetDiscussionByID(ctx, 1, 2, 3)
		if got != nil {
			t.Errorf("testNewRequestAndDoFailure %v = %#v, want nil", methodName, got)
		}
		return resp, err
	})
}

func TestTeamsService_GetDiscussionBySlug(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/orgs/o/teams/s/discussions/3", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, `{"number":3}`)
	})

	ctx := context.Background()
	discussion, _, err := client.Teams.GetDiscussionBySlug(ctx, "o", "s", 3)
	if err != nil {
		t.Errorf("Teams.GetDiscussionBySlug returned error: %v", err)
	}

	want := &TeamDiscussion{Number: Int(3)}
	if !cmp.Equal(discussion, want) {
		t.Errorf("Teams.GetDiscussionBySlug returned %+v, want %+v", discussion, want)
	}

	const methodName = "GetDiscussionBySlug"
	testBadOptions(t, methodName, func() (err error) {
		_, _, err = client.Teams.GetDiscussionBySlug(ctx, "o\no", "s\ns", -3)
		return err
	})

	testNewRequestAndDoFailure(t, methodName, client, func() (*Response, error) {
		got, resp, err := client.Teams.GetDiscussionBySlug(ctx, "o", "s", 3)
		if got != nil {
			t.Errorf("testNewRequestAndDoFailure %v = %#v, want nil", methodName, got)
		}
		return resp, err
	})
}

func TestTeamsService_CreateDiscussionByID(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	input := TeamDiscussion{Title: String("c_t"), Body: String("c_b")}

	mux.HandleFunc("/organizations/1/team/2/discussions", func(w http.ResponseWriter, r *http.Request) {
		v := new(TeamDiscussion)
		assertNilError(t, json.NewDecoder(r.Body).Decode(v))

		testMethod(t, r, "POST")
		if !cmp.Equal(v, &input) {
			t.Errorf("Request body = %+v, want %+v", v, input)
		}

		fmt.Fprint(w, `{"number":3}`)
	})

	ctx := context.Background()
	comment, _, err := client.Teams.CreateDiscussionByID(ctx, 1, 2, input)
	if err != nil {
		t.Errorf("Teams.CreateDiscussionByID returned error: %v", err)
	}

	want := &TeamDiscussion{Number: Int(3)}
	if !cmp.Equal(comment, want) {
		t.Errorf("Teams.CreateDiscussionByID returned %+v, want %+v", comment, want)
	}

	const methodName = "CreateDiscussionByID"
	testBadOptions(t, methodName, func() (err error) {
		_, _, err = client.Teams.CreateDiscussionByID(ctx, -1, -2, input)
		return err
	})

	testNewRequestAndDoFailure(t, methodName, client, func() (*Response, error) {
		got, resp, err := client.Teams.CreateDiscussionByID(ctx, 1, 2, input)
		if got != nil {
			t.Errorf("testNewRequestAndDoFailure %v = %#v, want nil", methodName, got)
		}
		return resp, err
	})
}

func TestTeamsService_CreateDiscussionBySlug(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	input := TeamDiscussion{Title: String("c_t"), Body: String("c_b")}

	mux.HandleFunc("/orgs/o/teams/s/discussions", func(w http.ResponseWriter, r *http.Request) {
		v := new(TeamDiscussion)
		assertNilError(t, json.NewDecoder(r.Body).Decode(v))

		testMethod(t, r, "POST")
		if !cmp.Equal(v, &input) {
			t.Errorf("Request body = %+v, want %+v", v, input)
		}

		fmt.Fprint(w, `{"number":3}`)
	})

	ctx := context.Background()
	comment, _, err := client.Teams.CreateDiscussionBySlug(ctx, "o", "s", input)
	if err != nil {
		t.Errorf("Teams.CreateDiscussionBySlug returned error: %v", err)
	}

	want := &TeamDiscussion{Number: Int(3)}
	if !cmp.Equal(comment, want) {
		t.Errorf("Teams.CreateDiscussionBySlug returned %+v, want %+v", comment, want)
	}

	const methodName = "CreateDiscussionBySlug"
	testBadOptions(t, methodName, func() (err error) {
		_, _, err = client.Teams.CreateDiscussionBySlug(ctx, "o\no", "s\ns", input)
		return err
	})

	testNewRequestAndDoFailure(t, methodName, client, func() (*Response, error) {
		got, resp, err := client.Teams.CreateDiscussionBySlug(ctx, "o", "s", input)
		if got != nil {
			t.Errorf("testNewRequestAndDoFailure %v = %#v, want nil", methodName, got)
		}
		return resp, err
	})
}

func TestTeamsService_EditDiscussionByID(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	input := TeamDiscussion{Title: String("e_t"), Body: String("e_b")}

	mux.HandleFunc("/organizations/1/team/2/discussions/3", func(w http.ResponseWriter, r *http.Request) {
		v := new(TeamDiscussion)
		assertNilError(t, json.NewDecoder(r.Body).Decode(v))

		testMethod(t, r, "PATCH")
		if !cmp.Equal(v, &input) {
			t.Errorf("Request body = %+v, want %+v", v, input)
		}

		fmt.Fprint(w, `{"number":3}`)
	})

	ctx := context.Background()
	comment, _, err := client.Teams.EditDiscussionByID(ctx, 1, 2, 3, input)
	if err != nil {
		t.Errorf("Teams.EditDiscussionByID returned error: %v", err)
	}

	want := &TeamDiscussion{Number: Int(3)}
	if !cmp.Equal(comment, want) {
		t.Errorf("Teams.EditDiscussionByID returned %+v, want %+v", comment, want)
	}

	const methodName = "EditDiscussionByID"
	testBadOptions(t, methodName, func() (err error) {
		_, _, err = client.Teams.EditDiscussionByID(ctx, -1, -2, -3, input)
		return err
	})

	testNewRequestAndDoFailure(t, methodName, client, func() (*Response, error) {
		got, resp, err := client.Teams.EditDiscussionByID(ctx, 1, 2, 3, input)
		if got != nil {
			t.Errorf("testNewRequestAndDoFailure %v = %#v, want nil", methodName, got)
		}
		return resp, err
	})
}

func TestTeamsService_EditDiscussionBySlug(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	input := TeamDiscussion{Title: String("e_t"), Body: String("e_b")}

	mux.HandleFunc("/orgs/o/teams/s/discussions/3", func(w http.ResponseWriter, r *http.Request) {
		v := new(TeamDiscussion)
		assertNilError(t, json.NewDecoder(r.Body).Decode(v))

		testMethod(t, r, "PATCH")
		if !cmp.Equal(v, &input) {
			t.Errorf("Request body = %+v, want %+v", v, input)
		}

		fmt.Fprint(w, `{"number":3}`)
	})

	ctx := context.Background()
	comment, _, err := client.Teams.EditDiscussionBySlug(ctx, "o", "s", 3, input)
	if err != nil {
		t.Errorf("Teams.EditDiscussionBySlug returned error: %v", err)
	}

	want := &TeamDiscussion{Number: Int(3)}
	if !cmp.Equal(comment, want) {
		t.Errorf("Teams.EditDiscussionBySlug returned %+v, want %+v", comment, want)
	}

	const methodName = "EditDiscussionBySlug"
	testBadOptions(t, methodName, func() (err error) {
		_, _, err = client.Teams.EditDiscussionBySlug(ctx, "o\no", "s\ns", -3, input)
		return err
	})

	testNewRequestAndDoFailure(t, methodName, client, func() (*Response, error) {
		got, resp, err := client.Teams.EditDiscussionBySlug(ctx, "o", "s", 3, input)
		if got != nil {
			t.Errorf("testNewRequestAndDoFailure %v = %#v, want nil", methodName, got)
		}
		return resp, err
	})
}

func TestTeamsService_DeleteDiscussionByID(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/organizations/1/team/2/discussions/3", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "DELETE")
	})

	ctx := context.Background()
	_, err := client.Teams.DeleteDiscussionByID(ctx, 1, 2, 3)
	if err != nil {
		t.Errorf("Teams.DeleteDiscussionByID returned error: %v", err)
	}

	const methodName = "DeleteDiscussionByID"
	testBadOptions(t, methodName, func() (err error) {
		_, err = client.Teams.DeleteDiscussionByID(ctx, -1, -2, -3)
		return err
	})

	testNewRequestAndDoFailure(t, methodName, client, func() (*Response, error) {
		return client.Teams.DeleteDiscussionByID(ctx, 1, 2, 3)
	})
}

func TestTeamsService_DeleteDiscussionBySlug(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/orgs/o/teams/s/discussions/3", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "DELETE")
	})

	ctx := context.Background()
	_, err := client.Teams.DeleteDiscussionBySlug(ctx, "o", "s", 3)
	if err != nil {
		t.Errorf("Teams.DeleteDiscussionBySlug returned error: %v", err)
	}

	const methodName = "DeleteDiscussionBySlug"
	testBadOptions(t, methodName, func() (err error) {
		_, err = client.Teams.DeleteDiscussionBySlug(ctx, "o\no", "s\ns", -3)
		return err
	})

	testNewRequestAndDoFailure(t, methodName, client, func() (*Response, error) {
		return client.Teams.DeleteDiscussionBySlug(ctx, "o", "s", 3)
	})
}

func TestTeamDiscussion_Marshal(t *testing.T) {
	testJSONMarshal(t, &TeamDiscussion{}, "{}")

	u := &TeamDiscussion{
		Author: &User{
			Login:       String("author"),
			ID:          Int64(0),
			URL:         String("https://api.github.com/users/author"),
			AvatarURL:   String("https://avatars1.githubusercontent.com/u/0?v=4"),
			GravatarID:  String(""),
			CreatedAt:   &Timestamp{referenceTime},
			SuspendedAt: &Timestamp{referenceTime},
		},
		Body:          String("test"),
		BodyHTML:      String("<p>test</p>"),
		BodyVersion:   String("version"),
		CommentsCount: Int(1),
		CommentsURL:   String("https://api.github.com/teams/2/discussions/3/comments"),
		CreatedAt:     &Timestamp{referenceTime},
		LastEditedAt:  &Timestamp{referenceTime},
		HTMLURL:       String("https://api.github.com/teams/2/discussions/3/comments"),
		NodeID:        String("A123"),
		Number:        Int(10),
		Pinned:        Bool(true),
		Private:       Bool(false),
		TeamURL:       String("https://api.github.com/teams/2/discussions/3/comments"),
		Title:         String("Test"),
		UpdatedAt:     &Timestamp{referenceTime},
		URL:           String("https://api.github.com/teams/2/discussions/3/comments"),
		Reactions: &Reactions{
			TotalCount: Int(1),
			PlusOne:    Int(2),
			MinusOne:   Int(-3),
			Laugh:      Int(4),
			Confused:   Int(5),
			Heart:      Int(6),
			Hooray:     Int(7),
			Rocket:     Int(8),
			Eyes:       Int(9),
			URL:        String("https://api.github.com/teams/2/discussions/3/comments"),
		},
	}

	want := `{
		"author": {
			"login": "author",
			"id": 0,
			"avatar_url": "https://avatars1.githubusercontent.com/u/0?v=4",
			"gravatar_id": "",
			"url": "https://api.github.com/users/author",
			"created_at": ` + referenceTimeStr + `,
			"suspended_at": ` + referenceTimeStr + `	
		},
		"body": "test",
		"body_html": "<p>test</p>",
		"body_version": "version",
		"comments_count": 1,
		"comments_url": "https://api.github.com/teams/2/discussions/3/comments",
		"created_at": ` + referenceTimeStr + `,
		"last_edited_at": ` + referenceTimeStr + `,
		"html_url": "https://api.github.com/teams/2/discussions/3/comments",
		"node_id": "A123",
		"number": 10,
		"pinned": true,
		"private": false,
		"team_url": "https://api.github.com/teams/2/discussions/3/comments",
		"title": "Test",
		"updated_at": ` + referenceTimeStr + `,
		"url": "https://api.github.com/teams/2/discussions/3/comments",
		"reactions": {
			"total_count": 1,
			"+1": 2,
			"-1": -3,
			"laugh": 4,
			"confused": 5,
			"heart": 6,
			"hooray": 7,
			"rocket": 8,
			"eyes": 9,
			"url": "https://api.github.com/teams/2/discussions/3/comments"
		}
	}`

	testJSONMarshal(t, u, want)
}
