package github

import (
	"context"
	"errors"

	"github.com/google/go-github/github"
)

type Github struct {
	cli *github.Client
}

func New(username, token string) *Github {
	tp := github.BasicAuthTransport{
		Username: username,
		Password: token,
	}
	return &Github{cli: github.NewClient(tp.Client())}
}

func (g *Github) EditGist(ctx context.Context, id string, name, content string) error {
	gist, _, err := g.cli.Gists.Get(ctx, id)
	if err != nil {
		return err
	}

	file, ok := gist.Files[github.GistFilename(name)]
	if !ok {
		return errors.New("Gist file NotFound. ")
	}

	file.Content = github.String(content)
	gist.Files[github.GistFilename(name)] = file

	_, _, err = g.cli.Gists.Edit(ctx, id, gist)
	if err != nil {
		return err
	}
	return nil
}
