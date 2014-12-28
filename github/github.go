package github

import (
	"code.google.com/p/goauth2/oauth"
	"github.com/google/go-github/github"
)

func NewGitHubClient(accessToken string) *github.Client {
	t := &oauth.Transport{
		Token: &oauth.Token{AccessToken: accessToken},
	}

	return github.NewClient(t.Client())
}
