package linkedlist

import (
	"fmt"
	"gitlab/pkg/rangelist"
	"gitlab/pkg/rangelist/base"
	"sync"
)

type linkedList struct {
	RWLock      sync.RWMutex
	DummyHeader *Node
	DummyTail   *Node
	Index       map[int]*Node
	rangelist.Range
}

func NewRangeList() rangelist.RangeList {
	dummyHeader := &Node{}
	dummyTail := &Node{Pre: dummyHeader}
	dummyHeader.Next = dummyTail
	return &linkedList{
		DummyHeader: dummyHeader,
		DummyTail:   dummyTail,
		Index:       make(map[int]*Node),
	}
}

func (l *linkedList) isEmpty() (bool, error) {
	if l == nil {
		return false, base.ErrorRangeListIsNil.Error()
	}
	return l.DummyHeader.Next == l.DummyTail, nil
}

func (l *linkedList) Add(r rangelist.Range) error {
	if l == nil {
		return base.ErrorRangeListIsNil.Error()
	}
	if !r.Valid() {
		return base.ErrorRangeIsUnavailable.Error()
	}
	if empty, _ := l.isEmpty(); empty {
		l.init(r)
		return nil
	}

	// 要插入的range在rangelist左侧
	if r.Left(l.Range) {
		l.addHeader(r)
		return nil
	}

	// 要插入的range在rangelist右侧
	if r.Right(l.Range) {
		l.addTail(r)
		return nil
	}

	// 处理有相交的情况
	intersect := l.findIntersectNode(r)
	switch len(intersect) {
	case 0:
		node := l.findMaxLowerBoundary(r.LowerBoundary())
		l.insertAfter(node, r)
	case 1:
		dst := intersect[0].Range.Union(r)
		l.modify(intersect[0], dst)
	case 2:
		l.mergeBidi(intersect[0], intersect[1], r)
	default:
		for i := 1; i < len(intersect)-1; i++ {
			l.del(intersect[i])
		}
		l.mergeBidi(intersect[0], intersect[len(intersect)-1], r)
	}

	return nil
}

func (l *linkedList) Remove(r rangelist.Range) error {
	if l == nil {
		return base.ErrorRangeListIsNil.Error()
	}
	if !r.Valid() {
		return base.ErrorRangeIsUnavailable.Error()
	}

	if empty, _ := l.isEmpty(); empty {
		return nil
	}
	// 范围外直接返回
	if r.Left(l.Range) || r.Right(l.Range) {
		return nil
	}
	// 处理有相交的情况
	intersect := l.findIntersectNode(r)
	switch len(intersect) {
	case 0:
	case 1:
		l.removeRange(intersect[0], r)
	case 2:
		l.removeBidiRange(intersect[0], intersect[1], r)
	default:
		for i := 1; i < len(intersect)-1; i++ {
			l.del(intersect[i])
		}
		l.removeBidiRange(intersect[0], intersect[len(intersect)-1], r)
	}

	return nil
}

func (l *linkedList) Print() {
	l.RWLock.RLock()
	defer l.RWLock.RUnlock()
	node := l.DummyHeader.Next
	for node != l.DummyTail {
		fmt.Printf("[%v, %v)", node.LowerBoundary(), node.UpperBoundary())
		node = node.Next
	}
	fmt.Printf("\n")
}
