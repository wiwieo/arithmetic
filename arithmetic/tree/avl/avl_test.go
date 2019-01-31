package avl

import (
	"fmt"
	"testing"
)

func TestNode(t *testing.T) {
	var root = &Node{data: 20}
	root.Insert(10)
	root.Insert(30)
	root.Insert(3)
	root.Insert(11)
	n, err := root.Search(3)
	if err != nil {
		t.Errorf(fmt.Sprintf("%s", err))
	}
	println("avl height: ", root.Height())
	println("avl balance factor: ", root.BalanceFactor())
	println(fmt.Sprintf("%+v", n))
	println("#############前序##############")
	root.PreOrder()
	println()
	println("#############中序##############")
	root.InOrder()
	println()
	println("#############后序##############")
	root.PostOrder()
	println()
}
func TestNode_InsertAndRotate(t *testing.T) {
	var root = &Node{data: 20}
	err := root.InsertAndRotate(10)
	if err != nil {
		t.Fatal(err)
	}
	err = root.InsertAndRotate(30)
	if err != nil {
		t.Fatal(err)
	}
	err = root.InsertAndRotate(3)
	if err != nil {
		t.Fatal(err)
	}
	err = root.InsertAndRotate(4)
	if err != nil {
		t.Fatal(err)
	}
	err = root.InsertAndRotate(8)
	if err != nil {
		t.Fatal(err)
	}
	err = root.InsertAndRotate(6)
	if err != nil {
		t.Fatal(err)
	}
	err = root.InsertAndRotate(5)
	if err != nil {
		t.Fatal(err)
	}
	err = root.InsertAndRotate(7)
	if err != nil {
		t.Fatal(err)
	}
	err = root.InsertAndRotate(2)
	if err != nil {
		t.Fatal(err)
	}

	err = root.DeleteAdnRotate(4)
	if err != nil {
		t.Fatal(err)
	}
	err = root.DeleteAdnRotate(2)
	if err != nil {
		t.Fatal(err)
	}
	err = root.DeleteAdnRotate(3)
	if err != nil {
		t.Fatal(err)
	}
	err = root.DeleteAdnRotate(30)
	if err != nil {
		t.Fatal(err)
	}

	println("#############前序##############")
	root.PreOrder()
	println()
	println("#############中序##############")
	root.InOrder()
	println()
	println("#############后序##############")
	root.PostOrder()
	println()
}

func TestNode_Rotate(t *testing.T) {
	var root = &Node{data: 20}
	root.Insert(10)
	root.Insert(11)
	println("############旋转前#前序##############")
	root.PreOrder()
	println()
	println("############旋转前#中序##############")
	root.InOrder()
	println()
	println("############旋转前#后序##############")
	root.PostOrder()
	println()
	root.Rotate()
	println("############旋转后#前序##############")
	root.PreOrder()
	println()
	println("############旋转后#中序##############")
	root.InOrder()
	println()
	println("############旋转后#后序##############")
	root.PostOrder()
	println()
}
func TestNode_RightRotate(t *testing.T) {
	var root = &Node{data: 20}
	root.Insert(10)
	root.Insert(1)
	println("############旋转前#前序##############")
	root.PreOrder()
	println()
	println("############旋转前#中序##############")
	root.InOrder()
	println()
	println("############旋转前#后序##############")
	root.PostOrder()
	println()
	root.RightRotate()
	println("############旋转后#前序##############")
	root.PreOrder()
	println()
	println("############旋转后#中序##############")
	root.InOrder()
	println()
	println("############旋转后#后序##############")
	root.PostOrder()
	println()
}
func TestNode_LeftRotate(t *testing.T) {
	var root = &Node{data: 20}
	root.Insert(30)
	root.Insert(40)
	println("############旋转前#前序##############")
	root.PreOrder()
	println()
	println("############旋转前#中序##############")
	root.InOrder()
	println()
	println("############旋转前#后序##############")
	root.PostOrder()
	println()
	root.LeftRotate()
	println("############旋转后#前序##############")
	root.PreOrder()
	println()
	println("############旋转后#中序##############")
	root.InOrder()
	println()
	println("############旋转后#后序##############")
	root.PostOrder()
	println()
}
