/**
 @author : ikkk
 @date   : 2023/9/5
 @text   : nil
**/

package gee

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type HandleFunc func(c *Context)

// 抽象出server接口，便于之后管理与拓展
type server interface {
	// Handler 使用匿名字段继承http.Handler的方法
	http.Handler
	// Start 启动操作
	Start(addr string) error
	// Stop 关闭服务操作
	Stop() error
	// addRouter 路由注册函数
	addRouter(method string, pattern string, handleFunc HandleFunc)
}

type HTTPOption func(h *HTTPServer)

type HTTPServer struct {
	srv *http.Server
	// 为了便于框架停止时使用者执行额外操作，提供了stop方法使得使用者可以自定义此类操作
	stop func() error
	// routers 存放路由位置以及对应处理函数
	*router
	//routers map[string]HandleFunc
}

// NewHTTP 使用可变参数构造，便于使用者为项目启动时为框架添加一些操作
func NewHTTP(opts ...HTTPOption) *HTTPServer {
	h := &HTTPServer{
		router: newRouter(),
	}
	for _, opt := range opts {
		opt(h)
	}
	return h
}

// WithHTTPServerStop 创建一个停止web框架时额外操作的方法
func WithHTTPServerStop(fn func() error) HTTPOption {
	return func(h *HTTPServer) {
		// 如果默认参数为空，就启用一个程序自定义的方法
		if fn == nil {
			fn = func() error {
				quit := make(chan os.Signal)
				signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
				<-quit
				log.Println("Shutdown Server ...")
				ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
				defer cancel()
				if err := h.srv.Shutdown(ctx); err != nil {
					log.Fatal("Server Shutdown:", err)
				}
				select {
				case <-ctx.Done():
					log.Println("timeout of 2 seconds.")

				}
				return nil
			}
		}
		h.stop = fn
	}
}

// ServerHTTP 接受转发请求
// 接受请求：接受前端发送过来请求
// 转发请求：将请求转发到web框架
func (h *HTTPServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	node, ok := h.getRouter(r.Method, r.URL.Path)
	fmt.Println("接口", r.URL.Path, "接受到", r.Method, "请求")
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("404 NOT FOUND"))
	} else {
		// 找到就转发请求
		c := NewContext(w, r)
		node.handleFunc(c)
	}
}

// Start 启动服务
func (h *HTTPServer) Start(addr string) error {
	//TODO implement me
	// 开启监听指定
	h.srv = &http.Server{
		Addr:    addr,
		Handler: h,
	}
	fmt.Println("启动监听端口", addr)
	return h.srv.ListenAndServe()
}

// Stop 停止服务
func (h *HTTPServer) Stop() error {
	//TODO implement me
	if h.stop == nil {
		return nil
	}
	return h.stop()
}

// addRouter 注册路由
// 注册路由时间点：框架启动前
//func (h *HTTPServer) addRouter(method string, pattern string, handleFunc HandleFunc) {
//	key := fmt.Sprintf("%s-%s", method, pattern)
//	fmt.Println(key, "设置")
//	h.routers[key] = handleFunc
//}

// GET 设置对应接口get请求响应
func (h *HTTPServer) GET(pattern string, handleFunc HandleFunc) {
	h.addRouter(http.MethodGet, pattern, handleFunc)
}

// POST 设置对应接口post请求响应
func (h *HTTPServer) POST(pattern string, handleFunc HandleFunc) {
	h.addRouter(http.MethodPost, pattern, handleFunc)
}

// DELETE 设置对应接口delete请求响应
func (h *HTTPServer) DELETE(pattern string, handleFunc HandleFunc) {
	h.addRouter(http.MethodDelete, pattern, handleFunc)
}

// PUT 设置对应接口put请求响应
func (h *HTTPServer) PUT(pattern string, handleFunc HandleFunc) {
	h.addRouter(http.MethodPut, pattern, handleFunc)
}
