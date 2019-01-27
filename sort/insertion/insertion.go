package insertion

func InsertionSort(a []int) []int {
	n := len(a)
	if n <= 1 {
		return a
	}

	for i := 1; i < n; i++ {
		v := a[i]
		j := i -1
		for ;j >= 0; j-- {
			if a[j] > v {
				a[j+1] = a[j]
			} else {
				break
			}
		}
		a[j+1] = v
	}


}