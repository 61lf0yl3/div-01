package main

import "fmt"

type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
}

func ListSize(l *List) int {
	counter := 0
	for l.Head != nil {
		counter++
		l.Head = l.Head.Next
	}
	return counter
}

func ListPushFront(link *List, data interface{}) {
	n := &NodeL{Data: data}
	if link.Head == nil {
		link.Head = n
		link.Tail = n
	} else {
		n.Next = link.Head
		link.Head = n
	}
}

func main() {
	link := &List{}

	ListPushFront(link, "Hello")
	ListPushFront(link, "2")
	ListPushFront(link, "you")
	ListPushFront(link, "man")

	fmt.Println(ListSize(link))
}
