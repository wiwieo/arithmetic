package bst

import "fmt"

// binary search tree
// Binary Search tree exhibits a special behavior.
// A node's left child must have a value less than its parent's value
// and the node's right child must have a value greater than its parent value.

type Node struct {
	data  int
	left  *Node
	right *Node
}

func (root *Node) Insert(data int) error {
	var temp = &Node{data: data}
	if root == nil {
		return fmt.Errorf("root is nil")
	}
	var parent = root
	for {
		if data <= parent.data {
			if parent.left == nil{
				parent.left = temp
				return nil
			}
			parent = parent.left
			continue
		}
		if data > parent.data {
			if parent.right == nil{
				parent.right = temp
				return nil
			}
			parent = parent.right
			continue
		}
	}
}

func (root *Node) Search(data int) (*Node, error){
	var parent = root
	for {
		if parent == nil{
			return nil, fmt.Errorf("node is not exist")
		}
		if data == parent.data{
			return parent, nil
		}
		if data < parent.data {
			parent = parent.left
			continue
		}
		if data > parent.data {
			parent = parent.right
			continue
		}
	}
}

// 中序
func (root *Node) InOrder(){
	parent := root
	if parent == nil{
		return
	}
	if parent.left != nil{
		parent.left.InOrder()
	}
	print(root.data, " ")
	if parent.right != nil{
		parent.right.InOrder()
	}
}

// 前序
func (root *Node) PreOrder(){
	parent := root
	if parent == nil{
		return
	}
	print(root.data, " ")
	if parent.left != nil{
		parent.left.PreOrder()
	}
	if parent.right != nil{
		parent.right.PreOrder()
	}
}

// 后序
func (root *Node) PostOrder(){
	parent := root
	if parent == nil{
		return
	}
	if parent.left != nil{
		parent.left.PostOrder()
	}
	if parent.right != nil{
		parent.right.PostOrder()
	}
	print(root.data, " ")
}