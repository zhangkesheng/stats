package weread

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"sort"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
)

type Stats struct {
	cookies map[string]*http.Cookie
}

func New() (*Stats, error) {
	cookie := os.Getenv("WEREAD_COOKIE")
	if len(cookie) == 0 {
		return nil, errors.New("weread cookie not found ")
	}

	cookies := make(map[string]*http.Cookie)
	for _, s := range strings.Split(cookie, ";") {
		c := strings.Split(s, "=")
		name := strings.TrimSpace(c[0])
		cookies[name] = &http.Cookie{Name: name, Value: c[1]}
	}

	return &Stats{
		cookies,
	}, nil
}

func (s *Stats) refresh() {
	if s.cookies["wr_skey"] != nil && !s.cookies["wr_skey"].Expires.IsZero() && s.cookies["wr_skey"].Expires.Before(time.Now()) {
		return
	}

	logrus.Info("Refresh rewead cookies.")

	req, _ := http.NewRequest("GET", "https://weread.qq.com/", nil)
	req.Header["accept"] = []string{"text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9"}
	req.Header["accept-encoding"] = []string{"gzip, deflate, br"}
	req.Header["accept-language"] = []string{"zh-CN,zh;q=0.9,en;q=0.8,en-GB;q=0.7,en-US;q=0.6"}
	req.Header["cache-control"] = []string{"max-age=0"}
	req.Header["sec-fetch-dest"] = []string{"document"}
	req.Header["sec-fetch-mode"] = []string{"navigate"}
	req.Header["sec-fetch-site"] = []string{"none"}
	req.Header["sec-fetch-user"] = []string{"?1"}
	req.Header["upgrade-insecure-requests"] = []string{"1"}
	req.Header["user-agent"] = []string{"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/84.0.4147.105 Safari/537.36 Edg/84.0.522.52"}

	for _, cookie := range s.cookies {
		req.AddCookie(cookie)
	}

	client := http.Client{
		Timeout: time.Minute,
	}
	resp, err := client.Do(req)
	if err != nil {
		logrus.Panic("Init cookie error.", err)
	}

	for _, c := range resp.Cookies() {
		s.cookies[c.Name] = c
	}
}

func (s *Stats) ShelfBooks() ([]Book, error) {
	s.refresh()

	req, _ := http.NewRequest("GET", fmt.Sprintf("https://i.weread.qq.com/shelf/sync"), nil)
	for _, cookie := range s.cookies {
		req.AddCookie(cookie)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)

	var shelf ShelfBookResp
	err = json.Unmarshal(body, &shelf)
	if err != nil {
		return nil, err
	}
	books := append(shelf.Books, shelf.LectureBooks...)
	sort.Slice(books, func(i, j int) bool {
		return books[i].ReadUpdateTime > books[j].ReadUpdateTime
	})

	var res []Book
	processMap := make(map[string]int, len(shelf.BookProgress))
	for _, p := range shelf.BookProgress {
		processMap[p.BookID] = p.Progress
	}
	for _, book := range books {
		res = append(res, Book{
			Id:             book.BookID,
			Title:          book.Title,
			Cover:          book.Cover,
			Author:         book.Author,
			FinishReading:  book.FinishReading,
			ReadUpdateTime: book.ReadUpdateTime,
			Progress:       processMap[book.BookID],
		})
	}

	return res, nil
}

type Book struct {
	Id             string
	Title          string
	Cover          string
	Author         string
	FinishReading  int
	ReadUpdateTime int
	Progress       int
}
