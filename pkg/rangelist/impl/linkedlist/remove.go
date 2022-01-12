package linkedlist

import (
	"gitlab/pkg/rangelist"
)

func (l *linkedList) del(node *Node) {
	l.RWLock.Lock()
	defer l.RWLock.Unlock()
	pre := node.Pre
	next := node.Next
	pre.Next = next
	next.Pre = pre
	delete(l.Index, node.LowerBoundary())
}

// @ removeRange 处理一种range的相交情况
// @ param node rangelist中的node
// @ param r 需要remove的range范围
// 分三种情况
// 情况1 : r 包含 node
// 情况2 : node 包含 r 左右差集
// 情况3 ：存在单差集
func (l *linkedList)removeRange(node *Node, r rangelist.Range){
	// r包含node
	if r.Contain(l.Range){
		l.del(node)
		return
	}

	// 情况2 : node 包含 r 左右差集。打成两段
	ld := node.Range.LeftDiff(r)
	rd := node.Range.RightDiff(r)
	if l.Contain(r) &&
		ld.Valid() && !ld.Equal(node.Range) && // 有左差集
		rd.Valid() && !rd.Equal(node.Range){   // 有右差集
		// 左差集
		 l.modify(node, ld)
		// 右差集
		l.insertAfter(node, rd)
		return
	}

	// 左差集
	if ld.Valid() && !ld.Equal(node.Range){
		l.modify(node, ld)
		return
	}

	// 右差集
	if rd.Valid() && !rd.Equal(node.Range){
		l.modify(node, rd)
	}

	return
}

// @ removeBidiRange 处理俩个range的相交情况
// @ param lnode rangelist 中左边的node
// @ param rnode rangelist 中右边的node
// @ param r 需要remove的range范围
func (l *linkedList)removeBidiRange(lnode , rnode *Node, r rangelist.Range){
	if r.Contain(lnode.Range){
		l.del(lnode)
	}
	if r.Contain(rnode.Range){
		l.del(rnode)
	}
	ld := lnode.LeftDiff(r)
	if ld.Valid(){
		l.modify(lnode, ld)
	}
	rd := rnode.RightDiff(r)
	if rd.Valid(){
		l.modify(rnode, rd)
	}
}
