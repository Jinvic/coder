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

func main() {
	defer writer.Flush()

	// logic
	n, q := NextInt(), NextInt()
	a := make([]int64, n)
	for i := 0; i < n; i++ {
		a[i] = NextInt64()
	}

	// 从大到小排序
	sort.Slice(a, func(i, j int) bool {
		return a[i] > a[j]
	})

	type Data struct {
		lnum, rnum int
	}
	data := make([]Data, n)
	for i := 0; i < q; i++ {
		l, r := NextInt(), NextInt()
		data[l-1].lnum++
		data[r-1].rnum++
	}

	mp := make(map[int]int) // 查询次数 -> 位置数量
	num := 0
	for i := 0; i < n; i++ {
		lnum := data[i].lnum
		rnum := data[i].rnum

		num += lnum // 查询该位置
		if num > 0 {
			mp[num]++ // 统计查询num次的位置数量
		}
		num -= rnum // 不再查询后续位置
	}

	type Data2 struct {
		num   int // 查询次数
		count int // 位置数量
	}
	data2 := make([]Data2, 0, len(mp))
	for k, v := range mp {
		data2 = append(data2, Data2{k, v})
	}

	// 按照查询次数从大到小排序
	sort.Slice(data2, func(i, j int) bool {
		return data2[i].num > data2[j].num
	})

	sum := int64(0)
	pos := 0
	for _, v := range data2 {
		for i := 0; i < v.count; i++ {
			sum += int64(v.num) * a[pos]
			pos++
		}
	}

	PrintInt64Ln(sum)
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
