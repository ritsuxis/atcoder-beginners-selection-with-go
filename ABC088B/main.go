package main

import (
	"fmt"
	"sort"
)

func main(){
    var n int
    var array[] int
    fmt.Scan(&n)

    for i := 0; i < n; i++ {
        var tmp int
        fmt.Scan(&tmp)
        array = append(array, tmp)
    }

    sort.Sort(sort.Reverse(sort.IntSlice(array)))
    a := 0 
    b := 0
    for i := 0; i < n; i++ {
        if i % 2 == 0 {
            a += array[i]
        } else {
            b += array[i]
        }
    }
    fmt.Println(a - b)

}