package main

import "fmt"

type NodeI struct {
	Data interface{}
	Next *NodeI
}

func ListPushBack(l *NodeI, data interface{}) *NodeI {
	n := &NodeI{Data: data}

	if l == nil {
		return n
	}
	iterator := l
	for iterator.Next != nil {
		iterator = iterator.Next
	}
	iterator.Next = n
	return l
}

func PrintList(l *NodeI) {
	it := l
	for it != nil {
		fmt.Print(it.Data, " -> ")
		it = it.Next
	}
	fmt.Print(nil, "\n")
}
func main() {

	link := &NodeI{}

	ListPushBack(link, "Hello")
	ListPushBack(link, "man")
	ListPushBack(link, "how are you")
	ListPushBack(link, "last word")

	PrintList(link)
	// 	for link.Head != nil {
	// 		fmt.Println(link.Head.Data)
	// 		link.Head = link.Head.Next
	// 	}
}
