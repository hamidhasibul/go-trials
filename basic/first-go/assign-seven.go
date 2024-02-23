package main

import "fmt"

// To calculate Fibonacci Series up to n numbers.

func calcFib(n int) {
	num1, num2 := 0, 1

	for i := 0; i < n; i++ {
		fmt.Println(num1 + num2)
		num1, num2 = num2, num1+num2
	}
}
