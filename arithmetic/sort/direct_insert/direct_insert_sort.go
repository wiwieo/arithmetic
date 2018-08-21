package sort

// 直接插入排序
// 1、从第二个元素开始，依次与前面的值进行比对（循环依次比较）
// 2、根据大小依次插入到对应的位置
func DirectInsertSort(ary []int) {
	// 外循环，用于取出每个待比较插入的值
	for i := 1; i < len(ary); i++ {
		// 待比较插入的值
		temp := ary[i]
		// 内循环，已排序好的序列，和temp一一比较后，插入到相应的位置
		j := i - 1
		for ; j >= 0; j-- {
			// 将大于temp的值往后移动一位
			if temp < ary[j] {
				ary[j+1] = ary[j]
				continue
			}
			// 因为是有序序列，找到位置后直接返回
			break
		}
		// 将temp插入指定的空出的位置
		ary[j+1] = temp
	}
}
