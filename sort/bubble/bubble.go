package bubble

func BubbleSort(a []int) []int {
	n := len(a)
	if n <= 1 {
		return a
	}

	for i := 0; i < n ; i++ {
		flag := false
		for j := 0 ;j< n-1; j++ {
			if a[j] > a[j+1] {
				a[j],a[j+1] = a[j+1],a[j]
				flag = true
			}
		}
		if flag == false {
			break
		}
	}

	return a

}