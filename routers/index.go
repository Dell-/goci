package routers

import (
	"github.com/Dell-/goci/pkg/context"
)

const INDEX = "index"

func Index(ctx *context.Context) {
	ctx.HTML(200, INDEX)
}
