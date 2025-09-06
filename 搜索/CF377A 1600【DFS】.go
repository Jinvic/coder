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

func main() {
	defer writer.Flush()

	// logic
	n, m, k := NextInt(), NextInt(), NextInt()
	grid := make([][]rune, n)
	for i := range n {
		s := NextString()
		grid[i] = []rune(s)
	}

	spaceNum := 0
	x, y := -1, -1
	for i := range n {
		for j := range m {
			if grid[i][j] == '.' {
				spaceNum++
				x, y = i, j
			}
		}
	}

	valid := func(x, y int) bool {
		return x >= 0 && x < n && y >= 0 && y < m
	}

	isSpace := func(x, y int) bool {
		return grid[x][y] == '.'
	}

	spaceNum -= k
	mark := make([][]bool, n)
	for i := range mark {
		mark[i] = make([]bool, m)
	}
	dx := []int{-1, 0, 1, 0}
	dy := []int{0, 1, 0, -1}

	var dfs func(x, y, cnt int) int
	dfs = func(x, y, cnt int) int {
		if cnt >= spaceNum {
			return cnt
		}

		cnt++
		mark[x][y] = true
		if cnt >= spaceNum {
			return cnt
		}

		for i := range 4 {
			nx, ny := x+dx[i], y+dy[i]
			if valid(nx, ny) && // 坐标合法
				isSpace(nx, ny) && // 是空坐标
				!mark[nx][ny] { // 未被标记过
				cnt = dfs(nx, ny, cnt)
				if cnt >= spaceNum {
					break
				}
			}
		}
		return cnt
	}

	dfs(x, y, 0)

	for i := range n {
		for j := range m {
			if isSpace(i, j) && !mark[i][j] {
				grid[i][j] = 'X'
			}
		}
		PrintStringLn(string(grid[i]))
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
