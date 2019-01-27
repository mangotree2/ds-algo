package quick

func QuickSort(a []int)  {
	n := len(a)
	if n <= 1 {
		return
	}

	quickSort(a,0,n-1)
}

func quickSort(a []int, start, end int) {
	if start >= end {
		return
	}

	pivot := partition(a,start,end)
	quickSort(a,start,pivot)
	quickSort(a,pivot+1,end)
}

func partition(a []int, start, end int) int {
	pivotv := a[start]

	for start < end {
		for start < end && a[end] > pivotv {
			end--
		}
		a[start] = a[end]
		for start < end && a[start] < pivotv {
			start++
		}
		a[end] = a[start]
	}

	a[start] = pivotv
	return start
}