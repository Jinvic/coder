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

func main() {
	defer writer.Flush()

	// logic

	tt := NextInt()
	for range tt {
		n, k := NextInt(), NextInt()
		s := make([]int, n)
		t := make([]int, n)
		for i := range n {
			x := NextInt() % k
			s[i] = min(x, k-x)
		}
		for i := range n {
			x := NextInt() % k
			t[i] = min(x, k-x)
		}

		sort.Ints(s)
		sort.Ints(t)

		res := true
		for i := 0; i < n; i++ {
			if s[i] != t[i] {
				res = false
				break
			}
		}

		if res {
			PrintStringLn("YES")
		} else {
			PrintStringLn("NO")
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
