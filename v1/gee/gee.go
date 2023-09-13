/**
 @author : ikkk
 @date   : 2023/9/11
 @text   : nil
**/

package gee

import (
	"fmt"
	"net/http"
)

type EngineOption func(engine *Engine)

type Engine struct {
	srv *http.Server
	// 为了便于框架停止时使用者执行额外操作，提供了stop方法使得使用者可以自定义此类操作
	stop func() error
	// routers 存放路由位置以及对应处理函数
	*router
	//routers map[string]HandleFunc
	//
	// RouterGroup 路由分组
	*RouterGroup
}

// New 使用可变参数构造，便于使用者为项目启动时为框架添加一些操作
func New(opts ...EngineOption) *Engine {
	engine := &Engine{
		router:      newRouter(),
		RouterGroup: newRouterGroup(),
	}
	engine.RouterGroup.engine = engine
	engine.RouterGroup.prefix = "/"
	for _, opt := range opts {
		opt(engine)
	}
	return engine
}

// ServerHTTP 接受转发请求
// 接受请求：接受前端发送过来请求
// 转发请求：将请求转发到web框架
func (e *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	node, params, ok := e.getRouter(r.Method, r.URL.Path)
	fmt.Println("接口", r.URL.Path, "接受到", r.Method, "请求")
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("404 NOT FOUND"))
	} else {
		// 找到就转发请求
		c := NewContext(w, r)
		c.params = params
		node.handleFunc(c)
		c.flashDataToResponse()
	}
}

// Run 启动服务
func (e *Engine) Run(addr string) error {
	//TODO implement me
	// 开启监听指定
	e.srv = &http.Server{
		Addr:    addr,
		Handler: e,
	}
	fmt.Println("启动监听端口", addr)
	return e.srv.ListenAndServe()
}

// Stop 停止服务
func (e *Engine) Stop() error {
	//TODO implement me
	//TODO implement me
	if e.stop == nil {
		return nil
	}
	return e.stop()
}
