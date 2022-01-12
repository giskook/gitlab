package linkedlist

import (
	"gitlib/pkg/rangelist"
)

func (l *linkedList) init(r rangelist.Range) {
	l.RWLock.Lock()
	defer l.RWLock.Unlock()
	node := new(Node)
	node.Range = r
	node.Pre = l.DummyHeader
	node.Next = l.DummyTail
	l.DummyHeader.Next = node
	l.DummyTail.Pre = node
	l.Index[r.LowerBoundary()] = node
	l.Range = r
}

func (l *linkedList) addHeader(r rangelist.Range) {
	l.RWLock.Lock()
	defer l.RWLock.Unlock()
	header := l.DummyHeader.Next
	node := new(Node)
	node.Range = r
	node.Pre = l.DummyHeader
	node.Next = header
	header.Pre = node
	l.DummyHeader.Next = node
	l.Index[r.LowerBoundary()] = node
	l.Range = l.Range.Union(r)
}

func (l *linkedList) addTail(r rangelist.Range) {
	l.RWLock.Lock()
	defer l.RWLock.Unlock()
	node := new(Node)
	node.Range = r
	oldLast := l.DummyTail.Pre
	oldLast.Next = node
	l.DummyTail.Pre = node
	node.Pre = oldLast
	node.Next = l.DummyTail
	l.Index[r.LowerBoundary()] = node
	l.Range = l.Range.Union(r)
}

func (l *linkedList) insertAfter(node *Node, r rangelist.Range) {
	l.RWLock.Lock()
	defer l.RWLock.Unlock()
	n := &Node{
		Range: r,
		Pre:   node,
		Next:  node.Next,
	}
	node.Next.Pre = node
	node.Next = n
	l.Index[r.LowerBoundary()] = n
}

func (l *linkedList) modify(node *Node, r rangelist.Range) {
	l.RWLock.Lock()
	defer l.RWLock.Unlock()
	oldRange := node.Range
	node.Range = r
	l.Range = l.Range.Union(r)
	if oldRange.LowerBoundary() != r.LowerBoundary() {
		delete(l.Index, oldRange.LowerBoundary())
		l.Index[r.LowerBoundary()] = node
	}
}

// @ mergeBidi 处理一个range 跨两个node的情况
// @ param lnode 相交左节点
// @ param rnode 相交右节点
// @ param r 使之相交的range
// 结果仅保留lnode。
func (l *linkedList) mergeBidi(lnode, rnode *Node, r rangelist.Range) {
	lr := r.Union(lnode.Range)
	rr := r.Union(rnode.Range)
	lrr := lr.Union(rr)
	dst := l.Range.Union(lrr)
	l.del(rnode)

	l.modify(lnode, dst)
}
