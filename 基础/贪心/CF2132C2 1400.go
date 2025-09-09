package main

import (
	"bufio"
	"os"
	"strconv"
)

// var reader *bufio.Reader
var writer *bufio.Writer
var scanner *bufio.Scanner

func init() {
	in := os.Stdin
	out := os.Stdout

	// go run . -t
	for _, arg := range os.Args {
		if arg == "-t" ||
			arg == "--test" {
			in, _ = os.Open("input.txt")
			out, _ = os.Create("output.txt")
		}
	}

	// reader = bufio.NewReader(in)
	writer = bufio.NewWriter(out)
	scanner = bufio.NewScanner(in)

	// 扩大缓冲区
	const MaxTokenLength = 300000
	scanner.Buffer(make([]byte, MaxTokenLength), MaxTokenLength)

	// 分词读入
	scanner.Split(bufio.ScanWords)
}

/*
- 首先打表单次交易购买1,3,9...3^x个西瓜的花费。
- 考虑交易次数最小的情况，则尽可能一次交易购买更多西瓜。逆序遍历打表的数据，统计不同交易的数量和总花费。
- 要使花费最低，则经可能将一次购买3^x个西瓜拆分成三次购买3^(x-1)次西瓜。每次拆分将使交易数+2.
- 逆序遍历交易次数最小的情况，将大单尽可能拆分为小单，直到交易次数达到上限.
*/

func main() {
	defer writer.Flush()

	// logic

	// 打表，购买num[i]=3^i个西瓜将花费cost[i]
	costs := make([]int64, 0)
	nums := make([]int64, 0)
	costs = append(costs, 3)
	nums = append(nums, 1)

	var x, val1, val2 int64
	x, val1, val2 = 1, 9, 1
	cost := val1 + x*val2
	for val2 <= 1e9 {
		costs = append(costs, cost)
		nums = append(nums, val2*3)
		val1 *= 3
		val2 *= 3
		x++
		cost = val1 + x*val2
	}

	t := NextInt()
	for range t {
		n, k := NextInt64(), NextInt64()
		cost := int64(0)
		deal := make([]int64, len(nums))
		numSum := int64(0)
		for i := len(nums) - 1; i >= 0; i-- {
			cnt := n / nums[i]
			if cnt > 0 {
				n %= nums[i]
				cost += costs[i] * cnt

				deal[i] = cnt
				numSum += cnt
			}
		}
		if numSum > k {
			PrintInt64Ln(-1)
		} else {
			rest := k - numSum
			for i := len(deal) - 1; i > 0; i-- {
				if rest < 2 {
					break
				}

				if deal[i] > 0 {
					cnt := min(rest/2, deal[i])
					rest -= 2 * cnt
					cost = cost - cnt*costs[i] + 3*cnt*costs[i-1]
					deal[i-1] += 3 * cnt
				}
			}
			PrintInt64Ln(cost)
		}
	}
}

func NextString() string {
	scanner.Scan()
	return scanner.Text()
}

func NextInt() int {
	scanner.Scan()
	x, _ := strconv.Atoi(scanner.Text())
	return x
}

func NextInt64() int64 {
	scanner.Scan()
	x, _ := strconv.ParseInt(scanner.Text(), 10, 64)
	return x
}

func NextFloat64() float64 {
	scanner.Scan()
	x, _ := strconv.ParseFloat(scanner.Text(), 64)
	return x
}

func PrintInt(n int) {
	writer.WriteString(strconv.Itoa(n))
}

func PrintInt64(n int64) {
	writer.WriteString(strconv.FormatInt(n, 10))
}

func PrintFloat64(f float64, prec int) {
	writer.WriteString(strconv.FormatFloat(f, 'f', prec, 64))
}

func PrintString(s string) {
	writer.WriteString(s)
}

func PrintIntLn(n int) {
	PrintInt(n)
	writer.WriteByte('\n')
}

func PrintInt64Ln(n int64) {
	PrintInt64(n)
	writer.WriteByte('\n')
}

func PrintFloat64Ln(f float64, prec int) {
	PrintFloat64(f, prec)
	writer.WriteByte('\n')
}

func PrintStringLn(s string) {
	PrintString(s)
	writer.WriteByte('\n')
}

func PrintIntSlice(ns []int, l, r int) {
	end := r - 1
	for i, n := range ns {
		PrintInt(n)
		if i != end {
			writer.WriteByte(' ')
		} else {
			writer.WriteByte('\n')
		}
	}
}

func PrintInt64Slice(ns []int64, l, r int) {
	end := r - 1
	for i, n := range ns {
		PrintInt64(n)
		if i != end {
			writer.WriteByte(' ')
		} else {
			writer.WriteByte('\n')
		}
	}
}

func PrintFloat64Slice(fs []float64, prec, l, r int) {
	end := r - 1
	for i, f := range fs {
		PrintFloat64(f, prec)
		if i != end {
			writer.WriteByte(' ')
		} else {
			writer.WriteByte('\n')
		}
	}
}

func PrintStringSlice(ss []string, l, r int) {
	end := r - 1
	for i, s := range ss {
		PrintString(s)
		if i != end {
			writer.WriteByte(' ')
		} else {
			writer.WriteByte('\n')
		}
	}
}
