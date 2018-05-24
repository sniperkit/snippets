// Copyright 2017 GoPic Authors. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package auth

import (
	"github.com/go-iris2/iris2"
	"strings"
)

type Auth struct {
}

func New() *Auth {
	return &Auth{}
}

func IsSessionAuthenticated(ctx *iris2.Context) bool {
	if val, _ := ctx.Session().GetBoolean("authenticated"); val {
		return true
	}
	return false
}

func SetSessionAuthenticated(ctx *iris2.Context, value bool) {
	ctx.Session().Set("authenticated", value)
}

func (a *Auth) Serve(ctx *iris2.Context) {
	if IsSessionAuthenticated(ctx) {
		ctx.Next()
	} else if strings.HasPrefix(ctx.RequestPath(false), "/static") ||
		strings.HasPrefix(ctx.RequestPath(false), "/login") ||
		strings.HasPrefix(ctx.RequestPath(false), "/signup") {
		ctx.Next()
	} else if !strings.HasPrefix(ctx.RequestPath(false), "/login") {
		ctx.Session().Set("redirect", ctx.RequestPath(false))
		ctx.Redirect("/login")
	}
}
