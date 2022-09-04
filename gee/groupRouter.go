package gee

import (
	"log"
	"net/http"
	"strings"
)

// RouterGroup
// 位于接受请求后，位与最顶级，相当于加一层。将url分组处理(分组路由)的同时，可以在这一层加中间件
type RouterGroup struct {
	prefix      string
	middlewares []HandlerFunc
	//parent      *RouterGroup //作用：支持分组嵌套的前提：知道父亲是谁
	child []*RouterGroup
	//作用：所有资源集中在Engine，调用engine存储group。Engine就类似与javaWeb中的ServletContext
	//group 最后还是调用engine 的router
	engine *Engine
}

func (group *RouterGroup) Use(handlerFunc ...HandlerFunc) {
	if handlerFunc != nil {
		group.middlewares = append(group.middlewares, handlerFunc...)
	}
}

// Group 创建新的group
func (group *RouterGroup) Group(prefix string) *RouterGroup {
	newGroup := &RouterGroup{
		prefix: group.prefix + prefix,
		//parent: group,
		engine: group.engine,
	}
	group.child = append(group.child, newGroup)
	//group.engine.groupRouters = append(group.engine.groupRouters, newGroup)
	return newGroup
}

func (group *RouterGroup) addRouter(method string, comp string, handler HandlerFunc) {
	pattern := group.prefix + comp
	log.Printf("Route %s - %s", method, pattern)
	group.engine.router.addRouter(method, pattern, handler)
}

func (group *RouterGroup) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	c := newContext(w, req)
	curGroup := group
	for curGroup != nil {
		if curGroup.middlewares != nil {
			c.handlers = append(c.handlers, curGroup.middlewares...)
		}
		tag := false
		for i := 0; i < len(curGroup.child); i++ {
			if strings.HasPrefix(c.Path, curGroup.child[i].prefix) {
				curGroup = curGroup.child[i]
				tag = true
				break
			}
		}
		if !tag {
			break
		}
	}
	group.engine.router.handle(c)
}

func (group *RouterGroup) Get(pattern string, handler HandlerFunc) {
	group.addRouter("GET", pattern, handler)
}

func (group *RouterGroup) POST(pattern string, handler HandlerFunc) {
	group.addRouter("POST", pattern, handler)
}
