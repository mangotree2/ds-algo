package slection

func SlectionSort(a []int) []int {
	n := len(a)
	if n <= 1 {
		return a
	}

	for i := 0; i < n ; i++ {
		minIndex := i
		for j:=i+1;j<n;j++{
			if a[j] < a[minIndex] {
				minIndex = j
			}

		}
		a[i],a[minIndex] = a[minIndex],a[i]
	}

}
