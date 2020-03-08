package sorting

//冒泡排序
func BubbleSort(s []int) []int {
	//将最大值排在末尾,i 轮数，j下标
	for i := 0; i < len(s); i++ {
		for j := 0; j < len(s)-1-i; j++ {
			if s[j] > s[j+1] {
				s[j], s[j+1] = s[j+1], s[j]
			}
		}
	}
	return s
}
