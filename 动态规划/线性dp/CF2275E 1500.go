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
核心结论：一旦开始走斜边，后续路径固定，即XXX...XXI，可算后缀和。
问题收缩至ai走直边bi还是斜边bi+1。
若走bi，下一步只能走ai+1，ai+1面临相同决策，可递推。
若走bi+1，如上讨论，路径固定，通过后缀和直接得到结果。
*/

func main() {
	defer writer.Flush()

	// logic
	t := NextInt()

	for range t {
		n := NextInt()
		a := make([]int64, n)
		b := make([]int64, n)
		for i := range n {
			a[i] = NextInt64()
		}
		for i := range n {
			b[i] = NextInt64()
		}

		dis := func(i int, j int) int64 {
			if a[i] == b[j] {
				return 2
			} else {
				return 1
			}
		}

		sufsum := make([]int64, n)
		sufsum[n-1] = dis(n-1, n-1)
		for i := n - 2; i >= 0; i-- {
			sufsum[i] = sufsum[i+1] + dis(i, i+1) + dis(i+1, i)
		}

		res := sufsum[0]
		// 只走直边时前i位的路径和
		dp := make([]int64, n)
		dp[0] = dis(0, 0)
		for i := 1; i < n; i++ {
			// 走斜边，结果为前i为走直边+b[i-1]到a[i]的边+后缀和
			res = max(res, dp[i-1]+dis(i, i-1)+sufsum[i])
			// 走直边，继续dp
			dp[i] = dp[i-1] + dis(i, i-1) + dis(i, i)
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
