package main

import (
	"context"
	"os"
	"weread/pkg/app"
	"weread/pkg/github"

	"github.com/sirupsen/logrus"
)

func main() {
	args := os.Args

	ctx := context.Background()
	githubCli := github.New(os.Getenv("GT_USER"), os.Getenv("GT_TOKEN"))
	gistId := os.Getenv("GIST_ID")

	switch args[1] {
	case "weread":
		stats := app.NewWeReadStats(gistId, githubCli)
		if err := stats.Run(ctx); err != nil {
			logrus.Panic("Run weRead stats error.", err)
		}
		break
	default:
		logrus.Fatal("Not supported command")
	}
}
