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

func PrintList(l *List) {
	link := l.Head
	for link != nil {
		fmt.Print(link.Data, " -> ")
		link = link.Next
	}
	fmt.Println(nil)
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

func ListClear(l *List) {
	if l.Head != nil {
		l.Head = nil
		l.Tail = nil
	} else {
		return
	}
}
\
func main() {
	link := &List{}

	ListPushBack(link, "I")
	ListPushBack(link, 1)
	ListPushBack(link, "something")
	ListPushBack(link, 2)

	fmt.Println("------list------")
	PrintList(link)
	ListClear(link)
	fmt.Println("------updated list------")
	PrintList(link)
}
