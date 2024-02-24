package main

import "fmt"

func sumInput() {
	sum := 0
	var n int

	for {

		fmt.Scan(&n)
		sum += n

		if n == 0 {
			break
		}
	}
	fmt.Printf("Sum %v", sum)
}
