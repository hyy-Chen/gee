/**
 @author : ikkk
 @date   : 2023/9/10
 @text   : nil
**/

package gee

import "net/http"

type RouterGroup struct {
	prefix string

	engine *Engine
}

func newRouterGroup() *RouterGroup {
	return &RouterGroup{}
}

// Group 注册路由组
func (r *RouterGroup) Group(prefix string) *RouterGroup {
	prefix = joinPaths(r.prefix, prefix)
	if err := checkRouterPath(prefix); err != nil {
		panic(err.Error())
	}
	return &RouterGroup{
		prefix: prefix,
		engine: r.engine,
	}
}

// GET 设置对应接口get请求响应
func (r *RouterGroup) GET(pattern string, handleFunc HandleFunc) {
	r.addRouter(http.MethodGet, pattern, handleFunc)
}

// POST 设置对应接口post请求响应
func (r *RouterGroup) POST(pattern string, handleFunc HandleFunc) {
	r.addRouter(http.MethodPost, pattern, handleFunc)
}

// DELETE 设置对应接口delete请求响应
func (r *RouterGroup) DELETE(pattern string, handleFunc HandleFunc) {
	r.addRouter(http.MethodDelete, pattern, handleFunc)
}

// PUT 设置对应接口put请求响应
func (r *RouterGroup) PUT(pattern string, handleFunc HandleFunc) {
	r.addRouter(http.MethodPut, pattern, handleFunc)
}

func (r *RouterGroup) addRouter(method string, comp string, handler HandleFunc) {
	pattern := joinPaths(r.prefix, comp)
	r.engine.router.addRouter(method, pattern, handler)
}
