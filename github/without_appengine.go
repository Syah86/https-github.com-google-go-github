// Copyright 2017 The go-github AUTHORS. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file provides glue for making github work without App Engine.

// +build !appengine

package github

import (
	"context"
	"net/http"
)

func addContext(ctx context.Context, req *http.Request) (context.Context, *http.Request) {
	return ctx, req.WithContext(ctx)
}
