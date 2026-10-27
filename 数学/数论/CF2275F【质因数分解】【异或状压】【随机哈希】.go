package main

/*
相关考点：
约数个数奇偶性；
完全平方数判定；
质因数分解；
线性筛；
异或状态压缩；
随机哈希；
前缀异或；
哈希表计数；
*/

import (
	"bufio"
	"os"
	"strconv"
	"time"
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

// randseed
var _seed = uint64(time.Now().UnixNano())

func rng() uint64 {
	_seed += 0x9e3779b97f4a7c15
	z := _seed
	z = (z ^ (z >> 30)) * 0xbf58476d1ce4e5b9
	z = (z ^ (z >> 27)) * 0x94d049bb133111eb
	return z ^ (z >> 31)
}

/*
核心结论：判断F(i,j)是否为完全平方数
对x的任意因子d，都有成对的另一个因子x/d。只有当d=x/d时因子数才为奇数
统计各因子数量的奇偶性，要求每个因子数量都为偶数
只需要考虑奇偶性，可以mod2进一步压缩为01两种状态。1表示奇数，0表示偶数
将a0-aj的累乘和ai分别处理为两个01串按位异或
如果结果为0，即两个状态集相等时，则表示所有因子都为偶数，符合条件

ai<=10^6，质因子数量很多64位装不下，需要压缩状态
单个ai质因子数量不超过7（2·3·5·7·11·13·17 = 510510 ≤ 1e6）
*/

func main() {
	defer writer.Flush()

	// logic

	// SPF最小质因子筛
	const maxa = 1000006
	spf := [maxa]int64{}
	primes := []int64{}
	for i := int64(2); i < maxa; i++ {
		if spf[i] == 0 {
			spf[i] = i
			primes = append(primes, i)
		}
		for _, p := range primes {
			if p > spf[i] || p*i >= maxa {
				break
			}
			spf[i*p] = p
		}

	}

	// 双哈希，每个质因子分配一个随机128位值
	h1 := make([]int64, maxa)
	h2 := make([]int64, maxa)
	for _, p := range primes {
		h1[p] = int64(rng())
		h2[p] = int64(rng())
	}

	hashMask := func(x int64) (int64, int64) {
		res1, res2 := int64(0), int64(0)
		for x > 1 {
			p := spf[x]
			// 	计数时只取奇偶性
			isOdd := false
			for x%p == 0 {
				x /= p
				isOdd = !isOdd
			}
			if isOdd {
				res1 ^= h1[p]
				res2 ^= h2[p]
			}
		}
		return res1, res2
	}

	type H struct{ h1, h2 int64 }

	t := NextInt()
	for range t {
		n := NextInt()
		mask1 := make([]int64, n)
		mask2 := make([]int64, n)
		count := make(map[H]int64)
		for i := range n {
			// 直接处理为掩码，不需要原始数据
			a := NextInt64()
			mask1[i], mask2[i] = hashMask(a)
			// 直接统计a[i]的数量避免遍历
			count[H{mask1[i], mask2[i]}]++
		}

		// 前j个数的累乘的因子奇偶性掩码
		multi1, multi2 := int64(0), int64(0)
		res := int64(0)
		for j := range n {
			multi1 = multi1 ^ mask1[j]
			multi2 = multi2 ^ mask2[j]
			res += count[H{multi1, multi2}]
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
