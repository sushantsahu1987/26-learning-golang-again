package main

import "fmt"

func isPrimeNumber(num int) bool {
	if num < 2 {
		return false
	}
	for i := 2; i*i <= num; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	var num int
	fmt.Print("Enter a number : ")
	fmt.Scan(&num)

	if isPrimeNumber(num) {
		fmt.Println("Is a Prime no.")
	} else {
		fmt.Println("Is not a Prime no.")
	}

}
