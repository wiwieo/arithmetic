package sort

func MergeSort(ary []int) {
	sort(ary, 0, len(ary)-1)
}

// 分组有序合并
func sort(ary []int, l, h int) {
	if l >= h {
		return
	}
	m := (l + h) / 2
	// 左
	sort(ary, l, m)
	// 右
	sort(ary, m+1, h)
	// 合
	merge(ary, l, m, h)
}

// 合并
func merge(ary []int, l, m, h int) {
	// 分别用于遍历左右子数组
	i, j := l, m+1
	// 用于暂时存放合并后的有序数组
	temp := []int{}

	// 循环分别将两个有序数组中的小的部分放入临时数组中
	for i <= m && j <= h {
		if ary[i] > ary[j] {
			temp = append(temp, ary[j])
			j++
		} else {
			temp = append(temp, ary[i])
			i++
		}
	}

	// 上面遍历完之后，剩下的需要再全部放入临时数组中（剩下的是其中一个数组中的，且大于临时数组中任何数据，直接追加）
	if i <= m {
		temp = append(temp, ary[i:m+1]...)
	}

	if j <= h {
		temp = append(temp, ary[j:h+1]...)
	}

	// 将临时数组再放回原数组
	for i = 0; i < len(temp); i++ {
		ary[l+i] = temp[i]
	}
}
