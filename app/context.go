package app

import "net/http"

type Context struct {
	Request        *http.Request
	ResponseWriter http.ResponseWriter
	engine         *Engine
	isAbort        bool
}

type HandlerFun func(ctx *Context)

type HandlerList []HandlerFun
