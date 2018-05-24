package routers

import  (
	"github.com/xor-gate/go-by-example/macaron/modules/context"
)

func Dashboard(ctx *context.Context) {
	ctx.Data["Title"] = "Boem!"
	ctx.Data["List"]  = []string{"hello", "world"}

	ctx.HTML(200, "dashboard")
}
