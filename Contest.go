package main

func stoneGameIII(stoneValue []int) string {
	n := len(stoneValue)
	dp := make([]int, n)
	res := helper(stoneValue, dp, 0)
	if res > 0 {
		return "Alice"
	}
	if res == 0 {
		return "Tie"
	}
	return "Bob"
}

func helper(stoneValue []int, dp []int, index int) int {
	if index >= len(stoneValue) {
		return 0
	}
	if dp[index] != 0 {
		return dp[index]
	}
	res := stoneValue[index] - helper(stoneValue, dp, index+1)
	if index+1 < len(stoneValue) {
		res = max(res, stoneValue[index]+stoneValue[index+1]-helper(stoneValue, dp, index+2))
	}
	if index+2 < len(stoneValue) {
		res = max(res, stoneValue[index]+stoneValue[index+1]+stoneValue[index+2]-helper(stoneValue, dp, index+3))
	}
	dp[index] = res
	return res
}

func main() {
	stoneGameIII([]int{1, 2, 3, 7})
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

func twodarray(m, n int) [][]int {
	res := make([][]int, m)
	for i := range res {
		res[i] = make([]int, n)
	}
	return res
}
