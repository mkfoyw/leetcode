package main

/*
 * @lc app=leetcode.cn id=131 lang=golang
 *
 * [131] 分割回文串
 */

// @lc code=start

// dp[i][]string: 表示以字符s[i]结尾的字符串有的分割方案

func partition(s string) [][]string {
	n := len(s)
	if n == 0 {
		return [][]string{}
	}

	//flag[i][j]:表示字符串s[i]到s[j]是否是回文串
	flag := make([][]int8, 0)
	for i := 0; i < n; i++ {
		flag = append(flag, make([]int8, n))
	}

	// 1 表示是回文串， -1 表示不是回文串， 0 表示未搜索
	var isPalindrome func(i, j int) int8
	isPalindrome = func(i, j int) int8 {
		if i >= j {
			return 1
		}
		if flag[i][j] != 0 {
			return flag[i][j]
		}

		flag[i][j] = -1
		if s[i] == s[j] {
			flag[i][j] = isPalindrome(i+1, j-1)
		}
		return flag[i][j]
	}

	dp := make([][][]string, 0)
	//dp[i]:前 i 个字符串可能的分割方案
	for i := 0; i < n; i++ {
		cur := [][]string{}
		for j := 0; j <= i; j++ {
			if isPalindrome(j, i) == 1 {
				if j == 0 {
					cur = append(cur, append([]string(nil), s[:i+1]))
				} else {
					for _, item := range dp[j-1] {
						cur = append(cur, append([]string(nil), item...))
						cur[len(cur)-1] = append(cur[len(cur)-1], s[j:i+1])
					}
				}
			}

		}
		dp = append(dp, cur)
	}
	return dp[n-1]
}

// Accepted
// 32/32 cases passed (284 ms)
// Your runtime beats 62.66 % of golang submissions
// Your memory usage beats 5.07 % of golang submissions (41.7 MB)

//official 代码

func partition2(s string) (ans [][]string) {
	n := len(s)

	flag := make([][]int8, 0)
	for i := 0; i < n; i++ {
		flag = append(flag, make([]int8, n))
	}
	var isPalindrome func(int, int) int8
	isPalindrome = func(i, j int) int8 {
		if i >= j {
			return 1
		}

		if flag[i][j] != 0 {
			return flag[i][j]
		}

		flag[i][j] = -1
		if s[i] == s[j] {
			flag[i][j] = isPalindrome(i+1, j-1)
		}
		return flag[i][j]
	}

	var splits []string
	var dfs func(int)

	dfs = func(i int) {
		if i == n {
			ans = append(ans, splits)
			return
		}

		for j := 0; j <= n; j++ {
			if isPalindrome(j, i) == 1 {
				splits = append(splits, s[j:i+1])
				dfs(i + 1)
				splits = splits[:len(splits)+1]
			}
		}
	}

	dfs(0)
	return
}

// @lc code=end
