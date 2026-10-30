package main

import (
	"bufio"
	"os"
	"slices"
	"strconv"
	"time"
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

// randseed
var _seed = uint64(time.Now().UnixNano())

func rng() uint64 {
	_seed += 0x9e3779b97f4a7c15
	z := _seed
	z = (z ^ (z >> 30)) * 0xbf58476d1ce4e5b9
	z = (z ^ (z >> 27)) * 0x94d049bb133111eb
	return z ^ (z >> 31)
}

/* 核心结论：b[i]=a[i]-i

a[i+1]=a[i]+1
设 b[i]=a[i]-i
b[i+1]=a[i+1]-(i+1)=(a[i]+1)-(i+1)=a[i]-i=b[i]

a[i]=a[j]-(j-i)
设 b[i]=a[i]-i
b[i]=a[i]-i=a[j]-(j-i)-i=a[j]-j=b[j]

e.g.

5 5 5 9 8 a[i]
5 4 3 6 4 a[i]-i

6 6 6 6 8 a[i]
6 5 4 3 4 a[i]-i

设b[i]=a[i]-i，原题转换为排序题
*/

func main() {
	defer writer.Flush()

	// logic
	t := NextInt()
	for range t {
		n := NextInt()
		b := make([]int, n)
		for i := range n {
			a := NextInt()
			b[i] = a - i
		}
		slices.Sort(b)

		res := 1
		count := 1
		pre := b[0]
		for i := 1; i < n; i++ {
			if b[i] == pre {
				continue
			} else if b[i] == pre+1 {
				count++
				pre = b[i]
			} else {
				res = max(res, count)
				count = 1
				pre = b[i]
			}
		}
		res = max(res, count)
		PrintIntLn(res)
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
