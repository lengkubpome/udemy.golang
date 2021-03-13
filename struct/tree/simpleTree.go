package main

import "fmt"

type BinaryTree struct {
	value int
	left  *BinaryTree
	right *BinaryTree
}

func main() {
	root := BinaryTree{value: 2}
	left := BinaryTree{value: 1}
	right := BinaryTree{value: 3}

	root.left = &left
	root.right = &right

	showDF(&root)
}

func showDF(node *BinaryTree) {
	if node != nil {
		showDF(node.left)
		fmt.Println(node.value)
		showDF(node.right)

	}

}
