module github.com/google/go-github/v52/example

go 1.17

require (
	github.com/bradleyfalzon/ghinstallation/v2 v2.0.4
	github.com/gofri/go-github-ratelimit v1.0.3
	github.com/google/go-github/v52 v52.0.0
	golang.org/x/crypto v0.7.0
	golang.org/x/oauth2 v0.7.0
	google.golang.org/appengine v1.6.7
)

require (
	github.com/golang-jwt/jwt/v4 v4.0.0 // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/google/go-github/v41 v41.0.0 // indirect
	github.com/google/go-querystring v1.1.0 // indirect
	golang.org/x/net v0.9.0 // indirect
	golang.org/x/sys v0.7.0 // indirect
	golang.org/x/term v0.7.0 // indirect
)

// Use version at HEAD, not the latest published.
replace github.com/google/go-github/v52 => ../
