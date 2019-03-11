package longestsequence


//我们有一个数字序列包含n个不同的数字，
// 如何求出这个序列中的最长递增子序列长度？
// 比如2, 9, 3, 6, 5, 1, 7这样一组数字序列，
// 它的最长递增子序列就是2, 3, 5, 7，所以最长递增子序列的长度是4。
//递推公式:
//a[0...i] 的最长子序列为: a[i] 之前所有比它小的元素中子序列长度最大的 + 1
func LongestSubSequence(a []int,n int) int {
	// 创建一个数组, 索引 i 对应考察元素的下标,
	// 存储 arr[0...i] 的最长上升子序列大小
	lengths := make([]int,n)

	//哨兵
	lengths[0] = 1

	for i := 1; i < n; i++ {
		max := 1
		for j := 0;j < i; j++{
			if a[i]>a[j] && lengths[j] >= max {
				max = lengths[j] + 1
			}
		}

		lengths[i] =max
	}

	return lengths[n-1]

}
