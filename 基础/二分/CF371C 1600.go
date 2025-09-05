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
	scanner.Split(bufio.ScanWords)

}

type Data struct {
	Char   string
	ReqNum int64 // 需求数量
	OwnNum int64 // 拥有数量
	Price  int64 // 价格
}

func main() {
	defer writer.Flush()

	// logic
	str := NextString()
	n := len(str)

	var data = make([]Data, 3)
	data[0].Char = "B"
	data[1].Char = "S"
	data[2].Char = "C"

	for i := range 3 {
		data[i].OwnNum = NextInt64()
	}
	for i := range 3 {
		data[i].Price = NextInt64()
	}

	cost := NextInt64()

	for i := range n {
		switch str[i] {
		case 'B':
			data[0].ReqNum++
		case 'S':
			data[1].ReqNum++
		case 'C':
			data[2].ReqNum++
		}
	}

	// 二分：能构成x个必然可以构成x-1个，二分寻找最大值

	totalPrice := int64(0)
	for i := range 3 {
		totalPrice += data[i].ReqNum * data[i].Price
	}
	validate := func(x int64) bool {
		reqcost := totalPrice * x
		c := cost
		for i := range 3 {
			reqNum := x * data[i].ReqNum
			OwnNum := min(reqNum, data[i].OwnNum)
			OwnCost := OwnNum * data[i].Price
			c += OwnCost
		}

		return c >= reqcost
	}

	l, r := int64(0), cost+105
	for l <= r {
		mid := (l + r) / 2
		if validate(mid) {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}

	PrintInt64Ln(r)
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
