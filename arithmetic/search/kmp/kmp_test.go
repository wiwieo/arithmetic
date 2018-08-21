package search

import (
	"testing"
	"fmt"
)

func TestQuickSort(t *testing.T) {
	idx := Kmp("你好，我是XXX", "我")
	fmt.Println(idx)
}