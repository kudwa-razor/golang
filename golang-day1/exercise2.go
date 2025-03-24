package main

import "fmt"

// struct for tree node
type Node struct {
	Value string
	Left  *Node
	Right *Node
}

// pre order traverse (VLR)
func Preorder(node *Node) {
	if node == nil {
		return
	}
	fmt.Print(node.Value, " ")
	Preorder(node.Left)
	Preorder(node.Right)
}

// post order traverse (LRV)
func Postorder(node *Node) {
	if node == nil {
		return
	}
	Postorder(node.Left)
	Postorder(node.Right)
	fmt.Print(node.Value, " ")
}

func main() {
	//manual constr the expression tree for "a+b-c"
	root := &Node{Value: "+"}
	root.Left = &Node{Value: "a"}
	root.Right = &Node{Value: "-"}
	root.Right.Left = &Node{Value: "b"}
	root.Right.Right = &Node{Value: "c"}

	//fmt.Println()
	fmt.Print("Preorder traversal: ")
	Preorder(root)
	fmt.Println()

	fmt.Print("Postorder traversal: ")
	Postorder(root)
	fmt.Println()
}
