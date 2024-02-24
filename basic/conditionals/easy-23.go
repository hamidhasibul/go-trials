package main

import "fmt"

// Input a number and print all the factors of that number (use loops). Printing all nums except 1 and n

func factors(n int) {

	for i := 2; i <= n/2; i++ {
		if n%i == 0 {
			fmt.Println(i)
		}
	}
}
