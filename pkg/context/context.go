package context

import (
	"gopkg.in/macaron.v1"
	log "github.com/go-clog/clog"
)

// Context represents context of a request
type Context struct {
	*macaron.Context
}

// HTML responses template with given status.
func (ctx *Context) HTML(status int, name string) {
	log.Trace("Template: %s", name)
	ctx.Context.HTML(status, name)
}
