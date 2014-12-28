package github

import (
	"github.com/google/go-github/github"
	"code.google.com/p/goauth2/oauth"
)

func NewGitHubClient(accessToken string) *github.Client {
	t := &oauth.Transport{
		Token: &oauth.Token{AccessToken: accessToken},
	}

	return github.NewClient(t.Client())
}
