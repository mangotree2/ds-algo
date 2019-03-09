package circulardqueue

//循环队列
type MyCircularDeque struct {
	cap int
	len int
	front int//头部下一个要插入的位置
	rear int//尾部下一个要插入的位置
	queue []int
}


/** Initialize your data structure here. Set the size of the deque to be k. */
func Constructor(k int) MyCircularDeque {
	return MyCircularDeque{
		cap:   k,
		len:   0,
		front: 1,
		rear:  0,
		queue: make([]int,k),
	}
}

func (this *MyCircularDeque) findIndex(index int) int {
	if index < 0 {
		return index + this.cap
	} else if index >= this.cap {
		return index%this.cap
	} else {
		return index
	}
}

/** Adds an item at the front of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) InsertFront(value int) bool {
	if this.IsFull() {
		return false
	}
	this.queue[this.front] = value
	this.front = this.findIndex(this.front+1)
	this.len++
	return true
}


/** Adds an item at the rear of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) InsertLast(value int) bool {
	if this.IsFull() {
		return false
	}
	this.queue[this.rear] = value
	this.rear = this.findIndex(this.rear-1)
	this.len++
	return true
}


/** Deletes an item from the front of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) DeleteFront() bool {
	if this.IsEmpty() {
		return false
	}

	index := this.findIndex(this.front-1)
	this.queue[index] = -1
	this.front=index
	this.len--
	return true

}


/** Deletes an item from the rear of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) DeleteLast() bool {
	if this.IsEmpty() {
		return false
	}

	index := this.findIndex(this.rear+1)
	this.queue[index] = -1
	this.rear = index
	this.len--
	return true
}


/** Get the front item from the deque. */
func (this *MyCircularDeque) GetFront() int {
	if this.IsEmpty() {
		return -1
	}

	return this.queue[this.findIndex(this.front-1)]
}


/** Get the last item from the deque. */
func (this *MyCircularDeque) GetRear() int {
	if this.IsEmpty() {
		return -1
	}

	return this.queue[this.findIndex(this.rear+1)]
}


/** Checks whether the circular deque is empty or not. */
func (this *MyCircularDeque) IsEmpty() bool {
	return  this.len ==0
}


/** Checks whether the circular deque is full or not. */
func (this *MyCircularDeque) IsFull() bool {

	return this.cap == this.len

}


/**
 * Your MyCircularDeque object will be instantiated and called as such:
 * obj := Constructor(k);
 * param_1 := obj.InsertFront(value);
 * param_2 := obj.InsertLast(value);
 * param_3 := obj.DeleteFront();
 * param_4 := obj.DeleteLast();
 * param_5 := obj.GetFront();
 * param_6 := obj.GetRear();
 * param_7 := obj.IsEmpty();
 * param_8 := obj.IsFull();
 */
