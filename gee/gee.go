package gee

import "net/http"

type Engine struct {
	router *router
}

func New() *Engine {
	return &Engine{
		router: newRouter(),
	}
}

func (e *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	c := newContext(w, req)
	e.router.handle(c)
}

func (e *Engine) Get(pattern string, handler HandlerFunc) {
	e.router.addRouter("GET", pattern, handler)
}

func (e *Engine) POST(pattern string, handler HandlerFunc) {
	e.router.addRouter("POST", pattern, handler)
}

func (e *Engine) RUN(addr string) {
	http.ListenAndServe(addr, e)
}
