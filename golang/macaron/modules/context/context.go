package context

import (
	"gopkg.in/macaron.v1"
)

type Context struct {
	*macaron.Context
}

func NotFound(ctx *Context) {
	ctx.Data["Title"] = "Page Not Found"
	ctx.HTML(404, "status/404")
}

func (ctx *Context) NotFound() {
	NotFound(ctx)
}

func Contexter() macaron.Handler {
	return func(c *macaron.Context) {
		ctx := &Context{
			Context: c,
		}
		c.Map(ctx)
	}
}
