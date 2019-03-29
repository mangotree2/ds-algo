package minstack


//设计一个支持 push，pop，top 操作，并能在常数时间内检索到最小元素的栈。
//
//push(x) -- 将元素 x 推入栈中。
//pop() -- 删除栈顶的元素。
//top() -- 获取栈顶元素。
//getMin() -- 检索栈中的最小元素。
//示例:
//
//MinStack minStack = new MinStack();
//minStack.push(-2);
//minStack.push(0);
//minStack.push(-3);
//minStack.getMin();   --> 返回 -3.
//minStack.pop();
//minStack.top();      --> 返回 0.
//minStack.getMin();   --> 返回 -2.

type MinStack struct {
	diff []int
	minValue int

}


/** initialize your data structure here. */
func Constructor() MinStack {

	return MinStack{
		diff:make([]int,0,16),
	}
}


func (this *MinStack) Push(x int)  {
	if len(this.diff) == 0 {
		this.diff = append(this.diff,0)
		this.minValue = x

	} else {
		diff := x-this.minValue
		this.diff = append(this.diff,diff)
		if diff < 0 {
			this.minValue = x
		}
	}

}


func (this *MinStack) Pop()  {
	if len(this.diff) >0 {
		diff := this.diff[len(this.diff)-1]
		if diff < 0 {
			this.minValue -=diff
		}
		this.diff = this.diff[:len(this.diff)-1]
	}

}


func (this *MinStack) Top() int {
	if len(this.diff) > 0 {
		diff := this.diff[len(this.diff)-1]
		if diff < 0 {
			return this.minValue
		} else {
			return this.minValue + diff
		}
	}
	return 0
}


func (this *MinStack) GetMin() int {
	return this.minValue
}


/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
