package main

import (
	"bufio"
	"os"
	"slices"
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

/* 核心结论：求x最大质因数maxp，使x/maxp最小
f(x)=0 x<=k
f(x)=1 x>k && x/maxp<=k
f(x)=1 + min(p*f(x/p)) x>k && x/maxp>k
*/

func main() {
	defer writer.Flush()

	// logic

	// SPF最小质因子筛
	const maxa = 200005
	spf := [maxa]int{}
	lpf := [maxa]int{}
	primes := []int{}
	for i := int(2); i < maxa; i++ {
		if spf[i] == 0 {
			spf[i] = i
			lpf[i] = i
			primes = append(primes, i)
		}
		for _, p := range primes {
			if p > spf[i] || p*i >= maxa {
				break
			}
			spf[i*p] = p
			lpf[i*p] = max(lpf[i], p)
		}
	}

	t := NextInt()
	for range t {
		n := NextInt()
		k := NextInt()
		a := make([]int, n)
		for i := range n {
			a[i] = NextInt()
		}
		slices.Sort(a)

		/*
			f(x)=0 x<=k
			f(x)=1 x>k && x/maxp<=k
			f(x)=1 + min(p*f(x/p)) x>k && x/maxp>k
		*/
		var f func(x int) int64
		fmp := make(map[int]int64) // 记忆结果
		f = func(x int) int64 {
			if x <= k {
				return 0
			}

			// 已经算过结果
			if fmp[x] != 0 {
				return fmp[x]
			}

			p := lpf[x]
			q := x / p
			if q <= k {
				fmp[x] = 1
			} else {
				fmp[x] = 1 + int64(p)*f(q)
				y := x
				for y > 1 {
					p := spf[y]
					q := x / p
					fmp[x] = min(fmp[x], 1+int64(p)*f(q))
					for y%p == 0 {
						y /= p
					}
				}
			}

			return fmp[x]
		}

		res := int64(0)
		for _, x := range a {
			res += f(x)
		}
		PrintInt64Ln(res)
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

func PrintIntSlice(ns []int, l, r int) {
	end := r - 1
	for i, n := range ns {
		PrintInt(n)
		if i != end {
			writer.WriteByte(' ')
		} else {
			writer.WriteByte('\n')
		}
	}
}

func PrintInt64Slice(ns []int64, l, r int) {
	end := r - 1
	for i, n := range ns {
		PrintInt64(n)
		if i != end {
			writer.WriteByte(' ')
		} else {
			writer.WriteByte('\n')
		}
	}
}

func PrintFloat64Slice(fs []float64, prec, l, r int) {
	end := r - 1
	for i, f := range fs {
		PrintFloat64(f, prec)
		if i != end {
			writer.WriteByte(' ')
		} else {
			writer.WriteByte('\n')
		}
	}
}

func PrintStringSlice(ss []string, l, r int) {
	end := r - 1
	for i, s := range ss {
		PrintString(s)
		if i != end {
			writer.WriteByte(' ')
		} else {
			writer.WriteByte('\n')
		}
	}
}
