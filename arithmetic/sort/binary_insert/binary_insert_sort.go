package sort

// 二分插入排序
// 1、从第二个元素开始，依次与前面的值进行比对（使用二分查找法进行比较）
// 2、根据大小依次插入到对应的位置
func BinaryInsertSort(ary []int){
	// 外循环，用于取出每个待比较插入的值
	for i := 1; i < len(ary); i++ {
		// 待比较插入的值
		temp := ary[i]
		// 内循环，已排序好的序列，和temp一一比较后，插入到相应的位置
		// 使用二分比较法，找到对应的位置
		thePlace := BinarySearch(ary, 0, i - 1, temp)

		// 找到可插入的位置后，将大于temp的值后移一位
		for j := i - 1; j >= thePlace; j-- {
			ary[j+1] = ary[j]
		}
		// 将temp插入指定的空出的位置
		ary[thePlace] = temp
	}
}

func BinarySearch(ary []int, left, right, temp int) int{
	mid := 0
	for{
		if left > right {
			break
		}
		mid = (left + right) / 2
		if ary[mid] < temp {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return left
}