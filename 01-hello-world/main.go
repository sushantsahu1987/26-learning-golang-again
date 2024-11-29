package main

import (
	"fmt"
	"strings"
)

func call_hello(s string) {
	fmt.Println(s)
}

func main() {
	call_hello("hello how do you do ?")
	s := "whoistheman?"
	words := strings.Split(s, "t")
	fmt.Println(len(words))
	fmt.Println(words[0])
	fmt.Println(words[1])

	fruits := []string{"apple", "banana", "kiwi"}

	for i, f := range fruits {
		fmt.Printf("index: %d, fruit %s\n", i, f)
	}

	capitals := map[string]string{"india": "delhi", "srilanka": "colombo", "bangladesh": "dhaka"}

	for country, capital := range capitals {
		fmt.Printf("%s %s\n", country, capital)
	}

	count := 4

	if count < 4 {
		fmt.Println("count < 4")
	} else if count == 4 {
		fmt.Println("count == 4")
	} else {
		fmt.Println("count > 4")
	}

}
