package main

import "fmt"

func toFarenheit(c float64) float64 {
	return (c * 9 / 5) + 32
}

func toCelcius(f float64) float64 {
	return (f - 32) * 5 / 9
}

func main() {

	var ch int
	var temp float64

	fmt.Println("1. To convert temperature to celcius")
	fmt.Println("2. To convert temperature to farenheit")
	fmt.Print("Enter choice : ")
	fmt.Scan(&ch)
	fmt.Print("Enter temperature : ")
	fmt.Scan(&temp)

	if ch == 1 {
		fmt.Printf("Temperature in celcius : %.2f", toCelcius(temp))
	} else if ch == 2 {
		fmt.Printf("Temperature in farenheit : %.2f", toFarenheit(temp))
	} else {
		fmt.Println("Invalid choice")
	}

}
