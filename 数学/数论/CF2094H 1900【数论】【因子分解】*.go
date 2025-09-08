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

/*
没懂
*/

func main() {
	defer writer.Flush()

	// logic
	t := NextInt()
	for t > 0 {
		t--
		n, q := NextInt(), NextInt()

		arr := make([]int, n+1)    // 1-indexed
		pos := make(map[int][]int) // pos[x] = 所有 x 出现的位置（递增）
		for i := 1; i <= n; i++ {
			arr[i] = NextInt()
			pos[arr[i]] = append(pos[arr[i]], i)
		}

		// 预处理：对每个值的位置排序（虽然已有序）
		for _, p := range pos {
			sort.Ints(p) // 确保有序
		}

		// ptr[x] 表示值 x 的当前指针（第一个 >= 当前 l 的位置的索引）
		ptr := make(map[int]int)

		// 存储查询：(l, r, k, id)
		type query struct{ l, r, k, id int }
		queries := make([]query, q)
		for i := 0; i < q; i++ {
			k, l, r := NextInt(), NextInt(), NextInt()
			queries[i] = query{l, r, k, i}
		}

		// 按 l 排序，用于离线处理
		sort.Slice(queries, func(i, j int) bool {
			return queries[i].l < queries[j].l
		})

		anss := make([]int64, q)
		prevL := 1

		// 处理每个查询
		for _, qr := range queries {
			l, r, k, id := qr.l, qr.r, qr.k, qr.id

			// 移动指针：将 [prevL, l-1] 的 arr[i] 的 ptr 前移
			for j := prevL; j < l; j++ {
				val := arr[j]
				// ptr[val] 指向 pos[val] 中第一个 >= j 的位置
				// 现在 j 被排除，所以 ptr[val]++
				ptr[val]++
			}
			prevL = l

			// 收集所有因子 d，使得 d 是 k 的因子，且 d 在 [l,r] 内有出现
			type event struct{ pos, val int }
			var events []event

			// 分解 k 的因子
			for d := 1; d*d <= k; d++ {
				if k%d != 0 {
					continue
				}

				// d 是因子
				checkFactor := func(f int) {
					if f == 1 {
						return
					}
					p, exists := pos[f]
					if !exists || ptr[f] >= len(p) {
						return
					}
					firstPos := p[ptr[f]]
					if firstPos > r {
						return
					}
					events = append(events, event{firstPos, f})
				}

				checkFactor(d)
				if d*d != k {
					checkFactor(k / d)
				}
			}

			// 按位置排序
			sort.Slice(events, func(i, j int) bool {
				return events[i].pos < events[j].pos
			})

			// 模拟过程
			var ans int64
			cur := k
			last := l
			for _, e := range events {
				// [last, e.pos) 区间内 cur 不变
				ans += int64(cur) * int64(e.pos-last)
				// 在 e.pos 处除以 e.val
				for cur%e.val == 0 {
					cur /= e.val
				}
				last = e.pos
			}
			// 剩余 [last, r]
			ans += int64(cur) * int64(r-last+1)

			if len(events) == 0 {
				ans = int64(k) * int64(r-l+1)
			}

			anss[id] = ans
		}

		// 输出答案
		for _, ans := range anss {
			PrintInt64Ln(ans)
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
