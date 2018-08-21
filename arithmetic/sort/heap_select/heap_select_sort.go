package sort

// 堆排序
// 利用堆的属性，来查找出最小/最大值
// 堆顶为最大/最小，将最后一个元素与堆顶交换即可
func HeapSelectSort(ary []int){
	for i := 0; i < len(ary); i ++{
		adjustMaxHeap(ary, len(ary) - i)
		ary[0], ary[len(ary) - 1 - i] = ary[len(ary) - 1 - i], ary[0]
	}
}

// 数组本身即是被打乱的堆，需要进行调整
// 叶子个数为：n/2 + 1
// 第一个非叶子序号为：n/2 - 1
func adjustMaxHeap(ary []int, last int){
	for i := last / 2 - 1; i >= 0 ; i --{
		if 2 * i + 1 < last && ary[2 * i + 1] > ary[i]{
			ary[2 * i + 1], ary[i] = ary[i], ary[2 * i + 1]
		}
		if 2 * i + 2 < last && ary[2 * i + 2] > ary[i]{
			ary[2 * i + 2], ary[i] = ary[i], ary[2 * i + 2]
		}
	}
}