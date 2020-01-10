package main

import (
	"fmt"
)

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

func main() {

	link := &List{}

	//ListPushFront(link, "Hello")
	ListPushFront(link, "Hello")
	ListPushFront(link, "man")
	ListPushFront(link, "how are you")
	//ListPushFront(link, "Hello")

	it := link.Head
	for it != nil {
		fmt.Print(it.Data, " ")
		it = it.Next
	}
	fmt.Println()
}
