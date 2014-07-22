package main

import (
	"github.com/arkors/log/handler"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"net/http"
)

func main() {
	m := martini.Classic()
	m.Use(render.Renderer())
	m.Group("/v1", func(r martini.Router) {
		r.Post("/log", handler.CreateLog)
	})
	http.ListenAndServe(":3001", m)
}
