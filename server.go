package main

import (
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
)

func StartServer(blog *Blog) {
	

	m := martini.Classic()

	// Setup middleware
	m.Use(render.Renderer())
	m.Use(martini.Static("assets"))
	
	// Routes
	m.Get("/", IndexHandler)
	m.Get("/articles", AllArticlesHandler)
	m.Get("/article/:id", ArticleHandler)
	m.Get("/:page", PageHandler)
	m.NotFound(NotFoundHandler)
	// Start Martini
	m.Run()
}

// Handlers
func IndexHandler(r render.Render) {
	r.HTML(200, "index", blog.NewestArticle)
}

func NotFoundHandler(r render.Render) {
	r.HTML(404, "404", "")
}

func AllArticlesHandler(r render.Render) {
	r.HTML(200, "archive", articles)
}

func ArticleHandler(params martini.Params, r render.Render) {
	a := GetArticle(params["id"])
	r.HTML(200, "article", a)
}

// PageHandler checks whether a page exists, if not it will serve a 404.
func PageHandler(params martini.Params, r render.Render) {
	page := params["page"]
	for _, p := range blog.Pages {
		if page == p {
			r.HTML(200, page, "")
		}
	}
	r.HTML(404, "404", "")
}