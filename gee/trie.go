package gee

import "strings"

type node struct {
	pattern  string  // 表示完整的路径，只有叶子节点才有该值，其它节点为空
	part     string  // 路径中的一部分
	children []*node // 子节点
	isFuzzy  bool    // 是否为 参数匹配(:) 或是 通配符(*)
}

// matchChild 找到当前节点下第一个符合条件的子节点
func (n *node) matchChild(part string) *node {
	for _, child := range n.children {
		if child.part == part || child.isFuzzy {
			return child
		}
	}

	return nil
}

// matchChildren 找到当前节点下所有符合条件的子节点
func (n *node) matchChildren(part string) []*node {
	var res []*node
	for _, child := range n.children {
		if child.part == part || child.isFuzzy {
			res = append(res, child)
		}
	}

	return res
}

// @brief: 插入节点
// @param: pattern 路径
// @param: parts 路径分割后的数组
// @param: height 当前处理的节点所在树的深度
func (n *node) insert(pattern string, parts []string, height int) {
	if len(parts) == height {
		return
	}

	part := parts[height]
	child := n.matchChild(part)
	// 节点不存在则创建
	if child == nil {
		child = &node{
			part:    part,
			isFuzzy: part[0] == ':' || part[0] == '*',
		}

		// 最后一个节点要给pattern赋值
		if height+1 == len(parts) {
			child.pattern = pattern
		}

		n.children = append(n.children, child)
	}

	// 处理下一层级的节点
	child.insert(pattern, parts, height+1)
}

// @brief: 根据路径，搜索对应节点
// @param: parts 路径数组
// @param: height 当前节点所在深度
func (n *node) search(parts []string, height int) *node {
	// 已经找到了路径的最后一层，或者当前树节点对应路径部分为通配符 *
	if len(parts) == height || strings.HasPrefix(n.part, "*") {
		// 非根节点，返回 nil
		if n.pattern == "" {
			return nil
		}
		return n
	}

	part := parts[height]
	children := n.matchChildren(part)
	for _, child := range children {
		res := child.search(parts, height+1)
		if res != nil {
			return res
		}
	}

	return nil
}
