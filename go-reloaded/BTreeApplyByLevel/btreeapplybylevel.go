package main

import "fmt"

func main() {

	root := &TreeNode{Data: "01"}
	BTreeInsertData(root, "07")
	BTreeInsertData(root, "12")
	BTreeInsertData(root, "10")
	BTreeInsertData(root, "05")
	BTreeInsertData(root, "02")
	BTreeInsertData(root, "03")
	BTreeApplyByLevel(root, fmt.Println)

	// 01
	// 07
	// 05
	// 12
	// 02
	// 10
	// 03

	// root := &TreeNode{Data: "01"}
	// BTreeInsertData(root, "07")
	// BTreeInsertData(root, "12")
	// BTreeInsertData(root, "10")
	// BTreeInsertData(root, "05")
	// BTreeInsertData(root, "02")
	// BTreeInsertData(root, "03")
	// BTreeApplyByLevel(root, fmt.Print)

	// 01070512021003

}

type TreeNode struct {
	Left, Right, Parent *TreeNode
	Data                string
}

func BTreeApplyByLevel(root *TreeNode, f func(...interface{}) (int, error)) {
	var slice []*TreeNode
	if root == nil {
		return
	}
	slice = append(slice, root)
	for root != nil {
		f(root.Data)
		if root.Left != nil {
			slice = append(slice, root.Left)
		}
		if root.Right != nil {
			slice = append(slice, root.Right)
		}
		if len(slice) == 1 {
			return
		}
		slice = slice[1:]
		root = slice[0]
	}
}

func BTreeInsertData(root *TreeNode, data string) *TreeNode {
	if root == nil {
		return &TreeNode{Data: data}
	}

	if data < root.Data {
		root.Left = BTreeInsertData(root.Left, data)
		root.Left.Parent = root
	} else {
		root.Right = BTreeInsertData(root.Right, data)
		root.Right.Parent = root
	}
	return root
}
