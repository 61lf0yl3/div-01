package main

import (
	"fmt"
)

func main() {
	link := &List{}
	ListPushBack(link, "hello")
	ListPushBack(link, "how are")
	ListPushBack(link, "you")
	ListPushBack(link, 1)

	//fmt.Println(ListAt(link.Head, 0).Data)
	fmt.Println(ListAt(link.Head, 1).Data)
	fmt.Println(ListAt(link.Head, 2).Data)
	fmt.Println(ListAt(link.Head, 3).Data)
	fmt.Println(ListAt(link.Head, 4).Data)
	//fmt.Println(ListAt(link.Head, 5))
}

type NodeL struct {
	Data interface{}
	Next *NodeL
}
type List struct {
	Head *NodeL
	Tail *NodeL
}

func ListAt(l *NodeL, pos int) *NodeL {

	for i := 1; i < pos; i++ {
		if l != nil {
			l = l.Next
		} else {
			return nil
		}
	}
	//fmt.Println(l.Data)
	return l
}

func ListPushBack(l *List, data interface{}) {

	n := &NodeL{Data: data}
	if l.Head == nil {
		l.Head = n
		l.Tail = n
	} else {
		l.Tail.Next = n
		l.Tail = n
	}
}
