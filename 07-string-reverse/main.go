package main

import "fmt"

func main() {

	var str string
	rvstr := ""
	fmt.Print("Enter a string :")
	fmt.Scan(&str)

	for _, ch := range str {
		rvstr = string(ch) + rvstr
	}

	fmt.Printf("Reversed string : %s", rvstr)

}
