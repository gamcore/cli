package api

import (
	"fmt"
	"net/http"
	"time"

	"github.com/google/go-github/v44/github"
)

var (
	httpClient   http.Client
	githubClient github.Client
)

type proxyTransport struct {
	transport http.RoundTripper
}

func (t *proxyTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	req.Header.Set("User-Agent", fmt.Sprintf("Goo App Manager v%s / Revision %s / Build Time %s", Version, Revision, GetTimestamp().UTC().Format(time.RFC822)))
	return t.transport.RoundTrip(req)
}

func init() {
	httpClient = http.Client{
		Transport: &proxyTransport{transport: http.DefaultTransport},
	}
	githubClient = *github.NewClient(&httpClient)
}
