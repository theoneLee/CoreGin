package core

import (
	"fmt"
	"net/http"
	"sync"
)

type Engine struct {
	RouterGroup
	router   map[string]HandlerList
	Handlers []HandlerFun
	pool     sync.Pool
}

/*
实现了Handler接口的对象可以注册到HTTP服务端，为特定的路径及其子树提供服务。

ServeHTTP应该将回复的头域和数据写入ResponseWriter接口然后返回。返回标志着该请求已经结束，
HTTP服务端可以转移向该连接上的下一个请求。
*/
func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	//panic("implement me")
	c := engine.pool.Get().(*Context)
	c.ResponseWriter = w
	c.Request = req
	engine.handleHTTPRequest(c)

	engine.pool.Put(c)
}

func NewEngine() *Engine {
	en := new(Engine)
	en.router = make(map[string]HandlerList)
	en.pool.New = func() interface{} {
		return en.allocateContext()
	}
	en.RouterGroup = RouterGroup{
		basePath: "/",
		Handlers: nil,
		engine:   en,
	}

	return en
}

func (engine *Engine) Run(addr string) (err error) {
	fmt.Println("Listening and serving HTTP on", addr)
	err = http.ListenAndServe(addr, engine)
	return
}

func (engine *Engine) handleHTTPRequest(c *Context) {
	httpMethod := c.Request.Method
	path := c.Request.URL.Path
	if handlers, ok := engine.router[httpMethod+"^"+path]; ok {
		for _, h := range handlers {
			h(c)
			if c.isAbort {
				return
			}
		}
	}
}

func (engine *Engine) allocateContext() *Context {
	return &Context{engine: engine}
}

func (engine *Engine) addRoute(httpMethod, absolutePath string, handlers HandlerList) {
	engine.router[httpMethod+"^"+absolutePath] = handlers
}
