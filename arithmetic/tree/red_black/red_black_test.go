package red_black

import (
	"fmt"
	"testing"
)

func TestNode(t *testing.T) {
	var root = &Node{data: 20, color: Black}
	root.Insert(10)
	root.Insert(30)
	root.Insert(3)
	root.Insert(11)
	n, err := root.Search(3)
	if err != nil {
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
func TestNode_InsertAndBalance(t *testing.T) {
	var root = &Node{data: 20}
	err := root.InsertAndBalance(10)
	if err != nil {
		t.Fatal(err)
	}
	err = root.InsertAndBalance(30)
	if err != nil {
		t.Fatal(err)
	}
	err = root.InsertAndBalance(3)
	if err != nil {
		t.Fatal(err)
	}
	err = root.InsertAndBalance(4)
	if err != nil {
		t.Fatal(err)
	}
	//err = root.InsertAndBalance(8)
	//if err != nil {
	//	t.Fatal(err)
	//}
	//err = root.InsertAndBalance(6)
	//if err != nil {
	//	t.Fatal(err)
	//}
	//err = root.InsertAndBalance(5)
	//if err != nil {
	//	t.Fatal(err)
	//}
	//err = root.InsertAndBalance(7)
	//if err != nil {
	//	t.Fatal(err)
	//}
	//err = root.InsertAndBalance(2)
	//if err != nil {
	//	t.Fatal(err)
	//}

	//err = root.DeleteAdnRotate(4)
	//if err != nil {
	//	t.Fatal(err)
	//}
	//err = root.DeleteAdnRotate(2)
	//if err != nil {
	//	t.Fatal(err)
	//}
	//err = root.DeleteAdnRotate(3)
	//if err != nil {
	//	t.Fatal(err)
	//}
	//err = root.DeleteAdnRotate(30)
	//if err != nil {
	//	t.Fatal(err)
	//}

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
