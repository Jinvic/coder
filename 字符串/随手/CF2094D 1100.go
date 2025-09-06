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

type Data struct {
	Start string
	Arr   []int
}

func InitData(s string) Data {
	start := s[:1]

	cnt := 1
	arr := []int{}
	for i := 1; i < len(s); i++ {
		if s[i] == s[i-1] {
			cnt++
		} else {
			arr = append(arr, cnt)
			cnt = 1
		}
	}
	arr = append(arr, cnt)

	return Data{
		Start: start,
		Arr:   arr,
	}
}

func CmpData(a, b Data) bool {
	if a.Start != b.Start {
		return false
	}

	if len(a.Arr) != len(b.Arr) {
		return false
	}

	n := len(a.Arr)
	for i := range n {
		if b.Arr[i] < a.Arr[i] ||
			b.Arr[i] > 2*a.Arr[i] {
			return false
		}
	}

	return true
}

func main() {
	defer writer.Flush()

	// logic
	t := NextInt()
	for range t {
		p, s := NextString(), NextString()
		dp, ds := InitData(p), InitData(s)
		if CmpData(dp, ds) {
			PrintStringLn("YES")
		} else {
			PrintStringLn("NO")
		}
	}

	// dp
	// const MAX int = 2e5
	// for i := range MAX {
	// 	if i%2 == 0 {
	// 		PrintString("L")
	// 	} else {
	// 		PrintString("R")
	// 	}
	// }

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
