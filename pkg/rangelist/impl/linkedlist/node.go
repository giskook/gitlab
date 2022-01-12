package linkedlist

import "gitlab/pkg/rangelist"

type Node struct {
	Pre  *Node
	Next *Node
	rangelist.Range
}
