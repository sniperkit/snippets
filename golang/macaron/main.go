package main

import (
	"net/http"

	"gopkg.in/macaron.v1"
	//"github.com/xor-gate/go-by-example/macaron/models"
	"github.com/xor-gate/go-by-example/macaron/routers"
	"github.com/xor-gate/go-by-example/macaron/modules/template"
	"github.com/xor-gate/go-by-example/macaron/modules/context"
)

func main() {
	m := macaron.Classic()
	m.Use(macaron.Static("static"))
	m.Use(macaron.Renderer(macaron.RenderOptions{
		Funcs:      template.NewFuncMap(),
		IndentJSON: macaron.Env != macaron.PROD,
	}))
	m.Use(context.Contexter())

	m.Get("/", func(ctx *macaron.Context) { ctx.Redirect("/dashboard") })
	m.Group("", func() {
		m.Get("/dashboard", routers.Dashboard)
	})

	m.NotFound(context.NotFound)

	http.ListenAndServe("0.0.0.0:1337", m)
}
