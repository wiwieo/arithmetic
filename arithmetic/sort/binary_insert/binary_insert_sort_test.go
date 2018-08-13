package sort

import (
	"testing"
	"fmt"
	"time"
	"math/rand"
)

func TestQuickSort(t *testing.T) {
	ary := []int{3,1,2,0,7,4,8}
	for i := 0; i < 10000;i ++{
		ary = append(ary, rand.Intn(int(time.Now().Unix())%1000))
	}
	BinaryInsertSort(ary)
	fmt.Println(ary)
}