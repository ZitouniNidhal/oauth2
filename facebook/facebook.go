// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package facebook provides constants for using OAuth2 to access Facebook.
package facebook // import "golang.org/x/oauth2/facebook"

import (
	"golang.org/x/oauth2"
)

// Endpoint is Facebook's OAuth 2.0 endpoint.
//
// The previous hard-coded endpoints targeted v3.2 of the Graph API.
// To avoid pinning to a specific old version and to follow common
// practice when you don't need a fixed Graph API version, these
// endpoints omit the version segment so the app will use Facebook's
// default/current version. If you need to target a specific Graph
// API version for compatibility reasons, use the versioned URLs
// (for example, "https://graph.facebook.com/v17.0/oauth/access_token").
var Endpoint = oauth2.Endpoint{
	AuthURL:  "https://www.facebook.com/dialog/oauth",
	TokenURL: "https://graph.facebook.com/oauth/access_token",
}
