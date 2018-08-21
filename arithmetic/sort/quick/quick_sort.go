package sort

// 快速排序
// 1、以某个值（此处以最后一个值）为基准，将数组分为左右两部分，左边均小于基准值，右边均大于基准值
// 2、循环进行上述操作
func QuickSort(ary []int) {
	partial(ary)
}

func partial(ary []int) {
	// 当数组内容只有一个值时，直接返回
	if len(ary) <= 1 {
		return
	}

	// 当数组只有两个值，如果前大后小则直接交换，无需走下面流程
	if len(ary) == 2 {
		if ary[0] > ary[1] {
			ary[0], ary[1] = ary[1], ary[0]
		}
		return
	}

	// 数组前后移动的下标
	var start, end int = 0, len(ary) - 1
	// 取数组最后一个值作为基准值，小于此值放在数组前面，大于此值放在数组后面
	temp := ary[end]
	for start < len(ary)-1 {
		if end < 0 {
			break
		}

		// 数组后部分存放的值大于基准值，所以当大于基准值时，则位置不变，
		// 且将下标向前移动一位，直到找到小于基准值的值，用于后续交换使用。
		if temp < ary[end] {
			end--
			continue
		}

		// 当前后下标相遇，表示数组遍历完成，返回
		if start >= end {
			break
		}

		// 数组后部分存放的值小于基准值，当大于基准值时，需要和后部分的数据进行交换
		if temp < ary[start] {
			ary[start], ary[end] = ary[end], ary[start]
			// 交换之后，将后部分下标前移和前部分下后移
			end--
		}
		start++
	}

	// 遍历完之后，end的位置表示为分界线位置
	// 即 end后部分为大于基准值的部分，可以将基准值和后部分第一个值进行交换
	// 如果end值为数组最后一位，则表明所有值均在左边，不需要交换且在重新分组时，不包含最后一位
	if end != len(ary)-1 {
		ary[len(ary)-1], ary[end+1] = ary[end+1], ary[len(ary)-1]
	} else {
		end--
	}
	// 以基准值为界，分成两个不同的区域，再分别比较
	partial(ary[:end+1])
	partial(ary[end+1:])
}

func partition1(a []int, lo, hi int) int {
	pivot := a[hi]
	i := lo - 1
	for j := lo; j < hi; j++ {
		if a[j] < pivot {
			i++
			a[j], a[i] = a[i], a[j]
		}
	}
	a[i+1], a[hi] = a[hi], a[i+1]
	return i + 1
}

func QuickSort1(a []int, lo, hi int) {
	if lo >= hi {
		return
	}
	p := partition1(a, lo, hi)
	QuickSort1(a, lo, p-1)
	QuickSort1(a, p+1, hi)
}
