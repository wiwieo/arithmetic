package sort

import (
	"testing"
	"time"
	"math/rand"
	"fmt"
)

func TestQuickSort(t *testing.T) {
	ary := []int{}
	for i := 0; i < 1000;i ++{
		ary = append(ary, rand.Intn(int(time.Now().Unix())%1000))
	}
	MergeSort(ary)
	fmt.Println(ary)
}