package main

import "fmt"

func main() {
	root := &TreeNode{Data: "01"}
	BTreeInsertData(root, "07")
	BTreeInsertData(root, "05")
	BTreeInsertData(root, "12")
	BTreeInsertData(root, "02")
	BTreeInsertData(root, "10")
	BTreeInsertData(root, "03")
	node := BTreeSearchItem(root, "02")
	fmt.Println("Before delete:")
	BTreeApplyInorder(root, fmt.Println)
	root = BTreeDeleteNode(root, node)
	fmt.Println("After delete:")
	BTreeApplyInorder(root, fmt.Println)

	// 01
	// 03
	// 05
	// 07
	// 10
	// 12

	// root := &TreeNode{Data: "33"}
	// BTreeInsertData(root, "5")
	// BTreeInsertData(root, "20")
	// BTreeInsertData(root, "31")
	// BTreeInsertData(root, "13")
	// BTreeInsertData(root, "11")
	// BTreeInsertData(root, "52")
	// node := BTreeSearchItem(root, "20")
	// fmt.Println("Before delete:")
	// BTreeApplyInorder(root, fmt.Println)
	// root = BTreeDeleteNode(root, node)
	// fmt.Println("After delete:")
	// BTreeApplyInorder(root, fmt.Println)

	// 11
	// 13
	// 31
	// 33
	// 5
	// 52

	// root := &TreeNode{Data: "03"}
	// BTreeInsertData(root, "39")
	// BTreeInsertData(root, "99")
	// BTreeInsertData(root, "44")
	// BTreeInsertData(root, "11")
	// BTreeInsertData(root, "14")
	// BTreeInsertData(root, "11")
	// node := BTreeSearchItem(root, "03")
	// fmt.Println("Before delete:")
	// BTreeApplyInorder(root, fmt.Println)
	// root = BTreeDeleteNode(root, node)
	// fmt.Println("After delete:")
	// BTreeApplyInorder(root, fmt.Println)

	// 11
	// 11
	// 14
	// 39
	// 44
	// 99

	// root := &TreeNode{Data: "03"}
	// BTreeInsertData(root, "03")
	// BTreeInsertData(root, "94")
	// BTreeInsertData(root, "19")
	// BTreeInsertData(root, "24")
	// BTreeInsertData(root, "111")
	// BTreeInsertData(root, "01")
	// node := BTreeSearchItem(root, "03")
	// fmt.Println("Before delete:")
	// BTreeApplyInorder(root, fmt.Println)
	// root = BTreeDeleteNode(root, node)
	// fmt.Println("After delete:")
	// BTreeApplyInorder(root, fmt.Println)

	// 01
	// 03
	// 111
	// 19
	// 24
	// 94
}

func MinVal(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	for root.Left != nil {
		root = root.Left
	}
	return root
}

func BTreeDeleteNode(root, node *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	if node.Data < root.Data {
		root.Left = BTreeDeleteNode(root.Left, node)
	} else if node.Data > root.Data {
		root.Right = BTreeDeleteNode(root.Right, node)
	} else {
		if root.Left == nil && root.Right == nil {
			return nil
		}
		if root.Left == nil {
			return root.Right
		} else if root.Right == nil {
			return root.Left
		}
		temp := MinVal(root.Right)
		root.Data = temp.Data
		root.Right = BTreeDeleteNode(root.Right, temp)
	}
	return root
}

type TreeNode struct {
	Left, Right, Parent *TreeNode
	Data                string
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

func BTreeSearchItem(root *TreeNode, data string) *TreeNode {
	if root == nil {
		return nil
	}
	if data < root.Data {
		root = BTreeSearchItem(root.Left, data)
	} else if data > root.Data {
		root = BTreeSearchItem(root.Right, data)
	} else if data == root.Data {
		return root
	}
	return root
}

func BTreeApplyInorder(root *TreeNode, f func(...interface{}) (int, error)) {
	if root != nil {
		BTreeApplyInorder(root.Left, f)
		f(root.Data)
		BTreeApplyInorder(root.Right, f)
	}
	if root == nil {
		return
	}
}
