package gee

import (
	"log"
	"net/http"
)

type HandlerFunc func(c *Context)

//router 实际路由的处理器
type router struct {
	handlers map[string]HandlerFunc
}

func newRouter() *router {
	return &router{
		handlers: make(map[string]HandlerFunc),
	}
}

func (r *router) addRouter(method string, pattern string, handler HandlerFunc) {
	log.Printf("Route %s - %s", method, pattern)
	key := method + "-" + pattern
	r.handlers[key] = handler
}

func (r *router) handle(c *Context) {
	key := c.Method + "-" + c.Path
	if handler, ok := r.handlers[key]; ok {
		c.handlers = append(c.handlers, handler)
		//真正开始处理请求
		c.Next()
	} else {
		c.String(http.StatusNotFound, "404 NOT FOUND %s", c.Path)
	}
}
