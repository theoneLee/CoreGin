package app

import "path"

type RouterGroup struct {
	Handlers []HandlerFun
	engine   *Engine
	basePath string
}

type IRouter interface {
	Use(...HandlerFun) IRouter
	Group(string, ...HandlerFun) *RouterGroup

	GET(string, ...HandlerFun) IRouter
	POST(string, ...HandlerFun) IRouter
	PUT(string, ...HandlerFun) IRouter
	DELETE(string, ...HandlerFun) IRouter
}

type Message struct {
	Message string
}

func (routerGroup *RouterGroup) Use(handlers ...HandlerFun) IRouter {
	routerGroup.Handlers = append(routerGroup.Handlers, handlers...)
	return routerGroup
}

func (routerGroup *RouterGroup) Group(path string, handlers ...HandlerFun) *RouterGroup {
	rg := RouterGroup{}
	rg.Handlers = routerGroup.CombineHandlers(handlers)
	rg.basePath = path
	rg.engine = routerGroup.engine
	return &rg
}

func (routerGroup *RouterGroup) GET(path string, handlers ...HandlerFun) IRouter {
	routerGroup.handler("GET", path, handlers)

	return routerGroup
}

func (routerGroup *RouterGroup) POST(string, ...HandlerFun) IRouter {
	panic("implement me")
}

func (routerGroup *RouterGroup) PUT(string, ...HandlerFun) IRouter {
	panic("implement me")
}

func (routerGroup *RouterGroup) DELETE(string, ...HandlerFun) IRouter {
	panic("implement me")
}

func (routerGroup *RouterGroup) handler(httpMethod string, relativePath string, handlers []HandlerFun) IRouter {
	absolutePath := routerGroup.calculateAbsolutePath(relativePath)
	handlers = routerGroup.CombineHandlers(handlers)
	routerGroup.engine.addRoute(httpMethod, absolutePath, handlers)
	return routerGroup
}

func (routerGroup *RouterGroup) CombineHandlers(handlers HandlerList) HandlerList {
	finalSize := len(routerGroup.Handlers) + len(handlers)
	mergedHandler := make(HandlerList, finalSize)
	copy(mergedHandler, routerGroup.Handlers)
	copy(mergedHandler[len(routerGroup.Handlers):], handlers)
	return mergedHandler
}

func (routerGroup *RouterGroup) calculateAbsolutePath(relativePath string) string {
	return joinPaths(routerGroup.basePath, relativePath)
}

func joinPaths(absolutePath, relativePath string) string {
	if relativePath == "" {
		return absolutePath
	}

	finalPath := path.Join(absolutePath, relativePath)

	appendSlash := lastChar(relativePath) == '/' && lastChar(finalPath) != '/'
	if appendSlash {
		return finalPath + "/"
	}

	return finalPath
}

//工具方法 获取字符串最后一个字符
func lastChar(str string) uint8 {
	if str == "" {
		panic("The length of the string can't be 0")
	}

	return str[len(str)-1]
}
