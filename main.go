package main

import "fmt"

var print = fmt.Println

func main() {
	var l, r int
	fmt.Scan(&l, &r)

	if l > r {
		print("the first is bigger")
	} else if l == r {
		print("equal")
	} else {
		print("the second is bigger")
	}
}
