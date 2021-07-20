package main

import (
	"fmt"
	// "strconv"
)

// func main() {
// 	var n, a, b int
// 	count := 0
// 	fmt.Scan(&n, &a, &b)

// 	for i := 1; i <= n; i++ {
// 		str := fmt.Sprintf("%05d", i)
// 		i1, _ := strconv.Atoi(str[:1])
// 		i2, _ := strconv.Atoi(str[1:2])
// 		i3, _ := strconv.Atoi(str[2:3])
// 		i4, _ := strconv.Atoi(str[3:4])
// 		i5, _ := strconv.Atoi(str[4:5])
// 		if ans := i1 + i2 + i3 + i4 + i5; ans >= a && ans <= b {
// 			count += i
// 		}
// 	}
// 	fmt.Println(count)
// }

func main() {
	var n, a, b int
	count := 0
	fmt.Scan(&n, &a, &b)

	for i := 1; i <= n; i++ {
		sum := 0
		for now := i; now > 0; now /= 10{
			sum += now % 10
		}
		if sum >= a && sum <= b {
			count += i
		}
	}
	fmt.Println(count)
}