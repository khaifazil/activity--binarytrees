package main

import "fmt"

type binaryNode struct {
	item  string
	left  *binaryNode
	right *binaryNode
}

type BST struct {
	root *binaryNode
}

func (bst *BST) insertNode(t **binaryNode, item string) error { //declaration of a type **

	if *t == nil {
		newNode := &binaryNode{
			item:  item,
			left:  nil,
			right: nil,
		}
		*t = newNode
		return nil

	}
	if item < (*t).item {
		bst.insertNode(&(*t).left, item) //dereferencing **t
	} else {
		bst.insertNode(&(*t).right, item)
	}
	return nil
}

func (bst *BST) insert(item string) error {
	bst.insertNode(&bst.root, item)
	return nil
}

func (bst *BST) inOrderTraversal(t *binaryNode) error {

	if t != nil {
		bst.inOrderTraversal(t.left)
		fmt.Println(t.item)
		bst.inOrderTraversal(t.right)
	}
	return nil
}

func (bst *BST) preOrderTraversal(t *binaryNode) error {

	if t != nil {
		fmt.Println(t.item)
		bst.inOrderTraversal(t.left)
		bst.inOrderTraversal(t.right)
	}
	return nil
}

func (bst *BST) searchNode(t *binaryNode, item string) *binaryNode {
	if t == nil {
		return nil
	} else {
		if t.item == item {
			return t
		} else {
			if t.item > item {
				return bst.searchNode(t.left, item)
			}else{
				return bst.searchNode(t.right, item)
			}
		}
	}
}


//utility functions
func (bst *BST) inOrder() {
	bst.inOrderTraversal(bst.root)
}

func (bst *BST) preOrder() {
	bst.preOrderTraversal(bst.root)
}

func (bst *BST) search(item string) *binaryNode {
	return bst.searchNode(bst.root, item)
}

func main() {
	bst := &BST{root: nil}
	bst.insert("Mary")
	bst.insert("Tom")
	bst.insert("Lucas")

	fmt.Println("inorder")
	bst.inOrder()
	fmt.Println("preorder")
	bst.preOrder()

	var name string = "John"
	fmt.Println("Searching for item...")
	t := bst.search(name)
	if t != nil {
		fmt.Println(name, "is present")
	}else {
		fmt.Println(name, "is not present")
	}
}