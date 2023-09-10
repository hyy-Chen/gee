/**
 @author : ikkk
 @date   : 2023/9/5
 @text   : nil
**/

package gee

import (
	"errors"
	"fmt"
	"log"
	"strings"
)

var (
	ErrRouterPathEmpty               = errors.New("web: 路由路径为空")
	ErrRouterPathPre                 = errors.New("web: 路由路径没有以 / 为开头")
	ErrRouterPathSuf                 = errors.New("web: 路由路径不能以 / 为结尾")
	ErrRouterPathContinuouslyEmpty   = errors.New("web: 路由路径出现了连续的 / ")
	ErrRouterPathParameterConflicts  = errors.New("web: 路由路径出现了重复的参数")
	ErrRouterPathStarInWrongPosition = errors.New("web: 路由路径中通配符位置异常")
)

type router struct {
	trees map[string]*node
}

func newRouter() *router {
	return &router{trees: map[string]*node{}}
}

func checkRouterPath(path string) error {
	if path == "" {
		return ErrRouterPathEmpty
	}
	if path == "/" {
		return nil
	}
	if !strings.HasPrefix(path, "/") {
		return ErrRouterPathPre
	}
	if strings.HasSuffix(path, "/") {
		return ErrRouterPathSuf
	}
	mp := map[string]struct{}{}
	paths := strings.Split(path[1:], "/")
	for id, p := range paths {
		if p == "" {
			return ErrRouterPathContinuouslyEmpty
		} else {
			if strings.HasPrefix(p, "*") {
				if id != len(paths)-1 {
					return ErrRouterPathStarInWrongPosition
				}
			} else if strings.HasPrefix(p, ":") {
				if _, ok := mp[p]; ok {
					return ErrRouterPathParameterConflicts
				}
				mp[p] = struct{}{}
			}
		}
	}
	return nil
}

func (r *router) addRouter(method string, pattern string, handleFunc HandleFunc) {
	log.Printf("try add router %s - %s\n", method, pattern)
	if err := checkRouterPath(pattern); err != nil {
		panic(err)
	}
	root, ok := r.trees[method]
	if !ok {
		root = &node{
			part: "/",
		}
		r.trees[method] = root
	}
	if pattern == "/" {
		root.handleFunc = handleFunc
		log.Printf("add router %s - %s success\n", method, pattern)
		return
	}
	// 切割得到路由路径
	parts := strings.Split(pattern[1:], "/")
	for _, part := range parts {
		root, ok = root.addNode(part)
		if !ok {
			panic(fmt.Sprintf("web: 路由冲突 - %s", pattern))
		}
	}
	root.handleFunc = handleFunc
	log.Printf("add router %s - %s success\n", method, pattern)
}

// 获取路由匹配对应处理函数
func (r *router) getRouter(method string, pattern string) (*node, map[string]string, bool) {
	params := make(map[string]string)
	// 对参数进行适量检验
	if pattern == "" {
		return nil, nil, false
	}
	// 获取前缀树根节点进行查找对应路由
	root, ok := r.trees[method]
	if !ok {
		return nil, nil, false
	}
	if pattern == "/" {
		return root, params, root.handleFunc != nil
	}
	// 切割得到路由路径
	parts := strings.Split(pattern[1:], "/")
	for _, part := range parts {
		if part == "" {
			return nil, nil, false
		}
		root = root.getNode(part)
		if root == nil {
			return nil, nil, false
		}
		if strings.HasPrefix(root.part, ":") {
			params[root.part[1:]] = part
		}
		if strings.HasPrefix(root.part, "*") {
			index := strings.Index(pattern, part)
			params[root.part[1:]] = pattern[index:]
			return root, params, root.handleFunc != nil
		}
	}
	return root, params, root.handleFunc != nil
}
