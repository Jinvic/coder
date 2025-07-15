# https://codeforces.com/contest/1722/problem/E

# 理解题意就是个基础的二维前缀和，不知道怎么能有1600

package main
 
import (
	"fmt"
)
 
func main() {
	var t int
	fmt.Scan(&t)
	for range t {
		var n int
		var q int
		fmt.Scan(&n, &q)
 
		area := [1001][1001]int64{}
		var h, w int64
		for i := 0; i < n; i++ {
			fmt.Scan(&h, &w)
			area[h][w] += h * w
		}
 
		// 前缀和
		for i := 1; i <= 1000; i++ {
			for j := 1; j <= 1000; j++ {
				area[i][j] += area[i][j-1] + area[i-1][j] - area[i-1][j-1]
			}
		}
 
		var hs, ws, hb, wb int64
		for range q {
			fmt.Scan(&hs, &ws, &hb, &wb)
 
			res := area[hb-1][wb-1] - area[hb-1][ws] - area[hs][wb-1] + area[hs][ws]
			fmt.Println(res)
		}
	}
}
