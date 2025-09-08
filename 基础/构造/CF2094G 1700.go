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
- 维护sum,sum2和prefix三个值。其中sum表示题目要求的特殊累加和，sum2为反向计算的特殊累加和。prefix为a1到a(n-1)的正常累加和。
- 当执行操作1和操作3时，计算新的sum，sum2和prefix。
- 当执行操作2时，根据反转次数输出sum或sum2
*/

func main() {
	defer writer.Flush()

	// logic
	t := NextInt()
	const MAX = 4e5 + 5
	a := make([]int64, MAX)
	for range t {
		var sum, sum2, total int64
		var head, tail int // 数组首尾指针
		len := int64(0)    // 数组长度
		revCount := 0      // 反转次数
		first := true

		q := NextInt()
		for range q {
			s := NextInt()
			switch s {
			case 1:
				if len == 1 {
					PrintInt64Ln(sum)
					continue
				}

				if revCount%2 == 0 { // 正向
					sum += total - a[tail]*len
					sum2 += a[tail]*len - total
					head--
					a[head] = a[tail]
					tail--
					PrintInt64Ln(sum)
				} else { // 反向
					sum += a[head]*len - total
					sum2 += total - a[head]*len
					tail++
					a[tail] = a[head]
					head++
					PrintInt64Ln(sum2)
				}

			case 2:
				revCount++
				if revCount%2 == 0 { // 正向
					PrintInt64Ln(sum)
				} else { // 反向
					PrintInt64Ln(sum2)
				}

			case 3:
				k := NextInt64()
				if first {
					// 从中间开始，方便向首尾添加元素
					head = 2e5
					tail = 2e5
					a[head] = k
					sum = k
					sum2 = k
					total = k
					len = 1
					first = false
					PrintInt64Ln(sum)
					continue
				}

				if revCount%2 == 0 { // 正向
					sum += k * (len + 1)
					sum2 += total + k
					total += k
					len++
					tail++
					a[tail] = k
					PrintInt64Ln(sum)
				} else { // 反向
					sum += total + k
					sum2 += k * (len + 1)
					total += k
					len++
					head--
					a[head] = k
					PrintInt64Ln(sum2)
				}

			}
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
