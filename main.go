package main

import "fmt"

var print = fmt.Println

func main() {
	var number int
	print("Enter the number: ")
	fmt.Scan(&number)

	if res := number * number; res >= 50 {
		print("min action")
	} else {
		print("extra action")
	}
}
