# https://codeforces.com/contest/1722/problem/G

# 按四个两两分组使其异或和为最大值，同时对多余的数特殊处理添加通解。
# 余2时需要六个数构建通解，注意特判

package main

import (
	"fmt"
)

const maxv = 1<<31 - 1
const offset = 16

var mode = "solve"

var adjust = [4][]int{
	{},
	{0},
	{0, 4, 1, 8, 3, 14},
	{1, 3, 2},
}

func main() {

	switch mode {
	case "solve":
		var t int
		fmt.Scan(&t)
		for range t {
			var n int
			fmt.Scan(&n)

			res := solve(n)
			// if !validate(res) {
			// 	fmt.Println("error", n)
			// 	break
			// }
			for _, v := range res {
				fmt.Printf("%d ", v)
			}
			fmt.Println()
		}
	case "validate":
		for i := 3; i <= 200000; i++ {
			fmt.Println("validating", i)
			res := solve(i)
			if !validate(res) {
				fmt.Println("error", i)
				break
			}
		}
	case "xor":
		for {
			var x, y int
			fmt.Scan(&x, &y)
			fmt.Println(x ^ y)
		}
	}

}

func solve(n int) []int {
	res := make([]int, 0, n)
	mod := n % 4
	pos := offset
	for i := 0; i < n/4; i++ {
		res = append(res, pos)
		res = append(res, pos+1)
		res = append(res, maxv-pos)
		res = append(res, maxv-pos-1)
		pos += 2
	}
	res = append(res, adjust[mod]...)
	if mod == 2 {
		res = res[4:]
	}
	return res
}

func validate(res []int) bool {
	var xor1, xor2 int
	for i := 0; i < len(res); i++ {
		if isOdd(i) {
			xor1 ^= res[i]
		} else {
			xor2 ^= res[i]
		}
	}
	if xor1 != xor2 {
		fmt.Println("xor1 != xor2", xor1, xor2, res)
		return false
	}
	return true
}

func isOdd(x int) bool {
	return x%2 == 1
}
