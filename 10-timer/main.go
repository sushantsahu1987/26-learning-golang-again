package main

import (
	"fmt"
	"time"
)

func main() {
	var t int
	fmt.Print("Enter seconds in timer")
	fmt.Scan(&t)

	for t > 0 {
		fmt.Println(t)
		time.Sleep(1 * time.Second)
		t--
	}

	fmt.Println("Time in seconds")

}
