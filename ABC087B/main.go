package main

import "fmt"

func main() {
	var a, b, c, x int
	fmt.Scan(&a, &b, &c, &x)
	count := 0

	for fiveth := 0; fiveth <= a; fiveth++ {
		for thou := 0; thou <= b; thou++ {
			if remain := x - (500 * fiveth + 100 * thou); remain >= 0 && remain % 50 == 0 && remain / 50 <= c {
				count++
			}
		}
	}
	fmt.Println(count)
}