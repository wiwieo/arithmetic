package avl

import (
	"fmt"
)

// 详情参见：
// https://www.tutorialspoint.com/data_structures_algorithms/avl_tree_algorithm.htm
// https://blog.csdn.net/jyy305/article/details/70949010
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
			if parent.left == nil {
				parent.left = temp
				return nil
			}
			parent = parent.left
			continue
		}
		if data > parent.data {
			if parent.right == nil {
				parent.right = temp
				return nil
			}
			parent = parent.right
			continue
		}
	}
}

func (root *Node) InsertAndRotate(data int) error {
	err := root.Insert(data)
	if err != nil {
		return err
	}
	factor := root.BalanceFactor()
	println(factor, data)
	if factor > 1 {
		return root.left.CheckAndRotate(root)
	} else if factor < -1 {
		return root.right.CheckAndRotate(root)
	}
	return nil
}

func (root *Node) Search(data int) (*Node, error) {
	var parent = root
	for {
		if parent == nil {
			return nil, fmt.Errorf("node is not exist")
		}
		if data == parent.data {
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

func (root *Node) CheckAndRotate(parent *Node) error {
	factor := root.BalanceFactor()
	if factor > 1 { // >1，说明是左子树不平衡
		parent = root
		return root.left.CheckAndRotate(parent)
	} else if factor < -1 { // 右子树不平衡
		parent = root
		return root.right.CheckAndRotate(parent)
	} else { // 找到，并旋转
		return parent.Rotate()
	}
	return fmt.Errorf("balance factor is wrong")
}

func (root *Node) Rotate() error {
	if root == nil {
		return fmt.Errorf("root is nil")
	}
	if root.BalanceFactor() == -2 {
		// 左-左旋一次即可
		if root.right.BalanceFactor() == -1 {
			return root.LeftRotate()
		}
		// 右-左旋
		if root.right.BalanceFactor() == 1 {
			err := root.right.RightRotate()
			if err != nil {
				return err
			}
			return root.LeftRotate()
		}
	}

	if root.BalanceFactor() == 2 {
		// 右旋一次即可
		if root.left.BalanceFactor() == 1 {
			return root.RightRotate()
		}
		// 左-右旋
		if root.left.BalanceFactor() == -1 {
			err := root.left.LeftRotate()
			if err != nil {
				return err
			}
			return root.RightRotate()
		}
	}
	return nil
}

func (root *Node) RightRotate() error {
	if root == nil {
		return fmt.Errorf("root is nil")
	}
	// 第一个节点
	first := *root
	// 第二个节点
	second := root.left
	first.left = nil
	// 将第二个节点升为第一个节点
	*root = *second
	// 将第一个节点降为第二个节点
	// 如果原第二个节点有右节点，则将该节点装在原根节点的左节点
	if root.right != nil {
		secondRightSon := *root.right
		first.left = &secondRightSon
	}
	root.right = &first
	return nil
}

func (root *Node) LeftRotate() error {
	if root == nil {
		return fmt.Errorf("root is nil")
	}
	// 第一个节点
	first := *root
	// 第二个节点
	second := root.right
	first.right = nil
	// 将第二个节点升为第一个节点
	*root = *second
	// 将第一个节点降为第二个节点
	//root.left = &first

	// 将第一个节点降为第二个节点
	// 如果原第二个节点有左节点，则将该节点装在原根节点的右节点
	if root.left != nil {
		secondRightSon := *root.left
		first.right = &secondRightSon
	}
	root.left = &first
	return nil
}

func (root *Node) BalanceFactor() int {
	if root == nil {
		return 0
	}
	return root.left.Height() - root.right.Height()
}

func (root *Node) Height() int {
	if root == nil {
		return -1
	}
	return max(root.left.Height(), root.right.Height()) + 1
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

// 中序（左中右）
func (root *Node) InOrder() {
	parent := root
	if parent == nil {
		return
	}
	if parent.left != nil {
		parent.left.InOrder()
	}
	print(root.data, " ")
	if parent.right != nil {
		parent.right.InOrder()
	}
}

// 前序（中左右）
func (root *Node) PreOrder() {
	parent := root
	if parent == nil {
		return
	}
	print(root.data, " ")
	if parent.left != nil {
		parent.left.PreOrder()
	}
	if parent.right != nil {
		parent.right.PreOrder()
	}
}

// 后序（左右中）
func (root *Node) PostOrder() {
	parent := root
	if parent == nil {
		return
	}
	if parent.left != nil {
		parent.left.PostOrder()
	}
	if parent.right != nil {
		parent.right.PostOrder()
	}
	print(root.data, " ")
}
