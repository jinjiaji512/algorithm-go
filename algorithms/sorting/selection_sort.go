package sorting

//选择排序
//每次选出最大的值，放在末尾
func SelectionSort(s []int) []int {
	if s == nil || len(s) < 2 {
		return s
	}
	for i := 0; i < len(s); i++ {
		max := 0
		for j := 0; j < len(s)-i; j++ {
			if s[j] > s[max] {
				max = j
			}
		}
		s[len(s)-i-1], s[max] = s[max], s[len(s)-i-1]
	}
	return s
}
