package mergesortcounting

func Count(a []int) int {

	num := 0
	mergeSortCounting(a,0,len(a)-1 ,&num)

	return num


}

func mergeSortCounting(a []int, p, r int,num *int) {

	if p >= r {
		return
	}

	q := (p + r) /2

	mergeSortCounting(a,p,q,num)
	mergeSortCounting(a,q+1,r,num)

	merge(a,p,q,r,num)

}

func merge(a []int,p,q,r int,num *int) {
	i := p
	j := q+1
	k := 0
	temp := make([]int,r-p)
	for i < q && j < r {
		if a[i] <= a[j] {
			temp[k] = a[i]
			k++
			i++
		} else {
			*num += (q-i)+1
			temp[k] = a[j]
			k++
			j++
		}

	}

	for i <= q {
		temp[k] = a[i]
		k++
		i++
	}

	for j <= r {
		temp[k] = a[j]
		k++
		j++
	}

	copy(a[p:r+1],temp)


}