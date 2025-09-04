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

// 全局图
var (
	n, m  int
	a     []int   // 节点值
	graph [][]int // 无权图，只存 to
)

func main() {
	defer writer.Flush()

	// logic
	n, m = NextInt(), NextInt()
	a = make([]int, n+1)
	graph = make([][]int, n+1)

	for i := 1; i <= n; i++ {
		a[i] = NextInt()
	}

	for i := 0; i < n-1; i++ {
		u, v := NextInt(), NextInt()
		graph[u] = append(graph[u], v)
		graph[v] = append(graph[v], u)
	}

	// DFS
	var dfs func(father, u, sum int) int
	dfs = func(father, u, sum int) int {
		if a[u] == 1 {
			sum++
		} else {
			sum = 0
		}
		if sum > m {
			return 0
		}
		if len(graph[u]) == 1 && u != 1 {
			return 1
		}
		res := 0
		for _, v := range graph[u] {
			if v == father {
				continue
			}
			res += dfs(u, v, sum)
		}
		return res
	}

	PrintIntLn(dfs(0, 1, 0))
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
