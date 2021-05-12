package main

import "fmt"

func main() {
	a := numTrees(5)
	fmt.Println(a)
}

func numTrees(n int) int {
	dp := make([]int, n+1)
	dp[0] = 1
	for x := 1; x <= n; x++ {
		for i := 1; i <= x; i++ {
			fmt.Printf("dp[%v] += %v * %v\n", x, dp[i-1], dp[x-i])
			// i represents the index of F(i,n)
			// G(n)=∑ G(i−1)⋅G(n−i)
			dp[x] += dp[i-1] * dp[x-i]
		}
	}
	fmt.Println(dp)
	return dp[n]
}
