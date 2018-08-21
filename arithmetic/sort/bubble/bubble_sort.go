package sort

// 冒泡
// 依次将所有元素与其他未排序元素进行比较
// 轻升重降
func BubbleSort(ary []int){
	for i := 0;i < len(ary); i ++{
		for j := 0; j < len(ary) - i - 1; j ++{
			if ary[j+1] < ary[j] {
				ary[j+1], ary[j] = ary[j], ary[j + 1]
			}
		}
	}
}
