package main

// Subtract the Product and Sum of Digits of an Integer

func subtractProductAndSum(n int) int {

	sum, prod := 0, 1

	for n > 0 {
		currDig := n % 10
		n = n / 10
		sum += currDig
		prod *= currDig
	}

	return prod - sum
}
