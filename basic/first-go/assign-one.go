package main

// Write a program to print whether a number is even or odd, also take input from the user.

func evenOdd(num int) string {
	if num == 0 || num%2 == 0 {
		return "even"
	}

	return "odd"
}
