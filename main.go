package main

import "fmt"

var print = fmt.Println

func main() {
	var number int
	fmt.Scan(&number)

	res := number%2 == 0

	if res && number > 0 {
		print("positive and even")
	} else if !res && number > 0 {
		print("positive and not even")
	} else if res && number < 0 {
		print("negative and even")
	} else if !res && number < 0 {
		print("negative and not even")
	} else {
		print("zero")
	}
}
