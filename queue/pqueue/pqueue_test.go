package pqueue

import (
	"container/heap"
	"fmt"
	"testing"
)

type user struct {
	tid int64
	uid int64
}

func (u user)String() string {
	return fmt.Sprintf("uid :%2d,tid :%2d",u.uid,u.tid)
}

func TestPriorityQueue(t *testing.T) {



	pq := New(15)

	us := []Item{
		{Value:user{tid :1, uid:1},Priority:0},{Value:user{tid :2, uid:2},Priority:1},{Value:user{tid :3, uid:3},Priority:0},
		{Value:user{tid :1, uid:4},Priority:0},{Value:user{tid :2, uid:5},Priority:1},{Value:user{tid :3, uid:6},Priority:0},
		{Value:user{tid :1, uid:7},Priority:0},{Value:user{tid :2, uid:8},Priority:1},{Value:user{tid :3, uid:9},Priority:0},
		{Value:user{tid :1, uid:10},Priority:0},{Value:user{tid :2, uid:11},Priority:1},{Value:user{tid :3, uid:12},Priority:0},
		{Value:user{tid :1, uid:13},Priority:0},{Value:user{tid :2, uid:14},Priority:1},{Value:user{tid :3, uid:15},Priority:0},

	}

	for i,_ := range us {
		heap.Push(&pq,&us[i])

	}

	var cnt int
	var lastTid, lastTid1 int64
	var tmpMap = make(map[int64]*Item)

	for {

		v:= heap.Pop(&pq)
		if cnt == 0 {
			lastTid =  v.(*Item).Value.(user).tid
			cnt ++

		} else if cnt == 1 {
			//lastTid1 = v.(Item).Value.(user).tid
			_v := v.(*Item).Value
			lastTid1 = _v.(user).tid

			cnt ++
			if lastTid == lastTid1 {
				v = heap.Pop(&pq)

				//heap.Push(&pq,&_v)
			}


		} else if cnt == 2 {
			_v := v.(*Item).Value
			TidOver := _v.(user).tid

			if lastTid == TidOver  || TidOver == lastTid1 {
				v = heap.Pop(&pq)

				//heap.Push(&pq,&_v)
			}
			cnt = 0
		}


		fmt.Println(v)

	}


}

func TestPriorityQueue1(t *testing.T) {
	c := 100
	pq := New(c)

	for i := 0; i < c+1; i++ {
		heap.Push(&pq, &Item{Value: i, Priority: int64(i)})
	}


	for i := 0; i < c+1; i++ {
		item := heap.Pop(&pq)
		fmt.Println(item)
	}

}