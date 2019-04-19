package _0_1bag

import "math"

//int maxW = Integer.MIN_VALUE; //存储背包中物品总重量的最大值
// cw表示当前已经装进去的物品的重量和；i表示考察到哪个物品了；
// w背包重量；items表示每个物品的重量；n表示物品个数
// 假设背包可承受重量100，物品个数10，物品重量存储在数组a中，那可以这样调用函数：
// f(0, 0, a, 10, 100)

















var maxW = math.MinInt64

func Find(i, cw int, item []int, n, w int) {
	if i == n || cw == w {
		if maxW < cw {
			maxW = cw
		}
		return
	}

	Find(i+1,cw,item,n,w)

	if cw + item[i] <=w {
		Find(i+1,cw+item[i],item,n,w)
	}


}