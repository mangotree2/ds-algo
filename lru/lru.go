package lru

//运用你所掌握的数据结构，设计和实现一个  LRU (最近最少使用) 缓存机制。它应该支持以下操作： 获取数据 get 和 写入数据 put 。
//
//获取数据 get(key) - 如果密钥 (key) 存在于缓存中，则获取密钥的值（总是正数），否则返回 -1。
//写入数据 put(key, value) - 如果密钥不存在，则写入其数据值。当缓存容量达到上限时，它应该在写入新数据之前删除最近最少使用的数据值，从而为新的数据值留出空间。
//
//进阶:
//
//你是否可以在 O(1) 时间复杂度内完成这两种操作？
//
//示例:
//
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



type LRUCache struct {
	Capacity	int
	Count		int
	Length		int
	Size		int
	Caches		[]*Cache
}

type Cache struct {
	Key			int
	Value		int
	Count		int
	Next		*Cache
}


func Constructor(capacity int) LRUCache {
	lc := LRUCache{
		Capacity:	capacity,
		Count:		0,
		Length:		0,
		Caches:		make([]*Cache, capacity / 10 + 1),
		Size:		capacity / 10 + 1,
	}
	return lc
}


func (this *LRUCache) Get(key int) int {
	cache := this.Caches[key % this.Size]
	result := -1
	var temp *Cache
	for cache != nil {
		if cache.Key == key {
			this.Count++
			cache.Count = this.Count
			result = cache.Value
			if cache.Next != nil {
				t := cache
				if temp == nil {//第一个
					this.Caches[key % this.Size] = cache.Next
				} else {//不是第一个则接到下一个
					temp.Next = cache.Next
				}
				for cache.Next != nil {//拿到最后一个key
					cache = cache.Next
				}
				//当前cache放到最后
				cache.Next = t
				t.Next = nil
			}
			break
		}
		//temp 上一个cache
		temp = cache
		cache = cache.Next
	}
	return result
}


func (this *LRUCache) Put(key int, value int)  {
	cache := this.Caches[key % this.Size]
	result := -1
	var temp *Cache
	for cache != nil {
		//key 已存在
		if cache.Key == key {
			this.Count++
			cache.Count = this.Count
			cache.Value = value
			result = value
			//放到最后
			if cache.Next != nil {
				t := cache
				if temp == nil {
					this.Caches[key % this.Size] = cache.Next
				} else {
					temp.Next = cache.Next
				}
				for cache.Next != nil {
					cache = cache.Next
				}
				cache.Next = t
				t.Next = nil
			}
			break
		}
		temp = cache
		cache = cache.Next
	}
	//存在返回
	if result != -1 {
		return
	}
	//不存在key ,但是有cache队列
	for temp != nil && temp.Next != nil {
		temp = temp.Next
	}
	//缓存已满
	if this.Length == this.Capacity {
		min := this.Count + 1
		pin := 0
		//找到最少访问的队列
		for i := 0; i < this.Size; i++ {
			if this.Caches[i] != nil && this.Caches[i].Count < min {
				min = this.Caches[i].Count
				pin = i
			}
		}
		//temp是和Key 一条队列的最后一个，如果满足这种情况 说明队列只有temp,
		if temp != nil && this.Caches[pin].Key == temp.Key {
			temp = nil
		}
		//删除最老最少使用的key
		this.Caches[pin] = this.Caches[pin].Next
	} else {
		this.Length++
	}
	this.Count++
	t := &Cache {
		Key:	key,
		Value:	value,
		Count:	this.Count,
	}
	//成为对应队列的第一个
	if temp == nil {
		this.Caches[key % this.Size] = t
	} else {
		temp.Next = t
	}
}

