package framework

import (
	"errors"
	"strings"
)

type Tree struct {
	root *node
}

type node struct {
	isLast  bool   // 该节点是否能成为一个独立的uri, 是否自身就是一个终极节点
	segment string // uri中的字符串
	handler ControllerHandler
	childs  []*node
}

func newNode() *node {
	return new(node)
}

func NewTree() *Tree {
	return &Tree{
		root: newNode(),
	}
}

// 判断一个segment是否是通用segment，即以:开头
func isWildSegment(segment string) bool {
	return strings.HasPrefix(segment, ":")
}

// 过滤下一层满足segment规则的子节点
func (n *node) filterChildNodes(segment string) []*node {
	if len(n.childs) == 0 {
		return nil
	}

	if isWildSegment(segment) {
		return n.childs
	}

	nodes := make([]*node, 0, len(n.childs))

	for _, cnodes := range n.childs {
		if isWildSegment(cnodes.segment) {
			// 如果下一层子节点有通配符，则满足需求
			nodes = append(nodes, cnodes)
		} else if cnodes.segment == segment {
			// 如果下一层子节点没有通配符，但是文本完全匹配，则满足需求
			nodes = append(nodes, cnodes)
		}
	}

	return nodes
}

// 匹配 route 节点
func (n *node) matchNode(url string) *node {
	segments := strings.SplitN(strings.Trim(url, "/"), "/", 2)
	segment := segments[0]
	isLast := len(segments) == 1

	if !isWildSegment(segment) {
		segment = strings.ToUpper(segment)
	}

	cnodes := n.filterChildNodes(segment)

	// 如果当前子节点没有一个符合，那么说明这个uri一定是之前不存在, 直接返回nil
	if cnodes == nil || len(cnodes) == 0 {
		return nil
	}

	for _, cn := range cnodes {
		if isLast {
			if cn.isLast {
				return cn
			}
		} else {
			tnMatch := cn.matchNode(segments[1])
			if tnMatch != nil {
				return tnMatch
			}
		}
	}
	return nil
}

// AddRouter 增加路由节点, 路由节点有先后顺序
/*
/book/list
/book/:id (冲突)
/book/:id/name
/book/:student/age
/:user/name
/:user/name/:age (冲突)
*/
func (tree *Tree) AddRouter(uri string, handler ControllerHandler) error {
	n := tree.root
	if n.matchNode(uri) != nil {
		return errors.New("route exist: " + uri)
	}

	segments := strings.Split(strings.Trim(uri, "/"), "/")
	// 对每个segment
	for index, segment := range segments {

		// 最终进入Node segment的字段
		if !isWildSegment(segment) {
			segment = strings.ToUpper(segment)
		}
		isLast := index == len(segments)-1

		var objNode *node // 标记是否有合适的子节点

		childNodes := n.filterChildNodes(segment)
		// 如果有匹配的子节点
		if len(childNodes) > 0 {
			// 如果有segment相同的子节点，则选择这个子节点
			for _, cnode := range childNodes {
				if cnode.segment == segment {
					objNode = cnode
					break
				}
			}
		}

		if objNode == nil {
			// 创建一个当前node的节点
			cnode := newNode()
			cnode.segment = segment
			if isLast {
				cnode.isLast = true
				cnode.handler = handler
			}
			n.childs = append(n.childs, cnode)
			objNode = cnode
		}

		n = objNode
	}

	return nil
}

// 匹配uri
func (tree *Tree) FindHandler(url string) ControllerHandler {
	matchNode := tree.root.matchNode(url)

	if matchNode == nil {
		return nil
	}
	return matchNode.handler
}
