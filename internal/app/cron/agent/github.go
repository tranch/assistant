package agent

import (
	"context"
	"crypto/sha1" // #nosec
	"github.com/tsundata/assistant/api/pb"
	"github.com/tsundata/assistant/internal/app/cron/pipeline/result"
	"github.com/tsundata/assistant/internal/pkg/rulebot"
	"github.com/tsundata/assistant/internal/pkg/util"
	"github.com/tsundata/assistant/internal/pkg/vendors/github"
)

func FetchGithubStarred(ctx context.Context, comp rulebot.IComponent) []result.Result {
	if comp.Middle() == nil {
		return []result.Result{result.EmptyResult()}
	}
	// get access token
	app, err := comp.Middle().GetAvailableApp(ctx, &pb.TextRequest{Text: github.ID})
	if err != nil {
		return []result.Result{result.ErrorResult(err)}
	}
	accessToken := app.GetToken()
	if accessToken == "" {
		return []result.Result{result.EmptyResult()}
	}

	// data
	client := github.NewGithub("", "", "", accessToken)
	user, err := client.GetUser()
	if err != nil {
		return []result.Result{result.ErrorResult(err)}
	}
	if *user.Login == "" {
		return []result.Result{result.EmptyResult()}
	}

	repos, err := client.GetStarred(*user.Login)
	if err != nil {
		return []result.Result{result.ErrorResult(err)}
	}
	var r []result.Result
	for _, item := range *repos {
		s := sha1.New()                           // #nosec
		s.Write(util.StringToByte(*item.HTMLURL)) // #nosec
		bs := s.Sum(nil)
		r = append(r, result.Result{
			ID:   util.ByteToString(bs),
			Kind: result.Repos,
			Content: map[string]string{
				"name":  *item.FullName,
				"owner": *item.Owner.Login,
				"repo":  *item.Name,
				"url":   *item.HTMLURL,
			},
		})
	}
	return r
}

func FetchGithubStargazers(_ context.Context, _ rulebot.IComponent) []result.Result {
	// todo
	return []result.Result{}
}
