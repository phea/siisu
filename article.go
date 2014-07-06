package main

import (
	"bytes"
	"html/template"
	"io/ioutil"

	"github.com/russross/blackfriday"
)

type Article struct {
	Title string
	Body  template.HTML
	Slug  string
	Date  string
}

// AllArticles returns a list of all articles in the ./articles folder.
func AllArticles() []Article {
	var articles []Article
	files, err := ioutil.ReadDir("./articles")
	if err != nil {
		return nil
	}

	for _, f := range files {
		a, _ := parseArticle(f.Name())
		if err != nil {
			return nil
		}
		articles = append(articles, a)
	}

	return articles
}

// parseArticle parses a .md file and returns an Article object.
// Date and Slug are taken from the filename, Date is in YYYY-MM-DD format.
// Title is the first line and Body is the rest.
func parseArticle(filename string) (Article, error) {
	var a Article
	a.Date = filename[:10]
	a.Slug = filename[11 : len(filename)-3]

	f, err := ioutil.ReadFile("./articles/" + filename)
	if err != nil {
		return a, err
	}

	index := bytes.Index(f, []byte("\n"))
	a.Title = string(f[:index])
	a.Body = template.HTML(blackfriday.MarkdownCommon(f[index+1:]))

	return a, nil
}
