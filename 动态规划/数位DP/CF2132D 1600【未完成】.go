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
题目：
有一个无限长的序列：123456789101112131415...
对每个测试样例，题目将给出一个正整数k(1≤k≤1e15)，求这个序列的前k位数位的累加和。

思路：
预处理若干位的位数和以及累加和，从而只需要处理最后特定位数的若干个数字。
但对这最后一部分数字的处理不能暴力计算，需要数位dp或记忆化搜索。
*/

func main() {
	defer writer.Flush()

	// logic

	// 打表
	sum := []int64{1: 45} // 前i位数的总和
	num := []int64{1: 9}  // 前i位数的总位数
	tmp := int64(1)
	for i := int64(2); i <= 13; i++ {
		tmp *= 10
		num = append(num, num[i-1]+9*tmp*i)
		sum = append(sum, sum[1]*tmp*i)
	}
	// PrintInt64Slice(sum, 0, len(sum))
	// PrintInt64Slice(num, 0, len(num))

	t := NextInt()
	for range t {
		k := NextInt64()
		res := int64(0)
		var pos int64
		for pos = int64(len(num) - 1); pos >= 0; pos-- {
			if k >= num[pos] { // 位数大于前pos位的位数和
				res += sum[pos] // 结果加上前pos位的总和
				k -= num[pos]   // 剩下的数为若干个pos+1位的数
				break
			}
		}

		d := pos + 1           // 剩余数的位数
		count := k / (pos + 1) // 剩多少个d位的数
		rest := k % (pos + 1)  // 剩多少个不足d的位数

		start := int64(1)
		for range pos {
			start *= 10
		}

		// // 计算剩余pos+1位的数的数位和
		// for i := int64(0); i < count; i++ {
		// 	res += culSum(start + i)
		// }

		// 计算剩余d位的数的数位和
		res += digitSumRange(d, start, start+count-1)

		last := start + count // 最后一个数，取前rest位
		for rest > 0 {
			res += last / start
			last /= start
			start /= 10
			rest--
		}

		PrintInt64Ln(res)
	}
}

// // 计算某个数的位数和
// func culSum(x int64) int64 {
// 	sum := int64(0)
// 	for x/10 > 0 {
// 		sum += x % 10
// 		x /= 10
// 	}
// 	sum += x
// 	return sum
// }

// 统计有d位数的[l,r]的数位累加和
func digitSumRange(d, l, r int64) int64 {
	// TODO：数位dp或记忆化搜索
	return 0
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
