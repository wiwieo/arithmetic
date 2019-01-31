package red_black

import (
	"fmt"
)

// 红黑树详情参见：
// https://towardsdatascience.com/red-black-binary-tree-maintaining-balance-e342f5aa6f5
// 将原有的性质改为如下，便于理解
// 1、The root node will always be black.
// 2、The incoming node is always red.
// 3、A double red violation will occur if a child and parent are both red.
// 4、Resolve a violation with recoloring(s) and/or a restructuring.

// 原性质：
// 1、Every node has a color either red or black.
// 2、Root of tree is always black.
// 3、There are no two adjacent red nodes (A red node cannot have a red parent or red child).
// 4、Every path from a node (including root) to any of its descendant NULL node has the same number of black nodes.
const (
	Red   color = false
	Black color = true
	Left  rank  = 1
	Right rank  = 2
)

type color bool
type rank int

type Node struct {
	data   int
	color  color
	rank   rank
	left   *Node
	right  *Node
	parent *Node
}

func (current *Node) Insert(data int) (*Node, error) {
	var temp = &Node{data: data}
	if current == nil {
		return nil, fmt.Errorf("current is nil")
	}
	var parent = current
	// 插入
	for {
		if data <= parent.data {
			if parent.left == nil {
				parent.left = temp
				temp.rank = Left
				temp.parent = parent
				break
			}
			parent = parent.left
			continue
		}
		if data > parent.data {
			if parent.right == nil {
				parent.right = temp
				temp.rank = Right
				temp.parent = parent
				break
			}
			parent = parent.right
			continue
		}
	}
	return temp, nil
}

func (current *Node) InsertAndBalance(data int) error {
	x, err := current.Insert(data)
	if err != nil {
		return err
	}
	x.ReBalance()
	return nil
}

// 重新平衡
// 如果叔节点是红色，则重新着色，并将爷节点设为当前节点，并重新平衡
// 如果叔节点是黑色，则进行旋转
func (current *Node) ReBalance() {
	// 根节点，直接变黑并退出
	if current.parent == nil {
		current.color = Black
		return
	}

	// 父节点如果是黑色，直接退出
	if current.parent.color == Black {
		return
	}

	parent := current.parent
	// 父节点为根节点，则将节点变黑退出
	if parent.parent == nil {
		parent.color = Black
		return
	}

	grandpa := current.parent.parent
	// 如果爷节点为根节点，则将节点变黑退出
	if grandpa.parent == nil {
		grandpa.color = Black
	}

	// 根据父节点位置，获取叔节点
	var uncle = func() *Node {
		switch parent.rank {
		case Left:
			return grandpa.right
		case Right:
			return grandpa.left
		}
		return nil
	}()

	if uncle == nil || uncle.color == Black { // 如果叔节点为黑色，则进行旋转变更结构
		current.Rotate()
	} else if uncle.color == Red { // 叔节点为红，则重新着色，并设置爷节点为当前节点再来一遍
		current.Colors(uncle)
		current.parent.parent.ReBalance()
	}
}

// 根据节点位置分为四种旋转
// 1、left-left：父为左，子为左，右旋
// 2、left-right：父为左，子为右，先左旋再右旋
// 3、right-right：父为右，子为右，左旋
// 4、right-left：父为右，子为左，先右旋再左旋
func (current *Node) Rotate() {
	// 以父节点为中心旋转
	parent := current.parent
	switch current.parent.rank {
	case Left:
		switch current.rank {
		case Left:
			parent.RightRotate()
		case Right:
			current.LeftRotate()
			parent.RightRotate()
		}
	case Right:
		switch current.rank {
		case Left:
			current.RightRotate()
			parent.LeftRotate()
		case Right:
			current.LeftRotate()
		}
	}
}

// 右旋规则
// 0、将当前节点变黑，父节点变红
// 1、将父节点（p）独立出来
// 2、将右子节点（r）独立出来
// 3、将r连接为p的左节点
// 4、将p连接为x的右节点
func (current *Node) RightRotate() {
	// 0:将当前节点变黑，父节点变红
	current.color = Black
	current.parent.color = Red

	// 1:将父节点（p）独立出来
	p := *current.parent
	p.left = nil
	p.parent = nil

	// 2:将右子节点（r）独立出来
	if current.right != nil {
		r := *current.right
		current.right = nil
		r.parent = nil

		// 3:将r连接为p的左节点
		r.rank = Left
		p.left = &r
		r.parent = &p
	}

	// 4:将p连接为x的右节点
	p.rank = Right
	current.right = &p
	p.parent = current
}

// 左旋规则
// 0、将当前节点变黑，父节点变红
// 1、将父节点（p）独立出来
// 2、将左子节点（r）独立出来
// 3、将r连接为p的右节点
// 4、将p连接为x的左节点
func (current *Node) LeftRotate() {
	// 0:将当前节点变黑，父节点变红
	current.color = Black
	current.parent.color = Red

	// 1:将父节点（p）独立出来
	p := *current.parent
	p.right = nil
	p.parent = nil

	// 2:将左子节点（r）独立出来
	if current.left != nil {
		r := *current.left
		current.left = nil
		r.parent = nil

		// 3:将r连接为p的右节点
		r.rank = Right
		p.right = &r
		r.parent = &p
	}

	// 4:将p连接为x的左节点
	p.rank = Left
	current.left = &p
	p.parent = current
	switch p.parent.rank {
	case Left:
		current.left = &p
	case Right:
		current.right = &p
	}
}

// 叔叔节点为红色，则进行重新着色
// 即：父叔节点变黑，爷节点变红（非根节点，否则变黑）
func (current *Node) Colors(uncle *Node) {
	current.parent.color = Black
	current.parent.parent.color = Red
	uncle.color = Black
}

func (current *Node) Search(data int) (*Node, error) {
	for {
		if current == nil {
			return nil, fmt.Errorf("node is not exist")
		}
		if data == current.data {
			return current, nil
		}
		if data < current.data {
			current = current.left
			continue
		}
		if data > current.data {
			current = current.right
			continue
		}
	}
}


// 中序（左中右）
func (current *Node) InOrder() {
	parent := current
	if parent == nil {
		return
	}
	if parent.left != nil {
		parent.left.InOrder()
	}
	print(fmt.Sprintf("%d(%s) ", current.data, func()string{if current.color == Red {return "Red"}else if current.color == Black {return "Black"}; return ""}()))
	if parent.right != nil {
		parent.right.InOrder()
	}
}

// 前序（中左右）
func (current *Node) PreOrder() {
	parent := current
	if parent == nil {
		return
	}
	print(fmt.Sprintf("%d(%s) ", current.data, func()string{if current.color == Red {return "Red"}else if current.color == Black {return "Black"}; return ""}()))
	if parent.left != nil {
		parent.left.PreOrder()
	}
	if parent.right != nil {
		parent.right.PreOrder()
	}
}

// 后序（左右中）
func (current *Node) PostOrder() {
	parent := current
	if parent == nil {
		return
	}
	if parent.left != nil {
		parent.left.PostOrder()
	}
	if parent.right != nil {
		parent.right.PostOrder()
	}
	print(fmt.Sprintf("%d(%s) ", current.data, func()string{if current.color == Red {return "Red"}else if current.color == Black {return "Black"}; return ""}()))
}
