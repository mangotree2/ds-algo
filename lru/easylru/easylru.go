package easylru

import "container/list"

//LRUCache cache = new LRUCache( 2 /* 缓存容量 */ );
//
//cache.put(1, 1);
//cache.put(2, 2);
//cache.get(1);       // 返回  1
//cache.put(3, 3);    // 该操作会使得密钥 2 作废
//cache.get(2);       // 返回 -1 (未找到)
//cache.put(4, 4);    // 该操作会使得密钥 1 作废
//cache.get(1);       // 返回 -1 (未找到)
//cache.get(3);       // 返回  3
//cache.get(4);       // 返回  4


//type LRUCache struct {
//    capacity int
//
//    ll *list.List
//    cache map[int]*list.Element
//}
//
//type e***y struct {
//    key int
//    value int
//}
//
//func Constructor(capacity int) LRUCache {
//    return LRUCache {
//        capacity: capacity,
//        ll: list.New(),
//        cache: make(map[int]*list.Element),
//    }
//}
//
//
//func (this *LRUCache) Get(key int) int {
//    elem, ok := this.cache[key]
//    if ok {
//        this.ll.MoveToFront(elem)
//        return elem.Value.(*e***y).value
//    }
//    return -1
//}
//
//
//func (this *LRUCache) Put(key int, value int)  {
//    elem, ok := this.cache[key]
//    if ok {
//        this.ll.MoveToFront(elem)
//        elem.Value.(*e***y).value = value
//        return
//    }
//
//    if this.capacity > 0 && this.ll.Len() >= this.capacity {
//        last := this.ll.Back()
//        this.ll.Remove(last)
//        delete(this.cache, last.Value.(*e***y).key)
//    }
//    this.cache[key] = this.ll.PushFront(&e***y{key, value})
//}

type LRUCache struct {
	cap int
	l list.List
	cache map[int]*list.Element

}

type entry struct {
	key int
	value int
}

func (c *LRUCache) Put(key int, value int) {
	if v, ok := c.cache[key]; ok {
		c.l.MoveToFront(v)
		v.Value.(*entry).value = value
		return
	}

	if c.l.Len() > 0 && c.l.Len() >= c.cap {
		last := c.l.Back()
		c.l.Remove(last)
		delete(c.cache,last.Value.(*entry).key)

	}

	c.cache[key] = c.l.PushFront(&entry{key,value})

}

func (c *LRUCache) Get(key int) int {

	if v, ok := c.cache[key];ok {
		c.l.MoveToFront(v)
		return v.Value.(*entry).value
	}

	return -1

}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		cap:   capacity,
		l:     list.List{},
		cache: make(map[int]*list.Element,capacity),
	}
}