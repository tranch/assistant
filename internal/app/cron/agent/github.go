package agent

import (
	"context"
	"github.com/tsundata/assistant/api/pb"
	"github.com/tsundata/assistant/internal/app/cron/pipeline/result"
	"github.com/tsundata/assistant/internal/pkg/rulebot"
	"github.com/tsundata/assistant/internal/pkg/utils"
	"github.com/tsundata/assistant/internal/pkg/vendors/github"
)

func FetchGithubStarred(b *rulebot.RuleBot) []result.Result {
	// get access token
	app, err := b.MidClient.GetAvailableApp(context.Background(), &pb.TextRequest{Text: github.ID})
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
		r = append(r, result.Result{
			ID:   utils.SHA1(*item.HTMLURL),
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
