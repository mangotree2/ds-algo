package lengthoflongestsubstring

//给定一个字符串，请你找出其中不含有重复字符的 最长子串 的长度。
//
//示例 1:
//
//输入: "abcabcbb"
//输出: 3
//解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
//示例 2:
//
//输入: "bbbbb"
//输出: 1
//解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。
//示例 3:
//
//输入: "pwwkew"
//输出: 3
//解释: 因为无重复字符的最长子串是 "wke"，所以其长度为 3。
//     请注意，你的答案必须是 子串 的长度，"pwke" 是一个子序列，不是子串。










func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

//滑动窗口解法
func lengthOfLongestSubstring(s string) int {
	left ,res := 0,0
	//上次字母出现的位置 如果是纯英文则用数组，其他用map
	lastHash := [256]int{}


	for i,v :=range s {
		if lastHash[v] == 0 || lastHash[v] < left {
			res = max(res,i-left+1)
		} else {//左缩
			left = lastHash[v]
		}

		//记录字符最后出现位置
		lastHash[v] = i+1
	}

	return res
}


