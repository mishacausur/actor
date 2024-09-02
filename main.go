package main

import "fmt"

var print = fmt.Println

func main() {
	var l, r int
	fmt.Scan(&l, &r)

	if l == 0 || r == 0 {
		print("one of them is zero")
	} else if l < 0 && r < 0 {
		print("both are negative")
	} else if l > 0 && r > 0 {
		print("both are positive")
	} else {
		print("one is positive and another on is negative")
	}
}
