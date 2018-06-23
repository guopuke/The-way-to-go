package main

import "fmt"

func main() {
	callback(1, Add)
}

func callback(y int, f func(int, int)) {
	fmt.Println("callback")
	f(y, 2)
}

func Add(a, b int) {
	fmt.Println("Add")
	fmt.Printf("The sum of %d and %d is: %d\n", a, b, a+b)
}
