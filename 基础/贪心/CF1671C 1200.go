package main

import (
	"bufio"
	"os"
	"sort"
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
初步思路：排序a后算前缀和sum[n]，然后二分找到sum[j]是当天能买到的最大值。
之后只考虑sum[0:j]这个区间，每天涨价sum[i]+=i，再二分找到新的sum[j]
累加每天的结果直到退出循环
误区：逐天模拟。x很大a很小时很多天结果一样。
优化思路：每日上涨幅度一致，考虑买j包最多买多少天。
第一步不变，排序a后算前缀和sum[n]，然后二分找到sum[j]是当天能买到的最大值。
sum[j]每天涨j，d=(x-sum[j])/j res+=d*j totald+=d
递减j，直到sum[j]+totald*j<x，重复如上步骤累加res和d
*/

func main() {
	defer writer.Flush()

	// logic
	t := NextInt()
	for range t {
		n := NextInt()
		x := NextInt64()
		a := make([]int64, n)
		for i := range n {
			a[i] = NextInt64()
		}
		sort.Slice(a, func(i, j int) bool { return a[i] < a[j] })
		sum := make([]int64, n)
		sum[0] = a[0]
		for i := 1; i < n; i++ {
			sum[i] = sum[i-1] + a[i]
		}

		res := int64(0)
		totalDay := int64(0)
		// 二分确定起点，或者直接从n开始也行
		num := int64(sort.Search(len(sum), func(i int) bool { return sum[i] > x }))
		for num > 0 {
			// 第 totalDay 天能买 num 包
			todaySum := sum[num-1] + totalDay*num
			if todaySum <= x {
				// 买num包能买多少天
				d := 1 + (x-todaySum)/num
				res += d * num
				totalDay += d
			}
			num--
		}
		PrintInt64Ln(res)
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
