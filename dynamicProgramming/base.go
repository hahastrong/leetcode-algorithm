package dynamicProgramming

import "fmt"

func fib(n int) int {
	if n == 0 {
		return 0
	}
	pre, cur, tmp := 0, 1, 0
	for i:=2; i<=n; i++ {
		tmp = cur + pre
		pre = cur
		cur = tmp
		fmt.Printf("%d - cur: %d, pre: %d\n", i, cur, pre)
	}
	return cur
}
