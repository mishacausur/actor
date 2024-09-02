package main

import "fmt"

var print = fmt.Println

func main() {
	var number int
	print("Enter the number: ")
	fmt.Scan(&number)

	if number > 0 {
		print("positive")
	} else if number == 0 {
		print("zero")
	} else {
		print("negative")
	}
}
