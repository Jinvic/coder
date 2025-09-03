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
	scanner.Split(bufio.ScanWords)

}

/*
- a[i]+a[j]>b[i]+b[j]可转换为a[i]-b[i]+a[j]-b[j]>0
- 设c[i]=a[i]-b[i]，则题目转换为求ij对满足c[i]+c[j]>0
- 排序c数组后，对每个元素二分查找满足条件的元素数量，时间复杂度nlogn。

似乎双指针会比二分更快一些。
*/

func main() {
	defer writer.Flush()

	// logic
	n := NextInt()
	a := make([]int, n)
	for i := range n {
		a[i] = NextInt()
	}
	for i := range n {
		a[i] -= NextInt()
	}

	sort.Ints(a)

	res := 0
	for i := range n {
		num := sort.Search(n, func(j int) bool {
			return j > i && a[i]+a[j] > 0
		})
		res += n - num

	}
	PrintIntLn(res)
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
