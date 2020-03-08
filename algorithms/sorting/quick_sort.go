package sorting

//快速排序
func QuickSort(s []int) []int {
	//快速排序，选定基准，小的放在基准左边，大的放在基准右边，重复递归
	//双指针
	//默认选定一个元素为基准，双向指针移动，左边slice最后一个元素，选定为新的基准元素进行递归

	if s == nil || len(s) < 2 {
		return s
	}

	oldPVal := s[0]
	newP := 0
	for i, j := 1, len(s)-1; i <= j; {
		if s[i] < oldPVal {
			newP = i
			i++
		} else if s[j] > oldPVal {
			j--
		} else {
			//左边比旧基准大，右边比旧基准小，交换
			s[i], s[j] = s[j], s[i]
			newP = i
			i++
		}
	}
	//旧基准和新基准交换，按新基准左右递归
	if newP == 0 {
		newP = 1
	} else {
		s[0], s[newP] = s[newP], s[0]
	}
	QuickSort(s[:newP])
	QuickSort(s[newP:])
	return s
}
