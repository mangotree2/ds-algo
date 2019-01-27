package counting

func ContingSort(a []int, n int) {
	if n <= 1 {
		return
	}

	max := a[0]

	for i := range a {
		if a[i] > max {
			max = a[i]
		}
	}

	c := make([]int,max+1)
	//每个元素的个数
	for i := range a {
		c[a[i]]++
	}

	for i := 1; i <= max; i++ {
		c[i] = c[i-1] + c[i]
	}

	tmp := make([]int,n)
	for i := n -1; i >= 0; i-- {
		index := c[a[i]] -1;
		tmp[index] = a[i]
		c[a[i]]--;

	}


	copy(a,tmp)
}
