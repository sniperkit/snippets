// Copyright 2017 GoPic Authors. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"net/http"
	"path/filepath"

	"github.com/go-iris2/iris2"
	"github.com/go-iris2/iris2/adaptors/sessions"
	"github.com/go-iris2/iris2/adaptors/view"
	"github.com/xor-gate/gopic/lib/config"
	"github.com/xor-gate/gopic/lib/db"
	"github.com/xor-gate/gopic/lib/http/auth"
	"github.com/xor-gate/gopic/lib/pic"
)

// httpHandler registers and runs a HTTP server in a goroutine
func httpHandler() {
	f := iris2.New()
	f.Use(auth.New())
	f.Adapt(sessions.New(sessions.Config{
		Cookie: "GoPicSessionID",
	}))

	f.StaticWeb("/static", "./static")
	f.Adapt(view.HTML("./templates", ".html"))

	f.OnError(http.StatusNotFound, func(ctx *iris2.Context) {
		ctx.Render("errors/404.html", nil)
	})

	f.OnError(http.StatusInternalServerError, func(ctx *iris2.Context) {
		ctx.Render("errors/500.html", nil)
	})

	f.Get("/", admin)
	f.Get("/admin/users", adminUsers)

	f.Get("/login", login)
	f.Post("/login", login)

	f.Get("/signup", signup)
	f.Post("/signup", signup)

	f.Get("/by-id/:id", byID)
	f.Get("/by-id-resized/:id/:size", byIDResized)
	f.Get("/by-folder-id/:id", byFolderID)

	host := config.Host()
	if host == "0.0.0.0" {
		host = ""
	}
	go f.Listen(host + ":" + config.Port())
}

func signup(ctx *iris2.Context) {
	switch ctx.Method() {
	case iris2.MethodGet:
		ctx.MustRender("signup/index.html", nil)
	case iris2.MethodPost:
		email := ctx.FormValue("email")
		password := ctx.FormValue("password")
		err := db.UserSave(email, password)
		if err != nil {
			fmt.Println(err)
			ctx.EmitError(http.StatusInternalServerError)
			return
		}
		ctx.MustRender("signup/approval.html", nil)
	}
}

func login(ctx *iris2.Context) {
	switch ctx.Method() {
	case iris2.MethodGet:
		ctx.MustRender("login.html", nil)
	case iris2.MethodPost:
		email := ctx.FormValue("email")
		password := ctx.FormValue("password")
		err := db.UserVerify(email, password)
		if err == nil || (err == db.ErrorUserIsInactive && email == "demo@demo.com") {
			auth.SetSessionAuthenticated(ctx, true)
			redirect := ctx.Session().GetString("redirect")
			ctx.Session().Set("redirect", "/")
			ctx.Redirect(redirect)
		}
		ctx.MustRender("signup/approval.html", nil)
	}
}

// admin handler for dashboard endpoint
func admin(ctx *iris2.Context) {
	paths, err := db.GetPicFilePaths()
	if err != nil {
		ctx.NotFound()
		return
	}

	// Preserve only the basename of the path
	for n, p := range *paths {
		(*paths)[n].Path = filepath.Base(p.Path)
	}

	// TODO: add picture count per path in view
	ctx.Render("admin/index.html", iris2.Map{"Paths": paths, "Title": "GoPic Admin Panel"})
}

// admin handler for dashboard endpoint
func adminUsers(ctx *iris2.Context) {
	users, err := db.GetUsers()
	if err != nil {
		ctx.NotFound()
		return
	}

	ctx.Render("admin/users.html", iris2.Map{"Users": users})
}

// byFolderID handler for folder index/gallery endpoint
func byFolderID(ctx *iris2.Context) {
	p, err := db.GetPicFilePath(ctx.Param("id"))
	if err != nil {
		ctx.NotFound()
		return
	}

	view := &struct {
		Title  string
		Images []pic.File
	}{
		Title: filepath.Base(p.Path),
	}

	for _, id := range p.HashList {
		f, err := db.GetPicFileByHash(id)
		if err == nil {
			view.Images = append(view.Images, *f)
		}
	}

	if len(view.Images) == 0 {
		ctx.NotFound()
		return
	}

	ctx.Render("by-folder-id.html", view)
}

// byID handler for single picture endpoint
func byID(ctx *iris2.Context) {
	f, err := db.GetPicFileByHash(ctx.Param("id"))
	if err != nil {
		ctx.NotFound()
		return
	}
	ctx.ServeFile(f.Filepath, false)
}

// byIDResized handler for single thumbnailed picture endpoint
func byIDResized(ctx *iris2.Context) {
	f, err := db.GetPicFileByHash(ctx.Param("id"))
	if err != nil {
		ctx.NotFound()
		return
	}

	maxWidth, err := ctx.ParamInt("size")
	if err != nil {
		ctx.NotFound()
		return
	}

	err = f.Thumbnail(maxWidth)
	if err != nil {
		ctx.EmitError(http.StatusInternalServerError)
		return
	}

	path := f.ThumbnailPath(maxWidth)
	ctx.ServeFile(path, false)
}
