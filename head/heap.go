package head

type Heap struct {
	a []int
	cap int
	count int
	//less
	//swap
}
//todo 小顶堆 前K大，优先队列，中位数

func NewHeap(cap int) *Heap {
	return &Heap{
		cap:cap,
		a:make([]int,cap+1),
	}
}

func (h *Heap) Insert(v int) {

	if h.count == h.cap {
		return
	}

	h.count++
	h.a[h.count] = v

	i:= h.count
	parent := i /2
	//less
	for parent > 0 && h.a[parent] < h.a[i] {
		//swap
		h.a[parent],h.a[i] = h.a[i],h.a[parent]
		i = parent
		parent = i/2

	}
}

func (h *Heap) removeMax() {
	if h.count == 0 {
		return
	}

	h.a[1],h.a[h.count] = h.a[h.count] , h.a[1]
	h.count--

}

func heapifyUptoDown(a []int, count int) {
	for i :=1 ;i <= count/2 ;{
		maxIndex := 1
		if a[i] < a[i*2] {
			maxIndex = i*2
		}

		if i*2+1 <= count && a[maxIndex] < a[i*2+1] {
			maxIndex = i*2+1
		}

		if maxIndex == i {
			break
		}

		a[i], a[maxIndex] = a[maxIndex],a[i]
		i = maxIndex

	}
}

