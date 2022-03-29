package main

import "fmt"

func main() {

	//Print from 0-1000
	for i := 0; i <= 1000; i++ {
		fmt.Println(i)
	}
	// Print only Odd number
	for i := 0; i <= 1000; i++ {

		if i%2 == 1 {
			fmt.Println(i)
		}

	}
	// Print only even number
	for i := 0; i <= 1000; i++ {

		if i%2 == 0 {
			fmt.Println(i)
		}

	}
}
