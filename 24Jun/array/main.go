package main

import (
	"fmt"
	"time"
)

func main() {
	a := [...]string{"a", "b", "c", "d"}
	for i, _ := range a {
		fmt.Println("Array item", i, "is", a[i])
	}
	fmt.Println()

	forArray()

	fmt.Println()
	fmt.Println("fibonacci")
	fibonacci()
}

func forArray() {
	var ar [16]int
	fmt.Println(ar)
	for i := 0; i < len(ar); i++ {
		ar[i] = i
	}
	fmt.Println(ar)
}

func fibonacci() {
	var ar [50]int
	start := time.Now()
	for i := 0; i < len(ar); i++ {
		if i < 2 {
			ar[i] = 1
		} else {
			ar[i] = ar[i-1] + ar[i-2]
		}
		fmt.Println(ar)
	}
	end := time.Now()
	de := end.Sub(start)
	fmt.Println(de)
}
