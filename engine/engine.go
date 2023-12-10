package engine

import (
	"fmt"
	"net/http"
)

type HandlerFunc func(ctx *PhoenixContext)

type Group struct {
	prefix      string
	middlewares []HandlerFunc
}

type Engine struct {
	routerMap IRouterMap
	group     *Group
}

func (e *Engine) Group(prefix string) *Engine {
	e.group.prefix = prefix
	return e
}
func NewPhoenixServer() *Engine {
	return &Engine{
		group:     &Group{},
		routerMap: NewCRouterPrefixTree(),
	}
}

func (e *Engine) GET(path string, handler HandlerFunc) {
	groupPath := e.group.prefix + path
	e.routerMap.InsertRoute(groupPath, handler)
}

func (e *Engine) Run(addr string) {
	http.ListenAndServe(addr, e)
}

func (e *Engine) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	newContext := NewPhoenixContext(res, req)
	fmt.Println("context", newContext)
	path := newContext.Path
	handlerSearched, params := e.routerMap.SearchRoute(path)
	if handlerSearched != nil {
		newContext.Params = params
		e.group.middlewares = append(e.group.middlewares, handlerSearched)
	} else {
		handlerNotFound := HandlerFunc(func(ctx *PhoenixContext) {
			ctx.JSON(404, "Not Found URL")
		})
		e.group.middlewares = append(e.group.middlewares, handlerNotFound)
	}
	newContext.middlewares = e.group.middlewares
	fmt.Println("run time address middleware: ", e.group.middlewares[0])
	newContext.Next()
}

func (e *Engine) Use(handlers ...HandlerFunc) {
	e.group.middlewares = append(e.group.middlewares, handlers...)
	fmt.Println("Add middleware success: ", e.group.middlewares[0])
}
