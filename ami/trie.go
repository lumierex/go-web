package ami

import "strings"

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

// 从高度(height)为0 递归 当高度等于 len(parts)时候结束递归 -> 当前节点赋值pattern
func (n *node) insert(pattern string, parts []string, height int) {
	if len(parts) == height {
		n.pattern = pattern
		return
	}
	part := parts[height]
	// 抽取匹配函数
	child := n.matchChild(part)

	// 找不到节点 进行插入
	if child == nil {
		child = &node{
			part: part,
			// /*test或者 /:id 进行模糊匹配
			isWild: part[0] == '*' || part[0] == ':',
		}
		n.children = append(n.children, child)
	}

	// 存在下一层, child插入对应的
	child.insert(pattern, parts, height+1)
}

func (n *node) search(parts []string, height int) *node {
	// 匹配到*就代表全都匹配到  & 如果只匹配到/:id的话 全部通过高度进行判断
	if len(parts) == height || strings.HasPrefix(n.part, "*") {
		if n.pattern == "" {
			return nil
		}
		return n
	}
	part := parts[height]
	// 抽取匹配函数
	children := n.matchChildren(part)
	for _, child := range children {
		result := child.search(parts, height+1)
		if result != nil {
			return result
		}
	}
	return nil
}
