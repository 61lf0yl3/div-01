package main

import (
	"fmt"
)

func main() {
	link := &List{}

	ListPushFront(link, "Hello")
	ListPushFront(link, "2")
	ListPushFront(link, "you")
	ListPushFront(link, "man")

	fmt.Println(ListSize(link))
}

type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
}

func ListPushFront(l *List, data interface{}) {
	n := &NodeL{Data: data}

	if l.Tail == nil {
		l.Tail = n
		l.Head = n
	} else {
		l.Tail.Next = n
		l.Tail = n
	}

}

func ListSize(l *List) int {
	it := l.Head
	counter := 0
	for it != nil {
		it = it.Next
		counter++
	}
	return counter
}
