package main

import "fmt"

func factorial(num int) int {
	result := 1
	for i := 2; i <= num; i++ {
		result *= i
	}
	return result
}

func main() {
	var num int
	fmt.Print("Enter a number")
	fmt.Scan(&num)
	result := factorial(num)
	fmt.Println(result)
}
