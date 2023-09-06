/**
 @author : ikkk
 @date   : 2023/9/5
 @text   : nil
**/

package gee

type node struct {
	part       string           // 当前节点的路由名称
	children   map[string]*node // 静态路由子节点信息
	handleFunc HandleFunc       //当前节点处理函数
}

// 插入一个路由节点信息
func (n *node) addNode(part string) (*node, bool) {
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
		return nil
	}
	child := n.children[part]
	return child
}
