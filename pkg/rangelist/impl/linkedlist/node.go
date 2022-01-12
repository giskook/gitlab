package linkedlist

import "gitlib/pkg/rangelist"

type Node struct {
	Pre  *Node
	Next *Node
	rangelist.Range
}
