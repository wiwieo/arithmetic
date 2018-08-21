package search

// 字符串匹配KMP
// 1、将需匹配字符串建立一张字符匹配表（第N个的匹配值为：1~N的前缀与后缀相同字符串的最大长度）
// 2、循环待匹配字符串，依次与需匹配字符串进行比较。如果不同，则将待匹配字串下标后移(已匹配的字符数 - 对应的部分匹配值)

// 建立匹配表
func buildMatchTable(str string) []int {
	var matchTable []int = make([]int, len(str))
	for idx := range str {
		matchTable[idx] = matchTableValue(str[:idx])
	}
	return matchTable
}

// 一个字符串的前缀与后缀相同字符串的最大长度
func matchTableValue(str string) int {
	prefixs := prefix(str)
	suffixs := suffix(str)

	maxL := 0
	for _, p := range prefixs {
		for _, s := range suffixs {
			if p == s {
				if maxL < len(p) {
					maxL = len(p)
				}
			}
		}
	}
	return maxL
}

// 获取字符串所有前缀
func prefix(str string) []string {
	var prefixs []string
	for idx := range str {
		// 前缀不包含自身
		if idx == len(str)-1 {
			break
		}
		prefixs = append(prefixs, str[:idx])
	}
	return prefixs
}

// 获取字符串所有后缀
func suffix(str string) []string {
	var prefixs []string
	for idx := range str {
		// 后缀不包含自身
		if idx == 0 {
			continue
		}
		prefixs = append(prefixs, str[idx:])
	}
	return prefixs
}

// 比较
func Kmp(ori, search string) int {
	// 待查找字符串长度大于原字符串，则必不存在
	if len(ori) < len(search){
		return -1
	}

	// 如果长度相同，则直接比较两个字符串是否相等
	if len(ori) == len(search){
		if ori == search {
			return 0
		}
		return -1
	}

	matchTable := buildMatchTable(search)

	// 标识是否查找成功
	match := true

	// 外层循环，控制待查找字符串的下标移动
	i := 0
	for i < len(ori) {
		// 待查找字符串的下标，用于和需匹配字符串同步移动
		var k = i
		for j := 0; j < len(search); j++ {
			if ori[k] == search[j] {
				// 字符相同，且是最后一位字符，则说明已经匹配成功
				if j >= len(search)-1 {
					return i
				}
				// 匹配成功，将待查询字符串下标移动一位
				k++
			} else {
				// 如果字符串不匹配，则将下标移动（已匹配的字符数 - 对应部分的匹配值）
				v := matchTable[j]
				if j <= v {
					v = j - 1
				}
				i = i + j - v
				match = false
				break
			}
		}
	}
	// 匹配不成功，则返回-1，表示没找到
	if !match {
		i = -1
	}
	return i
}
