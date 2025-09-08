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
 
func main() {
	defer writer.Flush()
 
	// logic
	a := make([]int64, 1e5+5)
	f := func(k, l, r int64) int64 {
		res := int64(0)
 
		for i := l; i <= r; i++ {
			for k%a[i] == 0 {
				k /= a[i]
			}
			res += k
 
			if k == 1 {
				res += r - i
				break
			}
 
		}
 
		return res
	}
 
	t := NextInt()
	for range t {
		n, q := NextInt(), NextInt()
		for i := range n {
			a[i] = NextInt64()
		}
		for range q {
			k, l, r := NextInt64(), NextInt64(), NextInt64()
			PrintInt64Ln(f(k, l-1, r-1))
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
