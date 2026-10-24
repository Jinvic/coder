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

func main() {
	defer writer.Flush()

	// logic
	t := NextInt()
	for range t {
		n := NextInt()
		k := NextInt64()
		as := make([]int64, n)
		bs := make([]int64, n)
		cs := make([]int64, n)
		sums := make([]int64, n)

		var l, r int64
		for i := range n {
			as[i] = NextInt64()
			bs[i] = NextInt64()
			cs[i] = NextInt64()

			sums[i] = as[i] + bs[i] + cs[i]
			if i == 0 {
				l = sums[i]
				r = sums[i]
			} else {
				if sums[i] < l {
					l = sums[i]
				}
				if sums[i] > r {
					r = sums[i]
				}
			}
		}

		// 实验室i最少需要几步满足s，-1标识不可行
		mincost := func(s int64, i int) int64 {
			a, b, c := as[i], bs[i], cs[i]
			sum := sums[i]

			if sum >= s {
				return 0
			}

			if sum+k < s {
				return -1
			}

			// abc存在递减，总和可以一直递增
			if a > b || a > c || b > c {
				return s - sum
			}

			// 总和不变
			if a == b && a == c && b == c {
				if sum < s {
					return -1
				} else {
					return 0
				}
			}

			// 单调不递减，先算diff造一个递减出来，之后可以一直递增
			diff := min(b-a, c-b)
			if sum+k-2*(diff+1) < s {
				return -1
			} else {
				return s - sum + 2*(diff+1)
			}
		}

		r += k
		res := l
		for l <= r {
			s := l + ((r - l) >> 1)
			available := true
			sumcost := int64(0)
			for i := range n {
				cost := mincost(s, i)
				if cost < 0 {
					available = false
					break
				} else {
					sumcost += cost
					if sumcost > k {
						available = false
						break
					}
				}
			}

			if available {
				res = s
				l = s + 1
			} else {
				r = s - 1
			}
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
