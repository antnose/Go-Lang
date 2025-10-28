package main

import "fmt"

func addWithReturn(num1 int, num2 int) int {
	sum := num1 + num2
	return sum
}

func getNumbers(num1 int, num2 int) (int, int) {
	sum := num1 + num2
	mul := num1 * num2
	return sum, mul
}

func main() {
	sum := addWithReturn(12, 40)
	fmt.Println(sum)

	su, mul := getNumbers(10, 20)
	fmt.Println(su, mul)
}
