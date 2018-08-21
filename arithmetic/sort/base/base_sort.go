package sort

// 基数排序
// 1、查出最大值
// 2、将数组中的元素按照位数依次排序
// 3、先按个位排，其次十位。。。
func BaseSort(ary []int){
	max := max(ary)
	for idx := 1; max/idx > 1; idx*=10{
		singleSort(ary, idx)
	}
}

func singleSort(ary []int, idx int){
	// 二维数组：
	// 第一维下标表示：位数的值
	// 第二维表示：当前位数的值对应的原数组的值的集合
	var quene [10][]int = [10][]int{}
	for _, v := range ary{
		i := (v/idx)%10
		quene[i] = append(quene[i], v)
	}

	// 将每个位数的值按顺序放回原数组
	i := 0
	for _, vs := range quene{
		if len(vs) == 0{
			continue
		}
		for _, v := range vs{
			ary[i] = v
			i ++
		}
	}
}

// 获取数组中最大值
func max(ary []int) int {
	max := 0
	for _, v := range ary{
		if max < v {
			max = v
		}
	}
	return max
}