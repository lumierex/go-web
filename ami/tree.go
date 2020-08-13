package ami

import (
	"fmt"
	"strings"
)

type node struct {
	pattern  string // 待匹配的路由
	part     string // 部分路由
	children []*node
	isWild   bool // 是否模糊匹配 遇到: 和*
}

// 第一个匹配成功的节点
func (n *node) matchChild(part string) *node {
	for _, child := range n.children {
		if child.part == part || child.isWild {
			return child
		}
	}
	return nil
}

func (n *node) matchChildren(part string) []*node {
	nodes := make([]*node, 0)
	for _, child := range n.children {
		if child.part == part || child.isWild {
			nodes = append(nodes, child)
		}
	}
	return nodes
}

// /admin/:di/info
// /admin/*test
// /admin/info

// 从高度为0 递归 当高度等于 len(parts)时候结束递归 -> 当前节点赋值pattern
func (n *node) insert(pattern string, parts []string, height int) {
	if len(parts) == height {
		n.pattern = pattern
		return
	}
	part := parts[height]
	child := n.matchChild(part)

	// 找不到节点 进行插入
	if child == nil {
		child = &node{
			part:   part,
			isWild: part[0] == '*' || part[0] == ':',
		}
		n.children = append(n.children, child)
	}

	// 下一层
	n.insert(pattern, parts, height+1)
}

func (n *node) search(parts []string, height int) {

}

//
func ParsePattern(pattern string) []string {
	ps := strings.Split(pattern, "/")
	fmt.Println(ps)
	parts := make([]string, 0)
	for _, item := range ps {
		if item != "" {
			parts = append(parts, item)
			if item[0] == '*' {
				break
			}
		}
	}
	return parts
}
