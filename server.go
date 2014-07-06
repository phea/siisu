package main

import (
	"io"
	"net/http"

	"github.com/zenazn/goji"
	"github.com/zenazn/goji/web"
)

func StartServer(blog *Blog) {
	// Routes
	goji.Get("/", IndexHandler)
	goji.Get("/articles", AllArticlesHandler)
	goji.Get("/article/:id", ArticleHandler)
	goji.Get("/:page", PageHandler)

	goji.Get("/assets/*", http.FileServer(http.Dir("public")))
	goji.NotFound(NotFoundHandler)

	// Start Goji
	goji.Serve()
}

// Handlers
func IndexHandler(c web.C, w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, blog.Pages["index"])
}

func NotFoundHandler(c web.C, w http.ResponseWriter, r *http.Request) {
	http.Error(w, blog.Pages["404"], 404)
}

func AllArticlesHandler(c web.C, w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, blog.Pages["archive"])
}

func ArticleHandler(c web.C, w http.ResponseWriter, r *http.Request) {
	id := c.URLParams["id"]
	article, ok := blog.Pages[id]
	if !ok {
		http.Error(w, blog.Pages["404"], 404)
		return
	}

	io.WriteString(w, article)
}

// PageHandler serves pages other than articles, ie home, about.
func PageHandler(c web.C, w http.ResponseWriter, r *http.Request) {
	p := c.URLParams["page"]
	page, ok := blog.Pages[p]
	if !ok {
		http.Error(w, blog.Pages["404"], 404)
		return
	}

	io.WriteString(w, page)
}
