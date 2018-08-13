package sort

func ShellInsertSort(ary []int) {
	//希尔排序
	// { 49, 38, 65, 97, 76, 13, 27, 49, 78, 34, 12, 64, 1 }
	d := len(ary)
	for {
		d = d / 2
		for x := 0; x < d; x++ {
			for i := x + d; i < len(ary); i = i + d {
				temp := ary[i]
				var j int
				for j = i - d; j >= 0 && ary[j] > temp; j = j - d {
					ary[j+d] = ary[j]
				}
				ary[j+d] = temp
			}
		}
		if d == 1 {
			break
		}
	}
}
