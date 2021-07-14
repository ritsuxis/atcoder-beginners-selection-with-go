package main

import "fmt"

func main(){
	var a, b int
	fmt.Scan(&a, &b)
	
	if ans := a * b; ans % 2 == 0 {
        fmt.Print("Even")
	} else {
        fmt.Print("Odd")
    }
}