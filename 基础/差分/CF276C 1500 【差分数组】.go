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

    // 排序：从大到小
    sort.Slice(a, func(i, j int) bool {
        return a[i] > a[j]
    })

    // 差分数组统计覆盖次数
    diff := make([]int, n+1)
    for i := 0; i < q; i++ {
        l, r := NextInt(), NextInt()
        diff[l-1]++
        diff[r]--
    }

    freq := make([]int, n)
    freq[0] = diff[0]
    for i := 1; i < n; i++ {
        freq[i] = freq[i-1] + diff[i]
    }

    // 频率从大到小排序
    sort.Sort(sort.Reverse(sort.IntSlice(freq)))

    // 贪心配对：高频率位置放最大值
    var sum int64
    for i := 0; i < n; i++ {
        sum += int64(freq[i]) * a[i]
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
