/**
 @author : ikkk
 @date   : 2023/9/5
 @text   : nil
**/

package gee

import (
	"fmt"
	"log"
	"strings"
)

type router struct {
	trees map[string]*node
}

func newRouter() *router {
	return &router{trees: map[string]*node{}}
}

func (r *router) addRouter(method string, pattern string, handleFunc HandleFunc) {
	log.Printf("try add router %s - %s\n", method, pattern)
	if pattern == "" {
		panic("web: 路由不能为空")
	}
	root, ok := r.trees[method]
	if !ok {
		root = &node{
			part: "/",
		}
		r.trees[method] = root
	}
	if pattern == "/" {
		if root.handleFunc != nil {
			panic(fmt.Sprintf("web：路由冲突 - %s", pattern))
		}
		root.handleFunc = handleFunc
		log.Printf("add router %s - %s success\n", method, pattern)
		return
	}

	if !strings.HasPrefix(pattern, "/") {
		panic("web: 路由必须由 / 开头")
	}
	if strings.HasSuffix(pattern, "/") {
		panic("web: 路由不能以 / 结尾")
	}

	// 切割得到路由路径
	parts := strings.Split(pattern[1:], "/")
	for _, part := range parts {
		if part == "" {
			panic("web: 路径错误，出现了连续的 / ")
		}
		root, ok = root.addNode(part)
		if !ok {
			panic(fmt.Sprintf("web: 路由冲突 - %s", pattern))
		}
	}
	if root.handleFunc != nil {
		panic(fmt.Sprintf("web：路由冲突 - %s", pattern))
	}
	root.handleFunc = handleFunc
	log.Printf("add router %s - %s success\n", method, pattern)
}

// 获取路由匹配对应处理函数
func (r *router) getRouter(method string, pattern string) (*node, bool) {
	// 对参数进行适量检验
	if pattern == "" {
		return nil, false
	}
	// 获取前缀树根节点进行查找对应路由
	root, ok := r.trees[method]
	if !ok {
		return nil, false
	}
	if pattern == "/" {
		return root, true
	}
	// 切割得到路由路径
	parts := strings.Split(pattern[1:], "/")
	for _, part := range parts {
		if part == "" {
			return nil, false
		}
		root = root.getNode(part)
		if root == nil {
			return nil, false
		}
	}
	return root, root.handleFunc != nil
}
