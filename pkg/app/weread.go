package app

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"time"
	"weread/pkg/card"
	"weread/pkg/github"
	"weread/pkg/stats/weread"

	"github.com/sirupsen/logrus"
)

const (
	ReadLimit = 5
)

type Stats interface {
	Run(ctx context.Context) error
}

type WeReadStats struct {
	gistId string
	gitCli *github.Github
}

func NewWeReadStats(
	gistId string,
	gitCli *github.Github,
) *WeReadStats {
	return &WeReadStats{
		gistId: gistId,
		gitCli: gitCli,
	}
}

func (s *WeReadStats) Run(ctx context.Context) error {
	stats, err := weread.New()
	if err != nil {
		logrus.Error("New stats error.", err)
		return err
	}

	books, err := stats.ShelfBooks()
	if err != nil {
		logrus.Error("Get shelf books error.", err)
		return err
	}
	if len(books) == 0 {
		logrus.Error("Book list empty", err)
		return nil
	}

	readCard := GenerateReadCard(books)
	err = s.gitCli.EditGist(ctx, s.gistId, "weread-read.svg", readCard)
	if err != nil {
		logrus.Error("Run error.", err)
		return err
	}

	readingCard := GenerateReadingCard(books)
	err = s.gitCli.EditGist(ctx, s.gistId, "weread-reading.svg", readingCard)
	if err != nil {
		logrus.Error("Run error.", err)
		return err
	}
	return nil
}

func GenerateReadCard(books []weread.Book) string {
	list, total := wereadList(books, ReadLimit, func(b weread.Book) bool {
		return b.FinishReading == 1
	})
	c := card.NewListCard(card.Config{
		Title:      "微信读书-已读",
		Width:      500,
		LineHeight: 25,
	})
	for _, book := range list {
		title := book.Title
		if len([]rune(title)) > 20 {
			title = string([]rune(title)[:20])
		}
		c.AddContent(fmt.Sprintf("%s  -- %s", title, book.Author))
	}
	if len(list) > ReadLimit {
		c.AddContent(fmt.Sprintf("...等 %d 本书", total))
	}

	return c.Svg()
}

func GenerateReadingCard(books []weread.Book) string {
	now := time.Now().AddDate(0, -1, 0).Unix()
	list, _ := wereadList(books, ReadLimit, func(b weread.Book) bool {
		return b.FinishReading != 1 && b.Progress > 0 && b.Progress < 100 && int64(b.ReadUpdateTime) > now
	})

	c := card.NewTableCard(
		card.Config{
			Title:      "微信读书-在读",
			Width:      500,
			LineHeight: 25,
		},
		[]card.Column{
			{
				Title: "",
				Width: .8,
			},
			{
				Title: "",
				Width: .2,
				Style: "font-size: 12px",
			},
		})
	for _, book := range list {
		title := book.Title
		if len([]rune(title)) > 20 {
			title = string([]rune(title)[:20])
		}
		c.AddContent([]string{strings.TrimSpace(title), strconv.Itoa(book.Progress) + "%"})
	}

	return c.Svg()
}

func wereadList(books []weread.Book, limit int, choose func(b weread.Book) bool) (res []weread.Book, total int) {
	for _, book := range books {
		if choose(book) {
			total += 1
			if len(res) > limit {
				continue
			}
			res = append(res, book)
		}
	}
	return res, total
}
