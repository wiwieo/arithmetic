package bst

import (
	"fmt"
	"testing"
)

func TestNode(t *testing.T) {
	var root = &Node{data:20}
	root.Insert(10)
	root.Insert(30)
	root.Insert(3)
	root.Insert(11)
	n, err := root.Search(3)
	if err != nil{
		t.Errorf(fmt.Sprintf("%s", err))
	}
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
