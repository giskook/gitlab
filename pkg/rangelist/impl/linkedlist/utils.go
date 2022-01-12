package linkedlist

import (
	"gitlab/pkg/rangelist"
	"gitlab/pkg/rangelist/base"
	"sort"
)

// @ findMaxLowerBoundary 查找小于lowerBoundary的最大的左节点
func (l *linkedList) findMaxLowerBoundary(lowerBoundary int) *Node {
	l.RWLock.RLock()
	defer l.RWLock.RUnlock()
	if node, ok := l.Index[lowerBoundary]; ok {
		return node
	}
	keys := make([]int, 0, len(l.Index))
	for k := range l.Index {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	index := base.BinarySearchInsert(keys, lowerBoundary)
	if index == 0 {
		return l.DummyHeader.Next
	}
	if index == len(keys) {
		return l.DummyTail.Pre
	}

	return l.Index[index]
}

// @ findIntersectNode range的up和low在整个rangelist的中间
func (l *linkedList) findIntersectNode(r rangelist.Range) []*Node {
	intersect := make([]*Node, 0)
	node := l.findMaxLowerBoundary(r.LowerBoundary())
	l.RWLock.RLock()
	defer l.RWLock.RUnlock()
	for node != nil &&
		node != l.DummyTail &&
		!node.Right(r) {
		if node.Insect(r) {
			intersect = append(intersect, node)
		}
		node = node.Next
	}

	return intersect
}
