package main

import "fmt"

var print = fmt.Println

func main() {
	var l int
	_, err := fmt.Scan(&l)

	if err != nil || l < 0 {
		print("Некорректный ввод")
		return
	}

	if l < 10 {
		print("Число меньше 10")
	} else if l < 100 {
		print("Число меньше 100")
	} else if l < 1000 {
		print("Число меньше 1000")
	} else {
		print("Число больше или равно 1000")
	}
}
