package gee

import "net/http"

// Engine 实际就是server的处理器，统一处理，调用router的handle，从map中根据url查找handler
// 再次封装，由groupRouter来实现路由的功能
type Engine struct {
	router *router
	//engine 也是groupRouter，拥有gr的字段和方法，可直接访问
	*RouterGroup
	groupRouters []*RouterGroup
}

func New() *Engine {
	engine := &Engine{
		router: newRouter(),
	}
	engine.RouterGroup = &RouterGroup{
		//指向自己，方便后面的继承
		engine: engine,
		prefix: "",
	}
	//engine的groupRouter是全局的filter（java）
	//append 即使没有make也会make后添加
	engine.groupRouters = append(engine.groupRouters, engine.RouterGroup)
	return engine
}

func (e *Engine) RUN(addr string) {
	http.ListenAndServe(addr, e)
}
