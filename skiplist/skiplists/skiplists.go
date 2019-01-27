package skiplists

import (
	"fmt"
	"math"
	"math/rand"
)

const MAX_LEVEL int = 16

type skipListNode struct {
	v interface{}

	score int
	level int
	forwards []*skipListNode
}

type SkipList struct {
	head *skipListNode
	level int
	length int
}

func newSkipListNode(v interface{},score, level int) *skipListNode {
	return &skipListNode{v :v,score:score,level:level,forwards:make([]*skipListNode,level,level)}
}

func NewSkipList() *SkipList {
	head := newSkipListNode(0,math.MinInt32,MAX_LEVEL)
	return &SkipList{head:head,level:1,length:0}
}

func (s *SkipList) Length() int {
	return s.length
}

func (s *SkipList)Level() int {
	return s.level
}

func (s *SkipList) Insert(v interface{}, score int) int {
	if nil == v {
		return 1
	}

	cur := s.head
	update := [MAX_LEVEL]*skipListNode{}
	i := MAX_LEVEL -1
	for ; i >=0 ;i-- {
		for nil != cur.forwards[i] {
			if cur.forwards[i].v == v {
				return 2
			}
			if cur.forwards[i].score > score {
				update[i] = cur
				break
			}
			cur = cur.forwards[i]
		}
		//若第一次插入 update 为heap update在[0-max-1] 都是相同的值
		if nil == cur.forwards[i] {
			update[i] = cur
		}
	}

	level := 1
	for i := 1; i < MAX_LEVEL; i++ {
		if rand.Int31() % 7 == 1 {
			level++
		}
	}

	newNode := newSkipListNode(v,score,level)

	for i:= 0; i <= level - 1;i++ {
		next := update[i].forwards[i]
		update[i].forwards[i] = newNode
		newNode.forwards[i] = next
	}

	//若果当前节点的层数大于之前跳表的层数 更新当前跳表层数
	if level > s.level {
		s.level = level
	}

	s.length++

	return 0


}

func (s *SkipList) Find(v interface{}, score int) *skipListNode {
	if v == nil || s.length == 0 {
		return nil
	}

	cur := s.head
	for i := s.level -1 ; i >= 0 ;i -- {
		for nil != cur.forwards[i] {
			if cur.forwards[i].score == score && cur.forwards[i].v == v {
				return cur.forwards[i]
			} else if cur.forwards[i].score > score {
				break
			}
			cur = cur.forwards[i]
		}
	}

	return nil

}

func (s *SkipList) Delete(v interface{}, score int) int {
	if nil == v {
		return 1
	}

	//查找前驱节点
	cur := s.head
	//记录前驱路径
	update := [MAX_LEVEL]*skipListNode{}
	for i := s.level - 1; i >= 0; i-- {
		update[i] = s.head
		for nil != cur.forwards[i] {
			if cur.forwards[i].score == score && cur.forwards[i].v == v {
				update[i] = cur
				break
			}
			cur = cur.forwards[i]
		}
	}

	cur = update[0].forwards[0]
	for i := cur.level - 1; i >= 0; i-- {
		if update[i] == s.head && cur.forwards[i] == nil {
			s.level = i
		}

		if nil == update[i].forwards[i] {
			update[i].forwards[i] = nil
		} else {
			update[i].forwards[i] = update[i].forwards[i].forwards[i]
		}
	}

	s.length--

	return 0
}

func (s *SkipList) String() string {
	return fmt.Sprintf("level:%+v, length:%+v", s.level, s.length)
}