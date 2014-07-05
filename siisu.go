package main

import (
	"io/ioutil"
	"log"
)

type Blog struct {
	Articles      map[string]Article
	Pages         []string
	NewestArticle Article // this should probably be a function.
}

// Global blog variable, pretty smelly.
var blog Blog
var articles []Article

func main() {
	// load articles in ./articles directory
	articles = AllArticles()

	// init the blog struct
	blog.Articles = make(map[string]Article)
	blog.NewestArticle = articles[len(articles)-1]
	// in a perfect world this will walk through the templates folder.
	blog.Pages = []string{"about"}
	for _, a := range articles {
		blog.Articles[a.Slug] = a
	}

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
	for _, page := range blog.Pages {
		url := "http://siisu.com/" + page + "\n"
		sitemap = append(sitemap, []byte(url)...)
	}

	for _, article := range articles {
		url := "http://siisu.com/" + article.Slug + "\n"
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
