package main

import "fmt"

func main() {
	var n, ans, i, j int64
	ans = 0
	fmt.Scanf("%d", &n)

	for i = 1; i*i <= n; i++ {
		if n%i == 0 {
			j = n / i
			ans = i + j - 2
		}
	}
	fmt.Printf("%d", ans)
}
