package started

// 给定一个 haystack 字符串和一个 needle 字符串,
// 在 haystack 字符串中找出 needle 字符串出现的第一个位置 (从 0 开始)。
// 如果不存在，则返回 -1。
func strStr(haystack, needle string) int {
	if len(needle) == 0 {
		return 0
	}
	if len(haystack) < len(needle) {
		return -1
	}
	// todo: 不需要遍历到haystack结束位置
	for i := 0; i < len(haystack)-len(needle)+1; i++ {
		for j := 0; j < len(needle); j++ {
			if haystack[j+i] != needle[j] {
				break
			}
			// todo: 找到子串
			if j == len(needle)-1 {
				return i
			}
		}
	}
	return -1
}

// 给定一组不含重复元素的整数数组 numA, 返回该数组所有可能的子集（幂集）。
// fixme 回溯算法学习下
func subsets(numA []int) [][]int {
	var result [][]int

	for i := 0; i < len(numA); i++ {
		for j := 0; j < len(numA)-i; j++ {
			if i+j > i {
				result = append(result, numA[i:i+j])
			}
		}
	}
	return result
}
