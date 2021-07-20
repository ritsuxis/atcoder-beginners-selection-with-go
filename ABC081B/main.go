package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	var ans float64 = 1000000000 + 10
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		var tmp int
		var count float64
		count = 0
		fmt.Scan(&tmp)
		for tmp%2 == 0 {
			count++
			tmp /= 2
		}
		ans = math.Min(count, ans)
	}
	fmt.Println(ans)
}
