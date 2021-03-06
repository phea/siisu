package main

import (
	"bytes"
	"html/template"
	"io/ioutil"
	"log"
)

type Blog struct {
	// Blog holds the pre-processed pages
	Pages map[string]string
}

var blog Blog
var articles []Article

func main() {
	blog.Pages = make(map[string]string)
	// load articles in ./articles directory
	articles = AllArticles()

	// Preprocess templates.
	data := make(map[string]interface{})
	for _, a := range articles {
		var buf bytes.Buffer
		data["Title"] = a.Title
		data["Date"] = a.Date
		data["Body"] = a.Body

		t, err := template.ParseFiles("templates/article.tmpl")
		if err != nil {
			panic(err)
		}

		err = t.Execute(&buf, data)
		if err != nil {
			panic(err)
		}

		blog.Pages[a.Slug] = buf.String()
	}

	// Find last article to display for index page
	latest := articles[len(articles)-1]
	blog.Pages["index"] = blog.Pages[latest.Slug]

	// Load pages
	blog.Pages["about"] = "About" // TODO: read the about.tmpl

	// create sitemap.xml
	err := writeSitemap(createSitemap())
	if err != nil {
		log.Println(err)
	}

	// Start Http server
	StartServer(&blog)
}

func createSitemap() []byte {
	sitemap := []byte("http://siisu.com\n")
	for k, _ := range blog.Pages {
		url := "http://siisu.com/" + k + "\n"
		sitemap = append(sitemap, []byte(url)...)
	}

	return sitemap
}

func writeSitemap(data []byte) error {
	err := ioutil.WriteFile("public/sitemap.xml", data, 0644)
	if err != nil {
		return err
	}

	return nil
}
