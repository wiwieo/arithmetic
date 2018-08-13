package sort

func DirectSelectSort(ary []int){
	// 	外层循环，用于记录有序数组。前面为有序
	for i := 0; i < len(ary); i ++{
		// 最小值
		min := ary[i]
		minIdx := i
		// 内层循环，用于找出剩下数组元素中的最小值
		for j := i + 1; j < len(ary);j ++{
			if min > ary[j]{
				min = ary[j]
				minIdx = j
			}
		}
		// 将最小值放在i的位置上
		ary[i], ary[minIdx] = ary[minIdx], ary[i]
	}
}
