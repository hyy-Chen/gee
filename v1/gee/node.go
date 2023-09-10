/**
 @author : ikkk
 @date   : 2023/9/5
 @text   : 存储前缀树节点的结构体，里面存储路由信息
	需要解决的是路由匹配优先级和路由冲突问题

**/

package gee

import "strings"

type node struct {
	part       string           // 当前节点的路由名称
	children   map[string]*node // 静态路由子节点信息
	handleFunc HandleFunc       //当前节点处理函数
	paramChild *node            // 参数路由
	starChild  *node            // 通配符路由

}

// 插入一个路由节点信息
func (n *node) addNode(part string) (*node, bool) {
	if strings.HasPrefix(part, "*") {
		// 如果要插入一个通配符子节点
		if n.paramChild != nil {
			// 如果之前已经有一个参数子节点，就会发生冲突, 插入失败
			return nil, false
		}
		if n.starChild == nil {
			n.starChild = &node{part: part}
		}
		if n.starChild.part != part {
			return nil, false
		}
		return n.starChild, true
	}
	if strings.HasPrefix(part, ":") {
		if n.starChild != nil {
			// 如果之前已经存在一个通配符子节点，就会发生插入冲突
			return nil, false
		}
		if n.paramChild == nil {
			n.paramChild = &node{part: part}
		}
		if n.paramChild.part != part {
			return nil, false
		}
		return n.paramChild, true
	}
	if n.children == nil {
		n.children = make(map[string]*node)
	}
	child, ok := n.children[part]
	if !ok {
		child = &node{
			part: part,
		}
		n.children[part] = child
	}
	return child, true
}

// 获取一个路由节点信息
func (n *node) getNode(part string) *node {
	if n.children == nil {
		if n.paramChild != nil {
			return n.paramChild
		}
		if n.starChild != nil {
			return n.starChild
		}
		return nil
	}
	child, ok := n.children[part]
	if !ok {
		if n.paramChild != nil {
			return n.paramChild
		}
		if n.starChild != nil {
			return n.starChild
		}
		return nil
	}
	return child
}
