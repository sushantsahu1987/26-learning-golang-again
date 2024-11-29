package main

import "fmt"

func main() {
	var num int
	fmt.Print("Enter a number : ")
	fmt.Scan(&num)

	if num%2 == 1 {
		fmt.Println("Odd")
	} else {
		fmt.Println("Even")
	}
}
