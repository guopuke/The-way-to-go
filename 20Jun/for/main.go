package main

import "fmt"

func main() {
	// for i := 0; ; i++ {
	// 	fmt.Println("Value of i")
	//
	// }

	// s := ""
	// for ; s != "aaa"; {
	// 	fmt.Println("Value of s:", s)
	// 	s = s + "a"
	// }

	// for i, j, s := 0, 5, "a"; i < 3 && j < 100 && s != "aaaaa"; i, j, s = i+1, j+1, s+"a" {
	// 	fmt.Println("Value of i, j, s:", i, j, s)
	// }

	// i := 0
	// for {
	// 	if i > 3 {
	// 		break
	// 	}
	// 	fmt.Println("Value of i is:", i)
	// 	i++
	// }

	for i := 0; i < 7; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println("Odd:", i)
	}
}
