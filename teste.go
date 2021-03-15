package main

import (
	"fmt"
	"sort"
	"time"
)

func main() {
	start := time.Now()
	// start--------------------
	arr1 := [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	arr2 := [10]int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}
	var arr3 []int

	for k := 0; k < 1000000; k++ {
		for i := 0; i < 10; i++ {
			arr3 = append(arr3, arr1[i]^i)
			arr3 = append(arr3, (arr1[i]+arr2[i])*i)
		}
	}

	sort.Ints(arr3)

	println(len(arr3))

	end := time.Now()
	// end----------------------
	println(fmt.Printf("The call took %v to run.\n", start.Sub(end)))
}
