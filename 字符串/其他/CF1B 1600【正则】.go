package main

import (
	"bufio"
	"os"
	"regexp"
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

	var sys1Pattern = regexp.MustCompile(`^[A-Z]+[0-9]+$`)
	var sys2Pattern = regexp.MustCompile(`^R[0-9]+C[0-9]+$`)

	for range n {
		s := NextString()

		if sys1Pattern.MatchString(s) {
			PrintStringLn(convertSys1ToSys2(s))
		} else if sys2Pattern.MatchString(s) {
			PrintStringLn(convertSys2ToSys1(s))
		}
	}
}

func convertSys1ToSys2(s string) string {
	var i int
	var char rune
	for i, char = range s {
		if unicode.IsDigit(char) {
			break
		}
	}

	res := "R" + s[i:] + "C"
	colVal := convertLetterToDigit(s[0:i])
	res += strconv.FormatInt(colVal, 10)

	return res
}
func convertSys2ToSys1(s string) string {
	var i int
	var char rune
	for i, char = range s {
		if char == 'C' {
			break
		}
	}

	colVal, _ := strconv.ParseInt(s[i+1:], 10, 64)
	res := convertDigitToLetter(colVal)
	res += s[1:i]
	return res
}

func convertLetterToDigit(s string) int64 {
	res := int64(0)
	for _, char := range s {
		res *= 26
		res += int64(char - 'A' + 1)
	}
	return res
}
func convertDigitToLetter(x int64) string {
	res := []rune{}
	// res := ""
	for x > 0 {
		y := x % 26
		x /= 26
		if y == 0 {
			res = append(res, 'Z')
			// res += "Z"
			x--
		} else {
			res = append(res, rune('A'+y-1))
		}
	}

	slices.Reverse(res)
	return string(res)
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
