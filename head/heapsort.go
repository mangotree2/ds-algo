package head

func BuildHeap(a []int, n int) {

	for i := n/2; i >= 1;i-- {
		heapifyUptoDownByIndex(a,i,n)
	}
}

func heapifyUptoDownByIndex(a []int, top, count int) {

	for i := top; i <= count; {
		maxIndex := i

		if a[i] <= a[i*2] {
			maxIndex = i*2
		}

		if i*2+1 <= count && a[maxIndex] < a[i*2+1] {
			maxIndex = i*2+1
		}

		if maxIndex == i {
			break
		}

		a[maxIndex],a[i] = a[i],a[maxIndex]
		i = maxIndex
	}
}

func Sort(a []int,n int) {

	BuildHeap(a,n)
	k := n
	for k >= 1 {
		a[1] ,a[k] = a[k],a[1]
		heapifyUptoDownByIndex(a,1,k-1)
		k--
	}

}