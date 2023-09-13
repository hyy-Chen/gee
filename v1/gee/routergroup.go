/**
 @author : ikkk
 @date   : 2023/9/10
 @text   : nil
**/

package gee

import "net/http"

type RouterGroup struct {
	prefix      string
	middlewares HandlersChain
	engine      *Engine
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
	rg := &RouterGroup{
		prefix: prefix,
		engine: r.engine,
	}
	r.engine.groups = append(r.engine.groups, rg)
	return rg
}

func (r *RouterGroup) Use(middlewares ...HandleFunc) {
	r.middlewares = append(r.middlewares, middlewares...)
}

// GET 设置对应接口get请求响应
func (r *RouterGroup) GET(pattern string, handleFunc ...HandleFunc) {
	r.addRouter(http.MethodGet, pattern, handleFunc)
}

// POST 设置对应接口post请求响应
func (r *RouterGroup) POST(pattern string, handlers ...HandleFunc) {
	r.addRouter(http.MethodPost, pattern, handlers)
}

// DELETE 设置对应接口delete请求响应
func (r *RouterGroup) DELETE(pattern string, handlers ...HandleFunc) {
	r.addRouter(http.MethodDelete, pattern, handlers)
}

// PUT 设置对应接口put请求响应
func (r *RouterGroup) PUT(pattern string, handlers ...HandleFunc) {
	r.addRouter(http.MethodPut, pattern, handlers)
}

func (r *RouterGroup) addRouter(method string, comp string, handlers HandlersChain) {
	pattern := joinPaths(r.prefix, comp)
	r.engine.router.addRouter(method, pattern, handlers)
}
