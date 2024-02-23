package main

// Take in two numbers and an operator (+, -, *, /) and calculate the value. (Use if conditions)

func calcVal(num1 int, num2 int, op string) int {

	if op == "+" {
		return num1 + num2
	} else if op == "-" {
		return num1 - num2
	} else if op == "*" {
		return num1 * num2
	} else {
		return num1 / num2
	}

}
