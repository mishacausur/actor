package main

import "fmt"

var print = fmt.Println

func main() {
	var l, r, t int
	_, err := fmt.Scan(&l, &r, &t)

	if err != nil {
		print("invalid enter")
		return
	}

	if l == r && r == t {
		print("equal")
	} else if l == r || r == t || t == l {
		print("two are equal")
	} else if l != r && r != t && l != t {
		print("different")
	} else {
		print("invalid enter")
	}
}
