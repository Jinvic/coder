package main

import (
	"bufio"
	"os"
	"slices"
	"strconv"
	"unicode"
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
	n := NextInt()
	for range n {
		s := []rune(NextString())
		rowStr, colStr := []rune{}, []rune{}

		const (
			InitMode    = iota // 处理第一个字符
			JudgingMode        // 正在判断编号系统，先按照Sys2RowMode处理，若再次出现字母则确认为Sys2，否则转换为Sys1
			Sys1ColMode        // 读取BC23的BC
			Sys1RowMode        // 读取BC23的23
			Sys2ColMode        // 读取R23C55的C55
			Sys2RowMode        // 读取R23C55的R23
		)

		mode := InitMode
		for i, char := range s {
			switch mode {
			case InitMode:
				if unicode.IsLetter(char) {
					if char == 'R' {
						mode = JudgingMode // 无法确认编号系统，暂时按照Sys2处理
						rowStr = append(rowStr, 'R')
					} else {
						mode = Sys1ColMode // 非R字母开头，确认为Sys1
						colStr = append(colStr, char)
					}
				}
			case JudgingMode:
				if unicode.IsDigit(char) {
					rowStr = append(rowStr, char) // 数字直接添加，相当于Sys2RowMode
				} else {
					if char == 'C' { // 再次出现C时
						if unicode.IsDigit(s[i-1]) { // RC之间有数字，说明是Sys2
							mode = Sys2ColMode
							colStr = append(colStr, 'C')
						} else { // RC之间没有数字，说明是Sys1
							mode = Sys1ColMode
							colStr = rowStr
							rowStr = []rune{}
							colStr = append(colStr, 'C')
						}
					} else { // 出现非C字母，说明是sys1
						mode = Sys1ColMode
						colStr = rowStr
						rowStr = []rune{}
						colStr = append(colStr, char)
					}
				}
			case Sys1ColMode:
				if unicode.IsLetter(char) { // 字母，说明仍是列
					colStr = append(colStr, char)
				} else { // 数字，说明切换到行
					mode = Sys1RowMode
					rowStr = append(rowStr, char)
				}
			case Sys1RowMode:
				rowStr = append(rowStr, char)
			case Sys2RowMode:
				// 实际在JudgingMode中已处理
				continue
			case Sys2ColMode:
				colStr = append(colStr, char)
			}
		}

		// 没有再出现字母，说明是sys1
		if mode == JudgingMode {
			mode = Sys1RowMode
			colStr = []rune{'R'}
			n := len(rowStr)
			rowStr = rowStr[1:n]
		}

		toInt := func(r rune) int64 {
			return int64(r-'A') + 1
		}

		pow := func(x int64, pow int) int64 {
			res := int64(1)
			for range pow {
				res *= x
			}
			return res
		}

		if mode == Sys1RowMode { // 处理sys1
			PrintString("R")
			PrintString(string(rowStr))

			res := int64(0)
			n := len(colStr)
			for i, char := range colStr {
				res += pow(26, n-i-1) * toInt(char)
			}
			PrintString("C")
			PrintInt64Ln(res)
		} else { // 处理sys2
			colVal, _ := strconv.ParseInt(string(colStr[1:]), 10, 64)

			colStr = []rune{}
			for colVal > 0 {
				v := colVal % 26
				colVal /= 26
				if v == 0 { // 进位特判
					colStr = append(colStr, 'Z')
					colVal--
				} else {
					char := rune('A' + v - 1)
					colStr = append(colStr, char)
				}
			}
			slices.Reverse(colStr)

			PrintString(string(colStr))
			PrintStringLn(string(rowStr[1:]))
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
